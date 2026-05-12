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
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		username := args[0]
		return github.SearchUserInfo(username)
	},
}

var userGetUrlCmd = &cobra.Command{
	Use:   "url [username]",
	Short: "Fetch url from a github user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		username := args[0]
		return github.SearchUserUrl(username)
	},
}

var userGetEmailCmd = &cobra.Command{
	Use:   "email [username]",
	Short: "Fetch email from a github user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		username := args[0]
		return github.SearchUserEmail(username)
	},
}

var userGetCreationCmd = &cobra.Command{
	Use:   "created-at [username]",
	Short: "Fetch creation date of the github user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		username := args[0]
		return github.SearchUserCreate(username)
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(userGetInfoCmd)
	userCmd.AddCommand(userGetUrlCmd)
	userCmd.AddCommand(userGetEmailCmd)
	userCmd.AddCommand(userGetCreationCmd)
}
