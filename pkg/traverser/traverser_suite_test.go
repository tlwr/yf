package traverser_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTraverser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Traverser Suite")
}
