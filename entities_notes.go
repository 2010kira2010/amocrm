package amocrm

type Notes struct {
	ID                int          `json:"id,omitempty"`
	EntityID          int          `json:"entity_id,omitempty"`
	CreatedBy         int          `json:"created_by,omitempty"`
	UpdatedBy         int          `json:"updated_by,omitempty"`
	CreatedAt         int          `json:"created_at,omitempty"`
	UpdatedAt         int          `json:"updated_at,omitempty"`
	ResponsibleUserID int          `json:"responsible_user_id,omitempty"`
	GroupID           int          `json:"group_id,omitempty"`
	NoteType          string       `json:"note_type,omitempty"`
	Params            *NotesParams `json:"params,omitempty"`
}

type NotesParams struct {
	Text            string `json:"text,omitempty"`
	Link            string `json:"link,omitempty"`
	Uniq            string `json:"uniq,omitempty"`
	Duration        int    `json:"duration,omitempty"`
	Source          string `json:"source,omitempty"`
	CallResponsible int    `json:"call_responsible,omitempty"`
	Service         string `json:"service,omitempty"`
	Status          string `json:"status,omitempty"`
	Address         string `json:"address,omitempty"`
	Longitude       string `json:"longitude,omitempty"`
	Latitude        string `json:"latitude,omitempty"`
	Phone           string `json:"phone,omitempty"`
	VersionUUID     string `json:"version_uuid,omitempty"`
	FileUUID        string `json:"file_uuid,omitempty"`
	FileName        string `json:"file_name,omitempty"`
}
