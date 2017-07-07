package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewSystemOutLoggerCreatesNewSystemOutLogger(t *testing.T) {
	result := NewSystemOutLogger()
	assert.NotNil(t, result)
}
