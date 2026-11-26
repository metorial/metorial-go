package templates

import (
	"encoding/json"
	"time"
)

// SkillsTemplatesListOutputItemsItemsIntegrationConfiguration - Which settings integration instances may customize.
type SkillsTemplatesListOutputItemsItemsIntegrationConfiguration struct {
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

// SkillsTemplatesListOutputItemsItemsIntegration represents the skills templates list output items items integration type.
type SkillsTemplatesListOutputItemsItemsIntegration struct {
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
	Configuration SkillsTemplatesListOutputItemsItemsIntegrationConfiguration `json:"configuration"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// SkillsTemplatesListOutputItemsItemsProvider represents the skills templates list output items items provider type.
type SkillsTemplatesListOutputItemsItemsProvider struct {
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

// SkillsTemplatesListOutputItemsItems represents the skills templates list output items items type.
type SkillsTemplatesListOutputItemsItems struct {
	Object      string                                          `json:"object"`
	Id          string                                          `json:"id"`
	Type        string                                          `json:"type"`
	Integration *SkillsTemplatesListOutputItemsItemsIntegration `json:"integration,omitempty"`
	Provider    *SkillsTemplatesListOutputItemsItemsProvider    `json:"provider,omitempty"`
	CreatedAt   time.Time                                       `json:"created_at"`
	UpdatedAt   time.Time                                       `json:"updated_at"`
}

// SkillsTemplatesListOutputItems represents the skills templates list output items type.
type SkillsTemplatesListOutputItems struct {
	Object      string                                `json:"object"`
	Id          string                                `json:"id"`
	Status      string                                `json:"status"`
	Owner       string                                `json:"owner"`
	Slug        string                                `json:"slug"`
	Name        string                                `json:"name"`
	Description *string                               `json:"description,omitempty"`
	Metadata    map[string]any                        `json:"metadata"`
	StoreId     string                                `json:"store_id"`
	Items       []SkillsTemplatesListOutputItemsItems `json:"items"`
	CreatedAt   time.Time                             `json:"created_at"`
	UpdatedAt   time.Time                             `json:"updated_at"`
}

// SkillsTemplatesListOutputPagination represents the skills templates list output pagination type.
type SkillsTemplatesListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SkillsTemplatesListOutput represents the skills templates list output type.
type SkillsTemplatesListOutput struct {
	// Object - String representing the object’s type.
	Object     string                              `json:"object"`
	Items      []SkillsTemplatesListOutputItems    `json:"items"`
	Pagination SkillsTemplatesListOutputPagination `json:"pagination"`
}

// MapSkillsTemplatesListOutputFromJSON deserializes JSON data into a SkillsTemplatesListOutput.
func MapSkillsTemplatesListOutputFromJSON(data []byte) (*SkillsTemplatesListOutput, error) {
	var v SkillsTemplatesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsTemplatesListOutputToJSON serializes a SkillsTemplatesListOutput to JSON.
func MapSkillsTemplatesListOutputToJSON(v *SkillsTemplatesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsTemplatesListQueryCreatedAt - Only return results whose skill template creation time is in this range. Set `gt`, `lt`, or both.
type SkillsTemplatesListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsTemplatesListQueryUpdatedAt - Only return results whose skill template last update time is in this range. Set `gt`, `lt`, or both.
type SkillsTemplatesListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsTemplatesListQuery represents the skills templates list query type.
type SkillsTemplatesListQuery struct {
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
	// Search - Only return skill templates matching this search text.
	Search *string `json:"search,omitempty"`
	// Status - Only return skill templates with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Owner - Only return skill templates with one of these owners.
	Owner *any `json:"owner,omitempty"`
	// Id - Only return skill templates with these IDs.
	Id *any `json:"id,omitempty"`
	// ProviderId - Only return skill templates using one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// IntegrationId - Only return skill templates using one of these integrations.
	IntegrationId *any `json:"integration_id,omitempty"`
	// CreatedAt - Only return results whose skill template creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *SkillsTemplatesListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose skill template last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *SkillsTemplatesListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapSkillsTemplatesListQueryFromJSON deserializes JSON data into a SkillsTemplatesListQuery.
func MapSkillsTemplatesListQueryFromJSON(data []byte) (*SkillsTemplatesListQuery, error) {
	var v SkillsTemplatesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsTemplatesListQueryToJSON serializes a SkillsTemplatesListQuery to JSON.
func MapSkillsTemplatesListQueryToJSON(v *SkillsTemplatesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
