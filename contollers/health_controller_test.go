package controllers

import (
	"testing"
	"io/ioutil"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

func TestHealthEndpointReturnsCorrectJson(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	HealthEndpoint(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, "{\"status\":\"ok\"}", string(body))
	assert.Equal(t, "application/json; charset=utf-8", resp.Header.Get("Content-Type"))
}

