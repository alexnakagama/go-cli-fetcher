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
