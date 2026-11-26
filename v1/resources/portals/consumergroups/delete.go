package consumergroups

import (
	"encoding/json"
	"time"
)

// PortalsConsumerGroupsDeleteOutput represents the portals consumer groups delete output type.
type PortalsConsumerGroupsDeleteOutput struct {
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

// MapPortalsConsumerGroupsDeleteOutputFromJSON deserializes JSON data into a PortalsConsumerGroupsDeleteOutput.
func MapPortalsConsumerGroupsDeleteOutputFromJSON(data []byte) (*PortalsConsumerGroupsDeleteOutput, error) {
	var v PortalsConsumerGroupsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerGroupsDeleteOutputToJSON serializes a PortalsConsumerGroupsDeleteOutput to JSON.
func MapPortalsConsumerGroupsDeleteOutputToJSON(v *PortalsConsumerGroupsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
