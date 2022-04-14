package gocord

const (
	MEMBERSHIP_STATE_INVITED  = 1
	MEMBERSHIP_STATE_ACCEPTED = 2
)

type TeamMember struct {
	membership_state int
	permissions      []string
	team_id          string
	user             User
}

type Team struct {
	icon          string
	id            string
	members       []TeamMember
	name          string
	owner_user_id string
}
