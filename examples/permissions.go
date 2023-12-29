package examples

import (
	"fmt"

	FrostAPI "github.com/g0dswisdom/frostgo"
)

func perm() {
	bot := FrostAPI.NewBot("Token")

	bot.On("ready", func() {
		fmt.Println("[+] Frost selfbot is ready to use!")
	})

	bot.On("messageCreate", func(message FrostAPI.Message) {
		if message.Content == "!ping" {
			bot.User.DeleteMessage(bot, message.ChannelID, message.ID)
			roles, err := bot.Guild.GetRolesForUser(bot, message.GuildID, message.Author.ID)

			if err != nil {
				fmt.Println(err)
				return
			}

			for _, role := range roles {
				if role.HasPermission(FrostAPI.Administrator) {
					bot.User.SendMessage(bot, message.ChannelID, "User has a role with Administrator permissions! :white_check_mark:")
				}
			}
		}
	})

	select {} // Keep the bot running
}
