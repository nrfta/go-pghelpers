package pghelpers

import (
	"database/sql"
	"log"
	"net/http/httptest"

	txdb "github.com/DATA-DOG/go-txdb"

	"github.com/jmoiron/sqlx"
)

var (
	server *httptest.Server
)

// RegisterPostgresTxDb registers postgres to txdb
func RegisterPostgresTxDb(pgConfig PostgresConfig) {
	addr := pgConfig.GenerateAddress()
	txdb.Register("txdb", "postgres", addr)
}

// NewTxInstance creates a new instance of tx-db.
// Make sure you Register first
func NewTxInstance(instanceID string) *sqlx.DB {
	db, err := sql.Open("txdb", instanceID)
	if err != nil {
		log.Fatal(err)
	}
	tdb := sqlx.NewDb(db, "postgres")
	return tdb
}
