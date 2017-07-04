package controllers

import (
	"encoding/json"
	"net/http"
)

const (
	// statusOk is the ok status message
	statusOk = "ok"
)

// HealthEndpoint returns a 200 success when the service is up
func HealthEndpoint(w http.ResponseWriter, r *http.Request) {
	appHealth := AppHealth{
		Status: statusOk,
	}

	response, err := json.Marshal(appHealth)
	if err != nil {
		w.Write([]byte("Oh No!")) // FIXME
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(response)
}

// AppHealth is the data struct to be returned from the HealthEndpoint
type AppHealth struct {
	Status string `json:"status"`
}
