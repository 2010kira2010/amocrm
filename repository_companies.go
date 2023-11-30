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

type Companies interface {
	GetСompany(companieID string) (*CompanyOne, error, int)
	GetCompanies(values url.Values) (*CompaniesArr, error, int)
	Create(сompanies []CompanyOne) ([]CompanyOne, error, int)
	Update(сompanies []CompanyOne) ([]CompanyOne, error, int)
}

// Verify interface compliance.
var _ Companies = сompanies{}

type сompanies struct {
	api *api
}

func newCompanies(api *api) Companies {
	return сompanies{api: api}
}

func (a сompanies) GetСompany(companieID string) (res *CompanyOne, err error, StatusCode int) {
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

	res = &CompanyOne{}
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, err, resp.StatusCode
	}

	return res, nil, resp.StatusCode
}

func (a сompanies) GetCompanies(values url.Values) (res *CompaniesArr, err error, StatusCode int) {
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

	res = &CompaniesArr{}
	if err := json.NewDecoder(r.Body).Decode(res); err != nil {
		return nil, err, r.StatusCode
	}

	return res, err, r.StatusCode
}

func (a сompanies) Create(сompanies []CompanyOne) (res []CompanyOne, err error, StatusCode int) {
	resp, rErr := a.api.do(companiesEndpoint, http.MethodPost, nil, nil, сompanies)
	if rErr != nil {
		return nil, fmt.Errorf("Create сompanies: %w", rErr), resp.StatusCode
	}

	var resСompanie struct {
		Embedded struct {
			Companies []CompanyOne `json:"сompanies"`
		} `json:"_embedded"`
	}
	if err := a.api.read(resp, &resСompanie); err != nil {
		return nil, err, resp.StatusCode
	}
	return resСompanie.Embedded.Companies, nil, resp.StatusCode
}

func (a сompanies) Update(сompanies []CompanyOne) (res []CompanyOne, err error, StatusCode int) {
	resp, rErr := a.api.do(companiesEndpoint, http.MethodPatch, nil, nil, сompanies)
	if rErr != nil {
		return nil, fmt.Errorf("Update сompanies: %w", rErr), resp.StatusCode
	}

	var resСompanie struct {
		Embedded struct {
			Companies []CompanyOne `json:"сompanies"`
		} `json:"_embedded"`
	}
	if err := a.api.read(resp, &resСompanie); err != nil {
		return nil, err, resp.StatusCode
	}
	return resСompanie.Embedded.Companies, nil, resp.StatusCode
}
