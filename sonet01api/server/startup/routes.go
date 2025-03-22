package startup

import (
	"net/http"
	health "sonet01api/api/healthcheck"
)

func RegisterRoutes(version string) {

	apiPrefix := "/api/" + version

	// create api instances
	hcr := health.HealthCheckAPI{}

	// register routes
	http.HandleFunc(apiPrefix+"/healthcheck", AuthMiddleware(hcr.CheckHealth))

}
