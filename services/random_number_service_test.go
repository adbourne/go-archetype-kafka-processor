package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomNumberServiceGenerateRandomNumberGeneratesARandomNumber(t *testing.T) {
	randomNumberService := DefaultRandomNumberService{}
	seed := 1
	expected := 5577006791947779410

	assert.Equal(t, expected, randomNumberService.GenerateRandomNumber(seed))
}
