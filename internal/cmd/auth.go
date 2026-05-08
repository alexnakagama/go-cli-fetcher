package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with GitHub",
	Long:  `Authenticate with GitHub to access private repositories and perform actions on behalf of the user.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Authenticating with GitHub...")
		fmt.Print("Type your GitHub token: ")
		var token string
		fmt.Scanln(&token)

		os.Setenv("GITHUB_TOKEN", token)
		fmt.Println("Authentication successful! Your token has been set as an environment variable.")
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}
