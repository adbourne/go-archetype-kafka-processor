package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSystemOutLoggerCreatesNewSystemOutLogger(t *testing.T) {
	result := NewSystemOutLogger()
	assert.NotNil(t, result)
}
