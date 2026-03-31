package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/consumers"
)

// ConsumersEndpoint provides access to manage instance consumers independently from portals and inspect the profiles linked to each consumer.
type ConsumersEndpoint struct {
	client *endpoint.Client
}

// NewConsumersEndpoint creates a new ConsumersEndpoint.
func NewConsumersEndpoint(client *endpoint.Client) *ConsumersEndpoint {
	return &ConsumersEndpoint{client: client}
}

// ConsumersEndpointListParams contains optional query parameters for List.
type ConsumersEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
}

// ConsumersEndpointCreateBody contains the request body for Create.
type ConsumersEndpointCreateBody struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ConsumersEndpointUpdateBody contains the request body for Update.
type ConsumersEndpointUpdateBody struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

// List returns a paginated list of consumers for an instance.
func (e *ConsumersEndpoint) List(params *ConsumersEndpointListParams) (*consumers.ConsumersListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"consumers"},
		Query: query,
	}
	var result consumers.ConsumersListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a consumer by ID.
func (e *ConsumersEndpoint) Get(consumerId string) (*consumers.ConsumersGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"consumers", consumerId},
	}
	var result consumers.ConsumersGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates or links a consumer for an instance.
func (e *ConsumersEndpoint) Create(body *ConsumersEndpointCreateBody) (*consumers.ConsumersCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"consumers"},
		Body: body,
	}
	var result consumers.ConsumersCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a consumer for an instance.
func (e *ConsumersEndpoint) Update(consumerId string, body *ConsumersEndpointUpdateBody) (*consumers.ConsumersUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"consumers", consumerId},
		Body: body,
	}
	var result consumers.ConsumersUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
