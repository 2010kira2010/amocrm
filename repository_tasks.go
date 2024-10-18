package amocrm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Tasks describes methods available for Tasks entity.
type Tasks interface {
	GetTask(taskID string) (*TaskOne, error, int)
	GetTasks(values url.Values) (*TasksArr, error, int)
	Create(tasks []TaskOne) ([]TaskOne, error, int)
	Update(tasks []TaskOne) ([]TaskOne, error, int)
}

// Verify interface compliance.
var _ Tasks = tasks{}

type tasks struct {
	api *api
}

func newTasks(api *api) Tasks {
	return tasks{api: api}
}

func (a tasks) GetTask(taskID string) (res *TaskOne, err error, StatusCode int) {
	urlTask := url.Values{}
	urlTask.Add("with", "contacts")
	r, errBody := a.api.do(tasksEndpoint+endpoint("/"+taskID), http.MethodGet, urlTask, nil, nil)
	if errBody != nil {
		//return nil, fmt.Errorf("get task: %w", err), r.StatusCode
		return nil, errBody, r.StatusCode
	}

	res = &TaskOne{}
	if err := json.NewDecoder(r.Body).Decode(res); err != nil {
		return nil, err, r.StatusCode
	}

	return res, nil, r.StatusCode
}

func (a tasks) GetTasks(values url.Values) (res *TasksArr, err error, StatusCode int) {
	r, err := a.api.do(tasksEndpoint, http.MethodGet, values, nil, nil)
	if err != nil {
		return nil, err, r.StatusCode
	}
	defer func() {
		if clErr := r.Body.Close(); clErr != nil {
			if err != nil {
				err = fmt.Errorf("close response body get tasks: %v: %v", clErr, err)
			} else {
				err = fmt.Errorf("close response body get tasks: %w", clErr)
			}
		}
	}()

	res = &TasksArr{}
	if err := json.NewDecoder(r.Body).Decode(res); err != nil {
		return nil, err, r.StatusCode
	}

	return res, err, r.StatusCode
}

func (a tasks) Create(tasks []TaskOne) (res []TaskOne, err error, StatusCode int) {
	resp, rErr := a.api.do(tasksEndpoint, http.MethodPost, nil, nil, tasks)
	if rErr != nil {
		return nil, fmt.Errorf("Create tasks: %w", rErr), resp.StatusCode
	}

	var resTask struct {
		Embedded struct {
			Tasks []TaskOne `json:"tasks"`
		} `json:"_embedded"`
	}

	if err := a.api.read(resp, &resTask); err != nil {
		return nil, err, resp.StatusCode
	}

	return resTask.Embedded.Tasks, nil, resp.StatusCode
}

func (a tasks) Update(tasks []TaskOne) (res []TaskOne, err error, StatusCode int) {
	resp, rErr := a.api.do(tasksEndpoint, http.MethodPatch, nil, nil, tasks)
	if rErr != nil {
		return nil, fmt.Errorf("Update tasks: %w", rErr), resp.StatusCode
	}

	var resTask struct {
		Embedded struct {
			Tasks []TaskOne `json:"tasks"`
		} `json:"_embedded"`
	}

	if err := a.api.read(resp, &resTask); err != nil {
		return nil, err, resp.StatusCode
	}

	return resTask.Embedded.Tasks, nil, resp.StatusCode
}
