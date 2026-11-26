package versions

import (
	"encoding/json"
	"time"
)

// ProvidersVersionsGetOutput represents the providers versions get output type.
type ProvidersVersionsGetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Version - Release label for the version.
	Version string `json:"version"`
	// ProviderId - Provider this version belongs to.
	ProviderId string `json:"provider_id"`
	// IsCurrent - The version deployments follow unless locked to another.
	IsCurrent bool `json:"is_current"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// SpecificationId - Specification this version belongs to.
	SpecificationId *string `json:"specification_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapProvidersVersionsGetOutputFromJSON deserializes JSON data into a ProvidersVersionsGetOutput.
func MapProvidersVersionsGetOutputFromJSON(data []byte) (*ProvidersVersionsGetOutput, error) {
	var v ProvidersVersionsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProvidersVersionsGetOutputToJSON serializes a ProvidersVersionsGetOutput to JSON.
func MapProvidersVersionsGetOutputToJSON(v *ProvidersVersionsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
