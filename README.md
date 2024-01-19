# FrostGo
FrostGo is a fast  Discord API wrapper developed in Go, for selfbots

# Installation
To install FrostGo, use the following command in your terminal:

`go get github.com/g0dswisdom/frostgo`
# Documentation
We're currently working on our documentation. For now, explore the example below to create new commands.

You can also check out the command examples in the `examples` folder.

You can also use [this](https://pkg.go.dev/github.com/g0dswisdom/frostgo) for API reference.
# Selfbot example
```go
package main

import (
	"fmt"

	FrostAPI "github.com/g0dswisdom/frostgo"
)

func main() {
	bot := FrostAPI.NewBot("Discord token")

	bot.On("ready", func() {
		fmt.Printf("[+] Logged in as %s\n", bot.Client.Username)
	})

	bot.On("messageCreate", func(message FrostAPI.Message) {
		if message.Content == "!ping" {
			bot.User.DeleteMessage(bot, message.ChannelID, message.ID)
			bot.User.SendMessage(bot, message.ChannelID, "Pong! :ping_pong:")
		}
	})

	select {} // Keep the bot running
}
```
# Note
Using a Discord selfbot is against the Discord TOS.

The author is not responsible for any damages done to the users account.
