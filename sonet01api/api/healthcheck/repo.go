package healthcheck

import (
	db "sonet01api/server/database"

	"gopkg.in/errgo.v2/fmt/errors"
)

// create repo interface and struct
type HealthCheckRepoInterface interface {
	HealthCheck() error
}

type HealthCheckRepo struct {
}

// create function to test a database connection
func (h HealthCheckRepo) HealthCheck() (err error) {

	// test database connection
	dberr := db.Ping()
	if dberr != nil {
		return errors.New("Database connection failed")
	}

	return err
}
