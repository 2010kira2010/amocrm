// The MIT License (MIT)
//
// Copyright (c) 2021 Alexey Khan
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

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

type CompanyOne struct {
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
	CustomFieldsValues []CustomsFields    `json:"custom_fields_values,omitempty"`
	AccountID          int                `json:"account_id,omitempty"`
	Embedded           *CompaniesEmbedded `json:"_embedded,omitempty"`
}

type CompaniesArr struct {
	Page     int `json:"_page,omitempty"`
	Embedded struct {
		Companies []struct {
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
			CustomFieldsValues []CustomsFields    `json:"custom_fields_values,omitempty"`
			AccountID          int                `json:"account_id,omitempty"`
			Embedded           *CompaniesEmbedded `json:"_embedded,omitempty"`
		} `json:"companies,omitempty"`
	} `json:"_embedded,omitempty"`
}

type CompaniesEmbedded struct {
	Tags []FieldValues `json:"tags,omitempty"`
}

type LeadOne struct {
	ID                 int             `json:"id,omitempty"`
	Name               string          `json:"name,omitempty"`
	Price              int             `json:"price,omitempty"`
	ResponsibleUserID  int             `json:"responsible_user_id,omitempty"`
	GroupID            int             `json:"group_id,omitempty"`
	StatusID           int             `json:"status_id,omitempty"`
	PipelineID         int             `json:"pipeline_id,omitempty"`
	LossReasonID       interface{}     `json:"loss_reason_id,omitempty"`
	CreatedBy          int             `json:"created_by,omitempty"`
	UpdatedBy          int             `json:"updated_by,omitempty"`
	CreatedAt          int             `json:"created_at,omitempty"`
	UpdatedAt          int             `json:"updated_at,omitempty"`
	ClosedAt           interface{}     `json:"closed_at,omitempty"`
	ClosestTaskAt      interface{}     `json:"closest_task_at,omitempty"`
	IsDeleted          bool            `json:"is_deleted,omitempty"`
	CustomFieldsValues []CustomsFields `json:"custom_fields_values,omitempty"`
	Score              interface{}     `json:"score,omitempty"`
	AccountID          int             `json:"account_id,omitempty"`
	LaborCost          interface{}     `json:"labor_cost,omitempty"`
	Embedded           *LeadEmbedded   `json:"_embedded,omitempty"`
}

type LeadsArr struct {
	Page  int `json:"_page,omitempty"`
	Links struct {
		Self struct {
			Href string `json:"href,omitempty"`
		} `json:"self,omitempty"`
	} `json:"_links,omitempty"`
	Embedded struct {
		Leads []struct {
			ID                 int             `json:"id,omitempty"`
			Name               string          `json:"name,omitempty"`
			Price              int             `json:"price,omitempty"`
			ResponsibleUserID  int             `json:"responsible_user_id,omitempty"`
			GroupID            int             `json:"group_id,omitempty"`
			StatusID           int             `json:"status_id,omitempty"`
			PipelineID         int             `json:"pipeline_id,omitempty"`
			LossReasonID       interface{}     `json:"loss_reason_id,omitempty"`
			CreatedBy          int             `json:"created_by,omitempty"`
			UpdatedBy          int             `json:"updated_by,omitempty"`
			CreatedAt          int             `json:"created_at,omitempty"`
			UpdatedAt          int             `json:"updated_at,omitempty"`
			ClosedAt           int             `json:"closed_at,omitempty"`
			ClosestTaskAt      interface{}     `json:"closest_task_at,omitempty"`
			IsDeleted          bool            `json:"is_deleted,omitempty"`
			CustomFieldsValues []CustomsFields `json:"custom_fields_values,omitempty"`
			Score              interface{}     `json:"score,omitempty"`
			AccountID          int             `json:"account_id,omitempty"`
			LaborCost          interface{}     `json:"labor_cost,omitempty"`
			Embedded           *LeadEmbedded   `json:"_embedded,omitempty"`
		} `json:"leads,omitempty"`
	} `json:"_embedded,omitempty"`
}

type LeadEmbedded struct {
	Tags      []FieldValues `json:"tags,omitempty"`
	Companies []FieldValues `json:"companies,omitempty"`
	Contacts  []FieldValues `json:"contacts,omitempty"`
}

type CustomsFields struct {
	FieldID   int                   `json:"field_id,omitempty"`
	FieldName string                `json:"field_name,omitempty"`
	FieldCode interface{}           `json:"field_code,omitempty"`
	FieldType string                `json:"field_type,omitempty"`
	Values    []CustomsFieldsValues `json:"values,omitempty"`
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

type Notes struct {
	ID                int         `json:"id,omitempty"`
	EntityID          int         `json:"entity_id,omitempty"`
	CreatedBy         int         `json:"created_by,omitempty"`
	UpdatedBy         int         `json:"updated_by,omitempty"`
	CreatedAt         int         `json:"created_at,omitempty"`
	UpdatedAt         int         `json:"updated_at,omitempty"`
	ResponsibleUserID int         `json:"responsible_user_id,omitempty"`
	GroupID           int         `json:"group_id,omitempty"`
	NoteType          string      `json:"note_type,omitempty"`
	Params            NotesParams `json:"params,omitempty"`
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

type Call struct {
	Direction         string `json:"direction"`                     // Направление звонка. inbound – входящий, outbound – исходящий. Обязательный параметр
	Uniq              string `json:"uniq,omitempty"`                // Уникальный идентификатор звонка. Необязательный параметр
	Duration          int    `json:"duration"`                      // Длительность звонка в секундах. Обязательный параметр
	Source            string `json:"source"`                        // Источник звонка. Обязательный параметр
	Link              string `json:"link,omitempty"`                // Ссылка на запись звонка. Необязательный параметр
	Phone             string `json:"phone"`                         // Номер телефона, по которому будет произведен поиск. Обязательный параметр
	CallResult        string `json:"call_result,omitempty"`         // Результат звонка. Необязательный параметр
	CallStatus        int    `json:"call_status,omitempty"`         // Статус звонка. Доступные варианты: 1 – оставил сообщение, 2 – перезвонить позже, 3 – нет на месте, 4 – разговор состоялся, 5 – неверный номер, 6 – Не дозвонился, 7 – номер занят. Необязательный параметр
	ResponsibleUserID int    `json:"responsible_user_id,omitempty"` // ID пользователя, ответственного за звонок
	CreatedBy         int    `json:"created_by,omitempty"`          // ID пользователя, создавший звонок
	UpdatedBy         int    `json:"updated_by,omitempty"`          // ID пользователя, изменивший звонок
	CreatedAt         int    `json:"created_at,omitempty"`          // Дата создания звонка, передается в Unix Timestamp
	UpdatedAt         int    `json:"updated_at,omitempty"`          // Дата изменения звонка, передается в Unix Timestamp
	RequestID         string `json:"request_id,omitempty"`          // Поле, которое вернется вам в ответе без изменений и не будет сохранено. Необязательный параметр
}

type CallError struct {
	Detail    string `json:"detail,omitempty"`
	RequestID string `json:"request_id,omitempty"`
	Status    int    `json:"status,omitempty"`
	Title     string `json:"title,omitempty"`
}

type EventAdd struct {
	Add []Event `json:"add"`
}

type Event struct {
	Type        string `json:"type"`         // Тип уведомления – phone_call
	PhoneNumber string `json:"phone_number"` // Номер телефона на который поступает звонок. Можно передавать в любом формате
	Users       []int  `json:"users"`        // Пользователи для которых будет отправлено уведомление. Если не передавать этот параметр, то уведомление будет отправлено для всех пользователей
}

type EventEmbeddedItem struct {
	ElementID   int    `json:"element_id"`
	ElementType int    `json:"element_type"`
	UID         string `json:"uid"`
	PhoneNumber string `json:"phone_number"`
}
