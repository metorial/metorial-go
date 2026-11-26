package items

import (
	"encoding/json"
	"time"
)

// SkillsTemplatesItemsListOutputItemsIntegrationConfiguration - Which settings integration instances may customize.
type SkillsTemplatesItemsListOutputItemsIntegrationConfiguration struct {
	// CanAttachCustomToolFilters - Lets integration instances attach their own tool filters.
	CanAttachCustomToolFilters bool `json:"can_attach_custom_tool_filters"`
	// CanAttachCustomProviderConfig - Lets integration instances use their own provider configs.
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	// CanOverrideToolFilters - Lets integration instances ignore the integration tool filters.
	CanOverrideToolFilters bool `json:"can_override_tool_filters"`
	// EnableCallbackTools - Makes callback event tools available for this integration.
	EnableCallbackTools bool `json:"enable_callback_tools"`
	// UseIntegrationNameInToolNames - Names session providers after the integration instead of the provider. `null` uses the default.
	UseIntegrationNameInToolNames *bool `json:"use_integration_name_in_tool_names,omitempty"`
}

// SkillsTemplatesItemsListOutputItemsIntegration represents the skills templates items list output items integration type.
type SkillsTemplatesItemsListOutputItemsIntegration struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Slug - URL-safe name, unique across Metorial.
	Slug string `json:"slug"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Configuration - Which settings integration instances may customize.
	Configuration SkillsTemplatesItemsListOutputItemsIntegrationConfiguration `json:"configuration"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// SkillsTemplatesItemsListOutputItemsProvider represents the skills templates items list output items provider type.
type SkillsTemplatesItemsListOutputItemsProvider struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Slug - URL-safe name for the provider.
	Slug string `json:"slug"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsTemplatesItemsListOutputItems represents the skills templates items list output items type.
type SkillsTemplatesItemsListOutputItems struct {
	Object      string                                          `json:"object"`
	Id          string                                          `json:"id"`
	Type        string                                          `json:"type"`
	Integration *SkillsTemplatesItemsListOutputItemsIntegration `json:"integration,omitempty"`
	Provider    *SkillsTemplatesItemsListOutputItemsProvider    `json:"provider,omitempty"`
	CreatedAt   time.Time                                       `json:"created_at"`
	UpdatedAt   time.Time                                       `json:"updated_at"`
}

// SkillsTemplatesItemsListOutputPagination represents the skills templates items list output pagination type.
type SkillsTemplatesItemsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SkillsTemplatesItemsListOutput represents the skills templates items list output type.
type SkillsTemplatesItemsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                   `json:"object"`
	Items      []SkillsTemplatesItemsListOutputItems    `json:"items"`
	Pagination SkillsTemplatesItemsListOutputPagination `json:"pagination"`
}

// MapSkillsTemplatesItemsListOutputFromJSON deserializes JSON data into a SkillsTemplatesItemsListOutput.
func MapSkillsTemplatesItemsListOutputFromJSON(data []byte) (*SkillsTemplatesItemsListOutput, error) {
	var v SkillsTemplatesItemsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsTemplatesItemsListOutputToJSON serializes a SkillsTemplatesItemsListOutput to JSON.
func MapSkillsTemplatesItemsListOutputToJSON(v *SkillsTemplatesItemsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsTemplatesItemsListQuery represents the skills templates items list query type.
type SkillsTemplatesItemsListQuery struct {
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

// MapSkillsTemplatesItemsListQueryFromJSON deserializes JSON data into a SkillsTemplatesItemsListQuery.
func MapSkillsTemplatesItemsListQueryFromJSON(data []byte) (*SkillsTemplatesItemsListQuery, error) {
	var v SkillsTemplatesItemsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsTemplatesItemsListQueryToJSON serializes a SkillsTemplatesItemsListQuery to JSON.
func MapSkillsTemplatesItemsListQueryToJSON(v *SkillsTemplatesItemsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
