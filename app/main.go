package main

import (
	"os"

	"gitlab.com/shaninalex/lumna/app/adapters/cli"
)

func main() {
	os.Exit(cli.Execute())
}
