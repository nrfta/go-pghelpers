package apm

import (
	"database/sql"
	"fmt"

	"github.com/neighborly/go-pghelpers"
	"go.elastic.co/apm/module/apmsql/v2"
	_ "go.elastic.co/apm/module/apmsql/v2/pq"
)

// Connect connects to postgres and adds APM instrumentation
func Connect(c pghelpers.PostgresConfig) (*sql.DB, error) {
	addr := c.GenerateAddress()

	db, err := apmsql.Open("postgres", addr)
	if err != nil {
		return nil, fmt.Errorf(
			"open postgres db %q: %w",
			fmt.Sprintf("%s:%d/%s", c.Host, c.Port, c.Database),
			err,
		)
	}

	maxOpenConnections := 10
	if c.MaxOpenConnections > 0 {
		maxOpenConnections = c.MaxOpenConnections
	}
	db.SetMaxOpenConns(maxOpenConnections)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf(
			"ping postgres db %q: %w",
			fmt.Sprintf("%s:%d/%s", c.Host, c.Port, c.Database),
			err,
		)
	}

	return db, nil
}
