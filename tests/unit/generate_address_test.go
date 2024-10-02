package pghelpers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	pgh "github.com/neighborly/go-pghelpers"
)

var (
	myTestString = "test"
	roTestString = "ro"
	testConfig   = pgh.PostgresConfig{
		Host:         myTestString,
		HostReadOnly: roTestString,
		Port:         1,
		Username:     myTestString,
		Password:     myTestString,
		Database:     myTestString,
		SSLEnabled:   false,
	}
)

var _ = Describe("Generate Address", func() {
	Context("RW", func() {
		It("Can generate a postgres address", func() {
			result := testConfig.GenerateAddress()
			Expect(result).To(Equal("host=test port=1 dbname=test user=test password=test sslmode=disable"))
		})
		It("should generate a postgres url", func() {
			result := testConfig.URL()
			Expect(result).To(Equal("postgresql://test:test@test:1/test?sslmode=disable"))
		})
	})

	Context("RO", func() {
		It("Can generate a postgres address", func() {
			result := testConfig.GenerateReadOnlyAddress()
			Expect(result).To(Equal("host=ro port=1 dbname=test user=test password=test sslmode=disable"))
		})
		It("should generate a postgres url", func() {
			result := testConfig.UrlReadOnly()
			Expect(result).To(Equal("postgresql://test:test@ro:1/test?sslmode=disable"))
		})
	})
})
