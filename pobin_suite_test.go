package pobin_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPobin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pobin Suite")
}
