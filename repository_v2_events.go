package amocrm

import (
	"fmt"
	"net/http"
)

// Events describes methods available for Events entity
type EventsV2 interface {
	Add(events []*Event) ([]*EventEmbeddedItem, error)
}

// Verify interface compliance.
var _ EventsV2 = eventsV2{}

type eventsV2 struct {
	api *api
}

func newEventsV2(api *api) EventsV2 {
	return eventsV2{api: api}
}

// Create returns an Contacts entity for successfully added Calls
func (a eventsV2) Add(events []*Event) ([]*EventEmbeddedItem, error) {

	resp, rErr := a.api.do(eventsV2endpoint, http.MethodPost, nil, nil, EventAdd{Add: events})
	if rErr != nil {
		return nil, fmt.Errorf("get calls: %w", rErr)
	}

	var res struct {
		Embedded struct {
			Items []*EventEmbeddedItem `json:"items"`
		} `json:"_embedded"`
	}
	if err := a.api.read(resp, &res); err != nil {
		return nil, err
	}

	return res.Embedded.Items, nil
}
