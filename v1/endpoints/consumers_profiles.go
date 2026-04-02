package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/consumers/profiles"
)

// ConsumersProfilesEndpoint provides access to manage instance consumers independently from portals and inspect the profiles linked to each consumer.
type ConsumersProfilesEndpoint struct {
	client *endpoint.Client
}

// NewConsumersProfilesEndpoint creates a new ConsumersProfilesEndpoint.
func NewConsumersProfilesEndpoint(client *endpoint.Client) *ConsumersProfilesEndpoint {
	return &ConsumersProfilesEndpoint{client: client}
}

// ConsumersProfilesEndpointListParams contains optional query parameters for List.
type ConsumersProfilesEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
}

// List returns a paginated list of profiles for a consumer in an instance.
func (e *ConsumersProfilesEndpoint) List(consumerId string, params *ConsumersProfilesEndpointListParams) (*profiles.ConsumersProfilesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"consumers", consumerId, "profiles"},
		Query: query,
	}
	var result profiles.ConsumersProfilesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a consumer profile by ID for a consumer.
func (e *ConsumersProfilesEndpoint) Get(consumerId string, consumerProfileId string) (*profiles.ConsumersProfilesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"consumers", consumerId, "profiles", consumerProfileId},
	}
	var result profiles.ConsumersProfilesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
