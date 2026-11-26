package items

import (
	"encoding/json"
	"time"
)

// SkillsItemsCreateOutputIntegrationConfiguration - Which settings integration instances may customize.
type SkillsItemsCreateOutputIntegrationConfiguration struct {
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

// SkillsItemsCreateOutputIntegration represents the skills items create output integration type.
type SkillsItemsCreateOutputIntegration struct {
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
	Configuration SkillsItemsCreateOutputIntegrationConfiguration `json:"configuration"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// SkillsItemsCreateOutputProvider represents the skills items create output provider type.
type SkillsItemsCreateOutputProvider struct {
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
	// Type - Attach an integration to the skill.
	Type *string `json:"type,omitempty"`
	// ProviderId - Provider to attach to the skill, by ID.
	ProviderId *string `json:"provider_id,omitempty"`
	// IntegrationId - Integration to attach to the skill, by ID.
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
