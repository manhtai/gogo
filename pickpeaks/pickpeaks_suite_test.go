package pickpeaks_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPickpeaks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pickpeaks Suite")
}
