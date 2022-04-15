package gocord

import "time"

type Overwrite struct {
	id            string
	overwriteType int
	allow         string
	deny          string
}

type Channel struct {
	id                            string
	channelType                   int
	guild_id                      string
	position                      int
	permission_overwrites         Overwrite
	name                          string
	topic                         string
	nsfw                          bool
	last_message_id               string
	bitrate                       int
	user_limit                    int
	rate_limit_per_user           int
	recipients                    []User
	icon                          string
	owner_id                      string
	application_id                string
	parent_id                     string
	last_pin_timestamp            time.Time
	rtc_region                    string
	video_quality_mode            int
	message_count                 int
	member_count                  int
	default_auto_archive_duration int
	permissions                   string
	flags                         int
	// thread_metadata       ThreadMetadata   <------- TODO: Add thread metadata and thread member structure
	// member 				  ThreadMember    <-------  /\
}
