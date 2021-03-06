package model

type Data struct {
	UserID string `dynamo:"UserID"`
	Name string `dynamo:"Name"`
	GitHubTodaysContributionCount int `dynamo:"GitHubTodaysContributionCount"`
	CodingTime string `dynamo:"CodingTime"`
	CreatedAt int64 `dynamo:"CreatedAt"` // unix timestamp
}

