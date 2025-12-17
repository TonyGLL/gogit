package main

import (
	"os"

	"github.com/TonyGLL/gogit/cmd/cli"
)

func main() {
	app := cli.DefaultApp()

	rootCmd := cli.NewRootCmd(app)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
