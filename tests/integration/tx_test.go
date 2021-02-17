package pghelpers_test

import (
	"database/sql"

	pgh "github.com/neighborly/go-pghelpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tx", func() {
	var (
		txDB *sql.DB
	)
	BeforeEach(func() {
		db, err := pgh.ConnectPostgres(getTestConfig())
		Expect(err).To(BeNil())
		txDB = db
		_, err = txDB.Exec("DROP TABLE IF EXISTS tx_test")
		Expect(err).To(BeNil())
		_, err = txDB.Exec("CREATE TABLE tx_test (value int)")
		Expect(err).To(BeNil())
	})
	Context("ExecInTx", func() {
		It("should commit tx", func() {
			var txErr error
			err := pgh.ExecInTx(txDB, func(tx *sql.Tx) bool {
				_, txErr = tx.Exec("INSERT INTO tx_test (value) VALUES (1)")
				return txErr == nil
			})
			Expect(err).To(BeNil())
			Expect(txErr).To(BeNil())

			row := txDB.QueryRow("SELECT value from tx_test")
			var result int
			err = row.Scan(&result)
			Expect(err).To(BeNil())
			Expect(result).To(Equal(1))
		})

		It("should rollback tx", func() {
			var txErr error
			err := pgh.ExecInTx(txDB, func(tx *sql.Tx) bool {
				_, txErr = tx.Exec("INSERT INTO tx_test (value) VALUES (1)")
				return false
			})
			Expect(err).To(BeNil())
			Expect(txErr).To(BeNil())

			row := txDB.QueryRow("SELECT count(*) from tx_test")
			var result int
			err = row.Scan(&result)
			Expect(err).To(BeNil())
			Expect(result).To(Equal(0))
		})
	})

	Context("Savepoints", func() {
		It("should commit release savepoint", func() {
			var txErr error
			err := pgh.ExecInTx(txDB, func(tx *sql.Tx) bool {
				_, txErr = tx.Exec("INSERT INTO tx_test (value) VALUES (1)")
				if txErr != nil {
					return false
				}
				txErr = pgh.SetSavepoint("test", tx)
				if txErr != nil {
					return false
				}
				_, txErr = tx.Exec("INSERT INTO tx_test (value) VALUES (2)")
				if txErr != nil {
					return false
				}
				txErr = pgh.ReleaseSavepoint("test", tx)
				return txErr == nil
			})
			Expect(err).To(BeNil())
			Expect(txErr).To(BeNil())

			expectSavepointRows(2, txDB)
		})
		It("should rollback savepoint", func() {
			var txErr error
			err := pgh.ExecInTx(txDB, func(tx *sql.Tx) bool {
				_, txErr = tx.Exec("INSERT INTO tx_test (value) VALUES (1)")
				if txErr != nil {
					return false
				}
				txErr = pgh.SetSavepoint("test", tx)
				if txErr != nil {
					return false
				}
				_, txErr = tx.Exec("INSERT INTO tx_test (value) VALUES (2)")
				if txErr != nil {
					return false
				}
				txErr = pgh.RollbackToSavepoint("test", tx)
				return txErr == nil
			})
			Expect(err).To(BeNil())
			Expect(txErr).To(BeNil())

			expectSavepointRows(1, txDB)
		})
		It("should rollback second savepoint", func() {
			var txErr error
			err := pgh.ExecInTx(txDB, func(tx *sql.Tx) bool {
				_, txErr = tx.Exec("INSERT INTO tx_test (value) VALUES (1)")
				if txErr != nil {
					return false
				}
				txErr = pgh.SetSavepoint("test1", tx)
				if txErr != nil {
					return false
				}
				_, txErr = tx.Exec("INSERT INTO tx_test (value) VALUES (2)")
				if txErr != nil {
					return false
				}
				txErr = pgh.SetSavepoint("test2", tx)
				if txErr != nil {
					return false
				}
				_, txErr = tx.Exec("INSERT INTO tx_test (value) VALUES (3)")
				if txErr != nil {
					return false
				}
				txErr = pgh.RollbackToSavepoint("test2", tx)
				return txErr == nil
			})
			Expect(err).To(BeNil())
			Expect(txErr).To(BeNil())

			expectSavepointRows(2, txDB)
		})
	})
})

func expectSavepointRows(numRows int, txDB *sql.DB) {
	rows, err := txDB.Query("SELECT value from tx_test")
	var result int
	expectedValue := 1
	for rows.Next() {
		err = rows.Scan(&result)
		Expect(err).To(BeNil())
		Expect(result).To(Equal(expectedValue))
		expectedValue++
	}
	Expect(expectedValue).To(Equal(numRows + 1))
}
