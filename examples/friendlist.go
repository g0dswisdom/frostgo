package examples

import (
	"fmt"

	FrostAPI "github.com/g0dswisdom/frostgo"
)

func friends() {
	bot := FrostAPI.NewBot("Token")

	bot.On("ready", func() {
		fmt.Printf("[+] Logged in as %s\n", bot.Client.Username)

		friends, err := bot.User.GetFriends(bot)
		if err != nil {
			fmt.Println(err)
			return
		}

		for i, friend := range friends {
			fmt.Printf("%d: %s | %s\n", i, friend.User.Username, friend.ID)
		}
	})

	select {}
}
