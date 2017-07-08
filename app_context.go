package main

import (
	"github.com/adbourne/go-archetype-kafka-processor/config"
	"github.com/adbourne/go-archetype-kafka-processor/processors"
	"github.com/adbourne/go-archetype-kafka-processor/services"
)

// AppContext is the application context
type AppContext struct {
	// AppConfig is the application config
	AppConfig config.AppConfig

	// Logger is the logger
	Logger services.Logger

	// RandomNumberService is the random number service
	RandomNumberService services.RandomNumberService

	// KafkaClient is the Kafka client
	KafkaClient services.KafkaClient

	// KafkaProcessor is the Kafka Processor
	KafkaProcessor services.KafkaProcessor

	// HTTPServer is the HTTP server
	HTTPServer services.HTTPServer
}

// NewAppContext creates the application context
func NewAppContext(appConfig config.AppConfig) *AppContext {
	logger := newLogger()
	randomNumberService := newRandomNumberService()
	return &AppContext{
		AppConfig:           appConfig,
		RandomNumberService: randomNumberService,
		KafkaClient:         newKafkaClient(appConfig, logger),
		KafkaProcessor:      newKafkaProcessor(randomNumberService),
		HTTPServer:          newHTTPServer(appConfig, logger),
	}
}

// Creates a new Logger
func newLogger() services.Logger {
	return &services.SystemOutLogger{}
}

// Creates a RandomNumberService
func newRandomNumberService() *services.DefaultRandomNumberService {
	return &services.DefaultRandomNumberService{}
}

// Creates a KafkaClient
func newKafkaClient(appConfig config.AppConfig, logger services.Logger) services.KafkaClient {
	return services.NewSaramaKafkaClient(appConfig, logger)
}

// Creates a new KafkaProcessor
func newKafkaProcessor(randomNumberService services.RandomNumberService) services.KafkaProcessor {
	return processors.RandomNumberProcessor{
		RandomNumberService: randomNumberService,
	}
}

// Creates the HTTP server
func newHTTPServer(appConfig config.AppConfig, logger services.Logger) services.HTTPServer {
	return services.NewDefaultHTTPServer(appConfig.Port, logger)
}
