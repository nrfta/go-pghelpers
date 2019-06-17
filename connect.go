package pghelpers

import (
	"database/sql"

	_ "github.com/lib/pq" // Postgres

	"contrib.go.opencensus.io/integrations/ocsql"
	"github.com/neighborly/gtoolbox/errors"
)

// ConnectPostgres connects to postgres
func ConnectPostgres(c PostgresConfig) (*sql.DB, errors.Error) {
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
			return nil, errors.New("Unable to enable tracing on postgres db").WithCause(err)
		}
	}

	db, err := sql.Open(driver, addr)
	if err != nil {
		return nil, errors.Newf("unable to open postgres db at %s:%d/%s", c.Host, c.Port, c.Database).WithCause(err)
	}

	if err := db.Ping(); err != nil {
		return nil, errors.Newf("unable to ping postgres db at %s:%d/%s", c.Host, c.Port, c.Database).WithCause(err)
	}

	return db, nil
}
