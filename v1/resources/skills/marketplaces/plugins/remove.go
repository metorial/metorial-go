package plugins

import (
	"encoding/json"
	"time"
)

// SkillsMarketplacesPluginsRemoveOutputSkillPluginSkills represents the skills marketplaces plugins remove output skill plugin skills type.
type SkillsMarketplacesPluginsRemoveOutputSkillPluginSkills struct {
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

// SkillsMarketplacesPluginsRemoveOutputSkillPlugin represents the skills marketplaces plugins remove output skill plugin type.
type SkillsMarketplacesPluginsRemoveOutputSkillPlugin struct {
	Object               string                                                   `json:"object"`
	Id                   string                                                   `json:"id"`
	Status               string                                                   `json:"status"`
	SyncStatus           string                                                   `json:"sync_status"`
	ImageUrl             string                                                   `json:"image_url"`
	Name                 string                                                   `json:"name"`
	Description          *string                                                  `json:"description,omitempty"`
	LongDescription      *string                                                  `json:"long_description,omitempty"`
	Category             *string                                                  `json:"category,omitempty"`
	Slug                 string                                                   `json:"slug"`
	SkillConfigurationId *string                                                  `json:"skill_configuration_id,omitempty"`
	Skills               []SkillsMarketplacesPluginsRemoveOutputSkillPluginSkills `json:"skills"`
	CreatedAt            time.Time                                                `json:"created_at"`
	UpdatedAt            time.Time                                                `json:"updated_at"`
}

// SkillsMarketplacesPluginsRemoveOutput represents the skills marketplaces plugins remove output type.
type SkillsMarketplacesPluginsRemoveOutput struct {
	Object               string                                            `json:"object"`
	Id                   string                                            `json:"id"`
	Status               string                                            `json:"status"`
	Identifier           string                                            `json:"identifier"`
	SkillConfigurationId *string                                           `json:"skill_configuration_id,omitempty"`
	SkillMarketplaceId   *string                                           `json:"skill_marketplace_id,omitempty"`
	SkillPluginId        *string                                           `json:"skill_plugin_id,omitempty"`
	SkillPlugin          *SkillsMarketplacesPluginsRemoveOutputSkillPlugin `json:"skill_plugin,omitempty"`
	CreatedAt            time.Time                                         `json:"created_at"`
	UpdatedAt            time.Time                                         `json:"updated_at"`
}

// MapSkillsMarketplacesPluginsRemoveOutputFromJSON deserializes JSON data into a SkillsMarketplacesPluginsRemoveOutput.
func MapSkillsMarketplacesPluginsRemoveOutputFromJSON(data []byte) (*SkillsMarketplacesPluginsRemoveOutput, error) {
	var v SkillsMarketplacesPluginsRemoveOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMarketplacesPluginsRemoveOutputToJSON serializes a SkillsMarketplacesPluginsRemoveOutput to JSON.
func MapSkillsMarketplacesPluginsRemoveOutputToJSON(v *SkillsMarketplacesPluginsRemoveOutput) ([]byte, error) {
	return json.Marshal(v)
}
