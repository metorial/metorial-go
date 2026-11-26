package consumergroups

import (
	"encoding/json"
	"time"
)

// PortalsConsumerGroupsCreateOutput represents the portals consumer groups create output type.
type PortalsConsumerGroupsCreateOutput struct {
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

// MapPortalsConsumerGroupsCreateOutputFromJSON deserializes JSON data into a PortalsConsumerGroupsCreateOutput.
func MapPortalsConsumerGroupsCreateOutputFromJSON(data []byte) (*PortalsConsumerGroupsCreateOutput, error) {
	var v PortalsConsumerGroupsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerGroupsCreateOutputToJSON serializes a PortalsConsumerGroupsCreateOutput to JSON.
func MapPortalsConsumerGroupsCreateOutputToJSON(v *PortalsConsumerGroupsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsConsumerGroupsCreateBody represents the portals consumer groups create body type.
type PortalsConsumerGroupsCreateBody struct {
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// IsDefault - Whether this is a default group on the portal.
	IsDefault *bool `json:"is_default,omitempty"`
}

// MapPortalsConsumerGroupsCreateBodyFromJSON deserializes JSON data into a PortalsConsumerGroupsCreateBody.
func MapPortalsConsumerGroupsCreateBodyFromJSON(data []byte) (*PortalsConsumerGroupsCreateBody, error) {
	var v PortalsConsumerGroupsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerGroupsCreateBodyToJSON serializes a PortalsConsumerGroupsCreateBody to JSON.
func MapPortalsConsumerGroupsCreateBodyToJSON(v *PortalsConsumerGroupsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
