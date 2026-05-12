package cmd

import (
	"github.com/alexnakagama/go-cli-fetcher/internal/github"
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage GitHub user",
	Long:  `Perform operations on GitHub user such as listing, searching, and getting details.`,
}

var userGetInfoCmd = &cobra.Command{
	Use:   "detail [username]",
	Short: "Fetch information of a github user",
	Long:  `Fetch detailed information about a GitHub user, including its name, description, followers/following, and more.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		username := args[0]
		return github.SearchUserInfo(username)
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(userGetInfoCmd)
}
