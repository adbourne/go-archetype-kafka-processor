package main

import (
	"github.com/adbourne/go-archetype-rest/config"
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

	MockHttpServer := new(MockHttpServer)
	MockHttpServer.On("RegisterEndpoint", "/health", mock.Anything).Return()
	MockHttpServer.On("Run").Return()

	mockAppContext := newTestAppContext(t, mockKafkaClient, MockHttpServer)
	RunApp(mockAppContext)

	mockKafkaClient.AssertExpectations(t)
	MockHttpServer.AssertExpectations(t)
}

func newTestAppContext(t *testing.T, mkc *MockKafkaClient, mhs *MockHttpServer) *config.AppContext {
	appConfig := newTestAppConfig()
	randomNumberService := newRandomNumberService()
	return &config.AppContext{
		AppConfig:           appConfig,
		RandomNumberService: randomNumberService,
		KafkaClient:         mkc,
		KafkaProcessor:      newKafkaProcessor(randomNumberService),
		HttpServer:          mhs,
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

type MockHttpServer struct {
	mock.Mock
}

func (mhs *MockHttpServer) RegisterEndpoint(endpoint string, handler config.Handler) {
	mhs.Called(endpoint, handler)
}

// Runs the HTTP server
func (mhs *MockHttpServer) Run() {
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
