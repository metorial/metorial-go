package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/consumers"
)

// ConsumersEndpoint provides access to a consumer is a person or system that uses an instance. Each consumer has profiles linking them to the surfaces they can access.
type ConsumersEndpoint struct {
	client *endpoint.Client
}

// NewConsumersEndpoint creates a new ConsumersEndpoint.
func NewConsumersEndpoint(client *endpoint.Client) *ConsumersEndpoint {
	return &ConsumersEndpoint{client: client}
}

// ConsumersEndpointListParams contains optional query parameters for List.
type ConsumersEndpointListParams struct {
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
	// Search - Search by name or email.
	Search *string `json:"search,omitempty"`
	// Email - Only return consumers with one of these emails.
	Email *any `json:"email,omitempty"`
	// Id - Only return consumers with one of these IDs.
	Id *any `json:"id,omitempty"`
}

// ConsumersEndpointCreateBody contains the request body for Create.
type ConsumersEndpointCreateBody struct {
	// Name - Display name.
	Name string `json:"name"`
	// Email - Email address. An existing consumer with this email is linked to the instance.
	Email string `json:"email"`
}

// ConsumersEndpointGetMemberConsumerBody contains the request body for GetMemberConsumer.
type ConsumersEndpointGetMemberConsumerBody struct {
	// SurfaceIdentifier - Surface to create the profile on.
	SurfaceIdentifier *string `json:"surface_identifier,omitempty"`
}

// ConsumersEndpointCreateMemberSessionBody contains the request body for CreateMemberSession.
type ConsumersEndpointCreateMemberSessionBody struct {
	// IntegrationId - Integration to grant the member's consumer profile access to.
	IntegrationId *string `json:"integration_id,omitempty"`
}

// ConsumersEndpointUpdateBody contains the request body for Update.
type ConsumersEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Email - Email address. Must be unique within the organization.
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

// Get retrieves a consumer.
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

// CreateMemberSession exchanges a CLI sign-in for a short-lived credential for the member’s own consumer profile.
func (e *ConsumersEndpoint) CreateMemberSession(instanceId string, body *ConsumersEndpointCreateMemberSessionBody) (*consumers.ConsumersCreateMemberSessionOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "member-consumer-session"},
		Body: body,
	}
	var result consumers.ConsumersCreateMemberSessionOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a consumer. Omitted fields keep their current value.
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
