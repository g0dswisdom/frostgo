package FrostAPI

import "fmt"

// Creates a new Discord webhook.
func (w *WebhookManager) CreateWebhook(b *Bot, ChannelID, Name string) {
	endpoint := fmt.Sprintf("https://discord.com/api/v9/channels/%s/webhooks", ChannelID)

	data := map[string]interface{}{
		"name": Name,
	}
	customRequest(b, "POST", endpoint, data, nil)
}

// Deletes a Discord webhook.
func (w *WebhookManager) DeleteWebhook(b *Bot, WebhookID string) {
	endpoint := fmt.Sprintf("https://discord.com/api/v9/webhooks/%s", WebhookID)
	customRequest(b, "DELETE", endpoint, nil, nil)
}
