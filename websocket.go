package FrostAPI

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Types

// This contains all of the webhook functions.
type WebhookManager struct{}

// This contains all of the user functions.
type UserManager struct{}

// This contains all of the guild functions.
type GuildManager struct{}

// Discord bot.
type Bot struct {
	Token          string
	Client         User
	eventListeners map[string][]interface{}
	conn           *websocket.Conn
	mutex          sync.Mutex
	User           UserManager
	Webhooks       WebhookManager
	Guild          GuildManager
}

// Bot functions and event emitter

func (b *Bot) getInfo() User {
	resp, err := b.Request(true, "GET", "https://discord.com/api/v9/users/@me", nil, nil)
	if err != nil {
		return User{}
	}
	var user User
	decode(resp, &user)
	return user
}

func NewBot(token string) *Bot {
	bot := &Bot{
		Token:          token,
		eventListeners: make(map[string][]interface{}),
	}
	userInfo := bot.getInfo()
	bot.Client = userInfo
	go bot.connectToDiscord()
	return bot
}

func (b *Bot) On(event string, listener interface{}) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.eventListeners[event] = append(b.eventListeners[event], listener)
}

func (b *Bot) Emit(event string, data interface{}) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if listeners, ok := b.eventListeners[event]; ok {
		for _, listener := range listeners {
			switch fn := listener.(type) {
			case func():
				fn()
			case func(Message):
				fn(data.(Message))
			case func(Channel):
				fn(data.(Channel))
			}
		}
	}
}

// Websocket

func (b *Bot) connectToDiscord() {
	http.DefaultTransport.(*http.Transport).MaxIdleConns = 100
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 100

	websocketURL := "wss://gateway.discord.gg/?v=9&encoding=json"
	conn, _, err := websocket.DefaultDialer.Dial(websocketURL, nil)
	if err != nil {
		log.Fatal("Could not connect to Discord's websocket!", err)
	}

	b.conn = conn
	b.onOpen()
	b.receiveMessages()
}

func (b *Bot) onOpen() {
	payload := map[string]interface{}{
		"op": 2,
		"d": map[string]interface{}{
			"token":   b.Token,
			"intents": 3241725,
			"properties": map[string]interface{}{
				"$os":      "linux",
				"$browser": "my_library",
				"$device":  "my_library",
			},
		},
	}
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Could not parse the auth payload to JSON!")
		return
	}
	if err := b.conn.WriteMessage(websocket.TextMessage, payloadJSON); err != nil {
		fmt.Println("Could not send auth message to the websocket! Error: ", err)
		return
	}
	b.Emit("ready", nil)
}

func (b *Bot) receiveMessages() {
	for {
		_, data, err := b.conn.ReadMessage()
		if err != nil {
			b.handleWebSocketError(err)
			return
		}

		b.onMessage(data)
	}
}

func (b *Bot) handleWebSocketError(err error) {
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		fmt.Println("Websocket connection closed")
	} else {
		fmt.Println(err)
		time.Sleep(5 * time.Second)
		go b.connectToDiscord()
	}
}

func (b *Bot) onMessage(data []byte) {
	var message DiscordMessage

	err := json.Unmarshal(data, &message)
	if err != nil {

	}
	switch message.Op {
	case 0:
		switch message.T {
		case "MESSAGE_CREATE":
			var messageData Message
			err := json.Unmarshal(message.D, &messageData)
			if err != nil {
				return
			}
			b.Emit("messageCreate", messageData)
		case "CHANNEL_CREATE":
			var channelData Channel
			err := json.Unmarshal(message.D, &channelData)
			if err != nil {
				return
			}
			b.Emit("channelCreate", channelData)
		case "CHANNEL_UPDATE":
			var channelData Channel
			err := json.Unmarshal(message.D, &channelData)
			if err != nil {
				return
			}
			b.Emit("channelUpdate", channelData)
		case "CHANNEL_DELETE":
			var channelData Channel
			err := json.Unmarshal(message.D, &channelData)
			if err != nil {
				return
			}
			b.Emit("channelDelete", channelData)
		}

	case 10:
		var heartbeatData HeartbeatPayloadData
		err := json.Unmarshal(message.D, &heartbeatData)
		if err != nil {
			fmt.Println("Error parsing heartbeat:", err)
			return
		}

		heartbeatInterval := heartbeatData.HeartbeatInterval

		go b.sendHeartbeat(heartbeatInterval)
	}
}

func (b *Bot) sendHeartbeat(interval int) {
	payload := `{"op": 1, "d": null}`
	ticker := time.NewTicker(time.Duration(interval) * time.Millisecond)
	defer ticker.Stop()
	for range ticker.C {

		if err := b.conn.WriteMessage(websocket.TextMessage, []byte(payload)); err != nil {
			return
		}
	}
}
