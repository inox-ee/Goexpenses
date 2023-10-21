package main

import (
	"bytes"
	"context"
	"encoding/json"
	"inox-ee/Goexpenses/util"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

type Handler func(ctx context.Context, event events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error)

type Slack struct {
	client *slack.Client
}

func NewMainHandler(slk Slack) Handler {
	return func(ctx context.Context, event events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
		var response events.LambdaFunctionURLResponse
		switch event.RawPath {
		case "/":
			eventApiEvent, err := slackevents.ParseEvent(json.RawMessage(event.Body), slackevents.OptionNoVerifyToken())
			if err != nil {
				return events.LambdaFunctionURLResponse{
					StatusCode: 500,
				}, err
			}

			innerEvent := eventApiEvent.InnerEvent

			switch event := innerEvent.Data.(type) {
			case *slackevents.AppMentionEvent:
				message := strings.Split(event.Text, " ")
				var slackCommand string
				slackCommand = message[1]

				switch slackCommand {
				case "ping":
					_, _, err := slk.client.PostMessage(event.Channel, slack.MsgOptionText("pong", false))
					if err != nil {
						return events.LambdaFunctionURLResponse{
							StatusCode: 500,
						}, err
					}
				}
				response = events.LambdaFunctionURLResponse{
					StatusCode: 200,
					Body:       "\"Hello from Lambda!\"",
				}
			}
		case "/challenge":
			response = events.LambdaFunctionURLResponse{
				StatusCode: 200,
				Body:       event.Body,
			}
		}
		return response, nil
	}
}

func slackVerificationMiddleware(config util.Config, next Handler) Handler {
	return func(ctx context.Context, event events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
		var r *http.Request
		for k, v := range event.Headers {
			r.Header.Add(k, v)
		}
		verifier, err := slack.NewSecretsVerifier(r.Header, config.SlackSigningSecret)
		if err != nil {
			return events.LambdaFunctionURLResponse{
				StatusCode: 500,
			}, err
		}
		bodyReader := io.TeeReader(r.Body, &verifier)
		body, err := ioutil.ReadAll(bodyReader)
		if err != nil {
			return events.LambdaFunctionURLResponse{
				StatusCode: 500,
			}, err
		}
		if err := verifier.Ensure(); err != nil {
			return events.LambdaFunctionURLResponse{
				StatusCode: 403,
			}, err
		}
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		return next(ctx, event)
	}
}

func main() {
	config, err := util.LoadConfig(".", "app")
	if err != nil {
		panic(err)
	}
	client := slack.New(config.SlackBotToken, slack.OptionDebug(true), slack.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)))
	slk := Slack{
		client: client,
	}
	handler := NewMainHandler(slk)
	lambda.Start(slackVerificationMiddleware(config, handler))
}
