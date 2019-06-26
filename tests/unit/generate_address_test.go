package pghelpers_test

import (
	pgh "github.com/neighborly/go-pghelpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	myTestString = "test"
	testConfig   = pgh.PostgresConfig{
		Host:       myTestString,
		Port:       1,
		Username:   myTestString,
		Password:   myTestString,
		Database:   myTestString,
		SSLEnabled: false,}
)

var _ = Describe("Generate Address", func() {
	It("Can generate a postgres address", func() {
		result := testConfig.GenerateAddress()
		Expect(result).To(Equal("host=test port=1 dbname=test user=test password=test sslmode=disable"))

	})

})
