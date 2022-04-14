package gocord

type Webhook struct {
	id             string
	type_webhook   int
	guild_id       string
	channel_id     string
	user           User
	name           string
	avatar         string
	token          string
	application_id string
	url            string
	source_guild   Guild // <--------------- TODO: Add guild structure
	// source_channel Channel <--------------- TODO: Add channel structure
}
