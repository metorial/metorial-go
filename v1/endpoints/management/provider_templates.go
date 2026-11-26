package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/providertemplates"
)

// ProviderTemplatesEndpoint provides access to a provider template is a consumer-facing view of an integration. It is created from an existing integration.
type ProviderTemplatesEndpoint struct {
	client *endpoint.Client
}

// NewProviderTemplatesEndpoint creates a new ProviderTemplatesEndpoint.
func NewProviderTemplatesEndpoint(client *endpoint.Client) *ProviderTemplatesEndpoint {
	return &ProviderTemplatesEndpoint{client: client}
}

// ProviderTemplatesEndpointListParams contains optional query parameters for List.
type ProviderTemplatesEndpointListParams struct {
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
	// Id - Only return provider templates with one of these IDs.
	Id *any `json:"id,omitempty"`
	// IntegrationId - Only return provider templates for one of these integrations.
	IntegrationId *any `json:"integration_id,omitempty"`
	// Search - Search provider templates by name or description.
	Search *string `json:"search,omitempty"`
	// Status - Only return provider templates with one of these statuses. `deleted` templates are never listed.
	Status *any `json:"status,omitempty"`
}

// ProviderTemplatesEndpointCreateBody contains the request body for Create.
type ProviderTemplatesEndpointCreateBody struct {
	// Name - Display name for the provider template.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the provider template.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// IntegrationId - Integration the template is created from.
	IntegrationId string `json:"integration_id"`
}

// ProviderTemplatesEndpointUpdateBody contains the request body for Update.
type ProviderTemplatesEndpointUpdateBody struct {
	// Name - Display name for the provider template.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the provider template. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// List returns a paginated list of provider templates.
func (e *ProviderTemplatesEndpoint) List(instanceId string, params *ProviderTemplatesEndpointListParams) (*providertemplates.ProviderTemplatesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "provider-templates"},
		Query: query,
	}
	var result providertemplates.ProviderTemplatesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a provider template.
func (e *ProviderTemplatesEndpoint) Get(instanceId string, providerTemplateId string) (*providertemplates.ProviderTemplatesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "provider-templates", providerTemplateId},
	}
	var result providertemplates.ProviderTemplatesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a provider template from an existing integration.
func (e *ProviderTemplatesEndpoint) Create(instanceId string, body *ProviderTemplatesEndpointCreateBody) (*providertemplates.ProviderTemplatesCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "provider-templates"},
		Body: body,
	}
	var result providertemplates.ProviderTemplatesCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a provider template. Omitted fields keep their current value.
func (e *ProviderTemplatesEndpoint) Update(instanceId string, providerTemplateId string, body *ProviderTemplatesEndpointUpdateBody) (*providertemplates.ProviderTemplatesUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "provider-templates", providerTemplateId},
		Body: body,
	}
	var result providertemplates.ProviderTemplatesUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a provider template.
func (e *ProviderTemplatesEndpoint) Delete(instanceId string, providerTemplateId string) (*providertemplates.ProviderTemplatesDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "provider-templates", providerTemplateId},
	}
	var result providertemplates.ProviderTemplatesDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
