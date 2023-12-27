package FrostAPI

import "fmt"

func (w *webhookManager) CreateWebhook(b *Bot, ChannelID, Name string) {
	endpoint := fmt.Sprintf("https://discord.com/api/v9/channels/%s/webhooks", ChannelID)

	data := map[string]interface{}{
		"name": Name,
	}
	customRequest(b, "POST", endpoint, data, nil)
}

func (w *webhookManager) DeleteWebhook(b *Bot, WebhookID string) {
	endpoint := fmt.Sprintf("https://discord.com/api/v9/webhooks/%s", WebhookID)
	customRequest(b, "DELETE", endpoint, nil, nil)
}
