package config

import (
	"github.com/adbourne/go-archetype-rest/services"
)

// AppContext is the application context
type AppContext struct {
	// AppConfig is the application config
	AppConfig AppConfig

	// RandomNumberService is the random number service
	RandomNumberService services.RandomNumberService

	// KafkaClient is the Kafka client
	KafkaClient KafkaClient

	// KafkaProcessor is the Kafka Processor
	KafkaProcessor KafkaProcessor

	// HTTPServer is the HTTP server
	HTTPServer HTTPServer
}
