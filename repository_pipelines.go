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

// Pipelines describes methods available for Pipelines entity.
type Pipelines interface {
	GetPipeline(pipelineID string) (*PipelineOne, error, int)
	GetPipelineStatuses(pipelineID string) (*PipelineStatuses, error, int)
	GetPipelines(values url.Values) (*PipelinesArr, error, int)
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

func (a pipelines) GetPipelines(values url.Values) (res *PipelinesArr, err error, StatusCode int) {
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

	res = &PipelinesArr{}
	if err := json.NewDecoder(r.Body).Decode(res); err != nil {
		return nil, err, r.StatusCode
	}

	return res, err, r.StatusCode
}
