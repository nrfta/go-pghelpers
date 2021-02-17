package pghelpers

import (
	"context"
	"database/sql"
	"fmt"
)

// ExecInTxFunc defines a function type for the ExecInTx function argument.
type ExecInTxFunc func(tx *sql.Tx) (commit bool)

// ExecInTx executes the provided function within a database transaction. The must return true for the
// transaction to be commit. Returning false will rollback the transaction. To pass variables into the
// transaction function or to return variables out, use a closure.
func ExecInTx(db *sql.DB, fn ExecInTxFunc) error {
	tx, err := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	if fn(tx) {
		return tx.Commit()
	}
	return tx.Rollback()
}

// SetSavepoint sets a named savepoint in the current transaction.
func SetSavepoint(name string, tx *sql.Tx) error {
	_, err := tx.Exec(fmt.Sprintf("SAVEPOINT %s", name))
	return err
}

// ReleaseSavepoint releases a named savepoint previously set in the transaction. This allows the commands
// executed after the savepoint to be committed.
func ReleaseSavepoint(name string, tx *sql.Tx) error {
	_, err := tx.Exec(fmt.Sprintf("RELEASE SAVEPOINT %s", name))
	return err
}

// RollbackToSavepoint rolls back the transaction to the named savepoint.
func RollbackToSavepoint(name string, tx *sql.Tx) error {
	_, err := tx.Exec(fmt.Sprintf("ROLLBACK TO SAVEPOINT %s", name))
	return err
}
