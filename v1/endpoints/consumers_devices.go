package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/consumers/devices"
)

// ConsumersDevicesEndpoint provides access to inspect and revoke the Metorial Desktop devices consumers have enrolled.
type ConsumersDevicesEndpoint struct {
	client *endpoint.Client
}

// NewConsumersDevicesEndpoint creates a new ConsumersDevicesEndpoint.
func NewConsumersDevicesEndpoint(client *endpoint.Client) *ConsumersDevicesEndpoint {
	return &ConsumersDevicesEndpoint{client: client}
}

// ConsumersDevicesEndpointListParams contains optional query parameters for List.
type ConsumersDevicesEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// ConsumerProfileId - Only return devices enrolled by these consumer profiles.
	ConsumerProfileId *any `json:"consumer_profile_id,omitempty"`
	// ConsumerId - Only return devices enrolled by any profile of these consumers.
	ConsumerId *any `json:"consumer_id,omitempty"`
	// Status - Only return devices with these statuses. Defaults to active.
	Status *any `json:"status,omitempty"`
}

// List returns the Metorial Desktop devices enrolled by consumers of an instance.
func (e *ConsumersDevicesEndpoint) List(params *ConsumersDevicesEndpointListParams) (*devices.ConsumersDevicesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"consumer-devices"},
		Query: query,
	}
	var result devices.ConsumersDevicesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a Metorial Desktop device enrolled by a consumer.
func (e *ConsumersDevicesEndpoint) Get(consumerDeviceId string) (*devices.ConsumersDevicesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"consumer-devices", consumerDeviceId},
	}
	var result devices.ConsumersDevicesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Revoke signs the device out and revokes its integration tokens. The app removes what it installed the next time it reaches Metorial.
func (e *ConsumersDevicesEndpoint) Revoke(consumerDeviceId string) (*devices.ConsumersDevicesRevokeOutput, error) {
	req := &endpoint.Request{
		Path: []string{"consumer-devices", consumerDeviceId, "revoke"},
	}
	var result devices.ConsumersDevicesRevokeOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
