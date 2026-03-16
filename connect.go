package pghelpers

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq" // Postgres

	"github.com/neighborly/go-errors"
)

// ConnectPostgres connects to postgres
func ConnectPostgres(c PostgresConfig) (*sql.DB, error) {
	addr := c.GenerateAddress()

	driver := "postgres"
	var err error

	db, err := sql.Open(driver, addr)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to open postgres db at %s:%d/%s", c.Host, c.Port, c.Database)
	}

	SetupPool(c, db)

	if err := db.Ping(); err != nil {
		return nil, errors.Wrapf(err, "unable to ping postgres db at %s:%d/%s", c.Host, c.Port, c.Database)
	}

	return db, nil
}

func SetupPool(c PostgresConfig, db *sql.DB) {
	maxOpenConnections := c.MaxOpenConnections
	if maxOpenConnections == 0 {
		maxOpenConnections = 10
	}
	db.SetMaxOpenConns(maxOpenConnections)

	maxIdleConnections := c.MaxIdleConnections
	if maxIdleConnections == 0 {
		maxIdleConnections = 2
	}
	db.SetMaxIdleConns(maxIdleConnections)

	maxConnectionLifetimeMinutes := c.MaxConnectionLifetimeMinutes
	if maxConnectionLifetimeMinutes == 0 {
		maxConnectionLifetimeMinutes = 5
	}
	db.SetConnMaxLifetime(time.Duration(maxConnectionLifetimeMinutes) * time.Minute)
}
