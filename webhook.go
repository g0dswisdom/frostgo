package FrostAPI

import (
	"fmt"
	"net/http"
)

// Creates a new Discord webhook. Returns a Webhook object, along with any encountered errors.
func (w *WebhookManager) CreateWebhook(b *Bot, ChannelID, Name string) (Webhook, error) {
	endpoint := fmt.Sprintf("channels/%s/webhooks", ChannelID)
	data := map[string]interface{}{
		"name": Name,
	}

	resp, err := b.Request(true, http.MethodPost, endpoint, data, nil)
	if err != nil {
		return Webhook{}, err
	}

	var webhook Webhook
	if err := decode(resp, &webhook); err != nil {
		return Webhook{}, err
	}
	return webhook, nil
}

// Deletes a Discord webhook. Requires MANAGE_WEBHOOKS.
func (w *WebhookManager) DeleteWebhook(b *Bot, WebhookID string) error {
	endpoint := fmt.Sprintf("webhooks/%s", WebhookID)

	_, err := b.Request(true, http.MethodDelete, endpoint, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

// Returns all of the webhooks from the specified channel, along with any encountered errors.
func (w *WebhookManager) GetChannelWebhooks(b *Bot, ChannelID string) ([]Webhook, error) {
	endpoint := fmt.Sprintf("channels/%s/webhooks", ChannelID)

	resp, err := b.Request(true, http.MethodGet, endpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	var webhooks []Webhook
	if err := decode(resp, &webhooks); err != nil {
		return nil, err
	}
	return webhooks, nil
}
