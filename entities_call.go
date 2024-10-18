package amocrm

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
