package management

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

// ConsumersEndpointGetMemberConsumerBody contains the request body for GetMemberConsumer.
type ConsumersEndpointGetMemberConsumerBody struct {
	SurfaceIdentifier *string `json:"surface_identifier,omitempty"`
}

// ConsumersEndpointUpdateBody contains the request body for Update.
type ConsumersEndpointUpdateBody struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

// List returns a paginated list of consumers for an instance.
func (e *ConsumersEndpoint) List(instanceId string, params *ConsumersEndpointListParams) (*consumers.ConsumersListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "consumers"},
		Query: query,
	}
	var result consumers.ConsumersListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a consumer by ID.
func (e *ConsumersEndpoint) Get(instanceId string, consumerId string) (*consumers.ConsumersGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "consumers", consumerId},
	}
	var result consumers.ConsumersGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates or links a consumer for an instance.
func (e *ConsumersEndpoint) Create(instanceId string, body *ConsumersEndpointCreateBody) (*consumers.ConsumersCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "consumers"},
		Body: body,
	}
	var result consumers.ConsumersCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMemberConsumer upserts and returns the consumer for the authenticated organization member.
func (e *ConsumersEndpoint) GetMemberConsumer(instanceId string, body *ConsumersEndpointGetMemberConsumerBody) (*consumers.ConsumersGetMemberConsumerOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "get-member-consumer"},
		Body: body,
	}
	var result consumers.ConsumersGetMemberConsumerOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a consumer for an instance.
func (e *ConsumersEndpoint) Update(instanceId string, consumerId string, body *ConsumersEndpointUpdateBody) (*consumers.ConsumersUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "consumers", consumerId},
		Body: body,
	}
	var result consumers.ConsumersUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
