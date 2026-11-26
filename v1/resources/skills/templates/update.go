package templates

import (
	"encoding/json"
	"github.com/metorial/metorial-go/v1/nullable"
	"time"
)

// SkillsTemplatesUpdateOutputItemsIntegrationConfiguration - Which settings integration instances may customize.
type SkillsTemplatesUpdateOutputItemsIntegrationConfiguration struct {
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

// SkillsTemplatesUpdateOutputItemsIntegration represents the skills templates update output items integration type.
type SkillsTemplatesUpdateOutputItemsIntegration struct {
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
	Configuration SkillsTemplatesUpdateOutputItemsIntegrationConfiguration `json:"configuration"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// SkillsTemplatesUpdateOutputItemsProvider represents the skills templates update output items provider type.
type SkillsTemplatesUpdateOutputItemsProvider struct {
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

// SkillsTemplatesUpdateOutputItems represents the skills templates update output items type.
type SkillsTemplatesUpdateOutputItems struct {
	Object      string                                       `json:"object"`
	Id          string                                       `json:"id"`
	Type        string                                       `json:"type"`
	Integration *SkillsTemplatesUpdateOutputItemsIntegration `json:"integration,omitempty"`
	Provider    *SkillsTemplatesUpdateOutputItemsProvider    `json:"provider,omitempty"`
	CreatedAt   time.Time                                    `json:"created_at"`
	UpdatedAt   time.Time                                    `json:"updated_at"`
}

// SkillsTemplatesUpdateOutput represents the skills templates update output type.
type SkillsTemplatesUpdateOutput struct {
	Object      string                             `json:"object"`
	Id          string                             `json:"id"`
	Status      string                             `json:"status"`
	Owner       string                             `json:"owner"`
	Slug        string                             `json:"slug"`
	Name        string                             `json:"name"`
	Description *string                            `json:"description,omitempty"`
	Metadata    map[string]any                     `json:"metadata"`
	StoreId     string                             `json:"store_id"`
	Items       []SkillsTemplatesUpdateOutputItems `json:"items"`
	CreatedAt   time.Time                          `json:"created_at"`
	UpdatedAt   time.Time                          `json:"updated_at"`
}

// MapSkillsTemplatesUpdateOutputFromJSON deserializes JSON data into a SkillsTemplatesUpdateOutput.
func MapSkillsTemplatesUpdateOutputFromJSON(data []byte) (*SkillsTemplatesUpdateOutput, error) {
	var v SkillsTemplatesUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsTemplatesUpdateOutputToJSON serializes a SkillsTemplatesUpdateOutput to JSON.
func MapSkillsTemplatesUpdateOutputToJSON(v *SkillsTemplatesUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsTemplatesUpdateBody represents the skills templates update body type.
type SkillsTemplatesUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description. `null` clears it.
	Description *nullable.Nullable[string] `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the skill template. Replaces existing metadata. `null` clears it.
	Metadata *nullable.Nullable[map[string]any] `json:"metadata,omitempty"`
}

// MapSkillsTemplatesUpdateBodyFromJSON deserializes JSON data into a SkillsTemplatesUpdateBody.
func MapSkillsTemplatesUpdateBodyFromJSON(data []byte) (*SkillsTemplatesUpdateBody, error) {
	var v SkillsTemplatesUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsTemplatesUpdateBodyToJSON serializes a SkillsTemplatesUpdateBody to JSON.
func MapSkillsTemplatesUpdateBodyToJSON(v *SkillsTemplatesUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
