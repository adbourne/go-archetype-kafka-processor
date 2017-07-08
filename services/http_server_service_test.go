package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDefaultHTTPServerReturnsNewHTTPServer(t *testing.T) {
	result := NewDefaultHTTPServer(1234, NewSystemOutLogger())
	assert.NotNil(t, result)
}
