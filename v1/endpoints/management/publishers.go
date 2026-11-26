package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/publishers"
)

// PublishersEndpoint provides access to a publisher is the organization or individual who created and maintains a provider.
type PublishersEndpoint struct {
	client *endpoint.Client
}

// NewPublishersEndpoint creates a new PublishersEndpoint.
func NewPublishersEndpoint(client *endpoint.Client) *PublishersEndpoint {
	return &PublishersEndpoint{client: client}
}

// PublishersEndpointListParams contains optional query parameters for List.
type PublishersEndpointListParams struct {
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

// List returns a paginated list of publishers.
func (e *PublishersEndpoint) List(instanceId string, params *PublishersEndpointListParams) (*publishers.PublishersListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "publishers"},
		Query: query,
	}
	var result publishers.PublishersListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific publisher by ID.
func (e *PublishersEndpoint) Get(instanceId string, publisherId string) (*publishers.PublishersGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "publishers", publisherId},
	}
	var result publishers.PublishersGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
