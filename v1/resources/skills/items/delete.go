package items

import (
	"encoding/json"
	"time"
)

// SkillsItemsDeleteOutputIntegrationConfiguration - Which settings integration instances may customize.
type SkillsItemsDeleteOutputIntegrationConfiguration struct {
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

// SkillsItemsDeleteOutputIntegration represents the skills items delete output integration type.
type SkillsItemsDeleteOutputIntegration struct {
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
	Configuration SkillsItemsDeleteOutputIntegrationConfiguration `json:"configuration"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// SkillsItemsDeleteOutputProvider represents the skills items delete output provider type.
type SkillsItemsDeleteOutputProvider struct {
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

// SkillsItemsDeleteOutput represents the skills items delete output type.
type SkillsItemsDeleteOutput struct {
	Object      string                              `json:"object"`
	Id          string                              `json:"id"`
	Status      string                              `json:"status"`
	Type        string                              `json:"type"`
	SkillId     string                              `json:"skill_id"`
	Integration *SkillsItemsDeleteOutputIntegration `json:"integration,omitempty"`
	Provider    *SkillsItemsDeleteOutputProvider    `json:"provider,omitempty"`
	CreatedAt   time.Time                           `json:"created_at"`
}

// MapSkillsItemsDeleteOutputFromJSON deserializes JSON data into a SkillsItemsDeleteOutput.
func MapSkillsItemsDeleteOutputFromJSON(data []byte) (*SkillsItemsDeleteOutput, error) {
	var v SkillsItemsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsItemsDeleteOutputToJSON serializes a SkillsItemsDeleteOutput to JSON.
func MapSkillsItemsDeleteOutputToJSON(v *SkillsItemsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
