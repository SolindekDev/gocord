package gocord

import "time"

type Voice struct {
	guild_id                   string
	channel_id                 string
	user_id                    string
	member                     User
	session_id                 string
	deaf                       bool
	mute                       bool
	self_deaf                  bool
	self_mute                  bool
	self_stream                bool
	self_video                 bool
	suppress                   bool
	request_to_speak_timestamp time.Time
}
