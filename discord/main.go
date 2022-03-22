package main

import (
	"context"
	"fmt"

	"github.com/DisgoOrg/disgo/discord"
	"github.com/DisgoOrg/disgo/webhook"
)

const (
	webHookID    = "937672879204618270"
	webHookToken = "d_tbjqQLt7qvvnrK93GwhYlFf_me9MRIpSc9d2KgWw6TsAghbYYruAA8LX9v0fXLiCbT"
)

func main() {
	client := webhook.NewClient(webHookID, webHookToken)
	defer client.Close(context.TODO())

	if _, err := client.CreateMessage(discord.NewWebhookMessageCreateBuilder().
		SetContentf("suck my digital dick").
		Build(),
	); err != nil {
		fmt.Printf("error sending message: %v", err)
	}
}
