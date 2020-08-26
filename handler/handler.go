package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/guregu/dynamo"
	localDynamoDB "github.com/ia17011/hollein/dynamodb"
	"github.com/ia17011/hollein/github"
)

const (
	region = endpoints.ApNortheast1RegionID
	timeFormat         = "20060102150405"
	dynamodbEndpoint = "http://dynamodb:8000"
	s3Endpoint = "http://s3:9000"
	githubTableName = "GitHubContributions"
)

func Handler(ctx context.Context, event events.CloudWatchEvent) (string, error) {
	xray.Configure(xray.Config{LogLevel: "trace"})

	dynamoDB := dynamodb.New(session.New(), localDynamoDB.Config(region, dynamodbEndpoint))	
	db := dynamo.NewFromIface(dynamoDB)

	// fetch GitHub Today's Contribution
	githubClient := github.New()
	contributionCount, err := githubClient.GetTodaysContributions("ia17011")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)	
	fmt.Println(contributionCount)

	return "success", nil
}
