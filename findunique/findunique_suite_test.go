package findunique_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFindunique(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Findunique Suite")
}
