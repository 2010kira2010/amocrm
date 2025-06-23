package amocrm

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Lang   string `json:"lang"`
	Rights struct {
		Leads struct {
			View   string `json:"view"`
			Edit   string `json:"edit"`
			Add    string `json:"add"`
			Delete string `json:"delete"`
			Export string `json:"export"`
		} `json:"leads"`
		Contacts struct {
			View   string `json:"view"`
			Edit   string `json:"edit"`
			Add    string `json:"add"`
			Delete string `json:"delete"`
			Export string `json:"export"`
		} `json:"contacts"`
		Companies struct {
			View   string `json:"view"`
			Edit   string `json:"edit"`
			Add    string `json:"add"`
			Delete string `json:"delete"`
			Export string `json:"export"`
		} `json:"companies"`
		Tasks struct {
			Edit   string `json:"edit"`
			Delete string `json:"delete"`
		} `json:"tasks"`
		MailAccess    bool `json:"mail_access"`
		CatalogAccess bool `json:"catalog_access"`
		StatusRights  []struct {
			EntityType string `json:"entity_type"`
			PipelineId int    `json:"pipeline_id"`
			StatusId   int    `json:"status_id"`
			Rights     struct {
				View   string `json:"view"`
				Edit   string `json:"edit"`
				Delete string `json:"delete"`
				Export string `json:"export,omitempty"`
			} `json:"rights"`
		} `json:"status_rights"`
		IsAdmin  bool        `json:"is_admin"`
		IsFree   bool        `json:"is_free"`
		IsActive bool        `json:"is_active"`
		GroupId  interface{} `json:"group_id"`
		RoleId   interface{} `json:"role_id"`
	} `json:"rights"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}

type Userss struct {
	TotalItems int `json:"_total_items"`
	Page       int `json:"_page"`
	PageCount  int `json:"_page_count"`
	Links      struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Embedded struct {
		Users []struct {
			Id     int    `json:"id"`
			Name   string `json:"name"`
			Email  string `json:"email"`
			Lang   string `json:"lang"`
			Rights struct {
				Leads struct {
					View   string `json:"view"`
					Edit   string `json:"edit"`
					Add    string `json:"add"`
					Delete string `json:"delete"`
					Export string `json:"export"`
				} `json:"leads"`
				Contacts struct {
					View   string `json:"view"`
					Edit   string `json:"edit"`
					Add    string `json:"add"`
					Delete string `json:"delete"`
					Export string `json:"export"`
				} `json:"contacts"`
				Companies struct {
					View   string `json:"view"`
					Edit   string `json:"edit"`
					Add    string `json:"add"`
					Delete string `json:"delete"`
					Export string `json:"export"`
				} `json:"companies"`
				Tasks struct {
					Edit   string `json:"edit"`
					Delete string `json:"delete"`
				} `json:"tasks"`
				MailAccess    bool `json:"mail_access"`
				CatalogAccess bool `json:"catalog_access"`
				StatusRights  []struct {
					EntityType string `json:"entity_type"`
					PipelineId int    `json:"pipeline_id"`
					StatusId   int    `json:"status_id"`
					Rights     struct {
						View   string `json:"view"`
						Edit   string `json:"edit"`
						Delete string `json:"delete"`
					} `json:"rights"`
				} `json:"status_rights"`
				IsAdmin  bool        `json:"is_admin"`
				IsFree   bool        `json:"is_free"`
				IsActive bool        `json:"is_active"`
				GroupId  interface{} `json:"group_id"`
				RoleId   interface{} `json:"role_id"`
			} `json:"rights"`
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
			Embedded struct {
				Roles []struct {
					Id    int    `json:"id"`
					Name  string `json:"name"`
					Links struct {
						Self struct {
							Href string `json:"href"`
						} `json:"self"`
					} `json:"_links"`
				} `json:"roles"`
				Groups []struct {
					Id   int    `json:"id"`
					Name string `json:"name"`
				} `json:"groups"`
			} `json:"_embedded"`
		} `json:"users"`
	} `json:"_embedded"`
}
