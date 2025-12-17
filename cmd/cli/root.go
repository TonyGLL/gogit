package cli

import (
	"github.com/spf13/cobra"
)

func NewRootCmd(app *App) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "gogit",
		Short: "gogit - a simplified Git replica written in Go",
		Long: `gogit is a minimalist version control system
	created as a learning project to understand the fundamental
	concepts of Git.`,
	}

	rootCmd.AddCommand(
		NewInitCmd(app),
		NewAddCmd(app),
		NewCommitCmd(app),
		NewLogCmd(app),
		NewStatusCmd(app),
		NewConfigCmd(app),
		NewCheckoutCmd(app),
		NewBranchCmd(app),
	)

	return rootCmd
}
