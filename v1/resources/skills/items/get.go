package items

import (
	"encoding/json"
	"time"
)

// SkillsItemsGetOutputIntegrationConfiguration - Which settings integration instances may customize.
type SkillsItemsGetOutputIntegrationConfiguration struct {
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

// SkillsItemsGetOutputIntegration represents the skills items get output integration type.
type SkillsItemsGetOutputIntegration struct {
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
	Configuration SkillsItemsGetOutputIntegrationConfiguration `json:"configuration"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// SkillsItemsGetOutputProvider represents the skills items get output provider type.
type SkillsItemsGetOutputProvider struct {
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

// SkillsItemsGetOutput represents the skills items get output type.
type SkillsItemsGetOutput struct {
	Object      string                           `json:"object"`
	Id          string                           `json:"id"`
	Status      string                           `json:"status"`
	Type        string                           `json:"type"`
	SkillId     string                           `json:"skill_id"`
	Integration *SkillsItemsGetOutputIntegration `json:"integration,omitempty"`
	Provider    *SkillsItemsGetOutputProvider    `json:"provider,omitempty"`
	CreatedAt   time.Time                        `json:"created_at"`
}

// MapSkillsItemsGetOutputFromJSON deserializes JSON data into a SkillsItemsGetOutput.
func MapSkillsItemsGetOutputFromJSON(data []byte) (*SkillsItemsGetOutput, error) {
	var v SkillsItemsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsItemsGetOutputToJSON serializes a SkillsItemsGetOutput to JSON.
func MapSkillsItemsGetOutputToJSON(v *SkillsItemsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
