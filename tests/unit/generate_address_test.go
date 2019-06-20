package pghelpers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	pgh "github.com/neighborly/go-pghelpers"
)

var (
	testTracingConfig = pgh.PostgresTracingConfig{true, true,
		true, true, true,
		true, true}
	myTestString = "test"
	testConfig   = pgh.PostgresConfig{
		Host:       myTestString,
		Port:       1,
		Username:   myTestString,
		Password:   myTestString,
		Database:   myTestString,
		SSLEnabled: false,
		Tracing:    testTracingConfig}
)

var _ = Describe("RatingsDimension", func() {
	It("Can generate ratings", func() {
		result := testConfig.GenerateAddress()
		Expect(result).To(Equal("host=test port=1 dbname=test user=test password=test sslmode=disable"))

	})

})
