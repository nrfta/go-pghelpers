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
		SSLEnabled: false,
	}
)

var _ = Describe("Generate Address", func() {
	Context("GenerateAddress", func() {
		It("Can generate a postgres address", func() {
			result := testConfig.GenerateAddress()
			Expect(result).To(Equal("host=test port=1 dbname=test user=test password=test sslmode=disable"))
		})
	})

	Context("URL", func() {
		It("should generate a postgres url", func() {
			result := testConfig.URL()
			Expect(result).To(Equal("postgresql://test:test@test:1/test?sslmode=disable"))
		})
	})
})
