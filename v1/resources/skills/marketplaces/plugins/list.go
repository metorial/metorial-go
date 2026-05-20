package plugins

import (
	"encoding/json"
	"time"
)

// SkillsMarketplacesPluginsListOutputItemsSkillPluginSkills represents the skills marketplaces plugins list output items skill plugin skills type.
type SkillsMarketplacesPluginsListOutputItemsSkillPluginSkills struct {
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

// SkillsMarketplacesPluginsListOutputItemsSkillPlugin represents the skills marketplaces plugins list output items skill plugin type.
type SkillsMarketplacesPluginsListOutputItemsSkillPlugin struct {
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
	Skills               []SkillsMarketplacesPluginsListOutputItemsSkillPluginSkills `json:"skills"`
	CreatedAt            time.Time                                                   `json:"created_at"`
	UpdatedAt            time.Time                                                   `json:"updated_at"`
}

// SkillsMarketplacesPluginsListOutputItems represents the skills marketplaces plugins list output items type.
type SkillsMarketplacesPluginsListOutputItems struct {
	Object               string                                               `json:"object"`
	Id                   string                                               `json:"id"`
	Status               string                                               `json:"status"`
	Identifier           string                                               `json:"identifier"`
	SkillConfigurationId *string                                              `json:"skill_configuration_id,omitempty"`
	SkillMarketplaceId   *string                                              `json:"skill_marketplace_id,omitempty"`
	SkillPluginId        *string                                              `json:"skill_plugin_id,omitempty"`
	SkillPlugin          *SkillsMarketplacesPluginsListOutputItemsSkillPlugin `json:"skill_plugin,omitempty"`
	CreatedAt            time.Time                                            `json:"created_at"`
	UpdatedAt            time.Time                                            `json:"updated_at"`
}

// SkillsMarketplacesPluginsListOutputPagination represents the skills marketplaces plugins list output pagination type.
type SkillsMarketplacesPluginsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SkillsMarketplacesPluginsListOutput represents the skills marketplaces plugins list output type.
type SkillsMarketplacesPluginsListOutput struct {
	Items      []SkillsMarketplacesPluginsListOutputItems    `json:"items"`
	Pagination SkillsMarketplacesPluginsListOutputPagination `json:"pagination"`
}

// MapSkillsMarketplacesPluginsListOutputFromJSON deserializes JSON data into a SkillsMarketplacesPluginsListOutput.
func MapSkillsMarketplacesPluginsListOutputFromJSON(data []byte) (*SkillsMarketplacesPluginsListOutput, error) {
	var v SkillsMarketplacesPluginsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMarketplacesPluginsListOutputToJSON serializes a SkillsMarketplacesPluginsListOutput to JSON.
func MapSkillsMarketplacesPluginsListOutputToJSON(v *SkillsMarketplacesPluginsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsMarketplacesPluginsListQueryCreatedAt - Filter skill marketplace plugin creation time by date range
type SkillsMarketplacesPluginsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for skill marketplace plugin creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for skill marketplace plugin creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsMarketplacesPluginsListQueryUpdatedAt - Filter skill marketplace plugin last update time by date range
type SkillsMarketplacesPluginsListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for skill marketplace plugin last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for skill marketplace plugin last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsMarketplacesPluginsListQuery represents the skills marketplaces plugins list query type.
type SkillsMarketplacesPluginsListQuery struct {
	Limit                *float64 `json:"limit,omitempty"`
	After                *string  `json:"after,omitempty"`
	Before               *string  `json:"before,omitempty"`
	Cursor               *string  `json:"cursor,omitempty"`
	Order                *string  `json:"order,omitempty"`
	Id                   *any     `json:"id,omitempty"`
	SkillPluginId        *any     `json:"skill_plugin_id,omitempty"`
	Status               *any     `json:"status,omitempty"`
	SkillConfigurationId *any     `json:"skill_configuration_id,omitempty"`
	// CreatedAt - Filter skill marketplace plugin creation time by date range
	CreatedAt *SkillsMarketplacesPluginsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter skill marketplace plugin last update time by date range
	UpdatedAt *SkillsMarketplacesPluginsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapSkillsMarketplacesPluginsListQueryFromJSON deserializes JSON data into a SkillsMarketplacesPluginsListQuery.
func MapSkillsMarketplacesPluginsListQueryFromJSON(data []byte) (*SkillsMarketplacesPluginsListQuery, error) {
	var v SkillsMarketplacesPluginsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMarketplacesPluginsListQueryToJSON serializes a SkillsMarketplacesPluginsListQuery to JSON.
func MapSkillsMarketplacesPluginsListQueryToJSON(v *SkillsMarketplacesPluginsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
