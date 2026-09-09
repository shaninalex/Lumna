package main

import (
	"os"

	"gitlab.com/shaninalex/lumna/app/cmd"
)

func main() {
	os.Exit(cmd.Execute())
}
