package examples

import (
	"fmt"

	FrostAPI "github.com/g0dswisdom/frostgo"
)

func hasrole() {
	bot := FrostAPI.NewBot("Token")
	roleId := "role id"

	bot.On("ready", func() {
		fmt.Printf("[+] Logged in as %s\n", bot.Client.Username)
	})

	bot.On("messageCreate", func(message FrostAPI.Message) {
		if message.Content == "!ping" {
			bot.User.DeleteMessage(bot, message.ChannelID, message.ID)
			if bot.Guild.HasRole(bot, message.GuildID, roleId, message.Author.ID) {
				bot.User.SendMessage(bot, message.ChannelID, "User has role!")
			} else {
				bot.User.SendMessage(bot, message.ChannelID, "User doesn't have role!")
			}
		}
	})

	select {} // Keep the bot running
}
