package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	t.Run("TestHandler", func(t *testing.T) {
		ctx := context.Background()
		request := events.APIGatewayProxyRequest{}
		expectedResponse := events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       "\"Hello from Lambda!\"",
		}

		response, err := handler(ctx, request)
		if err != nil {
			t.Errorf("handler returned an error: %v", err)
		}

		if response.StatusCode != expectedResponse.StatusCode {
			t.Errorf("handler returned unexpected status code: got %v want %v",
				response.StatusCode, expectedResponse.StatusCode)
		}

		if response.Body != expectedResponse.Body {
			t.Errorf("handler returned unexpected body: got %v want %v",
				response.Body, expectedResponse.Body)
		}
	})
}
