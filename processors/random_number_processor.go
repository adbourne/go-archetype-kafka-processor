package processors

import (
	"github.com/adbourne/go-archetype-kafka-processor/messages"
	"github.com/adbourne/go-archetype-kafka-processor/services"
)

// RandomNumberProcessor is a kafka processor that generates a random number
type RandomNumberProcessor struct {
	Logger              services.Logger
	RandomNumberService services.RandomNumberService
}

// NewRandomNumberProcessor creates a new RandomNumberProcessor
func NewRandomNumberProcessor(logger services.Logger, randomNumberService services.RandomNumberService) *RandomNumberProcessor {
	return &RandomNumberProcessor{
		Logger: logger,
		RandomNumberService: randomNumberService,
	}
}

// Process takes a SourceMessage, generates a random number and uses it to create a SinkMessage
func (rnp RandomNumberProcessor) Process(sourceMessage messages.SourceMessage) messages.SinkMessage {
	rnp.Logger.Debug("Generating new random number...")
	randomNumber := rnp.RandomNumberService.GenerateRandomNumber(sourceMessage.Seed)
	return messages.SinkMessage{
		RandomNumber: randomNumber,
	}
}
