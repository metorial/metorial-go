package marketplaces

import (
	"encoding/json"
	"time"
)

// SkillsMarketplacesArchiveOutputPluginsSkillPluginSkills represents the skills marketplaces archive output plugins skill plugin skills type.
type SkillsMarketplacesArchiveOutputPluginsSkillPluginSkills struct {
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

// SkillsMarketplacesArchiveOutputPluginsSkillPlugin represents the skills marketplaces archive output plugins skill plugin type.
type SkillsMarketplacesArchiveOutputPluginsSkillPlugin struct {
	Object               string                                                    `json:"object"`
	Id                   string                                                    `json:"id"`
	Status               string                                                    `json:"status"`
	SyncStatus           string                                                    `json:"sync_status"`
	ImageUrl             string                                                    `json:"image_url"`
	Name                 string                                                    `json:"name"`
	Description          *string                                                   `json:"description,omitempty"`
	LongDescription      *string                                                   `json:"long_description,omitempty"`
	Category             *string                                                   `json:"category,omitempty"`
	Slug                 string                                                    `json:"slug"`
	SkillConfigurationId *string                                                   `json:"skill_configuration_id,omitempty"`
	Skills               []SkillsMarketplacesArchiveOutputPluginsSkillPluginSkills `json:"skills"`
	CreatedAt            time.Time                                                 `json:"created_at"`
	UpdatedAt            time.Time                                                 `json:"updated_at"`
}

// SkillsMarketplacesArchiveOutputPlugins represents the skills marketplaces archive output plugins type.
type SkillsMarketplacesArchiveOutputPlugins struct {
	Object               string                                             `json:"object"`
	Id                   string                                             `json:"id"`
	Status               string                                             `json:"status"`
	Identifier           string                                             `json:"identifier"`
	SkillConfigurationId *string                                            `json:"skill_configuration_id,omitempty"`
	SkillMarketplaceId   *string                                            `json:"skill_marketplace_id,omitempty"`
	SkillPluginId        *string                                            `json:"skill_plugin_id,omitempty"`
	SkillPlugin          *SkillsMarketplacesArchiveOutputPluginsSkillPlugin `json:"skill_plugin,omitempty"`
	CreatedAt            time.Time                                          `json:"created_at"`
	UpdatedAt            time.Time                                          `json:"updated_at"`
}

// SkillsMarketplacesArchiveOutput represents the skills marketplaces archive output type.
type SkillsMarketplacesArchiveOutput struct {
	Object               string                                   `json:"object"`
	Id                   string                                   `json:"id"`
	Status               string                                   `json:"status"`
	SyncStatus           string                                   `json:"sync_status"`
	ImageUrl             string                                   `json:"image_url"`
	Name                 string                                   `json:"name"`
	Description          *string                                  `json:"description,omitempty"`
	Slug                 string                                   `json:"slug"`
	SkillConfigurationId *string                                  `json:"skill_configuration_id,omitempty"`
	Plugins              []SkillsMarketplacesArchiveOutputPlugins `json:"plugins"`
	CreatedAt            time.Time                                `json:"created_at"`
	UpdatedAt            time.Time                                `json:"updated_at"`
}

// MapSkillsMarketplacesArchiveOutputFromJSON deserializes JSON data into a SkillsMarketplacesArchiveOutput.
func MapSkillsMarketplacesArchiveOutputFromJSON(data []byte) (*SkillsMarketplacesArchiveOutput, error) {
	var v SkillsMarketplacesArchiveOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMarketplacesArchiveOutputToJSON serializes a SkillsMarketplacesArchiveOutput to JSON.
func MapSkillsMarketplacesArchiveOutputToJSON(v *SkillsMarketplacesArchiveOutput) ([]byte, error) {
	return json.Marshal(v)
}
