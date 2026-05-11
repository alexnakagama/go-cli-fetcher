package cmd

import (
	"github.com/alexnakagama/go-cli-fetcher/internal/github"
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:   "list",
	Short: "List information of the authenticated user repositories",
	Long:  `Fetch detailed information about a GitHub repository, including its name, description, stars, forks, and more.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return github.ListRepos()
	},
}

var repoDetailCmd = &cobra.Command{
	Use:   "detail [owner] [repo]",
	Short: "Fetch detailed information about a specific GitHub repository",
	Long:  `Fetch detailed information about a specific GitHub repository by providing the owner's username and the repository name.`,
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		owner := args[0]
		repoName := args[1]
		return github.GetRepo(owner, repoName)
	},
}

var repoCreateCmd = &cobra.Command{
	Use:   "create [owner] [repoName]",
	Short: "Create a repository of the authenticated user",
	Long:  "Create a repository of the authenticated user, command is 'create [owner] [repoName]'",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		owner := args[0]
		repoName := args[1]
		return github.CreateRepo(owner, repoName)
	},
}

func init() {
	rootCmd.AddCommand(repoCmd)
	repoCmd.AddCommand(repoDetailCmd)
}
