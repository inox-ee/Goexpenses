package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	var response events.LambdaFunctionURLResponse
	switch event.RawPath {
	case "/":
		response = events.LambdaFunctionURLResponse{
			StatusCode: 200,
			Body:       "\"Hello from Lambda!\"" + event.Body + event.RawPath + "hogehoge",
		}
	case "/challenge":
		response = events.LambdaFunctionURLResponse{
			StatusCode: 200,
			Body:       event.Body,
		}
	}
	return response, nil
}

func main() {
	lambda.Start(handler)
}
