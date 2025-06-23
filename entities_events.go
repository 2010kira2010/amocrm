package amocrm

type EventAdd struct {
	Add []*Event `json:"add,omitempty"`
}

type Event struct {
	Type        string `json:"type,omitempty"`         // Тип уведомления – phone_call
	PhoneNumber string `json:"phone_number,omitempty"` // Номер телефона на который поступает звонок. Можно передавать в любом формате
	Users       []int  `json:"users,omitempty"`        // Пользователи для которых будет отправлено уведомление. Если не передавать этот параметр, то уведомление будет отправлено для всех пользователей
}

type EventEmbeddedItem struct {
	ElementID   int    `json:"element_id,omitempty"`
	ElementType int    `json:"element_type,omitempty"`
	UID         string `json:"uid,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}
