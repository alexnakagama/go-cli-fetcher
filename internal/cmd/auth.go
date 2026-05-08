package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with GitHub",
	Long:  `Authenticate with GitHub to access private repositories and perform actions on behalf of the user.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Authenticating with GitHub...")
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}
