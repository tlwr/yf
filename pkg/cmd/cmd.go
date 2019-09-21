package cmd

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

func Entrypoint(r io.Reader, w io.WriteCloser, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf(
			"No arguments were provided, the first should be executable file",
		)
	}

	c := cli.NewApp()

	c.Name = "yf"
	c.Usage = "find things in YAML"

	c.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "path,p",
			Value: "/",
		},

		cli.StringFlag{
			Name: "file,f",
		},
	}

	c.Action = func(ctx *cli.Context) error {
		path := ctx.String("path")
		file := ctx.String("file")

		var content []byte
		var err error

		if file == "" {
			content, err = ioutil.ReadAll(r)

			if err != nil {
				return errors.Wrap(err, "Could not read from STDIN")
			}

			if len(content) == 0 {
				return fmt.Errorf("Did not read content from STDIN")
			}
		} else {
			content, err = ioutil.ReadFile(file)

			if err != nil {
				return errors.Wrap(
					err, fmt.Sprintf("Could not read file '%s'", file),
				)
			}
		}

		traversedContent, err := traverse(path, string(content))

		if err != nil {
			return err
		}

		w.Write([]byte(traversedContent))

		return nil
	}

	return c.Run(args)
}

func traverse(_ string, content string) (string, error) {
	return content, nil
}
