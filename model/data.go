package model

import "time"

type Data struct {
	ID int `dynamo:"ID"`
	UserID string `dynamo:"UserID"`
	Name string `dynamo:"Name"`
	GitHubTodaysContributionCount int `dynamo:"GitHubTodaysContributionCount"`
	CreatedAt time.Time `dynamo:"CreatedAt"`
}

