package FrostAPI

// Changes status.
func (u *UserManager) ChangeStatus(b *Bot, Status, Content string) {
	data := map[string]interface{}{
		"status": Status,
		"custom_status": map[string]string{
			"text": Content,
		},
	}
	customRequest(b, "PATCH", "https://discord.com/api/v9/users/@me/settings", data, nil)
}

// TODO: Settings

// Creates a Friend Invite. Returns a GuildInvite object.
func (u *UserManager) CreateFriendInvite(b *Bot) GuildInvite {
	var invite GuildInvite
	decode(customRequest(b, "POST", "https://discord.com/api/v9/users/@me/invites", nil, nil), &invite)
	return invite
}
