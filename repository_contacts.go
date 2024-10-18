package amocrm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Contacts interface {
	GetContact(contactID string) (*ContactOne, error, int)
	GetContacts(values url.Values) (*ContactsArr, error, int)
	Create(contacts []ContactOne) ([]ContactOne, error, int)
	Update(contacts []ContactOne) ([]ContactOne, error, int)
}

// Verify interface compliance.
var _ Contacts = contacts{}

type contacts struct {
	api *api
}

func newContacts(api *api) Contacts {
	return contacts{api: api}
}

func (a contacts) GetContact(contactID string) (res *ContactOne, err error, StatusCode int) {
	urlcontact := url.Values{}
	urlcontact.Add("with", "leads")
	resp, rErr := a.api.do(contactsEndpoint+endpoint("/"+contactID), http.MethodGet, urlcontact, nil, nil)
	if rErr != nil {
		return nil, fmt.Errorf("get contact: %w", rErr), resp.StatusCode
	}

	defer func() {
		if clErr := resp.Body.Close(); clErr != nil {
			if err != nil {
				err = fmt.Errorf("close response body get contact: %v: %v", clErr, err)
			} else {
				err = fmt.Errorf("close response body get contact: %w", clErr)
			}
		}
	}()

	res = &ContactOne{}
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, err, resp.StatusCode
	}

	return res, nil, resp.StatusCode
}

func (a contacts) GetContacts(values url.Values) (res *ContactsArr, err error, StatusCode int) {
	r, err := a.api.do(contactsEndpoint, http.MethodGet, values, nil, nil)
	if err != nil {
		return nil, err, r.StatusCode
	}
	defer func() {
		if clErr := r.Body.Close(); clErr != nil {
			if err != nil {
				err = fmt.Errorf("close response body get contacts: %v: %v", clErr, err)
			} else {
				err = fmt.Errorf("close response body get contacts: %w", clErr)
			}
		}
	}()

	res = &ContactsArr{}
	if err := json.NewDecoder(r.Body).Decode(res); err != nil {
		return nil, err, r.StatusCode
	}

	return res, err, r.StatusCode
}

func (a contacts) Create(contacts []ContactOne) (res []ContactOne, err error, StatusCode int) {
	resp, rErr := a.api.do(contactsEndpoint, http.MethodPost, nil, nil, contacts)
	if rErr != nil {
		return nil, fmt.Errorf("Create contacts: %w", rErr), resp.StatusCode
	}

	var resContact struct {
		Embedded struct {
			Contacts []ContactOne `json:"contacts"`
		} `json:"_embedded"`
	}
	if err := a.api.read(resp, &resContact); err != nil {
		return nil, err, resp.StatusCode
	}
	return resContact.Embedded.Contacts, nil, resp.StatusCode
}

func (a contacts) Update(contacts []ContactOne) (res []ContactOne, err error, StatusCode int) {
	resp, rErr := a.api.do(contactsEndpoint, http.MethodPatch, nil, nil, contacts)
	if rErr != nil {
		return nil, fmt.Errorf("Update contacts: %w", rErr), resp.StatusCode
	}

	var resContact struct {
		Embedded struct {
			Contacts []ContactOne `json:"contacts"`
		} `json:"_embedded"`
	}
	if err := a.api.read(resp, &resContact); err != nil {
		return nil, err, resp.StatusCode
	}
	return resContact.Embedded.Contacts, nil, resp.StatusCode
}
