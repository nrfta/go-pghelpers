package pghelpers

import (
	"database/sql"
	"log"
	"net/http/httptest"

	"github.com/DATA-DOG/go-txdb"
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
func NewTxInstance(instanceID string) *sql.DB {
	db, err := sql.Open("txdb", instanceID)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
