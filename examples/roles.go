package examples

import (
	"fmt"

	FrostAPI "github.com/g0dswisdom/frostgo"
)

func roles() {
	bot := FrostAPI.NewBot("Token")

	bot.On("ready", func() {
		fmt.Println("[+] Frost selfbot is ready to use!")
	})

	bot.On("messageCreate", func(message FrostAPI.Message) {
		if message.Content == "!roles" {
			bot.User.DeleteMessage(bot, message.ChannelID, message.ID)
			roles, err := bot.Guild.GetRolesForUser(bot, message.GuildID, message.Author.ID)

			if err != nil {
				fmt.Println(err)
				return
			}

			for _, role := range roles {
				fmt.Println(role.Name)
			}
		}
	})

	select {} // Keep the bot running
}
