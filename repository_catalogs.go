package amocrm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Catalogs describes methods available for Catalogs entity.
type Catalogs interface {
	GetCatalog(catalogID string) (*Catalog, error, int)
	GetCatalogs(values url.Values) (*Catalogss, error, int)
	Create(catalogs []*Catalog) ([]*Catalog, error, int)
	Update(catalogs []*Catalog) ([]*Catalog, error, int)
}

// Verify interface compliance.
var _ Catalogs = catalogs{}

type catalogs struct {
	api *api
}

func newCatalogs(api *api) Catalogs {
	return catalogs{api: api}
}

func (a catalogs) GetCatalog(catalogID string) (res *Catalog, err error, StatusCode int) {
	urlCatalog := url.Values{}
	urlCatalog.Add("with", "contacts")
	r, errBody := a.api.do(catalogsEndpoint+endpoint("/"+catalogID), http.MethodGet, urlCatalog, nil, nil)
	if errBody != nil {
		//return nil, fmt.Errorf("get catalog: %w", err), r.StatusCode
		return nil, errBody, r.StatusCode
	}

	res = &Catalog{}
	if err := json.NewDecoder(r.Body).Decode(res); err != nil {
		return nil, err, r.StatusCode
	}

	return res, nil, r.StatusCode
}

func (a catalogs) GetCatalogs(values url.Values) (res *Catalogss, err error, StatusCode int) {
	r, err := a.api.do(catalogsEndpoint, http.MethodGet, values, nil, nil)
	if err != nil {
		return nil, err, r.StatusCode
	}
	defer func() {
		if clErr := r.Body.Close(); clErr != nil {
			if err != nil {
				err = fmt.Errorf("close response body get catalogs: %v: %v", clErr, err)
			} else {
				err = fmt.Errorf("close response body get catalogs: %w", clErr)
			}
		}
	}()

	res = &Catalogss{}
	if err := json.NewDecoder(r.Body).Decode(res); err != nil {
		return nil, err, r.StatusCode
	}

	return res, err, r.StatusCode
}

func (a catalogs) Create(catalogs []*Catalog) (res []*Catalog, err error, StatusCode int) {
	resp, rErr := a.api.do(catalogsEndpoint, http.MethodPost, nil, nil, catalogs)
	if rErr != nil {
		return nil, fmt.Errorf("Create catalogs: %w", rErr), resp.StatusCode
	}

	var resCatalog struct {
		Embedded struct {
			Catalogs []*Catalog `json:"catalogs"`
		} `json:"_embedded"`
	}

	if err := a.api.read(resp, &resCatalog); err != nil {
		return nil, err, resp.StatusCode
	}

	return resCatalog.Embedded.Catalogs, nil, resp.StatusCode
}

func (a catalogs) Update(catalogs []*Catalog) (res []*Catalog, err error, StatusCode int) {
	resp, rErr := a.api.do(catalogsEndpoint, http.MethodPatch, nil, nil, catalogs)
	if rErr != nil {
		return nil, fmt.Errorf("Update catalogs: %w", rErr), resp.StatusCode
	}

	var resCatalog struct {
		Embedded struct {
			Catalogs []*Catalog `json:"catalogs"`
		} `json:"_embedded"`
	}

	if err := a.api.read(resp, &resCatalog); err != nil {
		return nil, err, resp.StatusCode
	}

	return resCatalog.Embedded.Catalogs, nil, resp.StatusCode
}
