package cli

import (
	"fmt"
	"os"

	"github.com/TonyGLL/gogit/internal/gogit"
	"github.com/spf13/cobra"
)

func NewInitCmd(app *App) *cobra.Command {
	return &cobra.Command{
		Use:   "init [directory]",
		Short: "Creates a new gogit repository",
		Args:  cobra.MaximumNArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			targetDir := "."
			if len(args) > 0 {
				targetDir = args[0]
			}

			if err := gogit.InitRepo(targetDir); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
		},
	}
}
