package main

import (
	"os"

	"github.com/TonyGLL/gogit/cmd/cli"
)

func main() {
	cmd := cli.NewRootCmd()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
