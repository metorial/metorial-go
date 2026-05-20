package plugins

import (
	"encoding/json"
	"time"
)

// SkillsMarketplacesPluginsAddOutputSkillPluginSkills represents the skills marketplaces plugins add output skill plugin skills type.
type SkillsMarketplacesPluginsAddOutputSkillPluginSkills struct {
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

// SkillsMarketplacesPluginsAddOutputSkillPlugin represents the skills marketplaces plugins add output skill plugin type.
type SkillsMarketplacesPluginsAddOutputSkillPlugin struct {
	Object               string                                                `json:"object"`
	Id                   string                                                `json:"id"`
	Status               string                                                `json:"status"`
	SyncStatus           string                                                `json:"sync_status"`
	ImageUrl             string                                                `json:"image_url"`
	Name                 string                                                `json:"name"`
	Description          *string                                               `json:"description,omitempty"`
	LongDescription      *string                                               `json:"long_description,omitempty"`
	Category             *string                                               `json:"category,omitempty"`
	Slug                 string                                                `json:"slug"`
	SkillConfigurationId *string                                               `json:"skill_configuration_id,omitempty"`
	Skills               []SkillsMarketplacesPluginsAddOutputSkillPluginSkills `json:"skills"`
	CreatedAt            time.Time                                             `json:"created_at"`
	UpdatedAt            time.Time                                             `json:"updated_at"`
}

// SkillsMarketplacesPluginsAddOutput represents the skills marketplaces plugins add output type.
type SkillsMarketplacesPluginsAddOutput struct {
	Object               string                                         `json:"object"`
	Id                   string                                         `json:"id"`
	Status               string                                         `json:"status"`
	Identifier           string                                         `json:"identifier"`
	SkillConfigurationId *string                                        `json:"skill_configuration_id,omitempty"`
	SkillMarketplaceId   *string                                        `json:"skill_marketplace_id,omitempty"`
	SkillPluginId        *string                                        `json:"skill_plugin_id,omitempty"`
	SkillPlugin          *SkillsMarketplacesPluginsAddOutputSkillPlugin `json:"skill_plugin,omitempty"`
	CreatedAt            time.Time                                      `json:"created_at"`
	UpdatedAt            time.Time                                      `json:"updated_at"`
}

// MapSkillsMarketplacesPluginsAddOutputFromJSON deserializes JSON data into a SkillsMarketplacesPluginsAddOutput.
func MapSkillsMarketplacesPluginsAddOutputFromJSON(data []byte) (*SkillsMarketplacesPluginsAddOutput, error) {
	var v SkillsMarketplacesPluginsAddOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMarketplacesPluginsAddOutputToJSON serializes a SkillsMarketplacesPluginsAddOutput to JSON.
func MapSkillsMarketplacesPluginsAddOutputToJSON(v *SkillsMarketplacesPluginsAddOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsMarketplacesPluginsAddBody represents the skills marketplaces plugins add body type.
type SkillsMarketplacesPluginsAddBody struct {
	SkillPluginId        string  `json:"skill_plugin_id"`
	SkillConfigurationId *string `json:"skill_configuration_id,omitempty"`
	Identifier           *string `json:"identifier,omitempty"`
}

// MapSkillsMarketplacesPluginsAddBodyFromJSON deserializes JSON data into a SkillsMarketplacesPluginsAddBody.
func MapSkillsMarketplacesPluginsAddBodyFromJSON(data []byte) (*SkillsMarketplacesPluginsAddBody, error) {
	var v SkillsMarketplacesPluginsAddBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMarketplacesPluginsAddBodyToJSON serializes a SkillsMarketplacesPluginsAddBody to JSON.
func MapSkillsMarketplacesPluginsAddBodyToJSON(v *SkillsMarketplacesPluginsAddBody) ([]byte, error) {
	return json.Marshal(v)
}
