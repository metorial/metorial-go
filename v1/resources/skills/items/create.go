package items

import (
	"encoding/json"
	"time"
)

// SkillsItemsCreateOutputIntegrationConfiguration represents the skills items create output integration configuration type.
type SkillsItemsCreateOutputIntegrationConfiguration struct {
	CanAttachCustomToolFilters    bool `json:"can_attach_custom_tool_filters"`
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	CanOverrideToolFilters        bool `json:"can_override_tool_filters"`
}

// SkillsItemsCreateOutputIntegration represents the skills items create output integration type.
type SkillsItemsCreateOutputIntegration struct {
	Object        string                                          `json:"object"`
	Id            string                                          `json:"id"`
	Slug          string                                          `json:"slug"`
	Name          string                                          `json:"name"`
	Description   *string                                         `json:"description,omitempty"`
	Metadata      *map[string]any                                 `json:"metadata,omitempty"`
	Configuration SkillsItemsCreateOutputIntegrationConfiguration `json:"configuration"`
	CreatedAt     time.Time                                       `json:"created_at"`
	UpdatedAt     time.Time                                       `json:"updated_at"`
	ArchivedAt    *time.Time                                      `json:"archived_at,omitempty"`
}

// SkillsItemsCreateOutputProvider represents the skills items create output provider type.
type SkillsItemsCreateOutputProvider struct {
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

// SkillsItemsCreateOutput represents the skills items create output type.
type SkillsItemsCreateOutput struct {
	Object      string                              `json:"object"`
	Id          string                              `json:"id"`
	Status      string                              `json:"status"`
	Type        string                              `json:"type"`
	SkillId     string                              `json:"skill_id"`
	Integration *SkillsItemsCreateOutputIntegration `json:"integration,omitempty"`
	Provider    *SkillsItemsCreateOutputProvider    `json:"provider,omitempty"`
	CreatedAt   time.Time                           `json:"created_at"`
}

// MapSkillsItemsCreateOutputFromJSON deserializes JSON data into a SkillsItemsCreateOutput.
func MapSkillsItemsCreateOutputFromJSON(data []byte) (*SkillsItemsCreateOutput, error) {
	var v SkillsItemsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsItemsCreateOutputToJSON serializes a SkillsItemsCreateOutput to JSON.
func MapSkillsItemsCreateOutputToJSON(v *SkillsItemsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsItemsCreateBody represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SkillsItemsCreateBody struct {
	Type          *string `json:"type,omitempty"`
	ProviderId    *string `json:"provider_id,omitempty"`
	IntegrationId *string `json:"integration_id,omitempty"`
}

// MapSkillsItemsCreateBodyFromJSON deserializes JSON data into a SkillsItemsCreateBody.
func MapSkillsItemsCreateBodyFromJSON(data []byte) (*SkillsItemsCreateBody, error) {
	var v SkillsItemsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsItemsCreateBodyToJSON serializes a SkillsItemsCreateBody to JSON.
func MapSkillsItemsCreateBodyToJSON(v *SkillsItemsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
