package items

import (
	"encoding/json"
	"time"
)

// SkillsItemsListOutputItemsIntegrationConfiguration - Which settings integration instances may customize.
type SkillsItemsListOutputItemsIntegrationConfiguration struct {
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

// SkillsItemsListOutputItemsIntegration represents the skills items list output items integration type.
type SkillsItemsListOutputItemsIntegration struct {
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
	Configuration SkillsItemsListOutputItemsIntegrationConfiguration `json:"configuration"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// SkillsItemsListOutputItemsProvider represents the skills items list output items provider type.
type SkillsItemsListOutputItemsProvider struct {
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

// SkillsItemsListOutputItems represents the skills items list output items type.
type SkillsItemsListOutputItems struct {
	Object      string                                 `json:"object"`
	Id          string                                 `json:"id"`
	Status      string                                 `json:"status"`
	Type        string                                 `json:"type"`
	SkillId     string                                 `json:"skill_id"`
	Integration *SkillsItemsListOutputItemsIntegration `json:"integration,omitempty"`
	Provider    *SkillsItemsListOutputItemsProvider    `json:"provider,omitempty"`
	CreatedAt   time.Time                              `json:"created_at"`
}

// SkillsItemsListOutputPagination represents the skills items list output pagination type.
type SkillsItemsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SkillsItemsListOutput represents the skills items list output type.
type SkillsItemsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                          `json:"object"`
	Items      []SkillsItemsListOutputItems    `json:"items"`
	Pagination SkillsItemsListOutputPagination `json:"pagination"`
}

// MapSkillsItemsListOutputFromJSON deserializes JSON data into a SkillsItemsListOutput.
func MapSkillsItemsListOutputFromJSON(data []byte) (*SkillsItemsListOutput, error) {
	var v SkillsItemsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsItemsListOutputToJSON serializes a SkillsItemsListOutput to JSON.
func MapSkillsItemsListOutputToJSON(v *SkillsItemsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsItemsListQueryCreatedAt - Only return results whose skill item creation time is in this range. Set `gt`, `lt`, or both.
type SkillsItemsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsItemsListQuery represents the skills items list query type.
type SkillsItemsListQuery struct {
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
	// Status - Only return skill items with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Type - Only return skill items of one of these types.
	Type *any `json:"type,omitempty"`
	// Id - Only return skill items with these IDs.
	Id *any `json:"id,omitempty"`
	// IntegrationId - Only return skill items for one of these integrations.
	IntegrationId *any `json:"integration_id,omitempty"`
	// ProviderId - Only return skill items for one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// CreatedAt - Only return results whose skill item creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *SkillsItemsListQueryCreatedAt `json:"created_at,omitempty"`
}

// MapSkillsItemsListQueryFromJSON deserializes JSON data into a SkillsItemsListQuery.
func MapSkillsItemsListQueryFromJSON(data []byte) (*SkillsItemsListQuery, error) {
	var v SkillsItemsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsItemsListQueryToJSON serializes a SkillsItemsListQuery to JSON.
func MapSkillsItemsListQueryToJSON(v *SkillsItemsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
