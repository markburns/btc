package base58

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "base58 Suite")
}
