package amocrm

type Task struct {
	Id                int    `json:"id,omitempty"`
	CreatedBy         int    `json:"created_by,omitempty"`
	UpdatedBy         int    `json:"updated_by,omitempty"`
	CreatedAt         int    `json:"created_at,omitempty"`
	UpdatedAt         int    `json:"updated_at,omitempty"`
	ResponsibleUserId int    `json:"responsible_user_id,omitempty"`
	GroupId           int    `json:"group_id,omitempty"`
	EntityId          int    `json:"entity_id,omitempty"`
	EntityType        string `json:"entity_type,omitempty"`
	Duration          int    `json:"duration,omitempty"`
	IsCompleted       bool   `json:"is_completed,omitempty"`
	TaskTypeId        int    `json:"task_type_id,omitempty"`
	Text              string `json:"text,omitempty"`
	Result            struct {
		Text string `json:"text,omitempty"`
	} `json:"result,omitempty"`
	CompleteTill int `json:"complete_till,omitempty"`
	AccountId    int `json:"account_id,omitempty"`
	Links        struct {
		Self struct {
			Href string `json:"href,omitempty"`
		} `json:"self,omitempty"`
	} `json:"_links,omitempty"`
}

type Taskss struct {
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
		Tasks []struct {
			ID                int    `json:"id,omitempty"`
			CreatedBy         int    `json:"created_by,omitempty"`
			UpdatedBy         int    `json:"updated_by,omitempty"`
			CreatedAt         int    `json:"created_at,omitempty"`
			UpdatedAt         int    `json:"updated_at,omitempty"`
			ResponsibleUserID int    `json:"responsible_user_id,omitempty"`
			GroupID           int    `json:"group_id,omitempty"`
			EntityID          int    `json:"entity_id,omitempty"`
			EntityType        string `json:"entity_type,omitempty"`
			Duration          int    `json:"duration,omitempty"`
			IsCompleted       bool   `json:"is_completed,omitempty"`
			TaskTypeID        int    `json:"task_type_id,omitempty"`
			Text              string `json:"text,omitempty"`
			Result            []any  `json:"result,omitempty"`
			CompleteTill      int    `json:"complete_till,omitempty"`
			AccountID         int    `json:"account_id,omitempty"`
			Links             struct {
				Self struct {
					Href string `json:"href,omitempty"`
				} `json:"self,omitempty"`
			} `json:"_links,omitempty"`
		} `json:"tasks,omitempty"`
	} `json:"_embedded,omitempty"`
}
