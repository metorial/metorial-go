package instances

import (
	"encoding/json"
	"time"
)

// InstancesListOutputItemsProject represents the instances list output items project type.
type InstancesListOutputItemsProject struct {
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

// InstancesListOutputItems represents the instances list output items type.
type InstancesListOutputItems struct {
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
	UpdatedAt time.Time                       `json:"updated_at"`
	Project   InstancesListOutputItemsProject `json:"project"`
}

// InstancesListOutput represents the instances list output type.
type InstancesListOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Items - The list of instances
	Items []InstancesListOutputItems `json:"items"`
}

// MapInstancesListOutputFromJSON deserializes JSON data into a InstancesListOutput.
func MapInstancesListOutputFromJSON(data []byte) (*InstancesListOutput, error) {
	var v InstancesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapInstancesListOutputToJSON serializes a InstancesListOutput to JSON.
func MapInstancesListOutputToJSON(v *InstancesListOutput) ([]byte, error) {
	return json.Marshal(v)
}
