package github

import (
	"context"
	"fmt"
)

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
	fmt.Printf("Location: %s\n", *user.Location)
	fmt.Printf("Bio: %s\n", *user.Bio)
	fmt.Printf("Created at: %s\n", *user.CreatedAt)
	fmt.Printf("Email: %s\n", *user.Email)
	fmt.Printf("Url: %s\n", *user.URL)

	return nil
}
