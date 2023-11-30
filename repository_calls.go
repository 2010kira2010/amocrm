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
)

// Calls describes methods available for Calls entity
type Calls interface {
	Create(calls []Call) ([]ContactOne, []CallError, error)
}

// Verify interface compliance.
var _ Calls = calls{}

type calls struct {
	api *api
}

func newCalls(api *api) Calls {
	return calls{api: api}
}

// Create returns an Contacts entity for successfully added Calls
func (a calls) Create(calls []Call) ([]ContactOne, []CallError, error) {
	resp, rErr := a.api.do(callsEndpoint, http.MethodPost, nil, nil, calls)
	if rErr != nil {
		return nil, nil, fmt.Errorf("get calls: %w", rErr)
	}

	var res struct {
		Errors   []CallError `json:"errors"`
		Embedded struct {
			Contacts []ContactOne `json:"calls"`
		} `json:"_embedded"`
	}
	if err := a.api.readCall(resp, &res); err != nil {
		return nil, nil, err
	}

	return res.Embedded.Contacts, res.Errors, nil
}

func (a *api) readCall(response *http.Response, target interface{}) (err error) {
	defer func() {
		if clErr := response.Body.Close(); clErr != nil {
			if err != nil {
				err = fmt.Errorf("close response body: %v: %v", clErr, err)
			} else {
				err = fmt.Errorf("close response body: %w", clErr)
			}
		}
	}()

	err = json.NewDecoder(response.Body).Decode(target)
	return
}
