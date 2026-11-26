package templates

import (
	"encoding/json"
	"time"
)

// SkillsTemplatesGetOutputItemsIntegrationConfiguration - Which settings integration instances may customize.
type SkillsTemplatesGetOutputItemsIntegrationConfiguration struct {
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

// SkillsTemplatesGetOutputItemsIntegration represents the skills templates get output items integration type.
type SkillsTemplatesGetOutputItemsIntegration struct {
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
	Configuration SkillsTemplatesGetOutputItemsIntegrationConfiguration `json:"configuration"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// SkillsTemplatesGetOutputItemsProvider represents the skills templates get output items provider type.
type SkillsTemplatesGetOutputItemsProvider struct {
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

// SkillsTemplatesGetOutputItems represents the skills templates get output items type.
type SkillsTemplatesGetOutputItems struct {
	Object      string                                    `json:"object"`
	Id          string                                    `json:"id"`
	Type        string                                    `json:"type"`
	Integration *SkillsTemplatesGetOutputItemsIntegration `json:"integration,omitempty"`
	Provider    *SkillsTemplatesGetOutputItemsProvider    `json:"provider,omitempty"`
	CreatedAt   time.Time                                 `json:"created_at"`
	UpdatedAt   time.Time                                 `json:"updated_at"`
}

// SkillsTemplatesGetOutput represents the skills templates get output type.
type SkillsTemplatesGetOutput struct {
	Object      string                          `json:"object"`
	Id          string                          `json:"id"`
	Status      string                          `json:"status"`
	Owner       string                          `json:"owner"`
	Slug        string                          `json:"slug"`
	Name        string                          `json:"name"`
	Description *string                         `json:"description,omitempty"`
	Metadata    map[string]any                  `json:"metadata"`
	StoreId     string                          `json:"store_id"`
	Items       []SkillsTemplatesGetOutputItems `json:"items"`
	CreatedAt   time.Time                       `json:"created_at"`
	UpdatedAt   time.Time                       `json:"updated_at"`
}

// MapSkillsTemplatesGetOutputFromJSON deserializes JSON data into a SkillsTemplatesGetOutput.
func MapSkillsTemplatesGetOutputFromJSON(data []byte) (*SkillsTemplatesGetOutput, error) {
	var v SkillsTemplatesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsTemplatesGetOutputToJSON serializes a SkillsTemplatesGetOutput to JSON.
func MapSkillsTemplatesGetOutputToJSON(v *SkillsTemplatesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
