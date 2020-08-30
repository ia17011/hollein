package handler

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/guregu/dynamo"
	localDynamoDB "github.com/ia17011/hollein/dynamodb"
	"github.com/ia17011/hollein/github"
	"github.com/ia17011/hollein/repository"
)

const (
	region = endpoints.ApNortheast1RegionID
	timeFormat         = "20060102150405"
	dynamodbEndpoint = "http://dynamodb:8000"
	s3Endpoint = "http://s3:9000"
	tableName = "DataTable"
)

func Handler(ctx context.Context, event events.CloudWatchEvent) (string, error) {
	log.Println("EVENT: GitHubCrawler")
	dynamoDB := dynamodb.New(session.New(), localDynamoDB.Config(region, dynamodbEndpoint))	
	db := dynamo.NewFromIface(dynamoDB)

	// fetch GitHub Today's Contribution
	log.Println("Before: fetching contributionCount")
	githubClient := github.New(nil)
	contributionCount, err := githubClient.GetTodaysPublicContributions("ia17011")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Fetched: %d Contribution", contributionCount)
	log.Println("After: fetched contributionCount")

	log.Println("Before: save data to DynamoDB")
	dataRepository := repository.Data{Table: db.Table(tableName)}
	dataRepository.Save(contributionCount)
	log.Println("After: saved data in DynamoDB")

	return "success", nil
}
