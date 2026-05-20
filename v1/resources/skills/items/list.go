package items

import (
	"encoding/json"
	"time"
)

// SkillsItemsListOutputItemsIntegrationConfiguration represents the skills items list output items integration configuration type.
type SkillsItemsListOutputItemsIntegrationConfiguration struct {
	CanAttachCustomToolFilters    bool `json:"can_attach_custom_tool_filters"`
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	CanOverrideToolFilters        bool `json:"can_override_tool_filters"`
}

// SkillsItemsListOutputItemsIntegration represents the skills items list output items integration type.
type SkillsItemsListOutputItemsIntegration struct {
	Object        string                                             `json:"object"`
	Id            string                                             `json:"id"`
	Slug          string                                             `json:"slug"`
	Name          string                                             `json:"name"`
	Description   *string                                            `json:"description,omitempty"`
	Metadata      *map[string]any                                    `json:"metadata,omitempty"`
	Configuration SkillsItemsListOutputItemsIntegrationConfiguration `json:"configuration"`
	CreatedAt     time.Time                                          `json:"created_at"`
	UpdatedAt     time.Time                                          `json:"updated_at"`
	ArchivedAt    *time.Time                                         `json:"archived_at,omitempty"`
}

// SkillsItemsListOutputItemsProvider represents the skills items list output items provider type.
type SkillsItemsListOutputItemsProvider struct {
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
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SkillsItemsListOutput represents the skills items list output type.
type SkillsItemsListOutput struct {
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

// SkillsItemsListQueryCreatedAt - Filter skill item creation time by date range
type SkillsItemsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for skill item creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for skill item creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsItemsListQuery represents the skills items list query type.
type SkillsItemsListQuery struct {
	Limit         *float64 `json:"limit,omitempty"`
	After         *string  `json:"after,omitempty"`
	Before        *string  `json:"before,omitempty"`
	Cursor        *string  `json:"cursor,omitempty"`
	Order         *string  `json:"order,omitempty"`
	Status        *any     `json:"status,omitempty"`
	Type          *any     `json:"type,omitempty"`
	Id            *any     `json:"id,omitempty"`
	IntegrationId *any     `json:"integration_id,omitempty"`
	ProviderId    *any     `json:"provider_id,omitempty"`
	// CreatedAt - Filter skill item creation time by date range
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
