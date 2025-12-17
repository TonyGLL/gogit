package cli

import (
	"fmt"
	"os"

	"github.com/TonyGLL/gogit/internal/gogit"
	"github.com/spf13/cobra"
)

func NewCommitCmd() *cobra.Command {
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
	if err := cmd.MarkFlagRequired("message"); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	return cmd
}
