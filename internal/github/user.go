package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v86/github"
)

func getUser(username string) (*github.User, error) {
	client, err := createClient()
	if err != nil {
		return nil, err
	}

	user, _, err := client.Users.Get(context.Background(), username)
	if err != nil {
		return nil, fmt.Errorf("Error fetching user: %v", err)
	}

	return user, nil
}

func SearchUserInfo(username string) error {
	user, err := getUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("Username: %s\n", *user.Login)
	fmt.Printf("Name: %s\n", *user.Name)
	fmt.Printf("Public repositories: %d\n", *user.PublicRepos)
	fmt.Printf("Followers: %d\n", *user.Followers)
	fmt.Printf("Following: %d\n", *user.Following)
	fmt.Printf("Location: %s\n", *user.Location)
	fmt.Printf("Bio: %s\n", *user.Bio)
	fmt.Printf("Created at: %s\n", *user.CreatedAt)
	fmt.Printf("Email: %s\n", *user.Email)
	fmt.Printf("Url: %s\n", *user.URL)

	return nil
}

func SearchUserUrl(username string) error {
	client, err := createClient()

	user, _, err := client.Users.Get(context.Background(), username)
	if err != nil {
		fmt.Printf("Error fetching the url: %v", err)
	}

	fmt.Printf("URL: %s", *user.URL)
	return nil
}

func SearchUserEmail(username string) error {
	client, err := createClient()

	user, _, err := client.Users.Get(context.Background(), username)
	if err != nil {
		fmt.Printf("Error fetching user: %v", err)
	}

	fmt.Printf("Email: %s", *user.Email)
	return nil
}

func SearchUserCreate(username string) error {
	client, err := createClient()

	user, _, err := client.Users.Get(context.Background(), username)
	if err != nil {
		fmt.Printf("Error fetching user: %v", err)
	}

	fmt.Printf("Created at: %s", *user.CreatedAt)
	return nil
}

func SearchPublicRepos(username string) error {
	client, err := createClient()

	user, _, err := client.Users.Get(context.Background(), username)
	if err != nil {
		fmt.Printf("Error fetching user: %v", err)
	}

	fmt.Printf("Number of public repositories: %d", *user.PublicRepos)
	return nil
}

func SearchUserLocation(username string) error {
	client, err := createClient()

	user, _, err := client.Users.Get(context.Background(), username)
	if err != nil {
		fmt.Printf("Error fetching user: %v", err)
	}

	fmt.Printf("Location: %s", *user.Location)
	return nil
}

func SearchUserFollowing(username string) error {
	client, err := createClient()

	listUser, _, err := client.Users.ListFollowing(context.Background(), username, nil)
	if err != nil {
		fmt.Printf("Error fetching user: %v", err)
	}

	fmt.Printf("Following:\n")
	for _, user := range listUser {
		fmt.Printf("%s: %s\n", *user.Login, *user.HTMLURL)
	}

	return nil
}
