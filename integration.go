package gocord

import "time"

type IntegrationApplication struct {
	id          string
	name        string
	icon        string
	description string
	bot         User
}

type Integration struct {
	id                  string
	name                string
	type_integration    string
	enabled             bool
	syncing             bool
	role_id             string
	enable_emoticons    bool
	expire_behavior     int // 0 - Remove role / 1 - Kick
	expire_grace_period int
	user                User
	account             Account
	synced_at           time.Time
	subscriber_count    int
	revoked             bool
	application         IntegrationApplication
}
