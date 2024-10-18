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
