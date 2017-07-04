package main

import (
	"github.com/adbourne/go-archetype-rest/config"
	"github.com/adbourne/go-archetype-rest/processors"
	"github.com/adbourne/go-archetype-rest/services"
)

// NewAppContext creates the application context
func NewAppContext(appConfig config.AppConfig) *config.AppContext {
	randomNumberService := newRandomNumberService()
	return &config.AppContext{
		AppConfig:           appConfig,
		RandomNumberService: randomNumberService,
		KafkaClient:         newKafkaClient(appConfig),
		KafkaProcessor:      newKafkaProcessor(randomNumberService),
		HttpServer:          newHTTPServer(appConfig),
	}
}

// Creates a RandomNumberService
func newRandomNumberService() *services.DefaultRandomNumberService {
	return &services.DefaultRandomNumberService{}
}

// Creates a KafkaClient
func newKafkaClient(appConfig config.AppConfig) config.KafkaClient {
	return config.NewSaramaKafkaClient(appConfig)
}

// Creates a new KafkaProcessor
func newKafkaProcessor(randomNumberService services.RandomNumberService) config.KafkaProcessor {
	return processors.RandomNumberProcessor{
		RandomNumberService: randomNumberService,
	}
}

// Creates the HTTP server
func newHTTPServer(appConfig config.AppConfig) config.HttpServer {
	return config.NewDefaultHttpServer(appConfig.Port)
}
