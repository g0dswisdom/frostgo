package FrostAPI

import "fmt"

// Sends a Discord message.
func (u *userManager) SendMessage(b *Bot, ChannelID, Content string) {
	nonce := newNonce()

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
func (u *userManager) SendMessageWithReply(b *Bot, ChannelID, MessageID, Content string) {
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
	customRequest(b, "POST", endpoint, data, nil)
}

// Deletes a Discord message.
func (u *userManager) DeleteMessage(b *Bot, ChannelID, MessageID string) {
	endpoint := fmt.Sprintf("https://discord.com/api/v9/channels/%s/messages/%s", ChannelID, MessageID)
	customRequest(b, "DELETE", endpoint, nil, nil)
}

// Edits a Discord message.
func (u *userManager) EditMessage(b *Bot, ChannelID, MessageID string, Content string) {
	endpoint := fmt.Sprintf("https://discord.com/api/v9/channels/%s/messages/%s", ChannelID, MessageID)

	data := map[string]interface{}{
		"content": Content,
	}
	customRequest(b, "PATCH", endpoint, data, nil)
}
