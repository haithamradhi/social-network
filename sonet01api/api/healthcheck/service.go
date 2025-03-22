package healthcheck

type HealthCheckService struct{}

type HealthCheckServiceInterface interface {
	HealthCheck() string
}

// HealthCheck is a simple health check handler
func (h HealthCheckService) HealthCheck() string {
	// get repo instance & call database ping method
	repo := HealthCheckRepo{}

	err := repo.HealthCheck()
	if err != nil {
		return "Database connection failed"
	}

	return "Server is running!"
}
