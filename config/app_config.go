package config

import (
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	// EnvVarPort is the environment variable of the HTTP port to use
	EnvVarPort = "KAFKA_PROCESSOR_ARCHETYPE_PORT"

	// EnvVarRandomSeed is the environment variable for the random seed to use
	EnvVarRandomSeed = "KAFKA_PROCESSOR_ARCHETYPE_RANDOM_SEED"

	// EnvVarKafkaBrokers is the environment variable for the comma delimited list of kafka brokers
	EnvVarKafkaBrokers = "KAFKA_PROCESSOR_ARCHETYPE_KAFKA_BROKERS"

	// EnvVarSourceTopic is the environment variable for the source Kafka topic
	EnvVarSourceTopic = "KAFKA_PROCESSOR_ARCHETYPE_SOURCE_TOPIC"

	// EnvVarSinkTopic is the environment variable for the Kafka sink topic
	EnvVarSinkTopic = "KAFKA_PROCESSOR_ARCHETYPE_SINK_TOPIC"
)

// AppConfig is the application configuration
type AppConfig struct {
	// Port is the port to run the rest server on
	Port int

	// RandomSeed is the seed to use when generating randomness
	RandomSeed int64

	// Brokers is the comma delimited list of kafka brokers
	Brokers string

	// SourceTopic is the name of the Kafka source topic
	SourceTopic string

	// SinkTopic is the name of the Kafka sink topic
	SinkTopic string
}

// NewAppConfig constructs a new AppConfig
func NewAppConfig() AppConfig {
	return AppConfig{
		Port:        loadEnvVarAsInt(EnvVarPort, 8080),
		RandomSeed:  int64(loadEnvVarAsInt(EnvVarRandomSeed, 1)),
		Brokers:     loadEnvVarAsString(EnvVarKafkaBrokers, "localhost:9092"),
		SourceTopic: loadEnvVarAsString(EnvVarSourceTopic, "source-topic"),
		SinkTopic:   loadEnvVarAsString(EnvVarSinkTopic, "sink-topic"),
	}
}

// GetBrokerList splits the comma delimited broker string into a slice of brokers
func (c *AppConfig) GetBrokerList() []string {
	return strings.Split(c.Brokers, ",")
}

// loadEnvVarAsInt is a utility function for loading an environment variable or use a default
func loadEnvVarAsInt(envVarName string, defaultVal int) int {
	ev, isFound := os.LookupEnv(envVarName)
	if !isFound {
		log.Printf("Environment variable '%s' not found, using default '%d'", envVarName, defaultVal)
		return defaultVal
	}

	evi, err := strconv.Atoi(ev)
	if err != nil {
		log.Fatalf("Environment variable '%s' was not a number", envVarName)
		return defaultVal
	}

	log.Printf("Environment variable '%s' found", envVarName)
	return evi
}

// loadEnvVarAsString is a utility function for loading an environment variable or using a default
func loadEnvVarAsString(envVarName string, defaultVal string) string {
	ev, isFound := os.LookupEnv(envVarName)
	if !isFound {
		log.Printf("Environment variable '%s' not found, using default '%s'", envVarName, defaultVal)
		return defaultVal
	}

	log.Printf("Environment variable '%s' found", envVarName)
	return ev
}
