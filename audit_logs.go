package gocord

// Audit log events
const (
	// Guild
	GUILD_UPDATE = 1

	// Channel
	CHANNEL_CREATE           = 10
	CHANNEL_UPDATE           = 11
	CHANNEL_DELETE           = 12
	CHANNEL_OVERWRITE_CREATE = 13
	CHANNEL_OVERWRITE_UPDATE = 14
	CHANNEL_OVERWRITE_DELETE = 15

	// Member
	MEMBER_KICK        = 20
	MEMBER_PRUNE       = 21
	MEMBER_BAN_ADD     = 22
	MEMBER_BAN_REMOVE  = 23
	MEMBER_UPDATE      = 24
	MEMBER_ROLE_UPDATE = 25
	MEMBER_MOVE        = 26
	MEMBER_DISCONNECT  = 27

	// Bot
	BOT_ADD = 28

	// Roles
	ROLE_CREATE = 30
	ROLE_UPDATE = 31
	ROLE_DELETE = 32

	// Invite
	INVITE_CREATE = 40
	INVITE_UPDATE = 41
	INVITE_DELETE = 42

	// Webhook
	WEBHOOK_CREATE = 50
	WEBHOOK_UPDATE = 51
	WEBHOOK_DELETE = 52

	// Emoji
	EMOJI_CREATE = 60
	EMOJI_UPDATE = 61
	EMOJI_DELETE = 62

	// Message
	MESSAGE_DELETE      = 72
	MESSAGE_BULK_DELETE = 73
	MESSAGE_PIN         = 74
	MESSAGE_UNPIN       = 75

	// Integrations
	INTEGRATION_CREATE = 80
	INTEGRATION_UPDATE = 81
	INTEGRATION_DELETE = 82

	// Stage
	STAGE_INSTANCE_CREATE = 83
	STAGE_INSTANCE_UPDATE = 84
	STAGE_INSTANCE_DELETE = 85

	// Sticker
	STICKER_CREATE = 90
	STICKER_UPDATE = 91
	STICKER_DELETE = 92

	// Guild scheduled
	GUILD_SCHEDULED_EVENT_CREATE = 100
	GUILD_SCHEDULED_EVENT_UPDATE = 101
	GUILD_SCHEDULED_EVENT_DELETE = 102

	// Thread
	THREAD_CREATE = 110
	THREAD_UPDATE = 111
	THREAD_DELETE = 112
)

type Mixed struct {
	// TODO: https://discord.com/developers/docs/resources/audit-log#audit-log-change-object-audit-log-change-key
}

type AuditLogChange struct {
	new_value Mixed
	old_value Mixed
	key       string
}

type AuditLogEntry struct {
	target_id   string
	changes     []AuditLogChange
	user_id     string
	id          string
	action_type int
	options     string
	reason      string // 0-512 characters
}

type AuditLog struct {
	audit_log_entries      []AuditLogEntry
	guild_scheduled_events []GuildScheduled
	integrations           []Integration
	users                  []User
	// threads []Thread                     <--------------------- TODO: Add thread structure
	// webhooks []Webhook                   <--------------------- TODO: Add webhook structure
}
