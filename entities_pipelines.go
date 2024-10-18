package amocrm

type PipelineOne struct {
	ID           int                `json:"id,omitempty"`
	Name         string             `json:"name,omitempty"`
	Sort         int                `json:"sort,omitempty"`
	IsMain       bool               `json:"is_main,omitempty"`
	IsUnsortedOn bool               `json:"is_unsorted_on,omitempty"`
	IsArchive    bool               `json:"is_archive,omitempty"`
	AccountID    int                `json:"account_id,omitempty"`
	Embedded     *PipelinesEmbedded `json:"_embedded,omitempty"`
}

type PipelinesArr struct {
	TotalItems int `json:"_total_items,omitempty"`
	Embedded   struct {
		Pipelines []struct {
			ID           int                `json:"id,omitempty"`
			Name         string             `json:"name,omitempty"`
			Sort         int                `json:"sort,omitempty"`
			IsMain       bool               `json:"is_main,omitempty"`
			IsUnsortedOn bool               `json:"is_unsorted_on,omitempty"`
			IsArchive    bool               `json:"is_archive,omitempty"`
			AccountID    int                `json:"account_id,omitempty"`
			Embedded     *PipelinesEmbedded `json:"_embedded,omitempty"`
		} `json:"pipelines,omitempty"`
	} `json:"_embedded,omitempty"`
}

type PipelineStatuses struct {
	TotalItems int `json:"_total_items,omitempty"`
	Embedded   struct {
		Statuses []struct {
			ID         int    `json:"id,omitempty"`
			Name       string `json:"name,omitempty"`
			Sort       int    `json:"sort,omitempty"`
			IsEditable bool   `json:"is_editable,omitempty"`
			PipelineID int    `json:"pipeline_id,omitempty"`
			Color      string `json:"color,omitempty"`
			Type       int    `json:"type,omitempty"`
			AccountID  int    `json:"account_id,omitempty"`
		} `json:"statuses,omitempty"`
	} `json:"_embedded,omitempty"`
}

type PipelinesEmbedded struct {
	Statuses []struct {
		ID         int    `json:"id,omitempty"`
		Name       string `json:"name,omitempty"`
		Sort       int    `json:"sort,omitempty"`
		IsEditable bool   `json:"is_editable,omitempty"`
		PipelineID int    `json:"pipeline_id,omitempty"`
		Color      string `json:"color,omitempty"`
		Type       int    `json:"type,omitempty"`
		AccountID  int    `json:"account_id,omitempty"`
	} `json:"statuses,omitempty"`
}
