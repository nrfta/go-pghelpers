package pghelpers

import "database/sql"

// ExecInTxFunc defines a function type for the ExecInTx function argument.
type ExecInTxFunc func(tx *sql.Tx) (commit bool)

// ExecInTx executes the provided function within a database transaction. The must return true for the
// transaction to be commit. Returning false will rollback the transaction. To pass variables into the
// transaction function or to return variables out, use a closure.
func ExecInTx(db *sql.DB, fn ExecInTxFunc) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if fn(tx) {
		return tx.Commit()
	}
	return tx.Rollback()
}
