package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/consumers/devices"
)

// ConsumersDevicesEndpoint provides access to a consumer device is a Metorial Desktop installation a consumer has enrolled. Revoking a device signs it out and removes its access.
type ConsumersDevicesEndpoint struct {
	client *endpoint.Client
}

// NewConsumersDevicesEndpoint creates a new ConsumersDevicesEndpoint.
func NewConsumersDevicesEndpoint(client *endpoint.Client) *ConsumersDevicesEndpoint {
	return &ConsumersDevicesEndpoint{client: client}
}

// ConsumersDevicesEndpointListParams contains optional query parameters for List.
type ConsumersDevicesEndpointListParams struct {
	// Limit - Page size, from 1 to 100. Defaults to 20.
	Limit *float64 `json:"limit,omitempty"`
	// After - Return the page after this point. Pass `pagination.after` if the previous response has it, otherwise the `id` of its last item.
	After *string `json:"after,omitempty"`
	// Before - Return the page before this point. Pass `pagination.before` if the previous response has it, otherwise the `id` of its first item.
	Before *string `json:"before,omitempty"`
	// Cursor - A `cur_` cursor from `pagination.after` or `pagination.before`. Alternative to `after` and `before`.
	Cursor *string `json:"cursor,omitempty"`
	// Order - Sort direction.
	Order *string `json:"order,omitempty"`
	// ConsumerProfileId - Only return devices enrolled by these consumer profiles.
	ConsumerProfileId *any `json:"consumer_profile_id,omitempty"`
	// ConsumerId - Only return devices enrolled by any profile of these consumers.
	ConsumerId *any `json:"consumer_id,omitempty"`
	// Status - Only return devices with these statuses. Defaults to active.
	Status *any `json:"status,omitempty"`
}

// List returns a paginated list of consumer devices.
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

// Get retrieves a consumer device.
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
