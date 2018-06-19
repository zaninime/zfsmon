package clif_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestClif(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "clif library suite")
}
