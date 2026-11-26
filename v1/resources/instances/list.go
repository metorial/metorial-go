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
	// AssistantEnabled - Whether Metorial Assistant is available in this project, including compliance restrictions.
	AssistantEnabled bool `json:"assistant_enabled"`
	// OrganizationId - The organization's unique identifier
	OrganizationId string `json:"organization_id"`
	// CreatedAt - The project's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The project's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// InstancesListOutputItems represents the instances list output items type.
type InstancesListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Slug - URL-safe name, unique across Metorial.
	Slug string `json:"slug"`
	// Name - Display name.
	Name string `json:"name"`
	// OrganizationId - Organization this instance belongs to.
	OrganizationId string `json:"organization_id"`
	// SandboxId - Sandbox backing this instance, if it is a sandbox.
	SandboxId *string `json:"sandbox_id,omitempty"`
	// Type - Whether this instance is for development or production use.
	Type string `json:"type"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time                       `json:"updated_at"`
	Project   InstancesListOutputItemsProject `json:"project"`
}

// InstancesListOutput represents the instances list output type.
type InstancesListOutput struct {
	// Object - String representing the object’s type.
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
