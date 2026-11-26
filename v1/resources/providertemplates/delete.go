package providertemplates

import (
	"encoding/json"
	"time"
)

// ProviderTemplatesDeleteOutput represents the provider templates delete output type.
type ProviderTemplatesDeleteOutput struct {
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

// MapProviderTemplatesDeleteOutputFromJSON deserializes JSON data into a ProviderTemplatesDeleteOutput.
func MapProviderTemplatesDeleteOutputFromJSON(data []byte) (*ProviderTemplatesDeleteOutput, error) {
	var v ProviderTemplatesDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderTemplatesDeleteOutputToJSON serializes a ProviderTemplatesDeleteOutput to JSON.
func MapProviderTemplatesDeleteOutputToJSON(v *ProviderTemplatesDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
