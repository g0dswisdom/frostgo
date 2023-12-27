package FrostAPI

import "fmt"

// Sends a Discord message.
func (u *UserManager) SendMessage(b *Bot, ChannelID, Content string) {
	nonce := NewNonce()

	endpoint := fmt.Sprintf("https://discord.com/api/v9/channels/%s/messages", ChannelID)

	data := map[string]interface{}{
		"content": Content,
		"nonce":   nonce,
		"tts":     false,
		"flags":   0,
	}
	customRequest(b, "POST", endpoint, data, nil)
}

// Replies to a Discord message.
func (u *UserManager) SendMessageWithReply(b *Bot, ChannelID, MessageID, Content string) {
	nonce := NewNonce()
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
	customRequest(b, "POST", endpoint, data, nil)
}

// Deletes a Discord message.
func (u *UserManager) DeleteMessage(b *Bot, ChannelID, MessageID string) {
	endpoint := fmt.Sprintf("https://discord.com/api/v9/channels/%s/messages/%s", ChannelID, MessageID)
	customRequest(b, "DELETE", endpoint, nil, nil)
}

// Edits a Discord message.
func (u *UserManager) EditMessage(b *Bot, ChannelID, MessageID string, Content string) {
	endpoint := fmt.Sprintf("https://discord.com/api/v9/channels/%s/messages/%s", ChannelID, MessageID)

	data := map[string]interface{}{
		"content": Content,
	}
	customRequest(b, "PATCH", endpoint, data, nil)
}
