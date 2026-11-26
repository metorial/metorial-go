package consumergroups

import (
	"encoding/json"
	"time"
)

// PortalsConsumerGroupsGetOutput represents the portals consumer groups get output type.
type PortalsConsumerGroupsGetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once you delete the group. Archived groups stay listable.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// IsDefault - Whether this is a default group on the portal.
	IsDefault bool `json:"is_default"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
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
