package cli

import (
	"fmt"
	"os"

	"github.com/TonyGLL/gogit/internal/gogit"
	"github.com/spf13/cobra"
)

func NewStatusCmd(app *App) *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Show commit status",
		Run: func(_ *cobra.Command, _ []string) {
			if err := gogit.StatusRepo(); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
		},
	}
}
