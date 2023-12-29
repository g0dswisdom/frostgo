package FrostAPI

import "net/http"

// Changes status.
func (u *UserManager) ChangeStatus(b *Bot, Options StatusOptions) error {
	data := map[string]interface{}{
		"status": Options.Status,
		"custom_status": map[string]string{
			"text": Options.Content,
		},
	}

	_, err := b.Request(true, http.MethodPatch, "users/@me/settings", data, nil)
	if err != nil {
		return err
	}
	return nil
	//customRequest(b, "PATCH", "https://discord.com/api/v9/users/@me/settings", data, nil)
}

// TODO: Settings

// Creates a Friend Invite. Returns a GuildInvite object, along with any encountered errors.
func (u *UserManager) CreateFriendInvite(b *Bot) (GuildInvite, error) {
	resp, err := b.Request(true, http.MethodPost, "users/@me/invites", nil, nil)
	if err != nil {
		return GuildInvite{}, err
	}

	var invite GuildInvite
	if err := decode(resp, &invite); err != nil {
		return GuildInvite{}, err
	}

	return invite, nil
}

func (u *UserManager) GetFriends(b *Bot) ([]Friend, error) {
	resp, err := b.Request(true, http.MethodGet, "users/@me/relationships", nil, nil)
	if err != nil {
		return nil, err
	}

	var friends []Friend
	if err := decode(resp, &friends); err != nil {
		return nil, err
	}

	return friends, nil
}
