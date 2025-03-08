package database

import (
	"database/sql"
	"errors"
	"fmt"

	migrate "github.com/golang-migrate/migrate/v4"
	d "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Sqlite struct {
	// contains filtered or unexported fields
}

var db *sql.DB

// connect to database
func Connect(dbpath string) {
	// connect to database
	newDB, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		fmt.Println("error opening db")
	}
	db = newDB
}

// run init database
func InitDB(dbpath string) error {
	// connect to database
	Connect(dbpath)
	// create database if not exists
	driver, err := d.WithInstance(db, &d.Config{})
	if err != nil {
		return err
	}

	// run migration
	m, err := migrate.NewWithDatabaseInstance("file://../server/database/migrations", "sqlite3", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil {
		return err
	}

	return nil

}

// ping database
func Ping() (err error) {

	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Database connection failed")
		}
	}()

	err = db.Ping()
	if err != nil {
		fmt.Println("database ping error")
		return err
	}
	return nil
}

// close database connection
func Close() {
	db.Close()
}

// db execute
func Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.Exec(query, args...)
}
