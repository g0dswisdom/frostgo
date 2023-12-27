# Frostgo
FrostGo is a customizable fully open-source Discord API wrapper created in Go, for selfbots.

FrostGo is still WIP (work in progress), which means its unfinished.
# Installation
Simply do `go get github.com/g0dswisdom/frostgo` in your terminal.
# Documentation
This project currently lacks documentation, but you can create more commands using the example below.
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
		fmt.Println("[+] Frost selfbot is ready to use!")
	})

	bot.On("messageCreate", func(message FrostAPI.Message) {
		if message.Content == "!ping" && message.Author.ID == "1043434358637342803" {
			bot.User.DeleteMessage(bot, message.ChannelID, message.ID)
			bot.User.SendMessage(bot, message.ChannelID, "Pong! :ping_pong:")
		}
	})

	select {} // Keep the bot running
}
```
