package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/gitpod-io/gitpod/components/public-api/go/client"
	v1 "github.com/gitpod-io/gitpod/components/public-api/go/experimental/v1"
	"github.com/google/go-github/v55/github"
	"golang.org/x/oauth2"
)

func ListAllOrgs() {
	token := os.Getenv("GITPOD_PAT_TOKEN")

	gitpod, err := client.New(client.WithCredentials(token))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to construct gitpod client %v", err)
		return
	}

	response, err := gitpod.Teams.ListTeams(context.Background(), connect.NewRequest(&v1.ListTeamsRequest{}))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to list teams %v", err)
		return
	}

	fmt.Fprintf(os.Stdout, "Retrieved teams %v", response.Msg.GetTeams())
}

func CreateProjectFromGitHubRepo(name, projectSlug, cloneUrl string) {
	token := os.Getenv("GITPOD_PAT_TOKEN")

	gitpod, err := client.New(client.WithCredentials(token))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to construct gitpod client %v", err)
		return
	}

	response, err := gitpod.Projects.CreateProject(context.Background(), connect.NewRequest(&v1.CreateProjectRequest{
		Project: &v1.Project{
			TeamId:   os.Getenv("GITPOD_ORG_ID"),
			Name:     name,
			Slug:     projectSlug,
			CloneUrl: cloneUrl,
		},
	}))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create project %v", err)
		return
	}

	fmt.Fprintf(os.Stdout, "Project has been successfully created\n %v\n", response.Msg)
}

func main() {
	/* If you don't have your OrgId, Get your teamId by uncomment the following */
	// ListAllOrgs()
	// return;

	/* Fetch All GitHub Repositories owned by a particular Org/User */
	startTime := time.Now()

	accessToken := os.Getenv("GITHUB_PAT_TOKEN")
	githubOwner := os.Getenv("GITHUB_USERNAME") // e.g., "Siddhant-K-code"

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	opts := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}

	var allRepos []*github.Repository

	for {
		repos, resp, err := client.Repositories.List(ctx, githubOwner, opts)
		if err != nil {
			fmt.Println("Error fetching repositories:", err)
			return
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	for _, repo := range allRepos {
		// fmt.Printf("Name: %v; Repo URL: %v\n", *repo.Name, *repo.CloneURL)
		CreateProjectFromGitHubRepo(*repo.Name, *repo.Name, *repo.CloneURL)
	}

	endTime := time.Now()

	fmt.Printf("\n\nTime taken to Fetch %v GitHub repos & add it to Gitpod projects: %v\n", len(allRepos), endTime.Sub(startTime))
}
