package processors

import (
	"github.com/adbourne/go-archetype-kafka-processor/config"
	"github.com/adbourne/go-archetype-kafka-processor/messages"
	"github.com/adbourne/go-archetype-kafka-processor/services"
)

// RandomNumberProcessor is a kafka processor that generates a random number
type RandomNumberProcessor struct {
	RandomNumberService services.RandomNumberService
}

// Process takes a SourceMessage, generates a random number and uses it to create a SinkMessage
func (rnp RandomNumberProcessor) Process(sourceMessage messages.SourceMessage) messages.SinkMessage {
	logger := config.NewLogger()
	logger.Debug("Generating new random number...")
	randomNumber := rnp.RandomNumberService.GenerateRandomNumber(sourceMessage.Seed)
	return messages.SinkMessage{
		RandomNumber: randomNumber,
	}
}
