package amocrm

type Lead struct {
	ID                 int              `json:"id,omitempty"`
	Name               string           `json:"name,omitempty"`
	Price              int              `json:"price,omitempty"`
	ResponsibleUserID  int              `json:"responsible_user_id,omitempty"`
	GroupID            int              `json:"group_id,omitempty"`
	StatusID           int              `json:"status_id,omitempty"`
	PipelineID         int              `json:"pipeline_id,omitempty"`
	LossReasonID       interface{}      `json:"loss_reason_id,omitempty"`
	CreatedBy          int              `json:"created_by,omitempty"`
	UpdatedBy          int              `json:"updated_by,omitempty"`
	CreatedAt          int              `json:"created_at,omitempty"`
	UpdatedAt          int              `json:"updated_at,omitempty"`
	ClosedAt           interface{}      `json:"closed_at,omitempty"`
	ClosestTaskAt      interface{}      `json:"closest_task_at,omitempty"`
	IsDeleted          bool             `json:"is_deleted,omitempty"`
	CustomFieldsValues []*CustomsFields `json:"custom_fields_values,omitempty"`
	Score              interface{}      `json:"score,omitempty"`
	AccountID          int              `json:"account_id,omitempty"`
	LaborCost          interface{}      `json:"labor_cost,omitempty"`
	Embedded           *LeadEmbedded    `json:"_embedded,omitempty"`
}

type Leadss struct {
	Page  int `json:"_page,omitempty"`
	Links struct {
		Self struct {
			Href string `json:"href,omitempty"`
		} `json:"self,omitempty"`
	} `json:"_links,omitempty"`
	Embedded struct {
		Leads []*Lead `json:"leads,omitempty"`
	} `json:"_embedded,omitempty"`
}

type LeadEmbedded struct {
	Tags      []*FieldValues `json:"tags,omitempty"`
	Companies []*FieldValues `json:"companies,omitempty"`
	Contacts  []*FieldValues `json:"contacts,omitempty"`
}

type CustomsFields struct {
	FieldID   int                    `json:"field_id,omitempty"`
	FieldName string                 `json:"field_name,omitempty"`
	FieldCode interface{}            `json:"field_code,omitempty"`
	FieldType string                 `json:"field_type,omitempty"`
	Values    []*CustomsFieldsValues `json:"values,omitempty"`
}

type CustomsFieldsValues struct {
	Value    interface{} `json:"value,omitempty"`
	EnumId   interface{} `json:"enum_id,omitempty"`
	EnumCode interface{} `json:"enum_code,omitempty"`
}

type CustomsField struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Type      string `json:"type,omitempty"`
	AccountID int    `json:"account_id,omitempty"`
	Code      string `json:"code,omitempty"`
	Sort      int    `json:"sort,omitempty"`
	IsAPIOnly bool   `json:"is_api_only,omitempty"`
	Enums     []struct {
		ID    int    `json:"id,omitempty"`
		Value string `json:"value,omitempty"`
		Sort  int    `json:"sort,omitempty"`
	} `json:"enums,omitempty"`
}
