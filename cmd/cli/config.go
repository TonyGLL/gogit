package cli

import (
	"fmt"

	"github.com/TonyGLL/gogit/internal/gogit"
	"github.com/spf13/cobra"
)

func NewConfigCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "Configure user name and email (like git config --global)",
		Long: `Set global configuration values for the current user.

Examples:
  gogit config user.name "John Doe"
  gogit config user.email "john@example.com"`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				fmt.Println("Usage: gogit config <key> <value>")
				fmt.Println("\nSupported keys:")
				fmt.Println("  user.name   Your full name")
				fmt.Println("  user.email  Your email address")
				fmt.Println("  list  	   Your list config")
				fmt.Println()
				if err := cmd.Usage(); err != nil {
					fmt.Printf("Error: %v\n", err)
					return
				}
				return
			}

			key := args[0]
			value := args[1]

			switch key {
			case "user.name":
				if err := gogit.SetName(value); err != nil {
					fmt.Printf("Error saving name: %v\n", err)
					return
				}

			case "user.email":
				if err := gogit.SetEmail(value); err != nil {
					fmt.Printf("Error saving email: %v\n", err)
					return
				}

			case "list":
				if err := gogit.GetConfig(value); err != nil {
					fmt.Printf("Error getting config: %v\n", err)
					return
				}

			default:
				fmt.Printf("Unknown config key: %s\n", key)
				fmt.Println("Supported keys are: user.name, user.email")
			}
		},
	}
}
