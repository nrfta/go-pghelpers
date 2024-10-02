package pghelpers_test

import (
	"os"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	pgh "github.com/neighborly/go-pghelpers"
)

var _ = Describe("Connection Test", func() {
	var (
		testConfig = getTestConfig()
	)

	It("should connect to a database", func() {
		db, err := pgh.ConnectPostgres(testConfig)
		Expect(db).To(Not(BeNil()))
		Expect(err).To(BeNil())
		Expect(db.Ping()).To(Succeed())
	})

	It("should connect to RO database", func() {
		db, err := pgh.ConnectPostgresReadOnly(testConfig)
		Expect(db).To(Not(BeNil()))
		Expect(err).To(BeNil())
		Expect(db.Ping()).To(Succeed())
	})
})

func getTestConfig() pgh.PostgresConfig {
	var port, _ = strconv.Atoi(getEnv("POSTGRES_PORT", "5432"))
	return pgh.PostgresConfig{
		ApplicationName: "test",
		Host:            getEnv("POSTGRES_HOST", "localhost"),
		HostReadOnly:    getEnv("POSTGRES_HOST_RO", "localhost"),
		Port:            port,
		Username:        getEnv("POSTGRES_USERNAME", "postgres"),
		Password:        getEnv("POSTGRES_PASSWORD", ""),
		Database:        getEnv("POSTGRES_DATABASE", "postgres"),
		SSLEnabled:      false,
	}
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
