package controllers

import (
	"encoding/json"
	"net/http"
)

const (
	StatusOk = "ok"
)

// GET /health
// Health check endpoint, returns a 200 success when the service is up
func HealthEndpoint(w http.ResponseWriter, r *http.Request) {
	appHealth := AppHealth{
		Status: StatusOk,
	}

	response, err := json.Marshal(appHealth)
	if err != nil {
		w.Write([]byte("Oh No!")) // FIXME
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(response)
}

type AppHealth struct {
	Status string `json:"status"`
}
