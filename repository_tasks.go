package amocrm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Tasks describes methods available for Tasks entity.
type Tasks interface {
	GetTask(taskID string) (*Task, error, int)
	GetTasks(values url.Values) (*Taskss, error, int)
	Create(tasks []*Task) ([]*Task, error, int)
	Update(tasks []*Task) ([]*Task, error, int)
}

// Verify interface compliance.
var _ Tasks = tasks{}

type tasks struct {
	api *api
}

func newTasks(api *api) Tasks {
	return tasks{api: api}
}

func (a tasks) GetTask(taskID string) (res *Task, err error, StatusCode int) {
	urlTask := url.Values{}
	urlTask.Add("with", "contacts")
	r, errBody := a.api.do(tasksEndpoint+endpoint("/"+taskID), http.MethodGet, urlTask, nil, nil)
	if errBody != nil {
		//return nil, fmt.Errorf("get task: %w", err), r.StatusCode
		return nil, errBody, r.StatusCode
	}

	res = &Task{}
	if err := json.NewDecoder(r.Body).Decode(res); err != nil {
		return nil, err, r.StatusCode
	}

	return res, nil, r.StatusCode
}

func (a tasks) GetTasks(values url.Values) (res *Taskss, err error, StatusCode int) {
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

	res = &Taskss{}
	if err := json.NewDecoder(r.Body).Decode(res); err != nil {
		return nil, err, r.StatusCode
	}

	return res, err, r.StatusCode
}

func (a tasks) Create(tasks []*Task) (res []*Task, err error, StatusCode int) {
	resp, rErr := a.api.do(tasksEndpoint, http.MethodPost, nil, nil, tasks)
	if rErr != nil {
		return nil, fmt.Errorf("Create tasks: %w", rErr), resp.StatusCode
	}

	var resTask struct {
		Embedded struct {
			Tasks []*Task `json:"tasks"`
		} `json:"_embedded"`
	}

	if err := a.api.read(resp, &resTask); err != nil {
		return nil, err, resp.StatusCode
	}

	return resTask.Embedded.Tasks, nil, resp.StatusCode
}

func (a tasks) Update(tasks []*Task) (res []*Task, err error, StatusCode int) {
	resp, rErr := a.api.do(tasksEndpoint, http.MethodPatch, nil, nil, tasks)
	if rErr != nil {
		return nil, fmt.Errorf("Update tasks: %w", rErr), resp.StatusCode
	}

	var resTask struct {
		Embedded struct {
			Tasks []*Task `json:"tasks"`
		} `json:"_embedded"`
	}

	if err := a.api.read(resp, &resTask); err != nil {
		return nil, err, resp.StatusCode
	}

	return resTask.Embedded.Tasks, nil, resp.StatusCode
}
