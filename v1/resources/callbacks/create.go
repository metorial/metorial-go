package callbacks

import (
	"encoding/json"
	"time"
)

// CallbacksCreateOutputProvider represents the callbacks create output provider type.
type CallbacksCreateOutputProvider struct {
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

// CallbacksCreateOutputTriggerGroups - A trigger group that accepts a selection
type CallbacksCreateOutputTriggerGroups struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Key - Trigger group key to use in `trigger_group_selections`
	Key string `json:"key"`
	// Name - Display name of the trigger group
	Name string `json:"name"`
	// Description - What the trigger group delivers
	Description *string `json:"description,omitempty"`
	// SelectionSchema - JSON Schema the selection for this trigger group must match
	SelectionSchema map[string]any `json:"selection_schema"`
	// Status - `selection_required` when the trigger group needs a selection before it delivers events.
	Status string `json:"status"`
}

// CallbacksCreateOutput represents the callbacks create output type.
type CallbacksCreateOutput struct {
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
	Provider              CallbacksCreateOutputProvider `json:"provider"`
	// TriggerGroupSelections - Selections for trigger groups that need to know what to watch, keyed by trigger group key
	TriggerGroupSelections *map[string]map[string]any `json:"trigger_group_selections,omitempty"`
	// TriggerGroups - Trigger groups that accept a selection, with their schema and whether one is still required
	TriggerGroups []CallbacksCreateOutputTriggerGroups `json:"trigger_groups"`
	// CreatedAt - Timestamp when the callback was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the callback was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapCallbacksCreateOutputFromJSON deserializes JSON data into a CallbacksCreateOutput.
func MapCallbacksCreateOutputFromJSON(data []byte) (*CallbacksCreateOutput, error) {
	var v CallbacksCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCallbacksCreateOutputToJSON serializes a CallbacksCreateOutput to JSON.
func MapCallbacksCreateOutputToJSON(v *CallbacksCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// CallbacksCreateBody represents the callbacks create body type.
type CallbacksCreateBody struct {
	// IntegrationId - Integration the integration provider belongs to
	IntegrationId string `json:"integration_id"`
	// IntegrationProviderId - Integration provider to enable callbacks for
	IntegrationProviderId string `json:"integration_provider_id"`
	// Name - Display name for the callback. Defaults to the name of the integration provider.
	Name *string `json:"name,omitempty"`
	// Description - Description for the callback. Defaults to the description of the integration provider.
	Description *string `json:"description,omitempty"`
	// TriggerGroupSelections - Resource selections keyed by trigger group key. Each value must match that group’s selection schema. Limited to 64 groups and 64 KiB of JSON.
	TriggerGroupSelections *map[string]map[string]any `json:"trigger_group_selections,omitempty"`
}

// MapCallbacksCreateBodyFromJSON deserializes JSON data into a CallbacksCreateBody.
func MapCallbacksCreateBodyFromJSON(data []byte) (*CallbacksCreateBody, error) {
	var v CallbacksCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCallbacksCreateBodyToJSON serializes a CallbacksCreateBody to JSON.
func MapCallbacksCreateBodyToJSON(v *CallbacksCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
