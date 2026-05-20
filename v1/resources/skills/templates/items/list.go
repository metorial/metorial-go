package items

import (
	"encoding/json"
	"time"
)

// SkillsTemplatesItemsListOutputItemsIntegrationConfiguration represents the skills templates items list output items integration configuration type.
type SkillsTemplatesItemsListOutputItemsIntegrationConfiguration struct {
	CanAttachCustomToolFilters    bool `json:"can_attach_custom_tool_filters"`
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	CanOverrideToolFilters        bool `json:"can_override_tool_filters"`
}

// SkillsTemplatesItemsListOutputItemsIntegration represents the skills templates items list output items integration type.
type SkillsTemplatesItemsListOutputItemsIntegration struct {
	Object        string                                                      `json:"object"`
	Id            string                                                      `json:"id"`
	Slug          string                                                      `json:"slug"`
	Name          string                                                      `json:"name"`
	Description   *string                                                     `json:"description,omitempty"`
	Metadata      *map[string]any                                             `json:"metadata,omitempty"`
	Configuration SkillsTemplatesItemsListOutputItemsIntegrationConfiguration `json:"configuration"`
	CreatedAt     time.Time                                                   `json:"created_at"`
	UpdatedAt     time.Time                                                   `json:"updated_at"`
	ArchivedAt    *time.Time                                                  `json:"archived_at,omitempty"`
}

// SkillsTemplatesItemsListOutputItemsProvider represents the skills templates items list output items provider type.
type SkillsTemplatesItemsListOutputItemsProvider struct {
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
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SkillsTemplatesItemsListOutput represents the skills templates items list output type.
type SkillsTemplatesItemsListOutput struct {
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
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
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
