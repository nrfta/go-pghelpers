package pghelpers_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPgHelpers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Reference Suite")
}
