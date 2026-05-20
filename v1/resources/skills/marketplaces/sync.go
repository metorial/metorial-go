package marketplaces

import (
	"encoding/json"
	"time"
)

// SkillsMarketplacesSyncOutputPluginsSkillPluginSkills represents the skills marketplaces sync output plugins skill plugin skills type.
type SkillsMarketplacesSyncOutputPluginsSkillPluginSkills struct {
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

// SkillsMarketplacesSyncOutputPluginsSkillPlugin represents the skills marketplaces sync output plugins skill plugin type.
type SkillsMarketplacesSyncOutputPluginsSkillPlugin struct {
	Object               string                                                 `json:"object"`
	Id                   string                                                 `json:"id"`
	Status               string                                                 `json:"status"`
	SyncStatus           string                                                 `json:"sync_status"`
	ImageUrl             string                                                 `json:"image_url"`
	Name                 string                                                 `json:"name"`
	Description          *string                                                `json:"description,omitempty"`
	LongDescription      *string                                                `json:"long_description,omitempty"`
	Category             *string                                                `json:"category,omitempty"`
	Slug                 string                                                 `json:"slug"`
	SkillConfigurationId *string                                                `json:"skill_configuration_id,omitempty"`
	Skills               []SkillsMarketplacesSyncOutputPluginsSkillPluginSkills `json:"skills"`
	CreatedAt            time.Time                                              `json:"created_at"`
	UpdatedAt            time.Time                                              `json:"updated_at"`
}

// SkillsMarketplacesSyncOutputPlugins represents the skills marketplaces sync output plugins type.
type SkillsMarketplacesSyncOutputPlugins struct {
	Object               string                                          `json:"object"`
	Id                   string                                          `json:"id"`
	Status               string                                          `json:"status"`
	Identifier           string                                          `json:"identifier"`
	SkillConfigurationId *string                                         `json:"skill_configuration_id,omitempty"`
	SkillMarketplaceId   *string                                         `json:"skill_marketplace_id,omitempty"`
	SkillPluginId        *string                                         `json:"skill_plugin_id,omitempty"`
	SkillPlugin          *SkillsMarketplacesSyncOutputPluginsSkillPlugin `json:"skill_plugin,omitempty"`
	CreatedAt            time.Time                                       `json:"created_at"`
	UpdatedAt            time.Time                                       `json:"updated_at"`
}

// SkillsMarketplacesSyncOutput represents the skills marketplaces sync output type.
type SkillsMarketplacesSyncOutput struct {
	Object               string                                `json:"object"`
	Id                   string                                `json:"id"`
	Status               string                                `json:"status"`
	SyncStatus           string                                `json:"sync_status"`
	ImageUrl             string                                `json:"image_url"`
	Name                 string                                `json:"name"`
	Description          *string                               `json:"description,omitempty"`
	Slug                 string                                `json:"slug"`
	SkillConfigurationId *string                               `json:"skill_configuration_id,omitempty"`
	Plugins              []SkillsMarketplacesSyncOutputPlugins `json:"plugins"`
	CreatedAt            time.Time                             `json:"created_at"`
	UpdatedAt            time.Time                             `json:"updated_at"`
}

// MapSkillsMarketplacesSyncOutputFromJSON deserializes JSON data into a SkillsMarketplacesSyncOutput.
func MapSkillsMarketplacesSyncOutputFromJSON(data []byte) (*SkillsMarketplacesSyncOutput, error) {
	var v SkillsMarketplacesSyncOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMarketplacesSyncOutputToJSON serializes a SkillsMarketplacesSyncOutput to JSON.
func MapSkillsMarketplacesSyncOutputToJSON(v *SkillsMarketplacesSyncOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsMarketplacesSyncBody represents the skills marketplaces sync body type.
type SkillsMarketplacesSyncBody struct{}

// MapSkillsMarketplacesSyncBodyFromJSON deserializes JSON data into a SkillsMarketplacesSyncBody.
func MapSkillsMarketplacesSyncBodyFromJSON(data []byte) (*SkillsMarketplacesSyncBody, error) {
	var v SkillsMarketplacesSyncBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMarketplacesSyncBodyToJSON serializes a SkillsMarketplacesSyncBody to JSON.
func MapSkillsMarketplacesSyncBodyToJSON(v *SkillsMarketplacesSyncBody) ([]byte, error) {
	return json.Marshal(v)
}
