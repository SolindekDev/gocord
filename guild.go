package gocord

import (
	"time"
)

type MemberGuild struct {
	user                         User
	nick                         string
	avatar                       string
	roles                        []string
	joined_at                    time.Time
	premium_since                time.Time
	deaf                         bool
	mute                         bool
	pending                      bool
	permissions                  string
	communication_disabled_until time.Time
}

type WelcomeScreenChannel struct {
	channel_id  string
	description string
	emoji_id    string
	emoji_name  string
}

type WelcomeScreen struct {
	description      string
	welcome_channels []WelcomeScreenChannel
}

type StageInstance struct {
	id                       string
	guild_id                 string
	channel_id               string
	topic                    string
	privacy_level            int
	discoverable_disabled    bool
	guild_scheduled_event_id string
}

type Guild struct {
	id                            string
	name                          string
	icon                          string
	icon_hash                     string
	splash                        string
	discovery_splash              string
	owner                         bool
	owner_id                      string
	permissions                   string
	region                        string
	afk_channel_id                string
	afk_timeout                   int
	widget_enabled                bool
	widget_channel_id             string
	verification_level            int
	default_message_notifications int
	explicit_content_filter       int
	roles                         []Role
	emojis                        []Emoji
	features                      []string // <- TODO: Add guild features strings https://discord.com/developers/docs/resources/guild#guild-object
	mfa_level                     int
	application_id                string
	system_channel_id             string
	system_channel_flags          int
	rules_channel_id              string
	joined_at                     time.Time
	large                         bool
	unavailable                   bool
	member_count                  int
	voice_states                  Voice
	members                       []MemberGuild
	channels                      []Channel
	threads                       []Channel
	max_presences                 int
	max_members                   int
	vanity_url_code               string
	description                   string
	banner                        string
	premium_tier                  int
	premium_subscription_count    int
	preferred_locale              string
	public_updates_channel_id     string
	max_video_channel_users       int
	approximate_member_count      int
	approximate_presence_count    int
	welcome_screen                WelcomeScreen
	nsfw_level                    int
	stage_instances               StageInstance
	stickers                      Sticker
	guild_scheduled_events        GuildScheduled
	premium_progress_bar_enabled  bool
	// presences []PresenceUpdate
}

const (
	GUILD_ONLY = iota + 2
)

const (
	SCHEDULED = iota + 1
	ACTIVE
	COMPLETED
	CANCELED
)

const (
	STAGE_INSTANCE = iota + 1
	VOICE
	EXTERNAL
)

type GuildScheduledEntityMetadata struct {
	location string // 0-100 characters
}

type GuildScheduled struct {
	id                   string
	guild_id             string
	channel_id           string
	creator_id           string
	name                 string
	description          string
	scheduled_start_time time.Time
	scheduled_end_time   time.Time
	privacy_level        int
	status               int
	entity_type          int
	entity_id            string
	entity_metadata      GuildScheduledEntityMetadata
	creator              User
	user_count           int
	image                string
}
