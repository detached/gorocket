package api

type Channel struct {
	Id           string   `json:"_id"`
	Name         string   `json:"name"`
	MessageCount int      `json:"msgs"`
	UserNames    []string `json:"usernames"`
	UsersCount   int      `json:"usersCount"`

	User User `json:"u"`

	ReadOnly  bool   `json:"ro"`
	Timestamp string `json:"ts"`
	T         string `json:"t"`
	UpdatedAt string `json:"_updatedAt"`
	SysMes    bool   `json:"sysMes"`
}

type Group struct {
	Id           string   `json:"_id"`
	Name         string   `json:"name"`
	MessageCount int      `json:"msgs"`
	UserNames    []string `json:"usernames"`

	User User `json:"u"`

	ReadOnly  bool   `json:"ro"`
	Timestamp string `json:"ts"`
	T         string `json:"t"`
	UpdatedAt string `json:"_updatedAt"`
	SysMes    bool   `json:"sysMes"`
}

type User struct {
	Id        string `json:"_id,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	Bcrypt    string `json:"bcrypt,omitempty"`
	Name      string `json:"name"`
	UserName  string `json:"username,omitempty"`
	Emails    []struct {
		Address  string `json:"adress,omitempty"`
		Verified bool   `json:"verified,omitempty"`
	} `json:"emails,omitempty"`
	UpdatedAt             string            `json:"updatedAt,omitempty"`
	Type                  string            `json:"type,omitempty"`
	Status                string            `json:"status,omitempty"`
	Email                 string            `json:"email,omitempty"`
	Password              string            `json:"password"`
	Active                bool              `json:"active,omitempty"`
	Roles                 []string          `json:"roles,omitempty"`
	JoinDefaultChannels   bool              `json:"joinDefaultChannels,omitempty"`
	RequirePasswordChange bool              `json:"requirePasswordChange,omitempty"`
	SendWelcomeEmail      bool              `json:"sendWelcomeEmail,omitempty"`
	Verified              bool              `json:"verified,omitempty"`
	CustomFields          map[string]string `json:"customFields,omitempty"`
}

type UserCredentials struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"pass"`
}

type Message struct {
	Id        string `json:"_id"`
	ChannelId string `json:"rid"`
	Text      string `json:"msg"`
	Timestamp string `json:"ts"`
	User      User   `json:"u"`
}

type Info struct {
	Version string `json:"version"`

	Build struct {
		Date        string `json:"date"`
		NodeVersion string `json:"nodeVersion"`
		Arch        string `json:"arch"`
		Platform    string `json:"platform"`
		OsRelease   string `json:"osRelease"`
		TotalMemory int64  `json:"totalMemory"`
		FreeMemory  int64  `json:"freeMemory"`
		CpuCount    int    `json:"cpus"`
	} `json:"build"`

	Travis struct {
		BuildNumber string `json:"buildNumber"`
		Branch      string `json:"branch"`
		Tag         string `json:"tag"`
	} `json:"travis"`

	Commit struct {
		Hash    string `json:"hash"`
		Date    string `json:"date"`
		Author  string `json:"author"`
		Subject string `json:"subject"`
		Tag     string `json:"tag"`
		Branch  string `json:"branch"`
	} `json:"commit"`

	GraphicsMagick struct {
		Enabled bool `json:"enabled"`
	} `json:"GraphicsMagick"`

	ImageMagick struct {
		Enabled bool   `json:"enabled"`
		Version string `json:"version"`
	} `json:"ImageMagick"`
}

type Setting struct {
	Id    string `json:"_id"`
	Value string `json:"value"`
}

type UserPreferences struct {
	NewRoomNotification         string   `json:"newRoomNotification,omitempty"`
	NewMessageNotification      string   `json:"newMessageNotification,omitempty"`
	MuteFocusedConversations    bool     `json:"muteFocusedConversations,omitempty"`
	UseEmojis                   bool     `json:"useEmojis,omitempty"`
	ConvertAsciiEmoji           bool     `json:"convertAsciiEmoji,omitempty"`
	SaveMobileBandwidth         bool     `json:"saveMobileBandwidth,omitempty"`
	CollapseMediaByDefault      bool     `json:"collapseMediaByDefault,omitempty"`
	AutoImageLoad               bool     `json:"autoImageLoad,omitempty"`
	EmailNotificationMode       string   `json:"emailNotificationMode,omitempty"`
	RoomsListExhibitionMode     string   `json:"roomsListExhibitionMode,omitempty"`
	UnreadAlert                 bool     `json:"unreadAlert,omitempty"`
	NotificationsSoundVolume    int      `json:"notificationsSoundVolume,omitempty"`
	DesktopNotifications        string   `json:"desktopNotifications,omitempty"`
	MobileNotifications         string   `json:"mobileNotifications,omitempty"`
	EnableAutoAway              bool     `json:"enableAutoAway,omitempty"`
	Highlights                  []string `json:"highlights,omitempty"`
	DesktopNotificationDuration int      `json:"desktopNotificationDuration,omitempty"`
	ViewMode                    int      `json:"viewMode,omitempty"`
	HideUsernames               bool     `json:"hideUsernames,omitempty"`
	HideRoles                   bool     `json:"hideRoles,omitempty"`
	HideAvatars                 bool     `json:"hideAvatars,omitempty"`
	HideFlexTab                 bool     `json:"hideFlexTab,omitempty"`
	SendOnEnter                 string   `json:"sendOnEnter,omitempty"`
	RoomCounterSidebar          bool     `json:"roomCounterSidebar,omitempty"`
}
