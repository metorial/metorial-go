package consumergroups

import (
	"encoding/json"
	"time"
)

// PortalsConsumerGroupsGetOutput represents the portals consumer groups get output type.
type PortalsConsumerGroupsGetOutput struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Status      string    `json:"status"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	IsDefault   bool      `json:"is_default"`
	SsoGroupIds []string  `json:"sso_group_ids"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MapPortalsConsumerGroupsGetOutputFromJSON deserializes JSON data into a PortalsConsumerGroupsGetOutput.
func MapPortalsConsumerGroupsGetOutputFromJSON(data []byte) (*PortalsConsumerGroupsGetOutput, error) {
	var v PortalsConsumerGroupsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerGroupsGetOutputToJSON serializes a PortalsConsumerGroupsGetOutput to JSON.
func MapPortalsConsumerGroupsGetOutputToJSON(v *PortalsConsumerGroupsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
