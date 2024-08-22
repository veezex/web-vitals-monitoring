package metric

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type metricImpl struct {
	id          string
	uri         string
	client      string
	name        string
	value       float64
	delta       float64
	attribution string
	rating      string
}

func Parse(in map[string]interface{}) (Metric, error) {
	if in["id"] == nil || in["client"] == nil || in["uri"] == nil || in["name"] == nil || in["value"] == nil || in["target"] == nil || in["rating"] == nil {
		return nil, fmt.Errorf("Invalid input: missing required fields")
	}

	return &metricImpl{
		id:          in["id"].(string),
		uri:         in["uri"].(string),
		client:      in["client"].(string),
		name:        in["name"].(string),
		value:       in["value"].(float64),
		attribution: in["target"].(string),
		rating:      in["rating"].(string),
	}, nil
}

func (m *metricImpl) GetName() string {
	return m.name
}

func (m *metricImpl) GetClient() string {
	return m.client
}

func (m *metricImpl) GetUri() string {
	return m.uri
}

func (m *metricImpl) GetValue() float64 {
	return m.value
}

func (m *metricImpl) GetRating() string {
	return m.rating
}

func (m *metricImpl) GetID() string {
	return m.id
}

func (m *metricImpl) GetAttribution() string {
	return m.attribution
}

func (m *metricImpl) GetDelta() float64 {
	return m.delta
}
