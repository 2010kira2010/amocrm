package amocrm

import (
	"fmt"
)

type endpoint string

func (e endpoint) path() string {
	if e == eventsV2endpoint {
		return fmt.Sprintf("/api/v2/%s/", e)
	}

	return fmt.Sprintf("/api/v%d/%s", apiVersion, e)
}

const (
	accountsEndpoint  endpoint = "account"
	leadsEndpoint     endpoint = "leads"
	contactsEndpoint  endpoint = "contacts"
	companiesEndpoint endpoint = "companies"
	pipelinesEndpoint endpoint = "pipelines"
	callsEndpoint     endpoint = "calls"
	catalogsEndpoint  endpoint = "catalogs"
	tasksEndpoint     endpoint = "tasks"
	usersEndpoint     endpoint = "users"
	eventsV2endpoint  endpoint = "events"
)
