package cli

import (
	"fmt"
	"os"

	"github.com/TonyGLL/gogit/internal/gogit"
	"github.com/spf13/cobra"
)

func NewLogCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "log",
		Short: "Show commits logs",
		Run: func(_ *cobra.Command, _ []string) {
			if err := gogit.LogRepo(); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
		},
	}
}
