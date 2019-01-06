package main

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rickar/cal"

	"github.com/cnicolov/krisbot/pkg/pagerduty"
	"github.com/cnicolov/krisbot/pkg/paramz"
	"github.com/cnicolov/krisbot/pkg/slack"
)

var (
	pagerdutyToken string
	slackToken     string
	pagerdutyEmail string
	scheduleID     string
	messageInput   *slack.MessageInput = &slack.MessageInput{
		Message: `On behalf of @knikolov: <!here> I'm L1 and my workday is almost over. Taking off soon.`,
	}
)

func init() {
	paramzConfig := &paramz.Config{
		Provider: paramz.SSMParameterProvider,
		Prefix:   "/krisbot",
	}
	paramzProvider := paramz.New(paramzConfig)
	pagerdutyToken = paramzProvider.MustGetString("pagerduty_token", true)
	slackToken = paramzProvider.MustGetString("slack_token", true)
	pagerdutyEmail = paramzProvider.MustGetString("pagerduty_email", true)
	scheduleID = paramzProvider.MustGetString("pagerduty_schedule_id", true)
	messageInput.Channel = paramzProvider.MustGetString("workday_reminder_slack_channel", true)
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) error {
	date := time.Now()

	calendar := cal.NewCalendar()
	cal.AddUsHolidays(calendar)

	if !calendar.IsWorkday(date) {
		log.Println("Nothing todo - not a workday")
		return nil
	}

	slackClient := slack.New(slackToken)

	pd := pagerduty.New(pagerdutyToken)

	userIsL1, err := pd.IsL1(pagerdutyEmail, scheduleID)
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	if userIsL1 {
		log.Printf("Sending reminder message to channel %s", messageInput.Channel)
		return slackClient.Send(messageInput)
	}
	return nil
}

func main() {
	lambda.Start(Handler)
}
