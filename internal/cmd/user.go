package cmd

import (
	"github.com/alexnakagama/go-cli-fetcher/internal/github"
	"github.com/spf13/cobra"
)

var userGetInfoCmd = &cobra.Command{
	Use:   "user",
	Short: "Fetch information of a github user",
	Long:  `Fetch detailed information about a GitHub user, including its name, description, followers/following, and more.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		username := args[0]
		return github.SearchUserInfo(username)
	},
}

func init() {
	rootCmd.AddCommand(userGetInfoCmd)
}
