package plugins

import (
	"encoding/json"
	"time"
)

// SkillsPluginsArchiveOutputSkills represents the skills plugins archive output skills type.
type SkillsPluginsArchiveOutputSkills struct {
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

// SkillsPluginsArchiveOutput represents the skills plugins archive output type.
type SkillsPluginsArchiveOutput struct {
	Object               string                             `json:"object"`
	Id                   string                             `json:"id"`
	Status               string                             `json:"status"`
	SyncStatus           string                             `json:"sync_status"`
	ImageUrl             string                             `json:"image_url"`
	Name                 string                             `json:"name"`
	Description          *string                            `json:"description,omitempty"`
	LongDescription      *string                            `json:"long_description,omitempty"`
	Category             *string                            `json:"category,omitempty"`
	Slug                 string                             `json:"slug"`
	SkillConfigurationId *string                            `json:"skill_configuration_id,omitempty"`
	Skills               []SkillsPluginsArchiveOutputSkills `json:"skills"`
	CreatedAt            time.Time                          `json:"created_at"`
	UpdatedAt            time.Time                          `json:"updated_at"`
}

// MapSkillsPluginsArchiveOutputFromJSON deserializes JSON data into a SkillsPluginsArchiveOutput.
func MapSkillsPluginsArchiveOutputFromJSON(data []byte) (*SkillsPluginsArchiveOutput, error) {
	var v SkillsPluginsArchiveOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsArchiveOutputToJSON serializes a SkillsPluginsArchiveOutput to JSON.
func MapSkillsPluginsArchiveOutputToJSON(v *SkillsPluginsArchiveOutput) ([]byte, error) {
	return json.Marshal(v)
}
