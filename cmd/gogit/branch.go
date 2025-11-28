package gogit

import (
	"fmt"
	"os"

	"github.com/TonyGLL/gogit/internal/gogit"
	"github.com/spf13/cobra"
)

var deleteFlag bool

var branchCmd = &cobra.Command{
	Use:   "branch [name]",
	Short: "Manage branches in the gogit repository",
	Long: `Create, list, delete, and switch branches in the gogit repository.
This command allows you to manage branches effectively.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		if deleteFlag {
			if len(args) < 1 {
				fmt.Fprintln(os.Stderr, "error: branch name required for deletion")
				return // O cmd.Usage()
			}
			if err := gogit.DeleteBranch(args[0]); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			}
		} else {
			if len(args) == 0 {
				if err := gogit.ListBranches(); err != nil {
					fmt.Fprintf(os.Stderr, "%v\n", err)
				}
			} else {
				if err := gogit.CreateBranch(args[0]); err != nil {
					fmt.Fprintf(os.Stderr, "%v\n", err)
				}
			}
		}
	},
}

func RegisterBranchCommand(root *cobra.Command) {
	branchCmd.Flags().BoolVarP(&deleteFlag, "delete", "d", false, "Delete a branch")
	root.AddCommand(branchCmd)
}
