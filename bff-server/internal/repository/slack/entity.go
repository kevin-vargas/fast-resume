package slack

type Channel struct {
	ID                      string   `json:"id"`
	Name                    string   `json:"name"`
	IsChannel               bool     `json:"is_channel"`
	IsGroup                 bool     `json:"is_group"`
	IsIm                    bool     `json:"is_im"`
	IsMpim                  bool     `json:"is_mpim"`
	IsPrivate               bool     `json:"is_private"`
	Created                 int      `json:"created"`
	IsArchived              bool     `json:"is_archived"`
	IsGeneral               bool     `json:"is_general"`
	Unlinked                int      `json:"unlinked"`
	NameNormalized          string   `json:"name_normalized"`
	IsShared                bool     `json:"is_shared"`
	IsOrgShared             bool     `json:"is_org_shared"`
	IsPendingExtShared      bool     `json:"is_pending_ext_shared"`
	PendingShared           []any    `json:"pending_shared"`
	ContextTeamID           string   `json:"context_team_id"`
	Updated                 int64    `json:"updated"`
	ParentConversation      any      `json:"parent_conversation"`
	Creator                 string   `json:"creator"`
	IsExtShared             bool     `json:"is_ext_shared"`
	SharedTeamIds           []string `json:"shared_team_ids"`
	PendingConnectedTeamIds []any    `json:"pending_connected_team_ids"`
	Topic                   struct {
		Value   string `json:"value"`
		Creator string `json:"creator"`
		LastSet int    `json:"last_set"`
	} `json:"topic"`
	Purpose struct {
		Value   string `json:"value"`
		Creator string `json:"creator"`
		LastSet int    `json:"last_set"`
	} `json:"purpose"`
	PreviousNames []any `json:"previous_names"`
}

type Channels struct {
	Ok       bool `json:"ok"`
	Channels []Channel
	Err      string `json:"error"`
}

type Message struct {
	ClientMsgID string `json:"client_msg_id,omitempty"`
	Type        string `json:"type"`
	Text        string `json:"text"`
	User        string `json:"user"`
	Ts          string `json:"ts"`
	Blocks      []struct {
		Type     string `json:"type"`
		BlockID  string `json:"block_id"`
		Elements []struct {
			Type     string `json:"type"`
			Elements []struct {
				Type string `json:"type"`
				Text string `json:"text"`
			} `json:"elements"`
		} `json:"elements"`
	} `json:"blocks,omitempty"`
	Team    string `json:"team,omitempty"`
	Subtype string `json:"subtype,omitempty"`
}

type Conversations struct {
	Ok       bool      `json:"ok"`
	Messages []Message `json:"messages"`
	Err      string    `json:"error"`
}

type User struct {
	ID       string `json:"id"`
	TeamID   string `json:"team_id"`
	Name     string `json:"name"`
	Deleted  bool   `json:"deleted"`
	Color    string `json:"color"`
	RealName string `json:"real_name"`
	Tz       string `json:"tz"`
	TzLabel  string `json:"tz_label"`
	TzOffset int    `json:"tz_offset"`
	Profile  struct {
		Title                  string `json:"title"`
		Phone                  string `json:"phone"`
		Skype                  string `json:"skype"`
		RealName               string `json:"real_name"`
		RealNameNormalized     string `json:"real_name_normalized"`
		DisplayName            string `json:"display_name"`
		DisplayNameNormalized  string `json:"display_name_normalized"`
		Fields                 any    `json:"fields"`
		StatusText             string `json:"status_text"`
		StatusEmoji            string `json:"status_emoji"`
		StatusEmojiDisplayInfo []any  `json:"status_emoji_display_info"`
		StatusExpiration       int    `json:"status_expiration"`
		AvatarHash             string `json:"avatar_hash"`
		ImageOriginal          string `json:"image_original"`
		IsCustomImage          bool   `json:"is_custom_image"`
		FirstName              string `json:"first_name"`
		LastName               string `json:"last_name"`
		Image24                string `json:"image_24"`
		Image32                string `json:"image_32"`
		Image48                string `json:"image_48"`
		Image72                string `json:"image_72"`
		Image192               string `json:"image_192"`
		Image512               string `json:"image_512"`
		Image1024              string `json:"image_1024"`
		StatusTextCanonical    string `json:"status_text_canonical"`
		Team                   string `json:"team"`
	} `json:"profile"`
	IsAdmin                bool   `json:"is_admin"`
	IsOwner                bool   `json:"is_owner"`
	IsPrimaryOwner         bool   `json:"is_primary_owner"`
	IsRestricted           bool   `json:"is_restricted"`
	IsUltraRestricted      bool   `json:"is_ultra_restricted"`
	IsBot                  bool   `json:"is_bot"`
	IsAppUser              bool   `json:"is_app_user"`
	Updated                int    `json:"updated"`
	IsEmailConfirmed       bool   `json:"is_email_confirmed"`
	Has2Fa                 bool   `json:"has_2fa"`
	WhoCanShareContactCard string `json:"who_can_share_contact_card"`
}

type UserInfo struct {
	Ok   bool   `json:"ok"`
	User User   `json:"user"`
	Err  string `json:"error"`
}
