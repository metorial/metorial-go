package consumergroups

import (
	"encoding/json"
	"time"
)

// PortalsConsumerGroupsUpdateOutput represents the portals consumer groups update output type.
type PortalsConsumerGroupsUpdateOutput struct {
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

// MapPortalsConsumerGroupsUpdateOutputFromJSON deserializes JSON data into a PortalsConsumerGroupsUpdateOutput.
func MapPortalsConsumerGroupsUpdateOutputFromJSON(data []byte) (*PortalsConsumerGroupsUpdateOutput, error) {
	var v PortalsConsumerGroupsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerGroupsUpdateOutputToJSON serializes a PortalsConsumerGroupsUpdateOutput to JSON.
func MapPortalsConsumerGroupsUpdateOutputToJSON(v *PortalsConsumerGroupsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsConsumerGroupsUpdateBody represents the portals consumer groups update body type.
type PortalsConsumerGroupsUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// IsDefault - Whether this is a default group on the portal.
	IsDefault *bool `json:"is_default,omitempty"`
}

// MapPortalsConsumerGroupsUpdateBodyFromJSON deserializes JSON data into a PortalsConsumerGroupsUpdateBody.
func MapPortalsConsumerGroupsUpdateBodyFromJSON(data []byte) (*PortalsConsumerGroupsUpdateBody, error) {
	var v PortalsConsumerGroupsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerGroupsUpdateBodyToJSON serializes a PortalsConsumerGroupsUpdateBody to JSON.
func MapPortalsConsumerGroupsUpdateBodyToJSON(v *PortalsConsumerGroupsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
