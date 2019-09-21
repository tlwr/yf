package main

import (
	"fmt"
	"os"

	"github.com/tlwr/yf/pkg/cmd"
)

func main() {
	err := cmd.Entrypoint(os.Stdin, os.Stdout, os.Args)

	if err != nil {
		fmt.Fprintf(os.Stderr, "yf: %s\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
