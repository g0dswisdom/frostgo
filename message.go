package FrostAPI

import "fmt"

// Sends a Discord message. Returns a Message object.
func (u *UserManager) SendMessage(b *Bot, ChannelID, Content string) Message {
	nonce := newNonce()
	endpoint := fmt.Sprintf("https://discord.com/api/v9/channels/%s/messages", ChannelID)
	data := map[string]interface{}{
		"content": Content,
		"nonce":   nonce,
		"tts":     false,
		"flags":   0,
	}
	var message Message
	decode(customRequest(b, "POST", endpoint, data, nil), &message)
	return message
}

// Replies to a Discord message. Returns a Message object.
func (u *UserManager) SendMessageWithReply(b *Bot, ChannelID, MessageID, Content string) Message {
	nonce := newNonce()
	endpoint := fmt.Sprintf("https://discord.com/api/v9/channels/%s/messages", ChannelID)
	data := map[string]interface{}{
		"content": Content,
		"nonce":   nonce,
		"tts":     false,
		"flags":   0,
		"message_reference": map[string]string{
			"channel_id": ChannelID,
			"message_id": MessageID,
		},
	}
	var message Message
	decode(customRequest(b, "POST", endpoint, data, nil), &message)
	return message
}

// Deletes a Discord message.
func (u *UserManager) DeleteMessage(b *Bot, ChannelID, MessageID string) {
	endpoint := fmt.Sprintf("https://discord.com/api/v9/channels/%s/messages/%s", ChannelID, MessageID)
	customRequest(b, "DELETE", endpoint, nil, nil)
}

// Edits a Discord message. Returns a Message object.
func (u *UserManager) EditMessage(b *Bot, ChannelID, MessageID string, Content string) Message {
	endpoint := fmt.Sprintf("https://discord.com/api/v9/channels/%s/messages/%s", ChannelID, MessageID)
	data := map[string]interface{}{
		"content": Content,
	}
	var message Message
	decode(customRequest(b, "PATCH", endpoint, data, nil), &message)
	return message
}

// Post a typing indicator for the specified channel.
func (u *UserManager) SendTyping(b *Bot, ChannelID string) {
	endpoint := fmt.Sprintf("https://discord.com/api/v9/channels/%s/typing", ChannelID)
	customRequest(b, "POST", endpoint, nil, nil)
}

// Pins a message.
func (g *GuildManager) PinMessage(b *Bot, ChannelID, MessageID string) {
	endpoint := fmt.Sprintf("https://discord.com/api/v9/channels/%s/pins/%s", ChannelID, MessageID)
	customRequest(b, "PUT", endpoint, nil, nil)
}

// Unpins a message.
func (g *GuildManager) RemoveMessageFromPins(b *Bot, ChannelID, MessageID string) {
	endpoint := fmt.Sprintf("https://discord.com/api/v9/channels/%s/pins/%s", ChannelID, MessageID)
	customRequest(b, "DELETE", endpoint, nil, nil)
}
