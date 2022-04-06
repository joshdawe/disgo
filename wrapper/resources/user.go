package resources

// User Object
// https://discord.com/developers/docs/resources/user#user-object
type User struct {
	ID            Snowflake `json:"id,omitempty"`
	Username      string    `json:"username,omitempty"`
	Discriminator string    `json:"discriminator,omitempty"`
	Avatar        string    `json:"avatar,omitempty"`
	Bot           bool      `json:"bot,omitempty"`
	System        bool      `json:"system,omitempty"`
	MFAEnabled    bool      `json:"mfa_enabled,omitempty"`
	Banner        string    `json:"banner,omitempty"`
	AccentColor   int       `json:"accent_color,omitempty"`
	Locale        string    `json:"locale,omitempty"`
	Verified      bool      `json:"verified,omitempty"`
	Email         *string   `json:"email,omitempty"`
	Flags         *BitFlag  `json:"flag,omitempty"`
	PremiumType   *Flag     `json:"premium_type,omitempty"`
	PublicFlags   BitFlag   `json:"public_flag,omitempty"`
}

// User Flags
// https://discord.com/developers/docs/resources/user#user-object-user-flags
const (
	FlagFlagsUserNONE                         = 0
	FlagFlagsUserSTAFF                        = 1 << 0
	FlagFlagsUserPARTNER                      = 1 << 1
	FlagFlagsUserHYPESQUAD                    = 1 << 2
	FlagFlagsUserBUG_HUNTER_LEVEL_1           = 1 << 3
	FlagFlagsUserHYPESQUAD_ONLINE_HOUSE_ONE   = 1 << 6
	FlagFlagsUserHYPESQUAD_ONLINE_HOUSE_TWO   = 1 << 7
	FlagFlagsUserHYPESQUAD_ONLINE_HOUSE_THREE = 1 << 8
	FlagFlagsUserPREMIUM_EARLY_SUPPORTER      = 1 << 9
	FlagFlagsUserTEAM_PSEUDO_USER             = 1 << 10
	FlagFlagsUserBUG_HUNTER_LEVEL_2           = 1 << 14
	FlagFlagsUserVERIFIED_BOT                 = 1 << 16
	FlagFlagsUserVERIFIED_DEVELOPER           = 1 << 17
	FlagFlagsUserCERTIFIED_MODERATOR          = 1 << 18
	FlagFlagsUserBOT_HTTP_INTERACTIONS        = 1 << 19
)

// Premium Types
// https://discord.com/developers/docs/resources/user#user-object-premium-types
const (
	FlagTypesPremiumNONE         = 0
	FlagTypesPremiumNITROCLASSIC = 1
	FlagTypesPremiumNITRO        = 2
)

// User Connection Object
// https://discord.com/developers/docs/resources/user#connection-object-connection-structure
type Connection struct {
	ID           string         `json:"id,omitempty"`
	Name         string         `json:"name,omitempty"`
	Type         string         `json:"type,omitempty"`
	Revoked      bool           `json:"revoked,omitempty"`
	Integrations []*Integration `json:"integrations,omitempty"`
	Verified     bool           `json:"verified,omitempty"`
	FriendSync   bool           `json:"friend_sync,omitempty"`
	ShowActivity bool           `json:"show_activity,omitempty"`
	Visibility   Flag           `json:"visibility,omitempty"`
}

// Visibility Types
// https://discord.com/developers/docs/resources/user#connection-object-visibility-types
const (
	FlagTypesVisibilityNONE     = 0
	FlagTypesVisibilityEVERYONE = 1
)
