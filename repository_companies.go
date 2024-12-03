package amocrm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Companies interface {
	GetСompany(companieID string) (*Company, error, int)
	GetCompanies(values url.Values) (*Companiess, error, int)
	Create(сompanies []Company) ([]Company, error, int)
	Update(сompanies []Company) ([]Company, error, int)
}

// Verify interface compliance.
var _ Companies = сompanies{}

type сompanies struct {
	api *api
}

func newCompanies(api *api) Companies {
	return сompanies{api: api}
}

func (a сompanies) GetСompany(companieID string) (res *Company, err error, StatusCode int) {
	urlcompanie := url.Values{}
	urlcompanie.Add("with", "leads")
	resp, rErr := a.api.do(companiesEndpoint+endpoint("/"+companieID), http.MethodGet, urlcompanie, nil, nil)
	if rErr != nil {
		return nil, fmt.Errorf("get companie: %w", rErr), resp.StatusCode
	}

	defer func() {
		if clErr := resp.Body.Close(); clErr != nil {
			if err != nil {
				err = fmt.Errorf("close response body get companie: %v: %v", clErr, err)
			} else {
				err = fmt.Errorf("close response body get companie: %w", clErr)
			}
		}
	}()

	res = &Company{}
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, err, resp.StatusCode
	}

	return res, nil, resp.StatusCode
}

func (a сompanies) GetCompanies(values url.Values) (res *Companiess, err error, StatusCode int) {
	r, err := a.api.do(companiesEndpoint, http.MethodGet, values, nil, nil)
	if err != nil {
		return nil, err, r.StatusCode
	}
	defer func() {
		if clErr := r.Body.Close(); clErr != nil {
			if err != nil {
				err = fmt.Errorf("close response body get сompanies: %v: %v", clErr, err)
			} else {
				err = fmt.Errorf("close response body get сompanies: %w", clErr)
			}
		}
	}()

	res = &Companiess{}
	if err := json.NewDecoder(r.Body).Decode(res); err != nil {
		return nil, err, r.StatusCode
	}

	return res, err, r.StatusCode
}

func (a сompanies) Create(сompanies []Company) (res []Company, err error, StatusCode int) {
	resp, rErr := a.api.do(companiesEndpoint, http.MethodPost, nil, nil, сompanies)
	if rErr != nil {
		return nil, fmt.Errorf("Create сompanies: %w", rErr), resp.StatusCode
	}

	var resСompanie struct {
		Embedded struct {
			Companies []Company `json:"сompanies"`
		} `json:"_embedded"`
	}
	if err := a.api.read(resp, &resСompanie); err != nil {
		return nil, err, resp.StatusCode
	}
	return resСompanie.Embedded.Companies, nil, resp.StatusCode
}

func (a сompanies) Update(сompanies []Company) (res []Company, err error, StatusCode int) {
	resp, rErr := a.api.do(companiesEndpoint, http.MethodPatch, nil, nil, сompanies)
	if rErr != nil {
		return nil, fmt.Errorf("Update сompanies: %w", rErr), resp.StatusCode
	}

	var resСompanie struct {
		Embedded struct {
			Companies []Company `json:"сompanies"`
		} `json:"_embedded"`
	}
	if err := a.api.read(resp, &resСompanie); err != nil {
		return nil, err, resp.StatusCode
	}
	return resСompanie.Embedded.Companies, nil, resp.StatusCode
}
