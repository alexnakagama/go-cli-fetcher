package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-cli-fetcher",
	Short: "A CLI tool to fetch data from GitHub repositories",
	Long: `go-cli-fetcher is a command-line tool that allows you to fetch data from GitHub repositories.
It provides various commands to interact with GitHub's API and retrieve information about repositories, issues, pull requests, and more.`,
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return nil
}
