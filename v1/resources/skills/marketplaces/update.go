package marketplaces

import (
	"encoding/json"
	"time"
)

// SkillsMarketplacesUpdateOutputPluginsSkillPluginSkills represents the skills marketplaces update output plugins skill plugin skills type.
type SkillsMarketplacesUpdateOutputPluginsSkillPluginSkills struct {
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

// SkillsMarketplacesUpdateOutputPluginsSkillPlugin represents the skills marketplaces update output plugins skill plugin type.
type SkillsMarketplacesUpdateOutputPluginsSkillPlugin struct {
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
	Skills               []SkillsMarketplacesUpdateOutputPluginsSkillPluginSkills `json:"skills"`
	CreatedAt            time.Time                                                `json:"created_at"`
	UpdatedAt            time.Time                                                `json:"updated_at"`
}

// SkillsMarketplacesUpdateOutputPlugins represents the skills marketplaces update output plugins type.
type SkillsMarketplacesUpdateOutputPlugins struct {
	Object               string                                            `json:"object"`
	Id                   string                                            `json:"id"`
	Status               string                                            `json:"status"`
	Identifier           string                                            `json:"identifier"`
	SkillConfigurationId *string                                           `json:"skill_configuration_id,omitempty"`
	SkillMarketplaceId   *string                                           `json:"skill_marketplace_id,omitempty"`
	SkillPluginId        *string                                           `json:"skill_plugin_id,omitempty"`
	SkillPlugin          *SkillsMarketplacesUpdateOutputPluginsSkillPlugin `json:"skill_plugin,omitempty"`
	CreatedAt            time.Time                                         `json:"created_at"`
	UpdatedAt            time.Time                                         `json:"updated_at"`
}

// SkillsMarketplacesUpdateOutput represents the skills marketplaces update output type.
type SkillsMarketplacesUpdateOutput struct {
	Object               string                                  `json:"object"`
	Id                   string                                  `json:"id"`
	Status               string                                  `json:"status"`
	SyncStatus           string                                  `json:"sync_status"`
	ImageUrl             string                                  `json:"image_url"`
	Name                 string                                  `json:"name"`
	Description          *string                                 `json:"description,omitempty"`
	Slug                 string                                  `json:"slug"`
	SkillConfigurationId *string                                 `json:"skill_configuration_id,omitempty"`
	Plugins              []SkillsMarketplacesUpdateOutputPlugins `json:"plugins"`
	CreatedAt            time.Time                               `json:"created_at"`
	UpdatedAt            time.Time                               `json:"updated_at"`
}

// MapSkillsMarketplacesUpdateOutputFromJSON deserializes JSON data into a SkillsMarketplacesUpdateOutput.
func MapSkillsMarketplacesUpdateOutputFromJSON(data []byte) (*SkillsMarketplacesUpdateOutput, error) {
	var v SkillsMarketplacesUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMarketplacesUpdateOutputToJSON serializes a SkillsMarketplacesUpdateOutput to JSON.
func MapSkillsMarketplacesUpdateOutputToJSON(v *SkillsMarketplacesUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsMarketplacesUpdateBody represents the skills marketplaces update body type.
type SkillsMarketplacesUpdateBody struct {
	Name                 *string `json:"name,omitempty"`
	Description          *string `json:"description,omitempty"`
	ImageFileId          *string `json:"image_file_id,omitempty"`
	SkillConfigurationId *string `json:"skill_configuration_id,omitempty"`
}

// MapSkillsMarketplacesUpdateBodyFromJSON deserializes JSON data into a SkillsMarketplacesUpdateBody.
func MapSkillsMarketplacesUpdateBodyFromJSON(data []byte) (*SkillsMarketplacesUpdateBody, error) {
	var v SkillsMarketplacesUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMarketplacesUpdateBodyToJSON serializes a SkillsMarketplacesUpdateBody to JSON.
func MapSkillsMarketplacesUpdateBodyToJSON(v *SkillsMarketplacesUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
