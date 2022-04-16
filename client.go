package gocord

type InstallParams struct {
	scopes      []string
	premissions string
}

type Application struct {
	id                     string
	name                   string
	icon                   string
	description            string
	rpc_origins            []string
	bot_public             bool
	bot_require_code_grant bool
	terms_of_service_url   string
	privacy_policy_url     string
	owner                  User
	verify_key             string
	team                   Team
	guild_id               string
	primary_sku_id         string
	slug                   string
	cover_image            string
	flags                  int
	tags                   []string
	install_params         InstallParams
	custom_install_url     string
}

type Client struct {
	token       string
	api_version string
	api_url     string
}

func NewClient(token string) Client {
	client := Client{token, "v9", "https://discord.com/api/"}

	return client
}
