package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v86/github"
)

func createClient() (*github.Client, error) {
	client, err := NewClient()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func ListRepos() error {
	client, err := createClient()

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
	client, err := createClient()

	repo, _, err := client.Repositories.Get(context.Background(), owner, repoName)
	if err != nil {
		return err
	}

	fmt.Printf("Name: %s\nDescription: %s\nStars: %d\nForks: %d\nURL: %s\n",
		*repo.Name, *repo.Description, *repo.StargazersCount, *repo.ForksCount, *repo.HTMLURL)
	return nil
}

func SearchUserInfo(username string) error {
	client, err := createClient()

	user, _, err := client.Users.Get(context.Background(), username)
	if err != nil {
		fmt.Printf("Error fetching user: %v\n", err)
	}

	fmt.Printf("Username: %s\n", *user.Login)
	fmt.Printf("Name: %s\n", *user.Name)
	fmt.Printf("Public repositories: %d\n", *user.PublicRepos)
	fmt.Printf("Followers: %d\n", *user.Followers)
	fmt.Printf("Following: %d\n", *user.Following)

	return nil
}
