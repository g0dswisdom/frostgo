package FrostAPI

func (u *UserManager) ChangeStatus(b *Bot, Status, Content string) {
	data := map[string]interface{}{
		"status": Status,
		"custom_status": map[string]string{
			"text": Content,
		},
	}
	customRequest(b, "PATCH", "https://discord.com/api/v9/users/@me/settings", data, nil)
}
