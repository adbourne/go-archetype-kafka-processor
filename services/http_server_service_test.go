package services

import (
	"testing"
	"github.com/adbourne/go-archetype-rest/config"
	"github.com/stretchr/testify/assert"
)

func TestNewDefaultHTTPServerReturnsNewHTTPServer(t *testing.T) {
	result := NewDefaultHTTPServer(1234, config.NewLogger())
	assert.NotNil(t, result)
}
