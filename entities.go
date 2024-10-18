package amocrm

type FieldValues map[string]interface{}

// Account represents amoCRM Account entity json DTO.
type Account struct {
	ID                      int    `json:"id,omitempty"`
	Name                    string `json:"name,omitempty"`
	Subdomain               string `json:"subdomain,omitempty"`
	CreatedAt               int    `json:"created_at,omitempty"`
	CreatedBy               int    `json:"created_by,omitempty"`
	UpdatedAt               int    `json:"updated_at,omitempty"`
	UpdatedBy               int    `json:"updated_by,omitempty"`
	CurrentUserID           int    `json:"current_user_id,omitempty"`
	Country                 string `json:"country,omitempty"`
	Currency                string `json:"currency,omitempty"`
	CustomersMode           string `json:"customers_mode,omitempty"`
	IsUnsortedOn            bool   `json:"is_unsorted_on,omitempty"`
	MobileFeatureVersion    int    `json:"mobile_feature_version,omitempty"`
	IsLossReasonEnabled     bool   `json:"is_loss_reason_enabled,omitempty"`
	IsHelpbotEnabled        bool   `json:"is_helpbot_enabled,omitempty"`
	IsTechnicalAccount      bool   `json:"is_technical_account,omitempty"`
	ContactNameDisplayOrder int    `json:"contact_name_display_order,omitempty"`
	AmojoID                 string `json:"amojo_id,omitempty"`
	UUID                    string `json:"uuid,omitempty"`
	Version                 int    `json:"version,omitempty"`
	Links                   struct {
		Self struct {
			Href string `json:"href,omitempty"`
		} `json:"self,omitempty"`
	} `json:"_links,omitempty"`
	Embedded struct {
		AmojoRights struct {
			CanDirect       bool `json:"can_direct,omitempty"`
			CanCreateGroups bool `json:"can_create_groups,omitempty"`
		} `json:"amojo_rights,omitempty"`
		UsersGroups []struct {
			ID   int         `json:"id,omitempty"`
			Name string      `json:"name,omitempty"`
			UUID interface{} `json:"uuid,omitempty"`
		} `json:"users_groups,omitempty"`
		TaskTypes []struct {
			ID     int         `json:"id,omitempty"`
			Name   string      `json:"name,omitempty"`
			Color  interface{} `json:"color,omitempty"`
			IconID interface{} `json:"icon_id,omitempty"`
			Code   string      `json:"code,omitempty"`
		} `json:"task_types,omitempty"`
		DatetimeSettings struct {
			DatePattern      string `json:"date_pattern,omitempty"`
			ShortDatePattern string `json:"short_date_pattern,omitempty"`
			ShortTimePattern string `json:"short_time_pattern,omitempty"`
			DateFormat       string `json:"date_format,omitempty"`
			TimeFormat       string `json:"time_format,omitempty"`
			Timezone         string `json:"timezone,omitempty"`
			TimezoneOffset   string `json:"timezone_offset,omitempty"`
		} `json:"datetime_settings,omitempty"`
	} `json:"_embedded,omitempty"`
}
