package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewDefaultHTTPServerReturnsNewHTTPServer(t *testing.T) {
	result := NewDefaultHTTPServer(1234, NewSystemOutLogger())
	assert.NotNil(t, result)
}
