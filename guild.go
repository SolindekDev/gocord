package gocord

import "time"

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
