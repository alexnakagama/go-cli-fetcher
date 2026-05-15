package cmd

import (
	"github.com/alexnakagama/go-cli-fetcher/internal/github"
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Manage GitHub repositories",
	Long:  `Perform operations on GitHub repositories such as listing, creating, deleting, and getting details.`,
}
var repoListCmd = &cobra.Command{
	Use:   "list",
	Short: "List information of the authenticated user repositories",
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
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		owner := args[0]
		repoName := args[1]
		return github.CreateRepo(owner, repoName)
	},
}

var repoDeleteCmd = &cobra.Command{
	Use:   "delete [owner] [repoName]",
	Short: "Delete a repository of the authenticated user",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		owner := args[0]
		repoName := args[1]
		return github.DeleteRepo(owner, repoName)
	},
}

var repoListBranchesCmd = &cobra.Command{
	Use:   "list-branches [owner] [repoName]",
	Short: "List a repository branch",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		owner := args[0]
		repoName := args[1]
		return github.ListBranches(owner, repoName)
	},
}

var repoListCommitsCmd = &cobra.Command{
	Use:   "list-commits [owner] [repoName]",
	Short: "List commits of a repository",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		owner := args[0]
		repoName := args[1]
		return github.ListCommits(owner, repoName)
	},
}

var repoSearchStarsCmd = &cobra.Command{
	Use:   "stars [owner] [repoName]",
	Short: "Search for stars in a user repository",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		owner := args[0]
		repoName := args[1]
		return github.SearchRepoStars(owner, repoName)
	},
}

var repoCreateIssue = &cobra.Command{
	Use:   "issue [owner] [repoName] [title] [body]",
	Short: "Create an issue in a repository",
	Args:  cobra.ExactArgs(4),
	RunE: func(cmd *cobra.Command, args []string) error {
		owner := args[0]
		repoName := args[1]
		title := args[2]
		body := args[3]
		return github.CreateIssue(owner, repoName, title, body)
	},
}

var repoListIssue = &cobra.Command{
	Use:   "list-issue [owner] [repoName]",
	Short: "List the issues of the repository",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		owner := args[0]
		repoName := args[1]
		return github.ListIssue(owner, repoName)
	},
}

func init() {
	rootCmd.AddCommand(repoCmd)
	repoCmd.AddCommand(repoListCmd)
	repoCmd.AddCommand(repoDetailCmd)
	repoCmd.AddCommand(repoCreateCmd)
	repoCmd.AddCommand(repoDeleteCmd)
	repoCmd.AddCommand(repoListBranchesCmd)
	repoCmd.AddCommand(repoListCommitsCmd)
	repoCmd.AddCommand(repoSearchStarsCmd)
	repoCmd.AddCommand(repoCreateIssue)
}
