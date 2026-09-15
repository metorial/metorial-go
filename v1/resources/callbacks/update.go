package callbacks

import (
	"encoding/json"
	"time"
)

// CallbacksUpdateOutputProvider represents the callbacks update output provider type.
type CallbacksUpdateOutputProvider struct {
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

// CallbacksUpdateOutput represents the callbacks update output type.
type CallbacksUpdateOutput struct {
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
	Provider              CallbacksUpdateOutputProvider `json:"provider"`
	// CreatedAt - Timestamp when the callback was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the callback was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapCallbacksUpdateOutputFromJSON deserializes JSON data into a CallbacksUpdateOutput.
func MapCallbacksUpdateOutputFromJSON(data []byte) (*CallbacksUpdateOutput, error) {
	var v CallbacksUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCallbacksUpdateOutputToJSON serializes a CallbacksUpdateOutput to JSON.
func MapCallbacksUpdateOutputToJSON(v *CallbacksUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// CallbacksUpdateBody represents the callbacks update body type.
type CallbacksUpdateBody struct {
	// Name - Updated display name
	Name *string `json:"name,omitempty"`
	// Description - Updated description
	Description *string `json:"description,omitempty"`
	// Metadata - Updated custom metadata
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// MapCallbacksUpdateBodyFromJSON deserializes JSON data into a CallbacksUpdateBody.
func MapCallbacksUpdateBodyFromJSON(data []byte) (*CallbacksUpdateBody, error) {
	var v CallbacksUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCallbacksUpdateBodyToJSON serializes a CallbacksUpdateBody to JSON.
func MapCallbacksUpdateBodyToJSON(v *CallbacksUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
