package gocord

type Sticker struct {
	id           string
	pack_id      string
	name         string
	description  string
	tags         string
	asset        string
	sticker_type int
	format_type  int
	available    bool
	guild_id     string
	user         User
	sort_value   int
}
