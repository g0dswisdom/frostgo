package FrostAPI

import (
	"fmt"
	"net/http"
)

// Sends a Discord message. Returns a Message object.
func (u *UserManager) SendMessage(b *Bot, ChannelID, Content string) (Message, error) {
	nonce := newNonce()
	endpoint := fmt.Sprintf("channels/%s/messages", ChannelID)
	data := map[string]interface{}{
		"content": Content,
		"nonce":   nonce,
		"tts":     false,
		"flags":   0,
	}
	resp, err := b.Request(true, http.MethodPost, endpoint, data, nil)
	if err != nil {
		return Message{}, err
	}

	var message Message
	if err := decode(resp, &message); err != nil {
		return Message{}, err
	}

	return message, nil
}

// Replies to a Discord message. Returns a Message object.
func (u *UserManager) SendMessageWithReply(b *Bot, ChannelID, MessageID, Content string) (Message, error) {
	nonce := newNonce()
	endpoint := fmt.Sprintf("channels/%s/messages", ChannelID)
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

	resp, err := b.Request(true, http.MethodPost, endpoint, data, nil)
	if err != nil {
		return Message{}, err
	}

	var message Message
	if err := decode(resp, &message); err != nil {
		return Message{}, err
	}

	return message, nil
}

// Deletes a Discord message.
func (u *UserManager) DeleteMessage(b *Bot, ChannelID, MessageID string) error {
	endpoint := fmt.Sprintf("channels/%s/messages/%s", ChannelID, MessageID)

	_, err := b.Request(true, http.MethodDelete, endpoint, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

// Edits a Discord message. Returns a Message object.
func (u *UserManager) EditMessage(b *Bot, ChannelID, MessageID string, Content string) (Message, error) {
	endpoint := fmt.Sprintf("channels/%s/messages/%s", ChannelID, MessageID)
	data := map[string]interface{}{
		"content": Content,
	}

	resp, err := b.Request(true, http.MethodPatch, endpoint, data, nil)
	if err != nil {
		return Message{}, err
	}

	var message Message
	if err := decode(resp, &message); err != nil {
		return Message{}, err
	}

	return message, nil
}

// Post a typing indicator for the specified channel.
func (u *UserManager) SendTyping(b *Bot, ChannelID string) error {
	endpoint := fmt.Sprintf("channels/%s/typing", ChannelID)

	_, err := b.Request(true, http.MethodPost, endpoint, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

// Pins a message.
func (g *GuildManager) PinMessage(b *Bot, ChannelID, MessageID string) error {
	endpoint := fmt.Sprintf("channels/%s/pins/%s", ChannelID, MessageID)

	_, err := b.Request(true, http.MethodPut, endpoint, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

// Unpins a message.
func (g *GuildManager) RemoveMessageFromPins(b *Bot, ChannelID, MessageID string) error {
	endpoint := fmt.Sprintf("channels/%s/pins/%s", ChannelID, MessageID)

	_, err := b.Request(true, http.MethodDelete, endpoint, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
