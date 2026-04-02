package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/consumersurfaces"
)

// ConsumerSurfacesEndpoint provides access to list and retrieve consumer surfaces for an instance.
type ConsumerSurfacesEndpoint struct {
	client *endpoint.Client
}

// NewConsumerSurfacesEndpoint creates a new ConsumerSurfacesEndpoint.
func NewConsumerSurfacesEndpoint(client *endpoint.Client) *ConsumerSurfacesEndpoint {
	return &ConsumerSurfacesEndpoint{client: client}
}

// ConsumerSurfacesEndpointListParams contains optional query parameters for List.
type ConsumerSurfacesEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
}

// List returns a paginated list of consumer surfaces for an instance.
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

// Get retrieves a consumer surface by ID.
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
