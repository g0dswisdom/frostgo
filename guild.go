package FrostAPI

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// Kicks an user.
func (g *GuildManager) KickUser(b *Bot, GuildID, User string) {
	endpoint := fmt.Sprintf("https://discord.com/api/v9/guilds/%s/members/%s", GuildID, User)
	customRequest(b, "DELETE", endpoint, nil, nil)
}

// Bans an user.
func (g *GuildManager) BanUser(b *Bot, GuildID, User string) {
	endpoint := fmt.Sprintf("https://discord.com/api/v9/guilds/%s/bans/%s", GuildID, User)
	customRequest(b, "PUT", endpoint, nil, nil)
}

// Creates a new channel.
func (g *GuildManager) CreateChannel(b *Bot, GuildID, CategoryID, Name string) Channel {
	endpoint := fmt.Sprintf("https://discord.com/api/v9/guilds/%s/channels", GuildID)
	data := map[string]interface{}{
		"name":      Name,
		"parent_id": CategoryID,
		"type":      0,
	}
	var channel Channel
	decode(customRequest(b, "POST", endpoint, data, nil), &channel)
	return channel
}

// Creates a new channel, with no category.
func (g *GuildManager) CreateChannelNoCategory(b *Bot, GuildID, Name string) Channel {
	endpoint := fmt.Sprintf("https://discord.com/api/v9/guilds/%s/channels", GuildID)
	data := map[string]interface{}{
		"name": Name,
		"type": 0,
	}
	var channel Channel
	decode(customRequest(b, "POST", endpoint, data, nil), &channel)
	return channel
}

// Sets a channel's topic.
func (g *GuildManager) SetChannelTopic(b *Bot, ChannelID, Topic string) {
	endpoint := fmt.Sprintf("https://discord.com/api/v9/channels/%s", ChannelID)
	data := map[string]interface{}{
		"topic": Topic,
	}
	customRequest(b, "PATCH", endpoint, data, nil)
}

// Sets an user's nickname.
func (g *GuildManager) SetUserNickname(b *Bot, GuildID, Member, Nickname string) User {
	endpoint := fmt.Sprintf("https://discord.com/api/v9/guilds/%s/members/%s", GuildID, Member)
	data := map[string]interface{}{
		"nick": Nickname,
	}
	var user User
	decode(customRequest(b, "PATCH", endpoint, data, nil), &user)
	return user
}

// Creates a Discord invite to the given channel.
// The invite options are specified using GuildInviteOptions.
func (g *GuildManager) CreateInvite(b *Bot, ChannelID string, options GuildInviteOptions) GuildInvite {
	endpoint := fmt.Sprintf("https://discord.com/api/v9/channels/%s/invites", ChannelID)
	data := map[string]interface{}{
		"max_age":     options.MaxAge,
		"max_uses":    options.MaxUses,
		"target_type": nil,
		"temporary":   false,
		"validate":    nil,
	}
	var invite GuildInvite
	decode(customRequest(b, "POST", endpoint, data, nil), &invite)
	return invite
}

// Creates a timeout.
// The timeout duration is specified using TimeoutOptions.
func (g *GuildManager) CreateTimeout(b *Bot, GuildID, UserID string, Options TimeoutOptions) {
	endpoint := fmt.Sprintf("https://discord.com/api/v9/guilds/%s/members/%s", GuildID, UserID)
	timestamp := time.Now().UTC().AddDate(0, 0, Options.DaysToAdd).Add(time.Minute * time.Duration(Options.MinutesToAdd)).Format(time.RFC3339)
	data := map[string]interface{}{
		"communication_disabled_until": timestamp,
	}
	customRequest(b, "PATCH", endpoint, data, nil)
}

// Removes a timeout.
func (g *GuildManager) RemoveTimeout(b *Bot, GuildID, UserID string) {
	endpoint := fmt.Sprintf("https://discord.com/api/v9/guilds/%s/members/%s", GuildID, UserID)
	data := map[string]interface{}{
		"communication_disabled_until": nil,
	}
	customRequest(b, "PATCH", endpoint, data, nil)
}

/*
	For some reason, I can't get the GetGuild functions to work with customRequest.
	customRequest also refuses to work with URLs that return arrays.
	Oh well, at least it works!
	I also have to improve on error handling someday lol
*/

// Returns a GuildMember object. Not fully functional.
func (g *GuildManager) GetGuildMember(b *Bot, GuildID, UserID string) GuildMember {
	endpoint := fmt.Sprintf("https://discord.com/api/v9/guilds/%s/members/%s", GuildID, UserID)
	resp, err := b.Request(true, "GET", endpoint, nil, nil)
	if err != nil {
		return GuildMember{}
	}

	var member GuildMember
	decode(resp, &member)

	return member
}

// Returns a Guild object.
func (g *GuildManager) GetGuild(b *Bot, GuildID string) Guild {
	endpoint := fmt.Sprintf("https://discord.com/api/v9/guilds/%s", GuildID)
	resp, err := b.Request(true, "GET", endpoint, nil, nil)
	if err != nil {
		return Guild{}
	}

	var guild Guild
	decode(resp, &guild)

	return guild
}

// Returns an array of Channel objects.
func (g *GuildManager) GetGuildChannels(b *Bot, GuildID string) []Channel {
	endpoint := fmt.Sprintf("https://discord.com/api/v9/guilds/%s/channels", GuildID)
	resp, err := b.Request(true, "GET", endpoint, nil, nil)
	if err != nil {
		return []Channel{}
	}

	var channels []Channel
	decode(resp, &channels)

	return channels
}

// Returns an array of role IDs.
func (g *GuildManager) GetRolesForUser(b *Bot, GuildID, UserID string) []Role {
	endpoint := fmt.Sprintf("https://discord.com/api/v9/guilds/%s/members/%s", GuildID, UserID)
	resp, err := b.Request(true, "GET", endpoint, nil, nil)
	if err != nil {
		return nil
	}

	var member GuildMember
	err = json.NewDecoder(resp.Body).Decode(&member)
	if err != nil {
		return nil
	}

	roles := g.GetGuild(b, GuildID).Roles
	var userRoles []Role
	for _, roleID := range member.Roles {
		for _, role := range roles {
			if role.ID == roleID {
				userRoles = append(userRoles, role)
				break
			}
		}
	}

	return userRoles
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
	roles := g.GetRolesForUser(b, GuildID, UserID)
	for _, role := range roles {
		if role.ID == RoleID {
			return true
		}
	}
	return false
}
