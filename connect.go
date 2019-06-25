package pghelpers

import (
	"database/sql"

	_ "github.com/lib/pq" // Postgres

	"github.com/neighborly/gtoolbox/errors"
)

// ConnectPostgres connects to postgres
func ConnectPostgres(c PostgresConfig) (*sql.DB, errors.Error) {
	addr := c.GenerateAddress()

	driver := "postgres"
	var err error

	db, err := sql.Open(driver, addr)
	if err != nil {
		return nil, errors.Newf("unable to open postgres db at %s:%d/%s", c.Host, c.Port, c.Database).WithCause(err)
	}

	if err := db.Ping(); err != nil {
		return nil, errors.Newf("unable to ping postgres db at %s:%d/%s", c.Host, c.Port, c.Database).WithCause(err)
	}

	return db, nil
}
