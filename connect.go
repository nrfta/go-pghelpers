package pghelpers

import (
	"database/sql"

	_ "github.com/lib/pq" // Postgres

	"contrib.go.opencensus.io/integrations/ocsql"
	"github.com/neighborly/go-errors"
)

// ConnectPostgres connects to postgres
func ConnectPostgres(c PostgresConfig) (*sql.DB, error) {
	addr := c.GenerateAddress()

	driver := "postgres"
	var err error

	if c.Tracing.Enabled {
		tracingConfig := ocsql.TraceOptions{
			AllowRoot:    true,
			Ping:         true,
			RowsNext:     c.Tracing.CreateRowsNextSpan,
			RowsClose:    c.Tracing.CreateRowsCloseSpan,
			RowsAffected: c.Tracing.CreateRowsAffectedSpan,
			LastInsertID: c.Tracing.CreateLastInsertedIDSpan,
			Query:        c.Tracing.AddQueryAttribute,
			QueryParams:  c.Tracing.AddQueryParamsAttributes,
		}

		driver, err = ocsql.Register(driver, ocsql.WithOptions(tracingConfig))
		if err != nil {
			return nil, errors.Wrap(err, "Unable to enable tracing on postgres db")
		}
	}

	db, err := sql.Open(driver, addr)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to open postgres db at %s:%d/%s", c.Host, c.Port, c.Database)
	}

	if err := db.Ping(); err != nil {
		return nil, errors.Wrapf(err, "unable to ping postgres db at %s:%d/%s", c.Host, c.Port, c.Database)
	}

	return db, nil
}
