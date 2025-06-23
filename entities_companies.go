package amocrm

type Company struct {
	ID                 int                `json:"id,omitempty"`
	Name               string             `json:"name,omitempty"`
	ResponsibleUserID  int                `json:"responsible_user_id,omitempty"`
	GroupID            int                `json:"group_id,omitempty"`
	CreatedBy          int                `json:"created_by,omitempty"`
	UpdatedBy          int                `json:"updated_by,omitempty"`
	CreatedAt          int                `json:"created_at,omitempty"`
	UpdatedAt          int                `json:"updated_at,omitempty"`
	ClosedAt           interface{}        `json:"closed_at,omitempty"`
	ClosestTaskAt      interface{}        `json:"closest_task_at,omitempty"`
	IsDeleted          bool               `json:"is_deleted,omitempty"`
	CustomFieldsValues []*CustomsFields   `json:"custom_fields_values,omitempty"`
	AccountID          int                `json:"account_id,omitempty"`
	Embedded           *CompaniesEmbedded `json:"_embedded,omitempty"`
}

type Companiess struct {
	Page     int `json:"_page,omitempty"`
	Embedded struct {
		Companies []*Company `json:"companies,omitempty"`
	} `json:"_embedded,omitempty"`
}

type CompaniesEmbedded struct {
	Tags []*FieldValues `json:"tags,omitempty"`
}
