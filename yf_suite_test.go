package main_test

import (
	"os/exec"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

func TestYf(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "yf Suite")
}

var _ = Describe("CLI", func() {
	var cliPath string

	var _ = BeforeSuite(func() {
		By("Compiling")

		var err error
		cliPath, err = Build("github.com/tlwr/yf")
		Expect(err).NotTo(HaveOccurred(), "Compilation should not fail")
	})

	var _ = AfterSuite(func() {
		CleanupBuildArtifacts()
	})

	Context("With -help as an argment", func() {
		Specify("It should print help", func() {
			command := exec.Command(cliPath, "-h")

			session, err := Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ShouldNot(HaveOccurred())

			Eventually(session).Should(
				SatisfyAll(
					Say("yf"),
					Exit(0),
				),
			)
		})
	})
})
