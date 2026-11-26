package providertemplates

import (
	"encoding/json"
	"time"
)

// ProviderTemplatesUpdateOutput represents the provider templates update output type.
type ProviderTemplatesUpdateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once you archive the template.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata map[string]any `json:"metadata"`
	// IntegrationId - Integration the template was created from. `null` when it wasn’t created from one.
	IntegrationId *string `json:"integration_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapProviderTemplatesUpdateOutputFromJSON deserializes JSON data into a ProviderTemplatesUpdateOutput.
func MapProviderTemplatesUpdateOutputFromJSON(data []byte) (*ProviderTemplatesUpdateOutput, error) {
	var v ProviderTemplatesUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderTemplatesUpdateOutputToJSON serializes a ProviderTemplatesUpdateOutput to JSON.
func MapProviderTemplatesUpdateOutputToJSON(v *ProviderTemplatesUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderTemplatesUpdateBody represents the provider templates update body type.
type ProviderTemplatesUpdateBody struct {
	// Name - Display name for the provider template.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the provider template. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// MapProviderTemplatesUpdateBodyFromJSON deserializes JSON data into a ProviderTemplatesUpdateBody.
func MapProviderTemplatesUpdateBodyFromJSON(data []byte) (*ProviderTemplatesUpdateBody, error) {
	var v ProviderTemplatesUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderTemplatesUpdateBodyToJSON serializes a ProviderTemplatesUpdateBody to JSON.
func MapProviderTemplatesUpdateBodyToJSON(v *ProviderTemplatesUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
