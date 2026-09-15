package marketplaces

import (
	"encoding/json"
	"time"
)

// SkillsMarketplacesGetOutputPluginsSkillPluginSkillsSkill represents the skills marketplaces get output plugins skill plugin skills skill type.
type SkillsMarketplacesGetOutputPluginsSkillPluginSkillsSkill struct {
	Object            string          `json:"object"`
	Id                string          `json:"id"`
	Status            string          `json:"status"`
	Slug              string          `json:"slug"`
	Name              string          `json:"name"`
	Description       *string         `json:"description,omitempty"`
	ImageUrl          string          `json:"image_url"`
	ClientName        string          `json:"client_name"`
	ClientDescription *string         `json:"client_description,omitempty"`
	ClientMetadata    *map[string]any `json:"client_metadata,omitempty"`
	License           *string         `json:"license,omitempty"`
	Compatibility     *string         `json:"compatibility,omitempty"`
	Metadata          *map[string]any `json:"metadata,omitempty"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
}

// SkillsMarketplacesGetOutputPluginsSkillPluginSkills represents the skills marketplaces get output plugins skill plugin skills type.
type SkillsMarketplacesGetOutputPluginsSkillPluginSkills struct {
	Object               string                                                   `json:"object"`
	Id                   string                                                   `json:"id"`
	Identifier           string                                                   `json:"identifier"`
	Status               string                                                   `json:"status"`
	ClientName           *string                                                  `json:"client_name,omitempty"`
	ClientDescription    *string                                                  `json:"client_description,omitempty"`
	ClientMetadata       *map[string]any                                          `json:"client_metadata,omitempty"`
	License              *string                                                  `json:"license,omitempty"`
	Compatibility        *string                                                  `json:"compatibility,omitempty"`
	SkillConfigurationId *string                                                  `json:"skill_configuration_id,omitempty"`
	SkillId              string                                                   `json:"skill_id"`
	Skill                SkillsMarketplacesGetOutputPluginsSkillPluginSkillsSkill `json:"skill"`
	CreatedAt            time.Time                                                `json:"created_at"`
	UpdatedAt            time.Time                                                `json:"updated_at"`
}

// SkillsMarketplacesGetOutputPluginsSkillPlugin represents the skills marketplaces get output plugins skill plugin type.
type SkillsMarketplacesGetOutputPluginsSkillPlugin struct {
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
	Skills               []SkillsMarketplacesGetOutputPluginsSkillPluginSkills `json:"skills"`
	CreatedAt            time.Time                                             `json:"created_at"`
	UpdatedAt            time.Time                                             `json:"updated_at"`
}

// SkillsMarketplacesGetOutputPlugins represents the skills marketplaces get output plugins type.
type SkillsMarketplacesGetOutputPlugins struct {
	Object               string                                         `json:"object"`
	Id                   string                                         `json:"id"`
	Status               string                                         `json:"status"`
	Identifier           string                                         `json:"identifier"`
	SkillConfigurationId *string                                        `json:"skill_configuration_id,omitempty"`
	SkillMarketplaceId   *string                                        `json:"skill_marketplace_id,omitempty"`
	SkillPlugin          *SkillsMarketplacesGetOutputPluginsSkillPlugin `json:"skill_plugin,omitempty"`
	CreatedAt            time.Time                                      `json:"created_at"`
	UpdatedAt            time.Time                                      `json:"updated_at"`
}

// SkillsMarketplacesGetOutput represents the skills marketplaces get output type.
type SkillsMarketplacesGetOutput struct {
	Object                string                               `json:"object"`
	Id                    string                               `json:"id"`
	Status                string                               `json:"status"`
	RepositoryAccessMode  string                               `json:"repository_access_mode"`
	ForceMergeOrPush      bool                                 `json:"force_merge_or_push"`
	MergeBeforeChecksPass bool                                 `json:"merge_before_checks_pass"`
	SyncStatus            string                               `json:"sync_status"`
	ImageUrl              string                               `json:"image_url"`
	Name                  string                               `json:"name"`
	Description           *string                              `json:"description,omitempty"`
	Slug                  string                               `json:"slug"`
	SkillConfigurationId  *string                              `json:"skill_configuration_id,omitempty"`
	Plugins               []SkillsMarketplacesGetOutputPlugins `json:"plugins"`
	CreatedAt             time.Time                            `json:"created_at"`
	UpdatedAt             time.Time                            `json:"updated_at"`
}

// MapSkillsMarketplacesGetOutputFromJSON deserializes JSON data into a SkillsMarketplacesGetOutput.
func MapSkillsMarketplacesGetOutputFromJSON(data []byte) (*SkillsMarketplacesGetOutput, error) {
	var v SkillsMarketplacesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMarketplacesGetOutputToJSON serializes a SkillsMarketplacesGetOutput to JSON.
func MapSkillsMarketplacesGetOutputToJSON(v *SkillsMarketplacesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
