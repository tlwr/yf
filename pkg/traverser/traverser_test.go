package traverser_test

import (
	. "github.com/tlwr/yf/pkg/traverser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

const (
	simpleHashYAML = "a: b"
)

var _ = Describe("Traverse", func() {
	Context("When given a simple hash", func() {
		Context("And given an invalid path", func() {
			Specify("It should return the value", func() {
				_, err := Traverse("/b", simpleHashYAML)

				Expect(err).To(HaveOccurred())

				Expect(err).To(
					MatchError(ContainSubstring("Bad path")), "Be descriptive",
				)
				Expect(err).To(
					MatchError(ContainSubstring("/b")), "Be helpful in diagnosis",
				)
			})
		})

		Context("And given a valid path which is just a key selector", func() {
			Specify("It should return the value", func() {
				traversed, err := Traverse("/a", simpleHashYAML)

				Expect(err).NotTo(HaveOccurred())

				Expect(traversed).To(MatchFields(IgnoreExtras, Fields{
					"Content":         Equal("b"),
					"OriginalContent": Equal(simpleHashYAML),
					"Line":            Equal(0),
				}))
			})
		})
	})
})
