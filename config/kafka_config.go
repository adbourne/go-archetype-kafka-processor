package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/adbourne/go-archetype-rest/messages"
	"log"
	"time"
)

// A KafkaProcessor turns a Source Message into a Sink Message
type KafkaProcessor interface {
	Process(messages.SourceMessage) messages.SinkMessage
}

// Abstraction away from the kafka client implementation
type KafkaClient interface {
	// Registers a Kafka Processor
	RegisterProcessor(kp KafkaProcessor)

	// Uses the registered KafkaProcessor to processes messages from the
	// source topic and publish to the sink topic
	// returns an error if no KafkaProcessor is registered
	Process() error

	// Closes the connection to Kafka
	Close()
}

// Sarama implementation of KafkaClient
type SaramaKafkaClient struct {
	Consumer  sarama.Consumer
	Producer  sarama.AsyncProducer
	AppConfig AppConfig
	processor KafkaProcessor
	Logger    Logger
}

// Registers a Kafka Processor
func (kc *SaramaKafkaClient) RegisterProcessor(kp KafkaProcessor) {
	if kc.Logger != nil {
		kc.Logger.Debug("Registering Kafka processor...")
	}
	kc.processor = kp
}

// Processes source messages into published sink messages
func (kc SaramaKafkaClient) Process() error {
	if kc.processor != nil {
		if kc.Logger != nil {
			kc.Logger.Debug("Starting to process messages...")
		}
		messagesInChan := kc.createConsumerMessageChannel(kc.Consumer, kc.AppConfig.SourceTopic)
		go processMessages(messagesInChan, kc.Producer, kc.AppConfig.SinkTopic, kc.processor)
		return nil
	}
	return errors.New("Processor was nil")
}

func processMessages(inMessages <-chan *sarama.ConsumerMessage, producer sarama.AsyncProducer, sinkTopic string, processor KafkaProcessor) {
	inMessage := <-inMessages
	inMessageValue := inMessage.Value

	var srcMsg *messages.SourceMessage

	err := json.Unmarshal(inMessageValue, &srcMsg)
	if err != nil {
		log.Printf("Received invalid message '%s'", inMessageValue)

	} else {
		logger := NewLogger()
		sinkMessage := processor.Process(*srcMsg)
		encodedMsg, _ := sinkMessage.Encode()
		logger.Debug(fmt.Sprintf("Sending sink message '%s'", encodedMsg))

		producer.Input() <- &sarama.ProducerMessage{
			Topic: sinkTopic,
			Key:   nil, // TODO: Choose a key?
			Value: sinkMessage,
		}
	}
}

// SaramaKafkaClient's implementation of KafkaClient's Close function
func (kc SaramaKafkaClient) Close() {
	logger := NewLogger()
	logger.Trace("Closing Sarama Kafka Client...")
	if kc.Consumer.Close() != nil {
		logger.Warn("Unable to close Sarama consumer")
	}
	if kc.Producer.Close() != nil {
		logger.Warn("Unable to close Sarama producer")
	}
}

func NewSaramaKafkaClient(appConfig AppConfig) *SaramaKafkaClient {
	return &SaramaKafkaClient{
		Consumer: newSaramaKafkaConsumer(appConfig.GetBrokerList()),
		Producer: newKafkaProducer(appConfig.GetBrokerList()),
		Logger:   NewLogger(),
	}
}

// Construct a new Kafka Consumer listening to the supplied brokers
func newSaramaKafkaConsumer(kafkaBrokers []string) sarama.Consumer {
	consumer, err := sarama.NewConsumer(kafkaBrokers, nil)
	if err != nil {
		log.Fatalln("Failed to start Sarama consumer:", err)
	}

	return consumer
}

func (skc SaramaKafkaClient) createConsumerMessageChannel(consumer sarama.Consumer, sourceTopic string) <-chan *sarama.ConsumerMessage {
	logger := NewLogger()
	logger.Debug(fmt.Sprintf("Finding partitions for source topic '%s'...", sourceTopic))
	partitionList := skc.getConsumerPartitions(consumer, sourceTopic)
	logger.Trace(fmt.Sprintf("Found %d partitions!", len(partitionList)))

	// Channel to publish messages to
	messageChan := make(chan *sarama.ConsumerMessage, 256)

	for index, partition := range partitionList {
		logger.Trace(fmt.Sprintf("Consuming from partition %d [%d] from source topic '%s'...", index, partition, sourceTopic))
		pc, err := consumer.ConsumePartition(sourceTopic, partition, sarama.OffsetOldest)
		if err != nil {
			panic(err)
		}
		go skc.consumeMessage(pc, messageChan)
	}

	return messageChan
}

func (skc SaramaKafkaClient) getConsumerPartitions(consumer sarama.Consumer, sourceTopic string) []int32 {
	logger := NewLogger()
	logger.Trace(fmt.Sprintf("Getting consumer partitions for topic '%s'", sourceTopic))
	partitionList, err := consumer.Partitions(sourceTopic) // get all partitions
	if err != nil {
		log.Fatalln("Failed to get consumer partitions:", err)
	}
	return partitionList
}

func (skc SaramaKafkaClient) consumeMessage(pc sarama.PartitionConsumer, messages chan<- *sarama.ConsumerMessage) {
	logger := NewLogger()
	for message := range pc.Messages() {
		if message != nil {
			logger.Debug(fmt.Sprintf("Received message %s:"+string(message.Value), message))
			messages <- message
		}
	}
}

// Construct a new Kafka Producer
func newKafkaProducer(brokers []string) sarama.AsyncProducer {
	producerConfig := sarama.NewConfig()
	producerConfig.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	producerConfig.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	producerConfig.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms

	producer, err := sarama.NewAsyncProducer(brokers, producerConfig)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}

	// We will just log to STDOUT if we're not able to produce messages.
	// Note: messages will only be returned here after all retry attempts are exhausted.
	go func() {
		for err := range producer.Errors() {
			log.Println("Failed to write message:", err)
		}
	}()

	return producer
}
