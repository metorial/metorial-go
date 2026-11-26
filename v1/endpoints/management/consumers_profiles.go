package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/consumers/profiles"
)

// ConsumersProfilesEndpoint provides access to a consumer is a person or system that uses an instance. Each consumer has profiles linking them to the surfaces they can access.
type ConsumersProfilesEndpoint struct {
	client *endpoint.Client
}

// NewConsumersProfilesEndpoint creates a new ConsumersProfilesEndpoint.
func NewConsumersProfilesEndpoint(client *endpoint.Client) *ConsumersProfilesEndpoint {
	return &ConsumersProfilesEndpoint{client: client}
}

// ConsumersProfilesEndpointListParams contains optional query parameters for List.
type ConsumersProfilesEndpointListParams struct {
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
}

// List returns a paginated list of profiles for a consumer.
func (e *ConsumersProfilesEndpoint) List(instanceId string, consumerId string, params *ConsumersProfilesEndpointListParams) (*profiles.ConsumersProfilesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "consumers", consumerId, "profiles"},
		Query: query,
	}
	var result profiles.ConsumersProfilesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a consumer profile.
func (e *ConsumersProfilesEndpoint) Get(instanceId string, consumerId string, consumerProfileId string) (*profiles.ConsumersProfilesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "consumers", consumerId, "profiles", consumerProfileId},
	}
	var result profiles.ConsumersProfilesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
