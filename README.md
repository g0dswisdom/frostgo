# FrostGo
FrostGo is a fast and efficient Discord API wrapper developed in Go, designed specifically for selfbots.

Please note: FrostGo is a work in progress (WIP), and some features might still be under development. Updates are made regularly by our single-person team.
# Features
* **Speed** - FrostGo is created in **Go**, which is one of the fastest programming languages.
* **User-Friendly** - FrostGo ensures a straightforward setup and smooth implementation for creating selfbots. It's designed to be user-friendly, even for those new to Go.
* **Stability** - FrostGo is built to provide a stable and reliable package for the Discord API.
# Installation
To install FrostGo, use the following command in your terminal:

`go get github.com/g0dswisdom/frostgo`
# Documentation
We're currently working on our documentation. For now, explore the example below to create new commands.

Additionally, check out the command examples in the `examples` folder for guidance.

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