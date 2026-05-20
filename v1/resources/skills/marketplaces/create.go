package marketplaces

import (
	"encoding/json"
	"time"
)

// SkillsMarketplacesCreateOutputPluginsSkillPluginSkills represents the skills marketplaces create output plugins skill plugin skills type.
type SkillsMarketplacesCreateOutputPluginsSkillPluginSkills struct {
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

// SkillsMarketplacesCreateOutputPluginsSkillPlugin represents the skills marketplaces create output plugins skill plugin type.
type SkillsMarketplacesCreateOutputPluginsSkillPlugin struct {
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
	Skills               []SkillsMarketplacesCreateOutputPluginsSkillPluginSkills `json:"skills"`
	CreatedAt            time.Time                                                `json:"created_at"`
	UpdatedAt            time.Time                                                `json:"updated_at"`
}

// SkillsMarketplacesCreateOutputPlugins represents the skills marketplaces create output plugins type.
type SkillsMarketplacesCreateOutputPlugins struct {
	Object               string                                            `json:"object"`
	Id                   string                                            `json:"id"`
	Status               string                                            `json:"status"`
	Identifier           string                                            `json:"identifier"`
	SkillConfigurationId *string                                           `json:"skill_configuration_id,omitempty"`
	SkillMarketplaceId   *string                                           `json:"skill_marketplace_id,omitempty"`
	SkillPluginId        *string                                           `json:"skill_plugin_id,omitempty"`
	SkillPlugin          *SkillsMarketplacesCreateOutputPluginsSkillPlugin `json:"skill_plugin,omitempty"`
	CreatedAt            time.Time                                         `json:"created_at"`
	UpdatedAt            time.Time                                         `json:"updated_at"`
}

// SkillsMarketplacesCreateOutput represents the skills marketplaces create output type.
type SkillsMarketplacesCreateOutput struct {
	Object               string                                  `json:"object"`
	Id                   string                                  `json:"id"`
	Status               string                                  `json:"status"`
	SyncStatus           string                                  `json:"sync_status"`
	ImageUrl             string                                  `json:"image_url"`
	Name                 string                                  `json:"name"`
	Description          *string                                 `json:"description,omitempty"`
	Slug                 string                                  `json:"slug"`
	SkillConfigurationId *string                                 `json:"skill_configuration_id,omitempty"`
	Plugins              []SkillsMarketplacesCreateOutputPlugins `json:"plugins"`
	CreatedAt            time.Time                               `json:"created_at"`
	UpdatedAt            time.Time                               `json:"updated_at"`
}

// MapSkillsMarketplacesCreateOutputFromJSON deserializes JSON data into a SkillsMarketplacesCreateOutput.
func MapSkillsMarketplacesCreateOutputFromJSON(data []byte) (*SkillsMarketplacesCreateOutput, error) {
	var v SkillsMarketplacesCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMarketplacesCreateOutputToJSON serializes a SkillsMarketplacesCreateOutput to JSON.
func MapSkillsMarketplacesCreateOutputToJSON(v *SkillsMarketplacesCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsMarketplacesCreateBody represents the skills marketplaces create body type.
type SkillsMarketplacesCreateBody struct {
	Name                 string  `json:"name"`
	Description          *string `json:"description,omitempty"`
	ImageFileId          *string `json:"image_file_id,omitempty"`
	SkillConfigurationId *string `json:"skill_configuration_id,omitempty"`
}

// MapSkillsMarketplacesCreateBodyFromJSON deserializes JSON data into a SkillsMarketplacesCreateBody.
func MapSkillsMarketplacesCreateBodyFromJSON(data []byte) (*SkillsMarketplacesCreateBody, error) {
	var v SkillsMarketplacesCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMarketplacesCreateBodyToJSON serializes a SkillsMarketplacesCreateBody to JSON.
func MapSkillsMarketplacesCreateBodyToJSON(v *SkillsMarketplacesCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
