package cli

import (
	"fmt"
	"os"

	"github.com/TonyGLL/gogit/internal/gogit"
	"github.com/spf13/cobra"
)

func NewCommitCmd(app *App) *cobra.Command {
	var msg string

	cmd := &cobra.Command{
		Use:   "commit",
		Short: "Add commit message to gogit repository",
		Run: func(_ *cobra.Command, _ []string) {
			if err := gogit.AddCommit(&msg); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVarP(&msg, "message", "m", "", "Mensaje del commit")
	cmd.MarkFlagRequired("message")

	return cmd
}
