package gethls

type Room struct {
	FanClubIsMember             bool     `json:"fan_club_is_member"`
	NeedsSupporterToPm          bool     `json:"needs_supporter_to_pm"`
	TokenBalance                int      `json:"token_balance"`
	SourceName                  string   `json:"source_name"`
	OfflineRoomChatPmEnabled    bool     `json:"offline_room_chat_pm_enabled"`
	SpyPrivateShowPrice         int      `json:"spy_private_show_price"`
	ChatSettings                Chat     `json:"chat_settings"`
	HideSatisfactionScore       bool     `json:"hide_satisfaction_score"`
	IsAgeVerified               bool     `json:"is_age_verified"`
	FollowNotificationFrequency string   `json:"follow_notification_frequency"`
	BroadcasterUsername         string   `json:"broadcaster_username"`
	HasStudio                   bool     `json:"has_studio"`
	TipsInPast24Hours           int      `json:"tips_in_past_24_hours"`
	BroadcasterUID              string   `json:"broadcaster_uid"`
	ShowMobileSiteBannerLink    bool     `json:"show_mobile_site_banner_link"`
	LastVoteInPast90Days        bool     `json:"last_vote_in_past_90_days_down"`
	ServerName                  string   `json:"server_name"`
	PrivateMinMinutes           int      `json:"private_min_minutes"`
	ChatUsername                string   `json:"chat_username"`
	IsTestbed                   bool     `json:"is_testbed"`
	ViewerGender                string   `json:"viewer_gender"`
	IsSupporter                 bool     `json:"is_supporter"`
	AllowAnonymousTipping       bool     `json:"allow_anonymous_tipping"`
	NumViewers                  int      `json:"num_viewers"`
	BroadcasterGender           string   `json:"broadcaster_gender"`
	HlsSource                   string   `json:"hls_source"`
	AllowShowRecordings         bool     `json:"allow_show_recordings"`
	BrowserId                   string   `json:"browser_id"`
	PrivateShowId               string   `json:"private_show_id"`
	IsWidescreen                bool     `json:"is_widescreen"`
	IsModerator                 bool     `json:"is_moderator"`
	RoomStatus                  string   `json:"room_status"`
	RoomUid                     string   `json:"room_uid"`
	BroadcasterOnNewChat        bool     `json:"broadcaster_on_new_chat"`
	EdgeAuth                    string   `json:"edge_auth"`
	PrivateShowPrice            int      `json:"private_show_price"`
	NumFollowedOnline           int      `json:"num_followed_online"`
	FanClubPaidWithTokens       bool     `json:"fan_club_paid_with_tokens"`
	AspAuthUrl                  string   `json:"asp_auth_url"`
	IsMobile                    bool     `json:"is_mobile"`
	OptOut                      bool     `json:"opt_out"`
	ChatPassword                string   `json:"chat_password"`
	RoomPass                    string   `json:"room_pass"`
	NumFollowed                 int      `json:"num_followed"`
	LowSatisfactionScore        bool     `json:"low_satisfaction_score"`
	ChatRules                   string   `json:"chat_rules"`
	Age                         int      `json:"age"`
	RoomTitle                   string   `json:"room_title"`
	SatisfactionScore           SatScore `json:"satisfaction_score"`
	ViewerUsername              string   `json:"viewer_username"`
	HidenMessage                string   `json:"hidden_message"`
	Following                   bool     `json:"following"`
	BroadcasterEnable_asp       bool     `json:"broadcaster_enable_asp"`
	WschatHost                  string   `json:"wschat_host"`
	RecommenderHmac             string   `json:"recommender_hmac"`
	ExploringHashtag            string   `json:"exploring_hashtag"`
	TfaEnabled                  bool     `json:"tfa_enabled"`
	AllowPrivateShows           bool     `json:"allow_private_shows"`
}

type Chat struct {
	SortUsersKey              string `json:"sort_users_key"`
	SilenceBroadcasters       string `json:"silence_broadcasters"`
	HighestTokenColor         string `json:"highest_token_color"`
	EmoticonAutocompleteDelay string `json:"emoticon_autocomplete_delay"`
	IgnoredUsers              string `json:"ignored_users"`
	ShowEmoticons             bool   `json:"show_emoticons"`
	FontSize                  string `json:"font_size"`
	BTipVol                   string `json:"b_tip_vol"`
	AllowedChat               string `json:"allowed_chat"`
	RoomLeaveFOr              string `json:"room_leave_for"`
	ModExpire                 int    `json:"mod_expire"`
	MaxPmAge                  int    `json:"max_pm_age"`
	FontColor                 string `json:"font_color"`
	FontFamily                string `json:"font_family"`
	RoomEntryFor              string `json:"room_entry_for"`
	VTipVol                   string `json:"v_tip_vol"`
}

type SatScore struct {
	DownVotes int `json:"down_votes"`
	UpVotes   int `json:"up_votes"`
	Percent   int `json:"percent"`
	Max       int `json:"max"`
}
