package healthcheck

import (
	"fmt"
	"net/http"
)

type HealthCheckAPI struct{}

type HealthCheckAPIInterface interface {
	CheckHealth(w http.ResponseWriter, r *http.Request)
}

// HealthCheck is a simple health check handler
func (h HealthCheckAPI) CheckHealth(w http.ResponseWriter, r *http.Request) {
	healthCheckService := HealthCheckService{}
	message := healthCheckService.HealthCheck()
	fmt.Fprintf(w, message)
}
