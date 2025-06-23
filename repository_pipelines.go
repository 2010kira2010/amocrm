package amocrm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Pipelines describes methods available for Pipelines entity.
type Pipelines interface {
	GetPipeline(pipelineID string) (*PipelineOne, error, int)
	GetPipelineStatuses(pipelineID string) (*PipelineStatuses, error, int)
	GetPipelines(values url.Values) (*Pipeliness, error, int)
}

// Verify interface compliance.
var _ Pipelines = pipelines{}

type pipelines struct {
	api *api
}

func newPipelines(api *api) Pipelines {
	return pipelines{api: api}
}

func (a pipelines) GetPipeline(pipelineID string) (res *PipelineOne, err error, StatusCode int) {
	urlPipeline := url.Values{}
	r, errBody := a.api.do(leadsEndpoint+"/"+pipelinesEndpoint+endpoint("/"+pipelineID), http.MethodGet, urlPipeline, nil, nil)

	if errBody != nil {
		return nil, errBody, r.StatusCode
	}

	res = &PipelineOne{}
	if err := json.NewDecoder(r.Body).Decode(res); err != nil {
		return nil, err, r.StatusCode
	}

	return res, nil, r.StatusCode
}

func (a pipelines) GetPipelineStatuses(pipelineID string) (res *PipelineStatuses, err error, StatusCode int) {
	urlPipeline := url.Values{}
	r, errBody := a.api.do(leadsEndpoint+"/"+pipelinesEndpoint+endpoint("/"+pipelineID)+"/statuses", http.MethodGet, urlPipeline, nil, nil)

	if errBody != nil {
		return nil, errBody, r.StatusCode
	}

	res = &PipelineStatuses{}
	if err := json.NewDecoder(r.Body).Decode(res); err != nil {
		return nil, err, r.StatusCode
	}

	return res, nil, r.StatusCode
}

func (a pipelines) GetPipelines(values url.Values) (res *Pipeliness, err error, StatusCode int) {
	r, err := a.api.do(leadsEndpoint+"/"+pipelinesEndpoint, http.MethodGet, values, nil, nil)
	if err != nil {
		return nil, err, r.StatusCode
	}
	defer func() {
		if clErr := r.Body.Close(); clErr != nil {
			if err != nil {
				err = fmt.Errorf("close response body get pipelines: %v: %v", clErr, err)
			} else {
				err = fmt.Errorf("close response body get pipelines: %w", clErr)
			}
		}
	}()

	res = &Pipeliness{}
	if err := json.NewDecoder(r.Body).Decode(res); err != nil {
		return nil, err, r.StatusCode
	}

	return res, err, r.StatusCode
}
