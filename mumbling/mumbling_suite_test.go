package mumbling_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMumbling(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mumbling Suite")
}
