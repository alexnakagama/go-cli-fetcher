package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Fetch information about a GitHub repository",
	Long:  `Fetch detailed information about a GitHub repository, including its name, description, stars, forks, and more.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fetching repository information...")
	},
}

func init() {
	rootCmd.AddCommand(repoCmd)
}
