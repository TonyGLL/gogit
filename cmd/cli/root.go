package cli

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "gogit",
		Short: "gogit - a simplified Git replica written in Go",
		Long: `gogit is a minimalist version control system
	created as a learning project to understand the fundamental
	concepts of Git.`,
	}

	rootCmd.AddCommand(
		NewInitCmd(),
		NewAddCmd(),
		NewCommitCmd(),
		NewLogCmd(),
		NewStatusCmd(),
		NewConfigCmd(),
		NewCheckoutCmd(),
		NewBranchCmd(),
	)

	return rootCmd
}
