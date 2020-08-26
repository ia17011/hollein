package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ia17011/hollein/handler"
)

func main() {
	lambda.Start(handler.Handler)
}
