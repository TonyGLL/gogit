package gogit

import (
	"fmt"
	"os"

	"github.com/TonyGLL/gogit/internal/gogit"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <file|directory>",
	Short: "Add a file or directory to the gogit repository",
	Long: `Adds the specified file or directory to the staging area (index).
When a directory is specified, it recursively adds all files within that
directory, excluding the .gogit directory itself.`,
	Args: cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		pathToAdd := args[0]

		if err := gogit.Add(pathToAdd); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("File(s) added successfully.")
	},
}

func RegisterAddCommand(root *cobra.Command) {
	root.AddCommand(addCmd)
}
