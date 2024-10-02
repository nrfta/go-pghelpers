package pghelpers

import (
	"database/sql"

	_ "github.com/lib/pq" // Postgres

	"github.com/neighborly/go-errors"
)

// ConnectPostgres connects to postgres
func ConnectPostgres(c PostgresConfig) (*sql.DB, error) {
	return c.connectPostgres(false)
}
func ConnectPostgresReadOnly(c PostgresConfig) (*sql.DB, error) {
	return c.connectPostgres(true)
}

func (c PostgresConfig) connectPostgres(ro bool) (*sql.DB, error) {
	addr := c.generateAddress(ro)

	driver := "postgres"
	var err error

	db, err := sql.Open(driver, addr)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to open postgres db at %s:%d/%s", c.Host, c.Port, c.Database)
	}

	maxOpenConnections := c.MaxOpenConnections
	if maxOpenConnections == 0 {
		maxOpenConnections = 10
	}
	db.SetMaxOpenConns(maxOpenConnections)

	if err := db.Ping(); err != nil {
		return nil, errors.Wrapf(err, "unable to ping postgres db at %s:%d/%s", c.Host, c.Port, c.Database)
	}

	return db, nil
}
