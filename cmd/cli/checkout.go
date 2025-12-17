package cli

import (
	"fmt"
	"os"

	"github.com/TonyGLL/gogit/internal/gogit"
	"github.com/spf13/cobra"
)

func NewCheckoutCmd(app *App) *cobra.Command {
	var bFlag bool

	return &cobra.Command{
		Use:   "checkout [branch-name]",
		Short: "Switch branches in the gogit repository",
		Long: `Switch to the specified branch in the gogit repository.
If the branch does not exist, it can be created with the -b flag.`,
		Args: cobra.ExactArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			if err := gogit.CheckoutBranch(args[0], bFlag); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
		},
	}
}
