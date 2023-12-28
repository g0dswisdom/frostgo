package FrostAPI

import (
	"fmt"
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

// Returns a GuildMember object. Not fully functional.
func (g *GuildManager) GetGuildMember(b *Bot, GuildID, UserID string) GuildMember {
	endpoint := fmt.Sprintf("https://discord.com/api/v9/guilds/%s/members/%s", GuildID, UserID)
	var member GuildMember
	decode(customRequest(b, "GET", endpoint, nil, nil), &member)
	return member
}
