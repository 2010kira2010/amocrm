package amocrm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Contacts interface {
	GetContact(contactID string) (*Contact, error, int)
	GetContacts(values url.Values) (*Contactss, error, int)
	Create(contacts []Contact) ([]Contact, error, int)
	Update(contacts []Contact) ([]Contact, error, int)
}

// Verify interface compliance.
var _ Contacts = contacts{}

type contacts struct {
	api *api
}

func newContacts(api *api) Contacts {
	return contacts{api: api}
}

func (a contacts) GetContact(contactID string) (res *Contact, err error, StatusCode int) {
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

	res = &Contact{}
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, err, resp.StatusCode
	}

	return res, nil, resp.StatusCode
}

func (a contacts) GetContacts(values url.Values) (res *Contactss, err error, StatusCode int) {
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

	res = &Contactss{}
	if err := json.NewDecoder(r.Body).Decode(res); err != nil {
		return nil, err, r.StatusCode
	}

	return res, err, r.StatusCode
}

func (a contacts) Create(contacts []Contact) (res []Contact, err error, StatusCode int) {
	resp, rErr := a.api.do(contactsEndpoint, http.MethodPost, nil, nil, contacts)
	if rErr != nil {
		return nil, fmt.Errorf("Create contacts: %w", rErr), resp.StatusCode
	}

	var resContact struct {
		Embedded struct {
			Contacts []Contact `json:"contacts"`
		} `json:"_embedded"`
	}
	if err := a.api.read(resp, &resContact); err != nil {
		return nil, err, resp.StatusCode
	}
	return resContact.Embedded.Contacts, nil, resp.StatusCode
}

func (a contacts) Update(contacts []Contact) (res []Contact, err error, StatusCode int) {
	resp, rErr := a.api.do(contactsEndpoint, http.MethodPatch, nil, nil, contacts)
	if rErr != nil {
		return nil, fmt.Errorf("Update contacts: %w", rErr), resp.StatusCode
	}

	var resContact struct {
		Embedded struct {
			Contacts []Contact `json:"contacts"`
		} `json:"_embedded"`
	}
	if err := a.api.read(resp, &resContact); err != nil {
		return nil, err, resp.StatusCode
	}
	return resContact.Embedded.Contacts, nil, resp.StatusCode
}
