package environments

import (
	"encoding/json"
	"time"
)

// CustomProvidersEnvironmentsGetOutput represents the custom providers environments get output type.
type CustomProvidersEnvironmentsGetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// CustomProviderId - Custom provider this environment belongs to.
	CustomProviderId string `json:"custom_provider_id"`
	// ProviderId - Associated provider, by ID. `null` when there is none.
	ProviderId *string `json:"provider_id,omitempty"`
	// CurrentProviderVersionId - Version currently live in this environment. `null` when no version is live.
	CurrentProviderVersionId *string `json:"current_provider_version_id,omitempty"`
	// InstanceId - ID of the instance this environment deploys into.
	InstanceId string `json:"instance_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapCustomProvidersEnvironmentsGetOutputFromJSON deserializes JSON data into a CustomProvidersEnvironmentsGetOutput.
func MapCustomProvidersEnvironmentsGetOutputFromJSON(data []byte) (*CustomProvidersEnvironmentsGetOutput, error) {
	var v CustomProvidersEnvironmentsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCustomProvidersEnvironmentsGetOutputToJSON serializes a CustomProvidersEnvironmentsGetOutput to JSON.
func MapCustomProvidersEnvironmentsGetOutputToJSON(v *CustomProvidersEnvironmentsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
