package marketplaces

import (
	"encoding/json"
	"time"
)

// SkillsMarketplacesListOutputItemsPluginsSkillPluginSkills represents the skills marketplaces list output items plugins skill plugin skills type.
type SkillsMarketplacesListOutputItemsPluginsSkillPluginSkills struct {
	Object               string          `json:"object"`
	Id                   string          `json:"id"`
	Identifier           string          `json:"identifier"`
	Status               string          `json:"status"`
	ClientName           *string         `json:"client_name,omitempty"`
	ClientDescription    *string         `json:"client_description,omitempty"`
	ClientMetadata       *map[string]any `json:"client_metadata,omitempty"`
	License              *string         `json:"license,omitempty"`
	Compatibility        *string         `json:"compatibility,omitempty"`
	SkillConfigurationId *string         `json:"skill_configuration_id,omitempty"`
	SkillId              string          `json:"skill_id"`
	CreatedAt            time.Time       `json:"created_at"`
	UpdatedAt            time.Time       `json:"updated_at"`
}

// SkillsMarketplacesListOutputItemsPluginsSkillPlugin represents the skills marketplaces list output items plugins skill plugin type.
type SkillsMarketplacesListOutputItemsPluginsSkillPlugin struct {
	Object               string                                                      `json:"object"`
	Id                   string                                                      `json:"id"`
	Status               string                                                      `json:"status"`
	SyncStatus           string                                                      `json:"sync_status"`
	ImageUrl             string                                                      `json:"image_url"`
	Name                 string                                                      `json:"name"`
	Description          *string                                                     `json:"description,omitempty"`
	LongDescription      *string                                                     `json:"long_description,omitempty"`
	Category             *string                                                     `json:"category,omitempty"`
	Slug                 string                                                      `json:"slug"`
	SkillConfigurationId *string                                                     `json:"skill_configuration_id,omitempty"`
	Skills               []SkillsMarketplacesListOutputItemsPluginsSkillPluginSkills `json:"skills"`
	CreatedAt            time.Time                                                   `json:"created_at"`
	UpdatedAt            time.Time                                                   `json:"updated_at"`
}

// SkillsMarketplacesListOutputItemsPlugins represents the skills marketplaces list output items plugins type.
type SkillsMarketplacesListOutputItemsPlugins struct {
	Object               string                                               `json:"object"`
	Id                   string                                               `json:"id"`
	Status               string                                               `json:"status"`
	Identifier           string                                               `json:"identifier"`
	SkillConfigurationId *string                                              `json:"skill_configuration_id,omitempty"`
	SkillMarketplaceId   *string                                              `json:"skill_marketplace_id,omitempty"`
	SkillPluginId        *string                                              `json:"skill_plugin_id,omitempty"`
	SkillPlugin          *SkillsMarketplacesListOutputItemsPluginsSkillPlugin `json:"skill_plugin,omitempty"`
	CreatedAt            time.Time                                            `json:"created_at"`
	UpdatedAt            time.Time                                            `json:"updated_at"`
}

// SkillsMarketplacesListOutputItems represents the skills marketplaces list output items type.
type SkillsMarketplacesListOutputItems struct {
	Object               string                                     `json:"object"`
	Id                   string                                     `json:"id"`
	Status               string                                     `json:"status"`
	SyncStatus           string                                     `json:"sync_status"`
	ImageUrl             string                                     `json:"image_url"`
	Name                 string                                     `json:"name"`
	Description          *string                                    `json:"description,omitempty"`
	Slug                 string                                     `json:"slug"`
	SkillConfigurationId *string                                    `json:"skill_configuration_id,omitempty"`
	Plugins              []SkillsMarketplacesListOutputItemsPlugins `json:"plugins"`
	CreatedAt            time.Time                                  `json:"created_at"`
	UpdatedAt            time.Time                                  `json:"updated_at"`
}

// SkillsMarketplacesListOutputPagination represents the skills marketplaces list output pagination type.
type SkillsMarketplacesListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SkillsMarketplacesListOutput represents the skills marketplaces list output type.
type SkillsMarketplacesListOutput struct {
	Items      []SkillsMarketplacesListOutputItems    `json:"items"`
	Pagination SkillsMarketplacesListOutputPagination `json:"pagination"`
}

// MapSkillsMarketplacesListOutputFromJSON deserializes JSON data into a SkillsMarketplacesListOutput.
func MapSkillsMarketplacesListOutputFromJSON(data []byte) (*SkillsMarketplacesListOutput, error) {
	var v SkillsMarketplacesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMarketplacesListOutputToJSON serializes a SkillsMarketplacesListOutput to JSON.
func MapSkillsMarketplacesListOutputToJSON(v *SkillsMarketplacesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsMarketplacesListQueryCreatedAt - Filter skill marketplace creation time by date range
type SkillsMarketplacesListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for skill marketplace creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for skill marketplace creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsMarketplacesListQueryUpdatedAt - Filter skill marketplace last update time by date range
type SkillsMarketplacesListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for skill marketplace last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for skill marketplace last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsMarketplacesListQuery represents the skills marketplaces list query type.
type SkillsMarketplacesListQuery struct {
	Limit                *float64 `json:"limit,omitempty"`
	After                *string  `json:"after,omitempty"`
	Before               *string  `json:"before,omitempty"`
	Cursor               *string  `json:"cursor,omitempty"`
	Order                *string  `json:"order,omitempty"`
	Id                   *any     `json:"id,omitempty"`
	Status               *any     `json:"status,omitempty"`
	SkillConfigurationId *any     `json:"skill_configuration_id,omitempty"`
	Search               *string  `json:"search,omitempty"`
	// CreatedAt - Filter skill marketplace creation time by date range
	CreatedAt *SkillsMarketplacesListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter skill marketplace last update time by date range
	UpdatedAt *SkillsMarketplacesListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapSkillsMarketplacesListQueryFromJSON deserializes JSON data into a SkillsMarketplacesListQuery.
func MapSkillsMarketplacesListQueryFromJSON(data []byte) (*SkillsMarketplacesListQuery, error) {
	var v SkillsMarketplacesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMarketplacesListQueryToJSON serializes a SkillsMarketplacesListQuery to JSON.
func MapSkillsMarketplacesListQueryToJSON(v *SkillsMarketplacesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
