package items

import (
	"encoding/json"
	"time"
)

// SkillsTemplatesItemsDeleteOutputIntegrationConfiguration - Which settings integration instances may customize.
type SkillsTemplatesItemsDeleteOutputIntegrationConfiguration struct {
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

// SkillsTemplatesItemsDeleteOutputIntegration represents the skills templates items delete output integration type.
type SkillsTemplatesItemsDeleteOutputIntegration struct {
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
	Configuration SkillsTemplatesItemsDeleteOutputIntegrationConfiguration `json:"configuration"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// SkillsTemplatesItemsDeleteOutputProvider represents the skills templates items delete output provider type.
type SkillsTemplatesItemsDeleteOutputProvider struct {
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

// SkillsTemplatesItemsDeleteOutput represents the skills templates items delete output type.
type SkillsTemplatesItemsDeleteOutput struct {
	Object      string                                       `json:"object"`
	Id          string                                       `json:"id"`
	Type        string                                       `json:"type"`
	Integration *SkillsTemplatesItemsDeleteOutputIntegration `json:"integration,omitempty"`
	Provider    *SkillsTemplatesItemsDeleteOutputProvider    `json:"provider,omitempty"`
	CreatedAt   time.Time                                    `json:"created_at"`
	UpdatedAt   time.Time                                    `json:"updated_at"`
}

// MapSkillsTemplatesItemsDeleteOutputFromJSON deserializes JSON data into a SkillsTemplatesItemsDeleteOutput.
func MapSkillsTemplatesItemsDeleteOutputFromJSON(data []byte) (*SkillsTemplatesItemsDeleteOutput, error) {
	var v SkillsTemplatesItemsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsTemplatesItemsDeleteOutputToJSON serializes a SkillsTemplatesItemsDeleteOutput to JSON.
func MapSkillsTemplatesItemsDeleteOutputToJSON(v *SkillsTemplatesItemsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
