package github

import (
	"context"
	"fmt"
	"log"

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
	events, response, err := gc.service.Activity.ListEventsPerformedByUser(context.Background(), "ia17011", true, nil)
	if _, ok := err.(*github.RateLimitError); ok {
		log.Println("hit late limit")
	}

	for i, event := range events {
		fmt.Printf("%v, %v\n", i+1, event.GetType())
	}


	fmt.Println("--------------------------------")
	fmt.Println(response)

	return "contribution nums"
}
