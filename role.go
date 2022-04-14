package gocord

type RoleTags struct {
	bot_id         string
	integration_id string
}

type Role struct {
	id            string
	name          string
	color         int
	hoist         bool
	icon          string
	unicode_emoji string
	position      int
	permissions   string
	managed       bool
	mentionable   bool
	tags          RoleTags
}
