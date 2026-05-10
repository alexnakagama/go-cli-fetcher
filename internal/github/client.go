package github

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v86/github"
	"golang.org/x/oauth2"
)

func NewClient() (*github.Client, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("GITHUB_TOKEN no está configurado")
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(context.Background(), ts)
	return github.NewClient(tc), nil
}

func createClient() (*github.Client, error) {
	client, err := NewClient()
	if err != nil {
		return nil, err
	}
	return client, nil
}
