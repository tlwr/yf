package cmd_test

import (
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/tlwr/yf/pkg/cmd"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
)

const (
	boringYAML = "a: b"
)

var _ = Describe("Entrypoint", func() {
	var (
		inStream  io.Reader
		outStream *Buffer
		args      []string
	)

	BeforeEach(func() {
		inStream = strings.NewReader("")
		outStream = NewBuffer()
		args = []string{"yf"}
	})

	Context("When we are using urfave/cli wrong", func() {
		Specify("It should return an error", func() {
			args = []string{}

			err := cmd.Entrypoint(inStream, outStream, args)

			Expect(err).To(MatchError(ContainSubstring("No arguments")))
		})
	})

	Context("When a path is not provided", func() {
		Context("And neither a file nor content at STDIN is present", func() {
			Specify("It should return an error", func() {
				inStream := strings.NewReader("")
				outStream := NewBuffer()

				args = []string{"yf"}

				err := cmd.Entrypoint(inStream, outStream, args)

				Expect(err).To(
					MatchError(ContainSubstring("Did not read content from STDIN")),
				)
			})
		})

		Context("And a file is provided", func() {
			Context("And the file is valid", func() {
				Specify("It should return the file contents", func() {

					tmp, err := ioutil.TempFile(os.TempDir(), "")
					Expect(err).NotTo(HaveOccurred())
					defer os.Remove(tmp.Name())

					_, err = tmp.Write([]byte(boringYAML))
					Expect(err).NotTo(HaveOccurred())

					inStream := strings.NewReader("")
					outStream := NewBuffer()

					args := []string{"yf", "--file", tmp.Name()}

					err = cmd.Entrypoint(inStream, outStream, args)

					Expect(err).NotTo(HaveOccurred())

					Expect(outStream.Contents()).To(ContainSubstring(boringYAML))
				})
			})

			Context("And the file is not valid", func() {
				Specify("It should return an error", func() {
					inStream := strings.NewReader("")
					outStream := NewBuffer()

					args := []string{"yf", "--file", "/path/to/file/not/exist"}

					err := cmd.Entrypoint(inStream, outStream, args)

					Expect(err).To(HaveOccurred())

					Expect(err).To(
						MatchError(ContainSubstring("Could not read file")),
						"The error should be helpful",
					)

					Expect(err).To(
						MatchError(ContainSubstring("/path/to/file/not/exist")),
						"The error should contain the file path",
					)
				})
			})
		})

		Context("And content at STDIN is present", func() {
			Specify("It should return the contents of STDIN", func() {
				inStream := strings.NewReader(boringYAML)
				outStream := NewBuffer()

				args := []string{"yf"}

				err := cmd.Entrypoint(inStream, outStream, args)

				Expect(err).NotTo(HaveOccurred())

				Expect(outStream.Contents()).To(ContainSubstring(boringYAML))
			})
		})
	})
})
