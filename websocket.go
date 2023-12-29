package FrostAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
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

// Gets information about the user.
//
// Used for client, so you can do bot.Client.Username or something.
func (b *Bot) getInfo() User {
	resp, err := b.Request(true, "GET", "users/@me", nil, nil)
	if err != nil {
		return User{}
	}
	var user User
	decode(resp, &user)
	return user
}

// Returns Discord's build number.
func getBuildNumber() (int, error) { // see https://github.com/Pixens/Discord-Build-Number
	// I couldn't get this working with b.Request. Oh well.
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://discord.com/login", nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Accept-Encoding", "identity")

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	page := string(body)
	re := regexp.MustCompile(`assets/(sentry\.\w+)\.js`)
	endpoint := fmt.Sprintf("https://discord.com/assets/%s.js", re.FindStringSubmatch(page)[1])

	req2, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return 0, err
	}
	req2.Header.Set("Accept-Encoding", "identity")

	resp2, err := client.Do(req2)
	if err != nil {
		return 0, err
	}
	defer resp2.Body.Close()

	buildBody, err := io.ReadAll(resp2.Body)
	if err != nil {
		return 0, err
	}

	regex := regexp.MustCompile(`buildNumber\D+(\d+)"`)
	builds := regex.FindStringSubmatch(string(buildBody))

	if len(builds) >= 2 {
		build, err := strconv.Atoi(builds[1])
		if err != nil {
			return 0, err
		}
		return build, nil
	}

	return 1, nil
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
	build, err := getBuildNumber()
	if err != nil {
		build = 256231 // current build number as of now
	}

	payload := map[string]interface{}{
		"op": 2,
		"d": map[string]interface{}{
			"token":   b.Token,
			"intents": 3241725,
			"properties": map[string]interface{}{
				"$os":                      "Windows",
				"$browser":                 "Firefox",
				"$device":                  "",
				"system_locale":            "en-US",
				"browser_user_agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:120.0) Gecko/20100101 Firefox/120.0",
				"browser_version":          "120.0",
				"os_version":               "10",
				"referrer":                 "",
				"referring_domain":         "",
				"referrer_current":         "",
				"referring_domain_current": "",
				"release_channel":          "stable",
				"client_build_number":      build,
				"client_event_source":      nil,
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
