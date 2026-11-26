package providerruns

import (
	"encoding/json"
	"time"
)

// ProviderRunsGetOutput represents the provider runs get output type.
type ProviderRunsGetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `running` while the provider handles messages. `stopped` once it finishes.
	Status string `json:"status"`
	// SessionId - Session the run belongs to.
	SessionId string `json:"session_id"`
	// SessionProviderId - Session provider that handled the run.
	SessionProviderId string `json:"session_provider_id"`
	// ProviderId - Provider that ran.
	ProviderId string `json:"provider_id"`
	// ConnectionId - Connection the run used.
	ConnectionId string `json:"connection_id"`
	// CompletedAt - When the run finished. `null` while it is still running.
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapProviderRunsGetOutputFromJSON deserializes JSON data into a ProviderRunsGetOutput.
func MapProviderRunsGetOutputFromJSON(data []byte) (*ProviderRunsGetOutput, error) {
	var v ProviderRunsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderRunsGetOutputToJSON serializes a ProviderRunsGetOutput to JSON.
func MapProviderRunsGetOutputToJSON(v *ProviderRunsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
