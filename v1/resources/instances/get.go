package instances

import (
	"encoding/json"
	"time"
)

// InstancesGetOutputProject represents the instances get output project type.
type InstancesGetOutputProject struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The project's unique identifier
	Id string `json:"id"`
	// Status - The project's status
	Status string `json:"status"`
	// Slug - The project's slug
	Slug string `json:"slug"`
	// Name - The project's name
	Name string `json:"name"`
	// OrganizationId - The organization's unique identifier
	OrganizationId string `json:"organization_id"`
	// CreatedAt - The project's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The project's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// InstancesGetOutput represents the instances get output type.
type InstancesGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The instance's unique identifier
	Id string `json:"id"`
	// Slug - The instance's slug
	Slug string `json:"slug"`
	// Name - The instance's name
	Name string `json:"name"`
	// OrganizationId - The organization's unique identifier
	OrganizationId string `json:"organization_id"`
	// Type - The instance's type
	Type string `json:"type"`
	// CreatedAt - The instance's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The instance's last update date
	UpdatedAt time.Time                 `json:"updated_at"`
	Project   InstancesGetOutputProject `json:"project"`
}

// MapInstancesGetOutputFromJSON deserializes JSON data into a InstancesGetOutput.
func MapInstancesGetOutputFromJSON(data []byte) (*InstancesGetOutput, error) {
	var v InstancesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapInstancesGetOutputToJSON serializes a InstancesGetOutput to JSON.
func MapInstancesGetOutputToJSON(v *InstancesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
