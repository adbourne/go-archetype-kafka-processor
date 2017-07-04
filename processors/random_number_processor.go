package processors

import (
	"github.com/adbourne/go-archetype-rest/config"
	"github.com/adbourne/go-archetype-rest/messages"
	"github.com/adbourne/go-archetype-rest/services"
)

type RandomNumberProcessor struct {
	RandomNumberService services.RandomNumberService
}

func (rnp RandomNumberProcessor) Process(sourceMessage messages.SourceMessage) messages.SinkMessage {
	logger := config.NewLogger()
	logger.Debug("Generating new random number...")
	randomNumber := rnp.RandomNumberService.GenerateRandomNumber(sourceMessage.Seed)
	return messages.SinkMessage{
		RandomNumber: randomNumber,
	}
}
