package cmd

import (
	"github.com/alexnakagama/go-cli-fetcher/internal/github"
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Fetch information about a GitHub repository",
	Long:  `Fetch detailed information about a GitHub repository, including its name, description, stars, forks, and more.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return github.ListRepos()
	},
}

func init() {
	rootCmd.AddCommand(repoCmd)
}
