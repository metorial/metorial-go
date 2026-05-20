package items

import (
	"encoding/json"
	"time"
)

// SkillsTemplatesItemsCreateOutputIntegrationConfiguration represents the skills templates items create output integration configuration type.
type SkillsTemplatesItemsCreateOutputIntegrationConfiguration struct {
	CanAttachCustomToolFilters    bool `json:"can_attach_custom_tool_filters"`
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	CanOverrideToolFilters        bool `json:"can_override_tool_filters"`
}

// SkillsTemplatesItemsCreateOutputIntegration represents the skills templates items create output integration type.
type SkillsTemplatesItemsCreateOutputIntegration struct {
	Object        string                                                   `json:"object"`
	Id            string                                                   `json:"id"`
	Slug          string                                                   `json:"slug"`
	Name          string                                                   `json:"name"`
	Description   *string                                                  `json:"description,omitempty"`
	Metadata      *map[string]any                                          `json:"metadata,omitempty"`
	Configuration SkillsTemplatesItemsCreateOutputIntegrationConfiguration `json:"configuration"`
	CreatedAt     time.Time                                                `json:"created_at"`
	UpdatedAt     time.Time                                                `json:"updated_at"`
	ArchivedAt    *time.Time                                               `json:"archived_at,omitempty"`
}

// SkillsTemplatesItemsCreateOutputProvider represents the skills templates items create output provider type.
type SkillsTemplatesItemsCreateOutputProvider struct {
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

// SkillsTemplatesItemsCreateOutput represents the skills templates items create output type.
type SkillsTemplatesItemsCreateOutput struct {
	Object      string                                       `json:"object"`
	Id          string                                       `json:"id"`
	Type        string                                       `json:"type"`
	Integration *SkillsTemplatesItemsCreateOutputIntegration `json:"integration,omitempty"`
	Provider    *SkillsTemplatesItemsCreateOutputProvider    `json:"provider,omitempty"`
	CreatedAt   time.Time                                    `json:"created_at"`
	UpdatedAt   time.Time                                    `json:"updated_at"`
}

// MapSkillsTemplatesItemsCreateOutputFromJSON deserializes JSON data into a SkillsTemplatesItemsCreateOutput.
func MapSkillsTemplatesItemsCreateOutputFromJSON(data []byte) (*SkillsTemplatesItemsCreateOutput, error) {
	var v SkillsTemplatesItemsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsTemplatesItemsCreateOutputToJSON serializes a SkillsTemplatesItemsCreateOutput to JSON.
func MapSkillsTemplatesItemsCreateOutputToJSON(v *SkillsTemplatesItemsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsTemplatesItemsCreateBody represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SkillsTemplatesItemsCreateBody struct {
	Type          *string `json:"type,omitempty"`
	ProviderId    *string `json:"provider_id,omitempty"`
	IntegrationId *string `json:"integration_id,omitempty"`
}

// MapSkillsTemplatesItemsCreateBodyFromJSON deserializes JSON data into a SkillsTemplatesItemsCreateBody.
func MapSkillsTemplatesItemsCreateBodyFromJSON(data []byte) (*SkillsTemplatesItemsCreateBody, error) {
	var v SkillsTemplatesItemsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsTemplatesItemsCreateBodyToJSON serializes a SkillsTemplatesItemsCreateBody to JSON.
func MapSkillsTemplatesItemsCreateBodyToJSON(v *SkillsTemplatesItemsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
