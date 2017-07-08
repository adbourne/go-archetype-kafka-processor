// +build integration
package main

import (
	"gopkg.in/ory-am/dockertest.v3"
	"log"
	"fmt"
	"os"
	"strconv"
	"time"
	"testing"
	"github.com/adbourne/go-archetype-kafka-processor/config"
)

const (
	DOCKER_IMAGE_NAME_KAFKA = "spotify/kafka"
	DOCKER_IMAGE_VERSION_KAFKA = "latest"
)

var kafkaPort int

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	kafkaResource, err := createKafkaContainer(pool)
	if err != nil {
		log.Fatalf("Could not start Kafka resource: %s", err)
	}
	time.Sleep(30 * time.Second)
	kafkaPort, _ := strconv.Atoi(kafkaResource.GetPort("9092/tcp"))
	log.Printf("Kafka port is %d", kafkaPort)


	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		log.Println("Waiting for services to start...")
		time.Sleep(30 * time.Second)
		return nil
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(kafkaResource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func createKafkaContainer(pool *dockertest.Pool) (kafkaResource *dockertest.Resource, err error) {
	log.Print("Starting Kafka...")
	runOptions := &dockertest.RunOptions{
		Repository: DOCKER_IMAGE_NAME_KAFKA,
		Tag: DOCKER_IMAGE_VERSION_KAFKA,
		Env: []string{
			"ADVERTISED_HOST=localhost",
			"ADVERTISED_PORT=9092",
		},
		//Entrypoint: ,
		//Cmd: ,
		//Mounts: ,
		//Links: ,
		//ExposedPorts: []string{"9092"},

	}
	kafkaResource, err = pool.RunWithOptions(runOptions)
	return
}

func TestAppStartsCorrectly(t *testing.T) {
	appConfig := &config.AppConfig{
		Port:        8080,
		RandomSeed:  int64(1),
		Brokers:     fmt.Sprintf("localhost:%d", kafkaPort),
		SourceTopic: "source-topic",
		SinkTopic:   "sink-topic",
	}

	appContext := NewAppContext(*appConfig)

	RunApp(appContext)
}
