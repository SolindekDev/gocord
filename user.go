package gocord

type Account struct {
	id   string
	name string
}

type User struct {
	id            string
	username      string
	discriminator string
	avatar        string
	bot           bool
	system        bool
	mfa_enabled   bool
	banner        string
	accent_color  int
	locale        string
	verified      bool
	email         string
	flags         int
	premium_type  int
	public_flags  int
}
