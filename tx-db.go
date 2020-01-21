package pghelpers

import (
	"database/sql"
	"log"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/lib/pq" // Postgres
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
