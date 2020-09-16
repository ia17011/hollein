package model

import "time"

type Data struct {
	UserID string `dynamo:"UserID"`
	Name string `dynamo:"Name"`
	GitHubTodaysContributionCount int `dynamo:"GitHubTodaysContributionCount"`
	CreatedAt time.Time `dynamo:"CreatedAt"`
}

