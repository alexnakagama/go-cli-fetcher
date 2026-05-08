package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-cli-fetcher",
	Short: "A CLI tool to fetch data from GitHub repositories",
	Long: `go-cli-fetcher is a command-line tool that allows you to fetch data from GitHub repositories.
It provides various commands to interact with GitHub's API and retrieve information about repositories, issues, pull requests, and more.`,
}

func Execute() error {
	if err := godotenv.Load(); err != nil {
		fmt.Fprintln(os.Stderr, "Warning: No .env file found. Make sure to set the GITHUB_TOKEN environment variable.")
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return nil
}

func init() {
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if os.Getenv("GITHUB_TOKEN") == "" {
			fmt.Fprintln(os.Stderr, "Error: GITHUB_TOKEN environment variable is not set. Please authenticate using the 'auth' command.")
			os.Exit(1)
		}
	}
}
