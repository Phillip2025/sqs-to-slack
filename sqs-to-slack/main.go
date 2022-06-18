package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/slack-go/slack"
)

func handler(ctx context.Context, request events.SQSEvent) error {

	fmt.Printf("Total records received to send to Slack: %d", len(request.Records))

	for i, r := range request.Records {
		if err := slack.PostWebhook(os.Getenv("SLACK_WEBHOOK"), &slack.WebhookMessage{
			Text: r.Body,
		}); err != nil {
			return fmt.Errorf("error posting message with webhook: %w", err)
		}
		fmt.Printf("Message %d sent", i+1)
	}

	fmt.Println("All messages sent successfully")

	return nil
}

func main() {
	lambda.Start(handler)
}
