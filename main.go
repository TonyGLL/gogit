package main

import (
	"github.com/TonyGLL/gogit/cmd/gogit"
)

func main() {
	gogit.RegisterInitCommand(gogit.RootCmd)
	gogit.RegisterCommitCommand(gogit.RootCmd)
	gogit.RegisterAddCommand(gogit.RootCmd)
	gogit.RegisterLogCommand(gogit.RootCmd)
	gogit.RegisterStatusCommand(gogit.RootCmd)

	gogit.Execute()
}
