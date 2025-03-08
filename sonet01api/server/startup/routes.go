package startup

import (
	"net/http"
	"sonet01api/api/auth"
	"sonet01api/api/healthcheck"
	"sonet01api/api/post"
)

func RegisterRoutes(version string) {

	apiPrefix := "/api/" + version

	healthCheckRoutes := healthcheck.HealthCheckAPI{}
	authRoutes := auth.AuthAPI{}
	postRoutes := post.PostAPI{}

	// register routes
	http.HandleFunc("GET "+apiPrefix+"/healthcheck", AuthMiddleware(healthCheckRoutes.CheckHealth))
	http.HandleFunc("POST "+apiPrefix+"/user/create", AuthMiddleware(authRoutes.CreateUser))
	http.HandleFunc("POST "+apiPrefix+"/post/create", AuthMiddleware(postRoutes.CreatePost))

}
