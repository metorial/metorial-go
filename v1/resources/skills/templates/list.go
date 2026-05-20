package templates

import (
	"encoding/json"
	"time"
)

// SkillsTemplatesListOutputItemsItemsIntegrationConfiguration represents the skills templates list output items items integration configuration type.
type SkillsTemplatesListOutputItemsItemsIntegrationConfiguration struct {
	CanAttachCustomToolFilters    bool `json:"can_attach_custom_tool_filters"`
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	CanOverrideToolFilters        bool `json:"can_override_tool_filters"`
}

// SkillsTemplatesListOutputItemsItemsIntegration represents the skills templates list output items items integration type.
type SkillsTemplatesListOutputItemsItemsIntegration struct {
	Object        string                                                      `json:"object"`
	Id            string                                                      `json:"id"`
	Slug          string                                                      `json:"slug"`
	Name          string                                                      `json:"name"`
	Description   *string                                                     `json:"description,omitempty"`
	Metadata      *map[string]any                                             `json:"metadata,omitempty"`
	Configuration SkillsTemplatesListOutputItemsItemsIntegrationConfiguration `json:"configuration"`
	CreatedAt     time.Time                                                   `json:"created_at"`
	UpdatedAt     time.Time                                                   `json:"updated_at"`
	ArchivedAt    *time.Time                                                  `json:"archived_at,omitempty"`
}

// SkillsTemplatesListOutputItemsItemsProvider represents the skills templates list output items items provider type.
type SkillsTemplatesListOutputItemsItemsProvider struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique provider identifier
	Id string `json:"id"`
	// Name - Display name of the provider
	Name string `json:"name"`
	// Description - Brief description of the provider
	Description *string `json:"description,omitempty"`
	// Slug - URL-friendly identifier
	Slug string `json:"slug"`
	// CreatedAt - Timestamp when the provider was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the provider was last updated
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
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SkillsTemplatesListOutput represents the skills templates list output type.
type SkillsTemplatesListOutput struct {
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

// SkillsTemplatesListQueryCreatedAt - Filter skill template creation time by date range
type SkillsTemplatesListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for skill template creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for skill template creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsTemplatesListQueryUpdatedAt - Filter skill template last update time by date range
type SkillsTemplatesListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for skill template last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for skill template last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsTemplatesListQuery represents the skills templates list query type.
type SkillsTemplatesListQuery struct {
	Limit         *float64 `json:"limit,omitempty"`
	After         *string  `json:"after,omitempty"`
	Before        *string  `json:"before,omitempty"`
	Cursor        *string  `json:"cursor,omitempty"`
	Order         *string  `json:"order,omitempty"`
	Search        *string  `json:"search,omitempty"`
	Status        *any     `json:"status,omitempty"`
	Owner         *any     `json:"owner,omitempty"`
	Id            *any     `json:"id,omitempty"`
	ProviderId    *any     `json:"provider_id,omitempty"`
	IntegrationId *any     `json:"integration_id,omitempty"`
	// CreatedAt - Filter skill template creation time by date range
	CreatedAt *SkillsTemplatesListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter skill template last update time by date range
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
