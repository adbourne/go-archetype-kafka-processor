package main

import (
	"github.com/adbourne/go-archetype-kafka-processor/config"
	"github.com/stretchr/testify/mock"
	"testing"
)

const (
	Broker      = "localhost"
	SourceTopic = "source-topic"
	SinkTopic   = "sink-topic"
)

func TestApplicationIsStartedCorrectly(t *testing.T) {
	mockKafkaClient := new(MockKafkaClient)
	mockKafkaClient.On("RegisterProcessor", mock.Anything).Return()
	mockKafkaClient.On("Process").Return(nil)

	MockHTTPServer := new(MockHTTPServer)
	MockHTTPServer.On("RegisterEndpoint", "/health", mock.Anything).Return()
	MockHTTPServer.On("Run").Return()

	mockAppContext := newTestAppContext(t, mockKafkaClient, MockHTTPServer)
	RunApp(mockAppContext)

	mockKafkaClient.AssertExpectations(t)
	MockHTTPServer.AssertExpectations(t)
}

func newTestAppContext(t *testing.T, mkc *MockKafkaClient, mhs *MockHTTPServer) *AppContext {
	appConfig := newTestAppConfig()
	randomNumberService := newRandomNumberService()
	return &AppContext{
		AppConfig:           appConfig,
		RandomNumberService: randomNumberService,
		KafkaClient:         mkc,
		KafkaProcessor:      newKafkaProcessor(randomNumberService),
		HTTPServer:          mhs,
	}
}

type MockKafkaClient struct {
	mock.Mock
}

func (mkc *MockKafkaClient) RegisterProcessor(kp config.KafkaProcessor) {
	mkc.Called(kp)
}

func (mkc *MockKafkaClient) Process() error {
	args := mkc.Called()
	return args.Error(0)
}

func (mkc *MockKafkaClient) Close() {
	mkc.Called()
}

type MockHTTPServer struct {
	mock.Mock
}

func (mhs *MockHTTPServer) RegisterEndpoint(endpoint string, handler config.Handler) {
	mhs.Called(endpoint, handler)
}

// Runs the HTTP server
func (mhs *MockHTTPServer) Run() {
	mhs.Called()
}

func newTestAppConfig() config.AppConfig {
	return config.AppConfig{
		Port:        8080,
		RandomSeed:  0,
		Brokers:     Broker,
		SourceTopic: SourceTopic,
		SinkTopic:   SinkTopic,
	}
}
