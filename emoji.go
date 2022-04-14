package gocord

type Emoji struct {
	id             string
	name           string
	roles          []Role
	user           User
	require_colons bool
	managed        bool
	animated       bool
	available      bool
}
