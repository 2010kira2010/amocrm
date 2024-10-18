package amocrm

type ContactOne struct {
	ID                 int               `json:"id,omitempty"`
	Name               string            `json:"name,omitempty"`
	FirstName          string            `json:"first_name,omitempty"`
	LastName           string            `json:"last_name,omitempty"`
	Price              int               `json:"price,omitempty"`
	ResponsibleUserID  int               `json:"responsible_user_id,omitempty"`
	GroupID            int               `json:"group_id,omitempty"`
	StatusID           int               `json:"status_id,omitempty"`
	PipelineID         int               `json:"pipeline_id,omitempty"`
	LossReasonID       interface{}       `json:"loss_reason_id,omitempty"`
	CreatedBy          int               `json:"created_by,omitempty"`
	UpdatedBy          int               `json:"updated_by,omitempty"`
	CreatedAt          int               `json:"created_at,omitempty"`
	UpdatedAt          int               `json:"updated_at,omitempty"`
	ClosedAt           interface{}       `json:"closed_at,omitempty"`
	ClosestTaskAt      interface{}       `json:"closest_task_at,omitempty"`
	IsDeleted          bool              `json:"is_deleted,omitempty"`
	CustomFieldsValues []CustomsFields   `json:"custom_fields_values,omitempty"`
	Score              interface{}       `json:"score,omitempty"`
	AccountID          int               `json:"account_id,omitempty"`
	LaborCost          interface{}       `json:"labor_cost,omitempty"`
	Embedded           *ContactsEmbedded `json:"_embedded,omitempty"`
}

type ContactsArr struct {
	Page     int `json:"_page,omitempty"`
	Embedded struct {
		Contacts []struct {
			ID                 int               `json:"id,omitempty"`
			Name               string            `json:"name,omitempty"`
			FirstName          string            `json:"first_name,omitempty"`
			LastName           string            `json:"last_name,omitempty"`
			ResponsibleUserID  int               `json:"responsible_user_id,omitempty"`
			GroupID            int               `json:"group_id,omitempty"`
			CreatedBy          int               `json:"created_by,omitempty"`
			UpdatedBy          int               `json:"updated_by,omitempty"`
			CreatedAt          int               `json:"created_at,omitempty"`
			UpdatedAt          int               `json:"updated_at,omitempty"`
			ClosestTaskAt      interface{}       `json:"closest_task_at,omitempty"`
			IsDeleted          bool              `json:"is_deleted,omitempty"`
			IsUnsorted         bool              `json:"is_unsorted,omitempty"`
			CustomFieldsValues []CustomsFields   `json:"custom_fields_values,omitempty"`
			AccountID          int               `json:"account_id,omitempty"`
			Embedded           *ContactsEmbedded `json:"_embedded,omitempty"`
		} `json:"contacts,omitempty"`
	} `json:"_embedded,omitempty"`
}

type ContactsEmbedded struct {
	Tags []FieldValues `json:"tags,omitempty"`
}
