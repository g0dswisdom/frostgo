package examples

import (
	"fmt"

	FrostAPI "github.com/g0dswisdom/frostgo"
)

func timeout() {
	bot := FrostAPI.NewBot("Token")
	userId := "user id"

	bot.On("ready", func() {
		fmt.Println("[+] Frost selfbot is ready to use!")
	})

	bot.On("messageCreate", func(message FrostAPI.Message) {
		if message.Content == "!timeout" {
			bot.User.DeleteMessage(bot, message.ChannelID, message.ID)
			options := FrostAPI.TimeoutOptions{
				MinutesToAdd: 30,
				DaysToAdd:    1,
			}
			bot.Guild.CreateTimeout(bot, message.GuildID, userId, options)
			bot.User.SendMessage(bot, message.ChannelID, "Created timeout!")
		}
	})

	select {} // Keep the bot running
}
