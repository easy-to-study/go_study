package main

import (
    "github.com/ashwanthkumar/slack-go-webhook"
    "os"
	"fmt"
)

const (
    WEBHOOKURL = "https://hooks.slack.com/services/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"
    CHANNEL    = "times_t-sataga"
    USERNAME   = "Go"
)

func main() {
    PostSlack("HelloWorld!!")
}

func PostSlack(msg string) {
    field1 := slack.Field{Title: "Message", Value: msg}
    field2 := slack.Field{Title: "AnythingKey", Value: "AnythingValue"}

    attachment := slack.Attachment{}
    attachment.AddField(field1).AddField(field2)
    color := "good"
    attachment.Color = &color
    payload := slack.Payload{
        Username:    USERNAME,
        Channel:     CHANNEL,
        Attachments: []slack.Attachment{attachment},
    }
    err := slack.Send(WEBHOOKURL, "", payload)
    if err != nil {
		fmt.Errorf("%s", err)
        os.Exit(1)
    }
}
