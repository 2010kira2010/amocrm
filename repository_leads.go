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

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Leads describes methods available for Leads entity.
type Leads interface {
	GetLead(leadID string) (*LeadOne, error, int)
	GetListCustomFieldsLeads(fieldID string) (*CustomsField, error, int)
	GetLeads(values url.Values) (*LeadsArr, error, int)
	Create(leads []LeadOne) ([]LeadOne, error, int)
	Update(leads []LeadOne) ([]LeadOne, error, int)
	AddNotes(notes []Notes) ([]Notes, error, int)
}

// Verify interface compliance.
var _ Leads = leads{}

type leads struct {
	api *api
}

func newLeads(api *api) Leads {
	return leads{api: api}
}

func (a leads) GetListCustomFieldsLeads(fieldID string) (res *CustomsField, err error, StatusCode int) {
	urlLead := url.Values{}
	r, errBody := a.api.do(leadsEndpoint+endpoint("/custom_fields/"+fieldID), http.MethodGet, urlLead, nil, nil)

	if errBody != nil {
		return nil, errBody, r.StatusCode
	}

	res = &CustomsField{}
	if err := json.NewDecoder(r.Body).Decode(res); err != nil {
		return nil, err, r.StatusCode
	}

	return res, nil, r.StatusCode
}

func (a leads) GetLead(leadID string) (res *LeadOne, err error, StatusCode int) {
	urlLead := url.Values{}
	urlLead.Add("with", "contacts")
	r, errBody := a.api.do(leadsEndpoint+endpoint("/"+leadID), http.MethodGet, urlLead, nil, nil)
	if errBody != nil {
		//return nil, fmt.Errorf("get lead: %w", err), r.StatusCode
		return nil, errBody, r.StatusCode
	}

	res = &LeadOne{}
	if err := json.NewDecoder(r.Body).Decode(res); err != nil {
		return nil, err, r.StatusCode
	}

	return res, nil, r.StatusCode
}

func (a leads) GetLeads(values url.Values) (res *LeadsArr, err error, StatusCode int) {
	r, err := a.api.do(leadsEndpoint, http.MethodGet, values, nil, nil)
	if err != nil {
		return nil, err, r.StatusCode
	}
	defer func() {
		if clErr := r.Body.Close(); clErr != nil {
			if err != nil {
				err = fmt.Errorf("close response body get leads: %v: %v", clErr, err)
			} else {
				err = fmt.Errorf("close response body get leads: %w", clErr)
			}
		}
	}()

	res = &LeadsArr{}
	if err := json.NewDecoder(r.Body).Decode(res); err != nil {
		return nil, err, r.StatusCode
	}

	return res, err, r.StatusCode
}

func (a leads) Create(leads []LeadOne) (res []LeadOne, err error, StatusCode int) {
	resp, rErr := a.api.do(leadsEndpoint, http.MethodPost, nil, nil, leads)
	if rErr != nil {
		return nil, fmt.Errorf("Create leads: %w", rErr), resp.StatusCode
	}

	var resLead struct {
		Embedded struct {
			Leads []LeadOne `json:"leads"`
		} `json:"_embedded"`
	}

	if err := a.api.read(resp, &resLead); err != nil {
		return nil, err, resp.StatusCode
	}

	return resLead.Embedded.Leads, nil, resp.StatusCode
}

func (a leads) Update(leads []LeadOne) (res []LeadOne, err error, StatusCode int) {
	resp, rErr := a.api.do(leadsEndpoint, http.MethodPatch, nil, nil, leads)
	if rErr != nil {
		return nil, fmt.Errorf("Update leads: %w", rErr), resp.StatusCode
	}

	var resLead struct {
		Embedded struct {
			Leads []LeadOne `json:"leads"`
		} `json:"_embedded"`
	}

	if err := a.api.read(resp, &resLead); err != nil {
		return nil, err, resp.StatusCode
	}

	return resLead.Embedded.Leads, nil, resp.StatusCode
}

func (a leads) AddNotes(notes []Notes) (res []Notes, err error, StatusCode int) {
	resp, rErr := a.api.do(leadsEndpoint+"/notes", http.MethodPost, nil, nil, notes)
	if rErr != nil {
		return nil, fmt.Errorf("Add notes to leads: %w", rErr), resp.StatusCode
	}

	var resNote struct {
		Embedded struct {
			Notes []Notes `json:"notes"`
		} `json:"_embedded"`
	}

	if err := a.api.read(resp, &resNote); err != nil {
		return nil, err, resp.StatusCode
	}

	return resNote.Embedded.Notes, nil, resp.StatusCode
}
