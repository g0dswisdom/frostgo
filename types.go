package FrostAPI

import "encoding/json"

// Websocket

type DiscordMessage struct {
	Op int             `json:"op"`
	D  json.RawMessage `json:"d"`
	T  string          `json:"t"`
}

type HeartbeatPayloadData struct {
	HeartbeatInterval int `json:"heartbeat_interval"`
}

// Discord Message

type MessageReference struct {
	MessageID string `json:"message_id"`
	ChannelID string `json:"channel_id"`
}

type Message struct {
	ID        string `json:"id"`
	Type      int    `json:"type"`
	Content   string `json:"content"`
	ChannelID string `json:"channel_id"`
	GuildID   string `json:"guild_id,omitempty"`
	Author    struct {
		ID            string `json:"id"`
		Username      string `json:"username"`
		Avatar        string `json:"avatar"`
		Discriminator string `json:"discriminator"`
	} `json:"author"`
	Timestamp         string            `json:"timestamp"`
	MessageReference  *MessageReference `json:"message_reference,omitempty"`
	ReferencedMessage *Message          `json:"referenced_message,omitempty"`
}

// Discord user

type User struct {
	ID               string  `json:"id"`
	Username         string  `json:"username"`
	Discriminator    string  `json:"discriminator"`
	GlobalName       *string `json:"global_name,omitempty"`
	Avatar           *string `json:"avatar,omitempty"`
	Bot              *bool   `json:"bot,omitempty"`
	System           *bool   `json:"system,omitempty"`
	MFAEnabled       *bool   `json:"mfa_enabled,omitempty"`
	Banner           *string `json:"banner,omitempty"`
	AccentColor      *int    `json:"accent_color,omitempty"`
	Locale           string  `json:"locale"`
	Verified         *bool   `json:"verified,omitempty"`
	Email            *string `json:"email,omitempty"`
	Flags            *int    `json:"flags,omitempty"`
	PremiumType      *int    `json:"premium_type,omitempty"`
	PublicFlags      *int    `json:"public_flags,omitempty"`
	AvatarDecoration *string `json:"avatar_decoration,omitempty"`
}

// Discord guild

type GuildInvite struct {
	Type      int     `json:"type"`
	Code      string  `json:"code"`
	Inviter   User    `json:"inviter"`
	MaxAge    int     `json:"max_age"`
	CreatedAt string  `json:"created_at"`
	ExpiresAt string  `json:"expires_at"`
	Guild     Guild   `json:"guild"`
	GuildID   string  `json:"guild_id"`
	Channel   Channel `json:"channel"`
	Uses      int     `json:"uses"`
	MaxUses   int     `json:"max_uses"`
	Temporary bool    `json:"temporary"`
}

type GuildInviteOptions struct {
	MaxAge  int `json:"max_age"`
	MaxUses int `json:"max_uses"`
}

type Guild struct {
	ID                       string   `json:"id"`
	Name                     string   `json:"name"`
	Splash                   *string  `json:"splash,omitempty"`
	Banner                   *string  `json:"banner,omitempty"`
	Description              *string  `json:"description,omitempty"`
	Icon                     *string  `json:"icon,omitempty"`
	Features                 []string `json:"features,omitempty"`
	VerificationLevel        int      `json:"verification_level"`
	VanityURLCode            *string  `json:"vanity_url_code,omitempty"`
	NSFWLevel                int      `json:"nsfw_level"`
	NSFW                     bool     `json:"nsfw"`
	PremiumSubscriptionCount int      `json:"premium_subscription_count"`
}

type GuildMember struct {
	User                       *User    `json:"user,omitempty"`
	Nick                       *string  `json:"nick,omitempty"`
	Avatar                     *string  `json:"avatar,omitempty"`
	Roles                      []string `json:"roles"`
	JoinedAt                   string   `json:"joined_at"`
	PremiumSince               *string  `json:"premium_since,omitempty"`
	Deaf                       bool     `json:"deaf"`
	Mute                       bool     `json:"mute"`
	Flags                      int      `json:"flags"`
	Pending                    *bool    `json:"pending,omitempty"`
	Permissions                *string  `json:"permissions,omitempty"`
	CommunicationDisabledUntil *string  `json:"communication_disabled_until,omitempty"`
}

type ThreadMetadata struct {
	Archived            bool   `json:"archived"`
	AutoArchiveDuration int    `json:"auto_archive_duration"`
	ArchiveTimestamp    string `json:"archive_timestamp"`
	Locked              bool   `json:"locked"`
	Invitable           bool   `json:"invitable,omitempty"`
	CreateTimestamp     string `json:"create_timestamp,omitempty"`
}

type ThreadMember struct {
	ID            *string      `json:"id,omitempty"`
	UserID        *string      `json:"user_id,omitempty"`
	JoinTimestamp string       `json:"join_timestamp"`
	Flags         int          `json:"flags"`
	Member        *GuildMember `json:"member,omitempty"`
}

type Overwrite struct {
	ID    string `json:"id"`
	Type  int    `json:"type"`
	Allow string `json:"allow"`
	Deny  string `json:"deny"`
}

type Tag struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Moderated bool    `json:"moderated"`
	EmojiID   *string `json:"emoji_id,omitempty"`
	EmojiName *string `json:"emoji_name,omitempty"`
}

type DefaultReaction struct {
	EmojiID   *string `json:"emoji_id,omitempty"`
	EmojiName *string `json:"emoji_name,omitempty"`
}

type Channel struct {
	ID                     string `json:"id"`
	Type                   int    `json:"type"`
	GuildID                string `json:"guild_id,omitempty"`
	Position               int    `json:"position,omitempty"`
	PermissionOverwrites   []Overwrite
	Name                   string `json:"name,omitempty"`
	Topic                  string `json:"topic,omitempty"`
	NSFW                   bool   `json:"nsfw,omitempty"`
	LastMessageID          string `json:"last_message_id,omitempty"`
	Bitrate                int    `json:"bitrate,omitempty"`
	UserLimit              int    `json:"user_limit,omitempty"`
	RateLimitPerUser       int    `json:"rate_limit_per_user,omitempty"`
	Recipients             []User
	Icon                   string `json:"icon,omitempty"`
	OwnerID                string `json:"owner_id,omitempty"`
	ApplicationID          string `json:"application_id,omitempty"`
	Managed                bool   `json:"managed,omitempty"`
	ParentID               string `json:"parent_id,omitempty"`
	LastPinTimestamp       string `json:"last_pin_timestamp,omitempty"`
	RTCRegion              string `json:"rtc_region,omitempty"`
	VideoQualityMode       int    `json:"video_quality_mode,omitempty"`
	MessageCount           int    `json:"message_count,omitempty"`
	MemberCount            int    `json:"member_count,omitempty"`
	ThreadMetadata         ThreadMetadata
	Member                 ThreadMember
	DefaultAutoArchiveDur  int    `json:"default_auto_archive_duration,omitempty"`
	Permissions            string `json:"permissions,omitempty"`
	Flags                  int    `json:"flags,omitempty"`
	TotalMessagesSent      int    `json:"total_message_sent,omitempty"`
	AvailableTags          []Tag
	AppliedTags            []string `json:"applied_tags,omitempty"`
	DefaultReactionEmoji   DefaultReaction
	DefaultThreadRateLimit int `json:"default_thread_rate_limit_per_user,omitempty"`
	DefaultSortOrder       int `json:"default_sort_order,omitempty"`
	DefaultForumLayout     int `json:"default_forum_layout,omitempty"`
}

type TimeoutOptions struct {
	DaysToAdd    int
	MinutesToAdd int
}
