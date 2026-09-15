package marketplaces

import (
	"encoding/json"
	"time"
)

// SkillsMarketplacesArchiveOutputPluginsSkillPluginSkillsSkill represents the skills marketplaces archive output plugins skill plugin skills skill type.
type SkillsMarketplacesArchiveOutputPluginsSkillPluginSkillsSkill struct {
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

// SkillsMarketplacesArchiveOutputPluginsSkillPluginSkills represents the skills marketplaces archive output plugins skill plugin skills type.
type SkillsMarketplacesArchiveOutputPluginsSkillPluginSkills struct {
	Object               string                                                       `json:"object"`
	Id                   string                                                       `json:"id"`
	Identifier           string                                                       `json:"identifier"`
	Status               string                                                       `json:"status"`
	ClientName           *string                                                      `json:"client_name,omitempty"`
	ClientDescription    *string                                                      `json:"client_description,omitempty"`
	ClientMetadata       *map[string]any                                              `json:"client_metadata,omitempty"`
	License              *string                                                      `json:"license,omitempty"`
	Compatibility        *string                                                      `json:"compatibility,omitempty"`
	SkillConfigurationId *string                                                      `json:"skill_configuration_id,omitempty"`
	SkillId              string                                                       `json:"skill_id"`
	Skill                SkillsMarketplacesArchiveOutputPluginsSkillPluginSkillsSkill `json:"skill"`
	CreatedAt            time.Time                                                    `json:"created_at"`
	UpdatedAt            time.Time                                                    `json:"updated_at"`
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
	SkillPlugin          *SkillsMarketplacesArchiveOutputPluginsSkillPlugin `json:"skill_plugin,omitempty"`
	CreatedAt            time.Time                                          `json:"created_at"`
	UpdatedAt            time.Time                                          `json:"updated_at"`
}

// SkillsMarketplacesArchiveOutput represents the skills marketplaces archive output type.
type SkillsMarketplacesArchiveOutput struct {
	Object                string                                   `json:"object"`
	Id                    string                                   `json:"id"`
	Status                string                                   `json:"status"`
	RepositoryAccessMode  string                                   `json:"repository_access_mode"`
	ForceMergeOrPush      bool                                     `json:"force_merge_or_push"`
	MergeBeforeChecksPass bool                                     `json:"merge_before_checks_pass"`
	SyncStatus            string                                   `json:"sync_status"`
	ImageUrl              string                                   `json:"image_url"`
	Name                  string                                   `json:"name"`
	Description           *string                                  `json:"description,omitempty"`
	Slug                  string                                   `json:"slug"`
	SkillConfigurationId  *string                                  `json:"skill_configuration_id,omitempty"`
	Plugins               []SkillsMarketplacesArchiveOutputPlugins `json:"plugins"`
	CreatedAt             time.Time                                `json:"created_at"`
	UpdatedAt             time.Time                                `json:"updated_at"`
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
