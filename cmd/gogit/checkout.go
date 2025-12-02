package gogit

import (
	"fmt"
	"os"

	"github.com/TonyGLL/gogit/internal/gogit"
	"github.com/spf13/cobra"
)

var bFlag bool

var checkCmd = &cobra.Command{
	Use:   "checkout [branch-name]",
	Short: "Switch branches in the gogit repository",
	Long: `Switch to the specified branch in the gogit repository.
If the branch does not exist, it can be created with the -b flag.`,
	Args: cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		if bFlag {
			if err := gogit.CreateBranch(args[0]); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
		}
		if err := gogit.CheckoutBranch(args[0]); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	},
}

func RegisterCheckoutCommand(root *cobra.Command) {
	checkCmd.Flags().BoolVarP(&bFlag, "create", "b", false, "Create branch if it does not exist")
	root.AddCommand(checkCmd)
}
