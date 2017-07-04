package main

import (
	"github.com/adbourne/go-archetype-rest/config"
	"github.com/adbourne/go-archetype-rest/contollers"
	"log"
)

func main() {
	// Load the application config
	appConfig := config.NewAppConfig()

	// Create the application context
	appContext := NewAppContext(appConfig)

	// Run the application
	RunApp(appContext)
}

// Runs the application
func RunApp(appContext *config.AppContext) {
	logger := config.NewLogger()

	printBanner()

	logger.Info("Starting application...")

	// Connect to Kafka
	logger.Debug("Connecting to Kafka...")
	kafkaClient := appContext.KafkaClient
	kafkaClient.RegisterProcessor(appContext.KafkaProcessor)
	err := kafkaClient.Process()
	if err != nil {
		log.Panic("A Kafka processor was not specified: ", err)
	}

	// Start the HTTP service
	logger.Debug("Starting HTTP server...")
	httpServer := appContext.HttpServer
	httpServer.RegisterEndpoint("/health", controllers.HealthEndpoint)
	httpServer.Run()
}

func printBanner() {
	log.Println(
		`
 ___ ___ ___ _____     _          _        _
| _ \ __/ __|_   _|   /_\  _ _ __| |_  ___| |_ _  _ _ __  ___
|   / _|\__ \ | |    / _ \| '_/ _| ' \/ -_)  _| || | '_ \/ -_)
|_|_\___|___/ |_|   /_/ \_\_| \__|_||_\___|\__|\_, | .__/\___|
                                               |__/|_|
		`)
}
