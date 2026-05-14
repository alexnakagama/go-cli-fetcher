package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v86/github"
)

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

func CreateRepo(owner string, repoName string) error {
	client, err := createClient()

	repo := &github.Repository{Name: github.Ptr(repoName)}
	_, _, err = client.Repositories.Create(context.Background(), owner, repo)
	if err != nil {
		return err
	}

	return nil
}

func DeleteRepo(owner string, repoName string) error {
	client, err := createClient()

	_, err = client.Repositories.Delete(context.Background(), owner, repoName)
	if err != nil {
		return err
	}

	return nil
}

func ListBranches(owner string, repoName string) error {
	client, err := createClient()

	branches, _, err := client.Repositories.ListBranches(context.Background(), owner, repoName, nil)
	if err != nil {
		return err
	}

	for _, branch := range branches {
		fmt.Println(*branch.Name)
	}
	return nil
}

func ListCommits(owner string, repoName string) error {
	client, err := createClient()

	commits, _, err := client.Repositories.ListCommits(context.Background(), owner, repoName, nil)
	if err != nil {
		return err
	}

	for _, commit := range commits {
		if commit.Author != nil && commit.Author.Login != nil {
			fmt.Printf("SHA: %s\n", *commit.SHA)
			fmt.Printf("URL: %s\n", *commit.HTMLURL)
			fmt.Printf("Author: %s\n", *commit.Author.Login)
			fmt.Printf("Email: %s\n", *commit.Commit.Author.Email)
			fmt.Printf("Message: %s\n", *commit.Commit.Message)
			fmt.Printf("Date: %s\n", *commit.Commit.Author.Date)
			fmt.Println("")
		} else {
			fmt.Println("No author info")
		}
	}
	return nil
}

func SearchRepoStars(owner string, repoName string) error {
	client, err := createClient()

	repos, _, err := client.Repositories.ListByAuthenticatedUser(context.Background(), nil)
	if err != nil {
		return err
	}

	for _, repo := range repos {
		fmt.Printf("Stars count: %d", *repo.StargazersCount)
		fmt.Printf("URL: %s", *repo.StargazersURL)
	}

	return nil
}

func CreateIssue(owner, repoName, title, body string) error {
	client, err := createClient()

	issueRequest := &github.IssueRequest{
		Title: github.Ptr(title),
		Body:  github.Ptr(body),
	}

	_, _, err = client.Issues.Create(context.Background(), owner, repoName, issueRequest)
	if err != nil {
		fmt.Printf("Error fetching the repository: %v", err)
	}

	return nil
}
