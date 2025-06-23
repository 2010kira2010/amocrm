package amocrm

type Catalog struct {
	ID              int    `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	CreatedBy       int    `json:"created_by,omitempty"`
	UpdatedBy       int    `json:"updated_by,omitempty"`
	CreatedAt       int    `json:"created_at,omitempty"`
	UpdatedAt       int    `json:"updated_at,omitempty"`
	Sort            int    `json:"sort,omitempty"`
	Type            string `json:"type,omitempty"`
	CanAddElements  bool   `json:"can_add_elements,omitempty"`
	CanShowInCards  bool   `json:"can_show_in_cards,omitempty"`
	CanLinkMultiple bool   `json:"can_link_multiple,omitempty"`
	CanBeDeleted    bool   `json:"can_be_deleted,omitempty"`
	SdkWidgetCode   any    `json:"sdk_widget_code,omitempty"`
	AccountID       int    `json:"account_id,omitempty"`
	Links           struct {
		Self struct {
			Href string `json:"href,omitempty"`
		} `json:"self,omitempty"`
	} `json:"_links,omitempty"`
}

type Catalogss struct {
	Page  int `json:"_page,omitempty"`
	Links struct {
		Self struct {
			Href string `json:"href,omitempty"`
		} `json:"self,omitempty"`
		Next struct {
			Href string `json:"href,omitempty"`
		} `json:"next,omitempty"`
	} `json:"_links,omitempty"`
	Embedded struct {
		Catalogs []struct {
			ID              int    `json:"id,omitempty"`
			Name            string `json:"name,omitempty"`
			CreatedBy       int    `json:"created_by,omitempty"`
			UpdatedBy       int    `json:"updated_by,omitempty"`
			CreatedAt       int    `json:"created_at,omitempty"`
			UpdatedAt       int    `json:"updated_at,omitempty"`
			Sort            int    `json:"sort,omitempty"`
			Type            string `json:"type,omitempty"`
			CanAddElements  bool   `json:"can_add_elements,omitempty"`
			CanShowInCards  bool   `json:"can_show_in_cards,omitempty"`
			CanLinkMultiple bool   `json:"can_link_multiple,omitempty"`
			CanBeDeleted    bool   `json:"can_be_deleted,omitempty"`
			SdkWidgetCode   any    `json:"sdk_widget_code,omitempty"`
			AccountID       int    `json:"account_id,omitempty"`
			Links           struct {
				Self struct {
					Href string `json:"href,omitempty"`
				} `json:"self,omitempty"`
			} `json:"_links,omitempty"`
		} `json:"catalogs,omitempty"`
	} `json:"_embedded,omitempty"`
}
