package config

import (
	"os"
	"log"
	"strconv"
	"strings"
)

const (
	// Environment variable of the HTTP port to use
	EnvVarPort = "REST_ARCHETYPE_PORT"

	// Environment variable for the random seed to use
	EnvVarRandomSeed = "REST_ARCHETYPE_RANDOM_SEED"

	// Environment variable for the comma delimited list of kafka brokers
	EnvVarKafkaBrokers = "REST_ARCHETYPE_KAFKA_BROKERS"

	// Environment variable for the source Kafka topic
	EnvVarSourceTopic = "REST_ARCHETYPE_SOURCE_TOPIC"

	// Environment variable for the Kafka sink topic
	EnvVarSinkTopic = "REST_ARCHETYPE_SINK_TOPIC"
)


// The application configuration
type AppConfig struct {
	// The port to run the rest server on
	Port        int

	// The seed to use when generating randomness
	RandomSeed  int64

	// The comma delimited list of kafka brokers
	Brokers     string

	// The name of the Kafka source topic
	SourceTopic string

	// The name of the Kafka sink topic
	SinkTopic   string
}

// Constructs a new AppConfig
func NewAppConfig() AppConfig {
	return AppConfig{
		Port: loadEnvVarAsInt(EnvVarPort, 8080),
		RandomSeed: int64(loadEnvVarAsInt(EnvVarRandomSeed, 1)),
		Brokers: loadEnvVarAsString(EnvVarKafkaBrokers, "localhost:9092"),
		SourceTopic: loadEnvVarAsString(EnvVarSourceTopic, "source-topic"),
		SinkTopic: loadEnvVarAsString(EnvVarSinkTopic, "sink-topic"),
	}
}

// Splits the comma delimited broker string into a slice of brokers
func (c *AppConfig) GetBrokerList() []string {
	return strings.Split(c.Brokers, ",")
}

// Utility function for loading an environment variable or use a default
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

// Utility function for loading an environment variable or using a default
func loadEnvVarAsString(envVarName string, defaultVal string) string {
	ev, isFound := os.LookupEnv(envVarName)
	if !isFound {
		log.Printf("Environment variable '%s' not found, using default '%s'", envVarName, defaultVal)
		return defaultVal
	}

	log.Printf("Environment variable '%s' found", envVarName)
	return ev
}