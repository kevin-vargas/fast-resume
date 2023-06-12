package entity

type SlackToken struct {
	Ok                  bool            `json:"ok"`
	AppID               string          `json:"app_id"`
	AuthedUser          SlackAuthedUser `json:"authed_user"`
	Team                SlackTeam       `json:"team"`
	Enterprise          any             `json:"enterprise"`
	IsEnterpriseInstall bool            `json:"is_enterprise_install"`
	Err                 string          `json:"error"`
}

type SlackAuthedUser struct {
	ID          string `json:"id"`
	Scope       string `json:"scope"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

type SlackTeam struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Claim struct {
	Token string `json:"token,omitempty"`
}

type Jwks struct {
	Keys []JwksKey `json:"keys"`
}
type JwksKey struct {
	N   string `json:"n"`
	E   string `json:"e"`
	Alg string `json:"alg"`
	Use string `json:"use"`
	Kid string `json:"kid"`
	Kty string `json:"kty"`
}
