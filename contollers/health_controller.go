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

	response, _ := json.Marshal(appHealth)

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(response)
}

// AppHealth is the data struct to be returned from the HealthEndpoint
type AppHealth struct {
	Status string `json:"status"`
}
