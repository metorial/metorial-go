package callbacks

import (
	"encoding/json"
	"time"
)

// CallbacksGetOutputProvider represents the callbacks get output provider type.
type CallbacksGetOutputProvider struct {
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

// CallbacksGetOutput represents the callbacks get output type.
type CallbacksGetOutput struct {
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
	IntegrationProviderId string                     `json:"integration_provider_id"`
	Provider              CallbacksGetOutputProvider `json:"provider"`
	// CreatedAt - Timestamp when the callback was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the callback was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapCallbacksGetOutputFromJSON deserializes JSON data into a CallbacksGetOutput.
func MapCallbacksGetOutputFromJSON(data []byte) (*CallbacksGetOutput, error) {
	var v CallbacksGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCallbacksGetOutputToJSON serializes a CallbacksGetOutput to JSON.
func MapCallbacksGetOutputToJSON(v *CallbacksGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
