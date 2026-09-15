package callbacks

import (
	"encoding/json"
	"time"
)

// CallbacksDeleteOutputProvider represents the callbacks delete output provider type.
type CallbacksDeleteOutputProvider struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique provider identifier
	Id string `json:"id"`
	// Name - Display name of the provider
	Name string `json:"name"`
	// Description - Brief description of the provider
	Description *string `json:"description,omitempty"`
	// Slug - URL-friendly identifier
	Slug string `json:"slug"`
	// CreatedAt - Timestamp when the provider was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the provider was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// CallbacksDeleteOutput represents the callbacks delete output type.
type CallbacksDeleteOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique callback identifier
	Id string `json:"id"`
	// Status - Callback lifecycle status. Archived once callbacks are disabled on the integration provider.
	Status string `json:"status"`
	// Name - Display name for the callback
	Name string `json:"name"`
	// Description - Optional callback description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional callback metadata
	Metadata *map[string]any `json:"metadata,omitempty"`
	// IntegrationId - Integration this callback belongs to
	IntegrationId string `json:"integration_id"`
	// IntegrationProviderId - Integration provider this callback was created for
	IntegrationProviderId string                        `json:"integration_provider_id"`
	Provider              CallbacksDeleteOutputProvider `json:"provider"`
	// CreatedAt - Timestamp when the callback was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the callback was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapCallbacksDeleteOutputFromJSON deserializes JSON data into a CallbacksDeleteOutput.
func MapCallbacksDeleteOutputFromJSON(data []byte) (*CallbacksDeleteOutput, error) {
	var v CallbacksDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCallbacksDeleteOutputToJSON serializes a CallbacksDeleteOutput to JSON.
func MapCallbacksDeleteOutputToJSON(v *CallbacksDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
