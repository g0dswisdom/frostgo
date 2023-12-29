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
	MentionEveryone   bool              `json:"mention_everyone"`
	Mentions          []User            `json:"mentions"`
	Timestamp         string            `json:"timestamp"`
	MessageReference  *MessageReference `json:"message_reference,omitempty"`
	ReferencedMessage *Message          `json:"referenced_message,omitempty"`
	TTS               bool              `json:"tts"`
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

type AvatarDecorationData struct {
	Asset string `json:"asset"`
	SKUID string `json:"sku_id"`
}

type Friend struct {
	ID       string  `json:"id"`
	Type     int     `json:"type"`
	Nickname *string `json:"nickname,omitempty"`
	User     struct {
		ID                   string                `json:"id"`
		Username             string                `json:"username"`
		GlobalName           string                `json:"global_name"`
		Avatar               string                `json:"avatar"`
		AvatarDecorationData *AvatarDecorationData `json:"avatar_decoration_data,omitempty"`
		Discriminator        string                `json:"discriminator"`
		PublicFlags          int                   `json:"public_flags"`
	} `json:"user"`
	Since string `json:"since"`
}

// Discord guild

type Webhook struct {
	ID            string  `json:"id"`
	Type          int     `json:"type"`
	GuildID       string  `json:"guild_id,omitempty"`
	ChannelID     string  `json:"channel_id,omitempty"`
	User          User    `json:"user,omitempty"`
	Name          string  `json:"name,omitempty"`
	Avatar        string  `json:"avatar,omitempty"`
	Token         string  `json:"token,omitempty"`
	ApplicationID string  `json:"application_id,omitempty"`
	SourceGuild   Guild   `json:"source_guild,omitempty"`
	SourceChannel Channel `json:"source_channel,omitempty"`
	URL           string  `json:"url,omitempty"`
}

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

type RoleTags struct {
	BotID                 string `json:"bot_id,omitempty"`
	IntegrationID         string `json:"integration_id,omitempty"`
	PremiumSubscriber     *bool  `json:"premium_subscriber,omitempty"`
	SubscriptionListingID string `json:"subscription_listing_id,omitempty"`
	AvailableForPurchase  *bool  `json:"available_for_purchase,omitempty"`
	GuildConnections      *bool  `json:"guild_connections,omitempty"`
}

type Role struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Color        int       `json:"color"`
	Hoist        bool      `json:"hoist"`
	Icon         *string   `json:"icon,omitempty"`
	UnicodeEmoji *string   `json:"unicode_emoji,omitempty"`
	Position     int       `json:"position"`
	Permissions  string    `json:"permissions"`
	Managed      bool      `json:"managed"`
	Mentionable  bool      `json:"mentionable"`
	Tags         *RoleTags `json:"tags,omitempty"`
	Flags        int       `json:"flags"`
}

type Permission int64

const (
	CreateInstantInvite Permission = 1 << iota
	KickMembers
	BanMembers
	Administrator
	ManageChannels
	ManageGuild
	AddReactions
	ViewAuditLog
	PrioritySpeaker
	Stream
	ViewChannel
	SendMessages
	SendTTSMessage
	ManageMessages
	EmbedLinks
	AttachFiles
	ReadMessageHistory
	MentionEveryone
	UseExternalEmojis
	ViewGuildInsights
	Connect
	Speak
	MuteMembers
	DeafenMembers
	MoveMembers
	UseVAD
	ChangeNickname
	ManageNicknames
	ManageRoles
	ManageWebhooks
	ManageGuildExpressions
	UseApplicationCommands
	RequestToSpeak
	ManageEvents
	ManageThreads
	CreatePublicThreads
	CreatePrivateThreads
	UseExternalStickers
	SendMessagesInThreads
	UseEmbeddedActivities
	ModerateMembers
	ViewCreatorMonetizationAnalytics
	UseSoundboard
	CreateGuildExpressions
	CreateEvents
	UseExternalSounds
	SendVoiceMessages
)

type WelcomeScreen struct {
	Description     *string                `json:"description,omitempty"`
	WelcomeChannels []WelcomeScreenChannel `json:"welcome_channels,omitempty"`
}

type WelcomeScreenChannel struct {
	ChannelID   string  `json:"channel_id"`
	Description string  `json:"description"`
	EmojiID     *string `json:"emoji_id,omitempty"`
	EmojiName   *string `json:"emoji_name,omitempty"`
}

type Emoji struct {
	ID            *string  `json:"id,omitempty"`
	Name          *string  `json:"name,omitempty"`
	Roles         []string `json:"roles,omitempty"`
	User          *User    `json:"user,omitempty"`
	RequireColons bool     `json:"require_colons,omitempty"`
	Managed       bool     `json:"managed,omitempty"`
	Animated      bool     `json:"animated,omitempty"`
	Available     bool     `json:"available,omitempty"`
}

type Guild struct {
	ID                          string         `json:"id"`
	Name                        string         `json:"name"`
	Icon                        *string        `json:"icon,omitempty"`
	Splash                      *string        `json:"splash,omitempty"`
	DiscoverySplash             *string        `json:"discovery_splash,omitempty"`
	Owner                       *bool          `json:"owner,omitempty"`
	OwnerID                     string         `json:"owner_id"`
	Permissions                 *string        `json:"permissions,omitempty"`
	Region                      *string        `json:"region,omitempty"`
	AfkChannelID                *string        `json:"afk_channel_id,omitempty"`
	AfkTimeout                  int            `json:"afk_timeout"`
	WidgetEnabled               *bool          `json:"widget_enabled,omitempty"`
	WidgetChannelID             *string        `json:"widget_channel_id,omitempty"`
	VerificationLevel           int            `json:"verification_level"`
	DefaultMessageNotifications int            `json:"default_message_notifications"`
	ExplicitContentFilter       int            `json:"explicit_content_filter"`
	Roles                       []Role         `json:"roles"`
	Emojis                      []Emoji        `json:"emojis"`
	Features                    []string       `json:"features"`
	MfaLevel                    int            `json:"mfa_level"`
	ApplicationID               *string        `json:"application_id,omitempty"`
	SystemChannelID             *string        `json:"system_channel_id,omitempty"`
	SystemChannelFlags          int            `json:"system_channel_flags"`
	RulesChannelID              *string        `json:"rules_channel_id,omitempty"`
	MaxPresences                *int           `json:"max_presences,omitempty"`
	MaxMembers                  int            `json:"max_members"`
	VanityURLCode               *string        `json:"vanity_url_code,omitempty"`
	Description                 *string        `json:"description,omitempty"`
	Banner                      *string        `json:"banner,omitempty"`
	PremiumTier                 int            `json:"premium_tier"`
	PremiumSubscriptionCount    *int           `json:"premium_subscription_count,omitempty"`
	PreferredLocale             string         `json:"preferred_locale"`
	PublicUpdatesChannelID      *string        `json:"public_updates_channel_id,omitempty"`
	MaxVideoChannelUsers        *int           `json:"max_video_channel_users,omitempty"`
	MaxStageVideoChannelUsers   *int           `json:"max_stage_video_channel_users,omitempty"`
	ApproximateMemberCount      *int           `json:"approximate_member_count,omitempty"`
	ApproximatePresenceCount    *int           `json:"approximate_presence_count,omitempty"`
	WelcomeScreen               *WelcomeScreen `json:"welcome_screen,omitempty"`
	NSFWLevel                   int            `json:"nsfw_level"`
	Stickers                    []Sticker      `json:"stickers,omitempty"`
	PremiumProgressBarEnabled   *bool          `json:"premium_progress_bar_enabled,omitempty"`
	SafetyAlertsChannelID       *string        `json:"safety_alerts_channel_id,omitempty"`
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

const (
	Online       = "online"
	Idle         = "idle"
	DoNotDisturb = "dnd"
	Invisible    = "invisible"
)

type StatusOptions struct {
	Status  string
	Content string
}

type Sticker struct {
	ID string `json:"id"`
}

type StickerPack struct {
	ID             string    `json:"id"`
	Stickers       []Sticker `json:"stickers"`
	Name           string    `json:"name"`
	SKUID          string    `json:"sku_id"`
	CoverStickerID string    `json:"cover_sticker_id"`
}

type StickerPacks struct {
	UserID             string      `json:"user_id"`
	PackID             string      `json:"pack_id"`
	EntitlementID      string      `json:"entitlement_id"`
	HasAccess          bool        `json:"has_access"`
	PremiumTypeReqired int         `json:"premium_type_required"`
	StickerPack        StickerPack `json:"sticker_pack"`
}
