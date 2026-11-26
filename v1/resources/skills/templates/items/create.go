package items

import (
	"encoding/json"
	"time"
)

// SkillsTemplatesItemsCreateOutputIntegrationConfiguration - Which settings integration instances may customize.
type SkillsTemplatesItemsCreateOutputIntegrationConfiguration struct {
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

// SkillsTemplatesItemsCreateOutputIntegration represents the skills templates items create output integration type.
type SkillsTemplatesItemsCreateOutputIntegration struct {
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
	Configuration SkillsTemplatesItemsCreateOutputIntegrationConfiguration `json:"configuration"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// SkillsTemplatesItemsCreateOutputProvider represents the skills templates items create output provider type.
type SkillsTemplatesItemsCreateOutputProvider struct {
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
	// Type - Attach an integration to the skill template.
	Type *string `json:"type,omitempty"`
	// ProviderId - Provider to attach to the skill template, by ID.
	ProviderId *string `json:"provider_id,omitempty"`
	// IntegrationId - Integration to attach to the skill template, by ID.
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
