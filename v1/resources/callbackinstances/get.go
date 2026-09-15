package callbackinstances

import (
	"encoding/json"
	"time"
)

// CallbackInstancesGetOutput represents the callback instances get output type.
type CallbackInstancesGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique callback instance identifier
	Id string `json:"id"`
	// Status - Callback instance lifecycle status. Archived once the callback or the integration instance provider goes away.
	Status string `json:"status"`
	// CallbackId - Callback this instance belongs to
	CallbackId string `json:"callback_id"`
	// IntegrationInstanceId - Integration instance this callback instance listens for
	IntegrationInstanceId string `json:"integration_instance_id"`
	// IntegrationInstanceProviderId - Integration instance provider this callback instance is registered on
	IntegrationInstanceProviderId string `json:"integration_instance_provider_id"`
	// CreatedAt - Timestamp when the callback instance was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the callback instance was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapCallbackInstancesGetOutputFromJSON deserializes JSON data into a CallbackInstancesGetOutput.
func MapCallbackInstancesGetOutputFromJSON(data []byte) (*CallbackInstancesGetOutput, error) {
	var v CallbackInstancesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCallbackInstancesGetOutputToJSON serializes a CallbackInstancesGetOutput to JSON.
func MapCallbackInstancesGetOutputToJSON(v *CallbackInstancesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
