package providertemplates

import (
	"encoding/json"
	"time"
)

// ProviderTemplatesGetOutput represents the provider templates get output type.
type ProviderTemplatesGetOutput struct {
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

// MapProviderTemplatesGetOutputFromJSON deserializes JSON data into a ProviderTemplatesGetOutput.
func MapProviderTemplatesGetOutputFromJSON(data []byte) (*ProviderTemplatesGetOutput, error) {
	var v ProviderTemplatesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderTemplatesGetOutputToJSON serializes a ProviderTemplatesGetOutput to JSON.
func MapProviderTemplatesGetOutputToJSON(v *ProviderTemplatesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
