package items

import (
	"encoding/json"
	"time"
)

// SkillsTemplatesItemsGetOutputIntegrationConfiguration - Which settings integration instances may customize.
type SkillsTemplatesItemsGetOutputIntegrationConfiguration struct {
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

// SkillsTemplatesItemsGetOutputIntegration represents the skills templates items get output integration type.
type SkillsTemplatesItemsGetOutputIntegration struct {
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
	Configuration SkillsTemplatesItemsGetOutputIntegrationConfiguration `json:"configuration"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// SkillsTemplatesItemsGetOutputProvider represents the skills templates items get output provider type.
type SkillsTemplatesItemsGetOutputProvider struct {
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

// SkillsTemplatesItemsGetOutput represents the skills templates items get output type.
type SkillsTemplatesItemsGetOutput struct {
	Object      string                                    `json:"object"`
	Id          string                                    `json:"id"`
	Type        string                                    `json:"type"`
	Integration *SkillsTemplatesItemsGetOutputIntegration `json:"integration,omitempty"`
	Provider    *SkillsTemplatesItemsGetOutputProvider    `json:"provider,omitempty"`
	CreatedAt   time.Time                                 `json:"created_at"`
	UpdatedAt   time.Time                                 `json:"updated_at"`
}

// MapSkillsTemplatesItemsGetOutputFromJSON deserializes JSON data into a SkillsTemplatesItemsGetOutput.
func MapSkillsTemplatesItemsGetOutputFromJSON(data []byte) (*SkillsTemplatesItemsGetOutput, error) {
	var v SkillsTemplatesItemsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsTemplatesItemsGetOutputToJSON serializes a SkillsTemplatesItemsGetOutput to JSON.
func MapSkillsTemplatesItemsGetOutputToJSON(v *SkillsTemplatesItemsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
