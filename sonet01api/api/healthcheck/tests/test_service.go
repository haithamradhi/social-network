package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"sonet01api/api/healthcheck"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(healthcheck.HealthCheckAPI{}.CheckHealth)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status code 200")
	assert.Equal(t, "OK", rr.Body.String(), "Expected response body 'OK'")
}
