package main

import (
	"github.com/adbourne/go-archetype-kafka-processor/config"
	"github.com/adbourne/go-archetype-kafka-processor/processors"
	"github.com/adbourne/go-archetype-kafka-processor/services"
)

// AppContext is the application context
type AppContext struct {
	// AppConfig is the application config
	AppConfig config.AppConfig

	// RandomNumberService is the random number service
	RandomNumberService services.RandomNumberService

	// KafkaClient is the Kafka client
	KafkaClient config.KafkaClient

	// KafkaProcessor is the Kafka Processor
	KafkaProcessor config.KafkaProcessor

	// HTTPServer is the HTTP server
	HTTPServer config.HTTPServer
}

// NewAppContext creates the application context
func NewAppContext(appConfig config.AppConfig) *AppContext {
	randomNumberService := newRandomNumberService()
	return &AppContext{
		AppConfig:           appConfig,
		RandomNumberService: randomNumberService,
		KafkaClient:         newKafkaClient(appConfig),
		KafkaProcessor:      newKafkaProcessor(randomNumberService),
		HTTPServer:          newHTTPServer(appConfig),
	}
}

// Creates a RandomNumberService
func newRandomNumberService() *services.DefaultRandomNumberService {
	return &services.DefaultRandomNumberService{}
}

// Creates a KafkaClient
func newKafkaClient(appConfig config.AppConfig) config.KafkaClient {
	return config.NewSaramaKafkaClient(appConfig)
}

// Creates a new KafkaProcessor
func newKafkaProcessor(randomNumberService services.RandomNumberService) config.KafkaProcessor {
	return processors.RandomNumberProcessor{
		RandomNumberService: randomNumberService,
	}
}

// Creates the HTTP server
func newHTTPServer(appConfig config.AppConfig) config.HTTPServer {
	return config.NewDefaultHTTPServer(appConfig.Port)
}
