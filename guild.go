package FrostAPI

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// Kicks an user.
func (g *GuildManager) KickUser(b *Bot, GuildID, User string) error {
	endpoint := fmt.Sprintf("guilds/%s/members/%s", GuildID, User)

	_, err := b.Request(true, http.MethodDelete, endpoint, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

// Bans an user.
func (g *GuildManager) BanUser(b *Bot, GuildID, User string) error {
	endpoint := fmt.Sprintf("guilds/%s/bans/%s", GuildID, User)

	_, err := b.Request(true, http.MethodPut, endpoint, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

// Creates a new channel. Returns a Channel object.
func (g *GuildManager) CreateChannel(b *Bot, GuildID, CategoryID, Name string) (Channel, error) {
	endpoint := fmt.Sprintf("guilds/%s/channels", GuildID)
	data := map[string]interface{}{
		"name":      Name,
		"parent_id": CategoryID,
		"type":      0,
	}

	resp, err := b.Request(true, http.MethodPost, endpoint, data, nil)
	if err != nil {
		return Channel{}, err
	}

	var channel Channel
	if err := decode(resp, &channel); err != nil {
		return Channel{}, err
	}
	return channel, nil
}

// Creates a new channel, with no category. Returns a Channel object.
func (g *GuildManager) CreateChannelNoCategory(b *Bot, GuildID, Name string) (Channel, error) {
	endpoint := fmt.Sprintf("guilds/%s/channels", GuildID)
	data := map[string]interface{}{
		"name": Name,
		"type": 0,
	}

	resp, err := b.Request(true, http.MethodPost, endpoint, data, nil)
	if err != nil {
		return Channel{}, err
	}

	var channel Channel
	if err := decode(resp, &channel); err != nil {
		return Channel{}, err
	}
	return channel, nil
}

// Sets a channel's topic.
func (g *GuildManager) SetChannelTopic(b *Bot, ChannelID, Topic string) error {
	endpoint := fmt.Sprintf("channels/%s", ChannelID)
	data := map[string]interface{}{
		"topic": Topic,
	}

	_, err := b.Request(true, http.MethodPatch, endpoint, data, nil)
	if err != nil {
		return err
	}

	return nil
}

// Sets an user's nickname. Returns an User object.
func (g *GuildManager) SetUserNickname(b *Bot, GuildID, Member, Nickname string) (User, error) {
	endpoint := fmt.Sprintf("guilds/%s/members/%s", GuildID, Member)
	data := map[string]interface{}{
		"nick": Nickname,
	}

	resp, err := b.Request(true, http.MethodPatch, endpoint, data, nil)
	if err != nil {
		return User{}, err
	}

	var user User
	if err := decode(resp, &user); err != nil {
		return User{}, err
	}

	return user, nil
}

// Creates a Discord invite to the given channel. Returns a GuildInvite object.
// The invite options are specified using GuildInviteOptions.
//
// MaxAge is specified in miliseconds.
func (g *GuildManager) CreateInvite(b *Bot, ChannelID string, options GuildInviteOptions) (GuildInvite, error) {
	endpoint := fmt.Sprintf("channels/%s/invites", ChannelID)
	data := map[string]interface{}{
		"max_age":     options.MaxAge,
		"max_uses":    options.MaxUses,
		"target_type": nil,
		"temporary":   false,
		"validate":    nil,
	}

	resp, err := b.Request(true, http.MethodPost, endpoint, data, nil)
	if err != nil {
		return GuildInvite{}, err
	}

	var invite GuildInvite
	if err := decode(resp, &invite); err != nil {
		return GuildInvite{}, err
	}

	return invite, nil
}

// Creates a timeout.
// The timeout duration is specified using TimeoutOpt ions.
func (g *GuildManager) CreateTimeout(b *Bot, GuildID, UserID string, Options TimeoutOptions) error {
	endpoint := fmt.Sprintf("guilds/%s/members/%s", GuildID, UserID)
	timestamp := time.Now().UTC().AddDate(0, 0, Options.DaysToAdd).Add(time.Minute * time.Duration(Options.MinutesToAdd)).Format(time.RFC3339)
	data := map[string]interface{}{
		"communication_disabled_until": timestamp,
	}

	_, err := b.Request(true, http.MethodPatch, endpoint, data, nil)
	if err != nil {
		return err
	}

	return nil
}

// Removes a timeout.
func (g *GuildManager) RemoveTimeout(b *Bot, GuildID, UserID string) error {
	endpoint := fmt.Sprintf("guilds/%s/members/%s", GuildID, UserID)
	data := map[string]interface{}{
		"communication_disabled_until": nil,
	}

	_, err := b.Request(true, http.MethodPatch, endpoint, data, nil)
	if err != nil {
		return err
	}

	return nil
}

// Returns a GuildMember object. Not fully functional.
func (g *GuildManager) GetGuildMember(b *Bot, GuildID, UserID string) (GuildMember, error) {
	endpoint := fmt.Sprintf("guilds/%s/members/%s", GuildID, UserID)
	resp, err := b.Request(true, http.MethodGet, endpoint, nil, nil)
	if err != nil {
		return GuildMember{}, err
	}

	var member GuildMember
	if err := decode(resp, &member); err != nil {
		return GuildMember{}, err
	}

	return member, nil
}

// Returns a Guild object.
func (g *GuildManager) GetGuild(b *Bot, GuildID string) (Guild, error) {
	endpoint := fmt.Sprintf("guilds/%s", GuildID)
	resp, err := b.Request(true, http.MethodGet, endpoint, nil, nil)
	if err != nil {
		return Guild{}, err
	}

	var guild Guild
	if err := decode(resp, &guild); err != nil {
		return Guild{}, err
	}

	return guild, nil
}

// Returns an array of Channel objects.
func (g *GuildManager) GetGuildChannels(b *Bot, GuildID string) ([]Channel, error) {
	endpoint := fmt.Sprintf("guilds/%s/channels", GuildID)
	resp, err := b.Request(true, http.MethodGet, endpoint, nil, nil)
	if err != nil {
		return []Channel{}, err
	}

	var channels []Channel
	if err := decode(resp, &channels); err != nil {
		return []Channel{}, err
	}

	return channels, nil
}

// Returns an array of role IDs.
func (g *GuildManager) GetRolesForUser(b *Bot, GuildID, UserID string) ([]Role, error) {
	endpoint := fmt.Sprintf("guilds/%s/members/%s", GuildID, UserID)
	resp, err := b.Request(true, http.MethodGet, endpoint, nil, nil)
	if err != nil {
		return []Role{}, err
	}

	var member GuildMember
	if err := decode(resp, &member); err != nil {
		return []Role{}, err
	}

	guild, err := g.GetGuild(b, GuildID)
	if err != nil {
		return []Role{}, err
	}

	var userRoles []Role
	for _, roleID := range member.Roles {
		for _, role := range guild.Roles {
			if role.ID == roleID {
				userRoles = append(userRoles, role)
				break
			}
		}
	}

	return userRoles, nil
}

func (r *Role) HasPermission(permission Permission) bool {
	rolePermissions, err := strconv.ParseInt(r.Permissions, 10, 64)
	if err != nil {
		return false
	}
	return (rolePermissions & int64(permission)) == int64(permission)
}

// Checks if an user has the specified role, in the given guild.
func (g *GuildManager) HasRole(b *Bot, GuildID, RoleID, UserID string) bool {
	roles, err := g.GetRolesForUser(b, GuildID, UserID)

	if err != nil {
		return false
	}

	for _, role := range roles {
		if role.ID == RoleID {
			return true
		}
	}

	return false
}

// Returns a GuildInvite object.
func (g *GuildManager) GetInvite(b *Bot, Invite string) (GuildInvite, error) {
	endpoint := fmt.Sprintf("invites/%s", Invite)
	resp, err := b.Request(true, http.MethodGet, endpoint, nil, nil)
	if err != nil {
		return GuildInvite{}, err
	}

	var invite GuildInvite
	if err := decode(resp, &invite); err != nil {
		return GuildInvite{}, err
	}

	return invite, nil
}

// Returns a Channel object.
func (g *GuildManager) GetChannel(b *Bot, ChannelID string) (Channel, error) {
	endpoint := fmt.Sprintf("channels/%s", ChannelID)
	resp, err := b.Request(true, http.MethodGet, endpoint, nil, nil)
	if err != nil {
		return Channel{}, err
	}

	var channel Channel
	if err := decode(resp, &channel); err != nil {
		return Channel{}, err
	}

	return channel, nil
}
