package instructure

type Student_Info_Instructure struct {
	ID              int         `json:"id"`
	Name            string      `json:"name"`
	CreatedAt       string      `json:"created_at"`
	SortableName    string      `json:"sortable_name"`
	ShortName       string      `json:"short_name"`
	AvatarURL       string      `json:"avatar_url"`
	Locale          interface{} `json:"locale"`
	EffectiveLocale string      `json:"effective_locale"`
	Permissions     struct {
		CanUpdateName           bool `json:"can_update_name"`
		CanUpdateAvatar         bool `json:"can_update_avatar"`
		LimitParentAppWebAccess bool `json:"limit_parent_app_web_access"`
	} `json:"permissions"`
}
