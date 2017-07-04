package config

import (
	"fmt"
	"net/http"
)

// A HTTP endpoint handler
type Handler func(w http.ResponseWriter, r *http.Request)

// Abstraction from a HTTP server
type HttpServer interface {
	// Registers a HTTP Endpoint
	RegisterEndpoint(endpoint string, handler Handler)

	// Runs the HTTP server
	Run()
}

type DefaultHttpServer struct {
	// The struct logger
	logger Logger

	// The port to run on
	Port int
}

func (dhs DefaultHttpServer) RegisterEndpoint(endpoint string, handler Handler) {
	dhs.logger.Debug(fmt.Sprintf("Registering handler on endpoint '%s'", endpoint))
	http.HandleFunc(endpoint, handler)
}

func (dhs DefaultHttpServer) Run() {
	dhs.logger.Info(fmt.Sprintf("Starting HTTP server on port '%d'...", dhs.Port))
	http.ListenAndServe(fmt.Sprintf(":%d", dhs.Port), nil)
}

func NewDefaultHttpServer(port int) *DefaultHttpServer {
	return &DefaultHttpServer{
		logger: NewLogger(),
		Port: port,
	}
}
