package handler

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/guregu/dynamo"
	localDynamoDB "github.com/ia17011/hollein/dynamodb"
	"github.com/ia17011/hollein/github"
)

const (
	region = endpoints.ApNortheast1RegionID
	timeFormat         = "20060102150405"
	dynamodbEndpoint = "http://dynamodb:8000"
	s3Endpoint = "http://s3:9000"
)

var (
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"
	ErrNoIP = errors.New("No IP in HTTP response")
	ErrNon200Response = errors.New("Non 200 Response found")
)

func Handler(ctx context.Context, event events.CloudWatchEvent) (string, error) {
	dynamoDB := dynamodb.New(session.New(), localDynamoDB.Config(region, dynamodbEndpoint))	
	db := dynamo.NewFromIface(dynamoDB)
	fmt.Println(db)

	githubClient := github.New()
	contributions := githubClient.GetTodaysContributions()
	fmt.Println(contributions)

	return "success", nil
}
