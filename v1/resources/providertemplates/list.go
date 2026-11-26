package providertemplates

import (
	"encoding/json"
	"time"
)

// ProviderTemplatesListOutputItems represents the provider templates list output items type.
type ProviderTemplatesListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once you archive the template.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata map[string]any `json:"metadata"`
	// IntegrationId - Integration the template was created from. `null` when it wasn’t created from one.
	IntegrationId *string `json:"integration_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderTemplatesListOutputPagination represents the provider templates list output pagination type.
type ProviderTemplatesListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ProviderTemplatesListOutput represents the provider templates list output type.
type ProviderTemplatesListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                `json:"object"`
	Items      []ProviderTemplatesListOutputItems    `json:"items"`
	Pagination ProviderTemplatesListOutputPagination `json:"pagination"`
}

// MapProviderTemplatesListOutputFromJSON deserializes JSON data into a ProviderTemplatesListOutput.
func MapProviderTemplatesListOutputFromJSON(data []byte) (*ProviderTemplatesListOutput, error) {
	var v ProviderTemplatesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderTemplatesListOutputToJSON serializes a ProviderTemplatesListOutput to JSON.
func MapProviderTemplatesListOutputToJSON(v *ProviderTemplatesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderTemplatesListQuery represents the provider templates list query type.
type ProviderTemplatesListQuery struct {
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

// MapProviderTemplatesListQueryFromJSON deserializes JSON data into a ProviderTemplatesListQuery.
func MapProviderTemplatesListQueryFromJSON(data []byte) (*ProviderTemplatesListQuery, error) {
	var v ProviderTemplatesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderTemplatesListQueryToJSON serializes a ProviderTemplatesListQuery to JSON.
func MapProviderTemplatesListQueryToJSON(v *ProviderTemplatesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
