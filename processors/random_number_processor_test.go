package processors

import (
	"fmt"
	"github.com/adbourne/go-archetype-kafka-processor/messages"
	"github.com/adbourne/go-archetype-kafka-processor/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessWithSeedGeneratesARandomNumber(t *testing.T) {
	expectedSeed := 1
	mockRandomNumber := 1234

	mockRandomNumberService := &MockRandomNumberService{}
	mockRandomNumberService.ExpectSeed(expectedSeed).ThenReturn(mockRandomNumber)

	randomNumberProcessor := NewRandomNumberProcessor(services.NewSystemOutLogger(), mockRandomNumberService)

	expected := messages.SinkMessage{
		RandomNumber: mockRandomNumber,
	}

	result := randomNumberProcessor.Process(messages.SourceMessage{
		Seed: expectedSeed,
	})

	assert.Equal(t, expected, result)
}

// Mock implementing the RandomNumberService interface
type MockRandomNumberService struct {
	expectedSeed int
	randomNumber int
}

func (rns *MockRandomNumberService) ExpectSeed(seed int) *MockRandomNumberService {
	rns.expectedSeed = seed
	return rns
}

// The "random" number to return
func (rns *MockRandomNumberService) ThenReturn(toReturn int) *MockRandomNumberService {
	rns.randomNumber = toReturn
	return rns
}

// Returns the "random" number only if the seed matches what is expected
func (rns *MockRandomNumberService) GenerateRandomNumber(seed int) int {
	if seed == rns.expectedSeed {
		return rns.randomNumber
	}
	panic(fmt.Sprintf("Provided seed '%d' did not match expected seed '%d'", seed, rns.expectedSeed))
}
