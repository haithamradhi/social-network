package healthcheck

type HealthCheckService struct {
	healthRepo HealthCheckRepo
}

type HealthCheckServiceInterface interface {
	HealthCheck() string
}

// HealthCheck is a simple health check handler
func (h HealthCheckService) HealthCheck() string {
	// get repo instance & call database ping method
	repo := h.healthRepo

	err := repo.HealthCheck()
	if err != nil {
		return "Database connection failed"
	}

	return "Server is running!"
}
