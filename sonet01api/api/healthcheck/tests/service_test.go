package healthcheck_tests

import (
	"testing"

	"sonet01api/api/healthcheck"
	"sonet01api/server/database"
)

func TestMain(m *testing.M) {
	database.InitDB("../../../server/database/sonet01_test.db")
	m.Run()
}

func TestHealthCheckHandler(t *testing.T) {

	// database.Close()
	mock := healthcheck.HealthCheckService{}

	database.Close()

	t.Run("Database is down", func(t *testing.T) {

		response := mock.HealthCheck()

		if response != "Database connection failed" {
			t.Fail()
		}

		t.Log("Expected: Database connection failed")
		t.Log("Actual:", response)
	})

	// create mock database connection
	database.Connect("../../../server/database/sonet01_test.db")

	t.Run("Database is up", func(t *testing.T) {
		response := mock.HealthCheck()

		if response != "Server is running!" {
			t.Fail()
		}

		//t.Error("Expected: Database connection failed")
		t.Log("Actual:", response)
	})

}
