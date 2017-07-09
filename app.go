package main

import (
	"fmt"
	"github.com/adbourne/go-archetype-kafka-processor/config"
	"github.com/adbourne/go-archetype-kafka-processor/contollers"
	"io/ioutil"
	"log"
	"path/filepath"
)

const (
	// BannerFileName is the file containing the banner
	BannerFileName = "banner.txt"

	// DefaultBannerText is the default text to be displayed when there is no banner
	DefaultBannerText = "Kafka Archetype starting..."
)

func main() {
	printBanner()

	// Load the application config
	appConfig := config.NewAppConfig()

	// Create the application context
	appContext := NewAppContext(appConfig)

	// Run the application
	RunApp(appContext)
}

// RunApp runs the application
func RunApp(appContext *AppContext) {
	logger := appContext.Logger

	logger.Debug(fmt.Sprintf("%+v\n", appContext.AppConfig))

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
	httpServer := appContext.HTTPServer
	httpServer.RegisterEndpoint("/health", controllers.HealthEndpoint)
	httpServer.Run()
}

func printBanner() {
	relativePath, err := filepath.Abs(BannerFileName)
	if err != nil {
		log.Println(DefaultBannerText)
		return
	}

	contents, err := ioutil.ReadFile(relativePath)
	if err != nil {
		log.Println(DefaultBannerText)
		return
	}

	log.Println(string(contents))

}
