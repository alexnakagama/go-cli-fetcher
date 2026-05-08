package github

import (
	"context"
	"fmt"
)

func ListRepos() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	repos, _, err := client.Repositories.ListByAuthenticatedUser(context.Background(), nil)
	if err != nil {
		return err
	}

	for _, repo := range repos {
		fmt.Printf("- %s: %s\n", *repo.Name, *repo.HTMLURL)
	}
	return nil
}

func GetRepo(owner, repoName string) error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	repo, _, err := client.Repositories.Get(context.Background(), owner, repoName)
	if err != nil {
		return err
	}

	fmt.Printf("Name: %s\nDescription: %s\nStars: %d\nForks: %d\nURL: %s\n",
		*repo.Name, *repo.Description, *repo.StargazersCount, *repo.ForksCount, *repo.HTMLURL)
	return nil
}
