package github

import (
	"github.com/google/go-github/v32/github"
)

type GitHubClient struct {
	service *github.Client
}

func New() GitHubClient {
	return GitHubClient{service: github.NewClient(nil)}
}

// TODO: design Response
type TodaysContributionResponse string

// TODO: write function
func (gc *GitHubClient) GetTodaysContributions() TodaysContributionResponse {	
	return "contribution nums"
}
