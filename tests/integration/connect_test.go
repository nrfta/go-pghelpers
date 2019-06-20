package pghelpers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	pgh "github.com/neighborly/go-pghelpers"
)

var _ = Describe("Connection Test", func() {

	var (
		testTracingConfig = pgh.PostgresTracingConfig{true, true,
			true, true, true,
			true, true}
		testConfig = pgh.PostgresConfig{
			Host:       "localhost",
			Port:       5432,
			Username:   "postgres",
			Password:   "",
			Database:   "postgres",
			SSLEnabled: false,
			Tracing:    testTracingConfig}
	)

	It("should connect to a database", func() {
		db, err := pgh.ConnectPostgres(testConfig)
		Expect(db).To(Not(BeNil()))
		Expect(err).To(BeNil())
		Expect(db.Ping()).To(Succeed())
	})
})
