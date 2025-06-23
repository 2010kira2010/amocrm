package amocrm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Users describes methods available for Users entity.
type Users interface {
	GetUser(userID string) (*User, error, int)
	GetUsers(values url.Values) (*Userss, error, int)
	Create(users []*User) ([]*User, error, int)
	Update(users []*User) ([]*User, error, int)
}

// Verify interface compliance.
var _ Users = users{}

type users struct {
	api *api
}

func newUsers(api *api) Users {
	return users{api: api}
}

func (a users) GetUser(userID string) (res *User, err error, StatusCode int) {
	urlUser := url.Values{}
	urlUser.Add("with", "contacts")
	r, errBody := a.api.do(usersEndpoint+endpoint("/"+userID), http.MethodGet, urlUser, nil, nil)
	if errBody != nil {
		//return nil, fmt.Errorf("get user: %w", err), r.StatusCode
		return nil, errBody, r.StatusCode
	}

	res = &User{}
	if err := json.NewDecoder(r.Body).Decode(res); err != nil {
		return nil, err, r.StatusCode
	}

	return res, nil, r.StatusCode
}

func (a users) GetUsers(values url.Values) (res *Userss, err error, StatusCode int) {
	r, err := a.api.do(usersEndpoint, http.MethodGet, values, nil, nil)
	if err != nil {
		return nil, err, r.StatusCode
	}
	defer func() {
		if clErr := r.Body.Close(); clErr != nil {
			if err != nil {
				err = fmt.Errorf("close response body get users: %v: %v", clErr, err)
			} else {
				err = fmt.Errorf("close response body get users: %w", clErr)
			}
		}
	}()

	res = &Userss{}
	if err := json.NewDecoder(r.Body).Decode(res); err != nil {
		return nil, err, r.StatusCode
	}

	return res, err, r.StatusCode
}

func (a users) Create(users []*User) (res []*User, err error, StatusCode int) {
	resp, rErr := a.api.do(usersEndpoint, http.MethodPost, nil, nil, users)
	if rErr != nil {
		return nil, fmt.Errorf("Create users: %w", rErr), resp.StatusCode
	}

	var resUser struct {
		Embedded struct {
			Users []*User `json:"users"`
		} `json:"_embedded"`
	}

	if err := a.api.read(resp, &resUser); err != nil {
		return nil, err, resp.StatusCode
	}

	return resUser.Embedded.Users, nil, resp.StatusCode
}

func (a users) Update(users []*User) (res []*User, err error, StatusCode int) {
	resp, rErr := a.api.do(usersEndpoint, http.MethodPatch, nil, nil, users)
	if rErr != nil {
		return nil, fmt.Errorf("Update users: %w", rErr), resp.StatusCode
	}

	var resUser struct {
		Embedded struct {
			Users []*User `json:"users"`
		} `json:"_embedded"`
	}

	if err := a.api.read(resp, &resUser); err != nil {
		return nil, err, resp.StatusCode
	}

	return resUser.Embedded.Users, nil, resp.StatusCode
}
