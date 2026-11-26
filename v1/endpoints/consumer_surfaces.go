package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/consumersurfaces"
)

// ConsumerSurfacesEndpoint provides access to a consumer surface is a portal, member directory, or CLI entry point consumers sign in through.
type ConsumerSurfacesEndpoint struct {
	client *endpoint.Client
}

// NewConsumerSurfacesEndpoint creates a new ConsumerSurfacesEndpoint.
func NewConsumerSurfacesEndpoint(client *endpoint.Client) *ConsumerSurfacesEndpoint {
	return &ConsumerSurfacesEndpoint{client: client}
}

// ConsumerSurfacesEndpointListParams contains optional query parameters for List.
type ConsumerSurfacesEndpointListParams struct {
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

// List returns a paginated list of consumer surfaces.
func (e *ConsumerSurfacesEndpoint) List(params *ConsumerSurfacesEndpointListParams) (*consumersurfaces.ConsumerSurfacesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"consumer-surfaces"},
		Query: query,
	}
	var result consumersurfaces.ConsumerSurfacesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a consumer surface.
func (e *ConsumerSurfacesEndpoint) Get(consumerSurfaceId string) (*consumersurfaces.ConsumerSurfacesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"consumer-surfaces", consumerSurfaceId},
	}
	var result consumersurfaces.ConsumerSurfacesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
