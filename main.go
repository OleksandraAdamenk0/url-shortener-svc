package main

import (
	"os"

	"gitlab.com/tokend/url-shortener-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
