package amocrm

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
