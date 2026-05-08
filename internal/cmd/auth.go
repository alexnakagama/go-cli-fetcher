package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Configure your GitHub token",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Enter your GitHub token: ")
		var token string
		fmt.Scanln(&token)

		envContent := fmt.Sprintf("GITHUB_TOKEN=%s\n", token)
		if err := os.WriteFile(".env", []byte(envContent), 0600); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return
		}

		fmt.Println("Token saved to .env file. You're ready to go!")
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}
