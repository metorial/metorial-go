package plugins

import (
	"encoding/json"
	"time"
)

// SkillsPluginsGetOutputSkills represents the skills plugins get output skills type.
type SkillsPluginsGetOutputSkills struct {
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

// SkillsPluginsGetOutput represents the skills plugins get output type.
type SkillsPluginsGetOutput struct {
	Object               string                         `json:"object"`
	Id                   string                         `json:"id"`
	Status               string                         `json:"status"`
	SyncStatus           string                         `json:"sync_status"`
	ImageUrl             string                         `json:"image_url"`
	Name                 string                         `json:"name"`
	Description          *string                        `json:"description,omitempty"`
	LongDescription      *string                        `json:"long_description,omitempty"`
	Category             *string                        `json:"category,omitempty"`
	Slug                 string                         `json:"slug"`
	SkillConfigurationId *string                        `json:"skill_configuration_id,omitempty"`
	Skills               []SkillsPluginsGetOutputSkills `json:"skills"`
	CreatedAt            time.Time                      `json:"created_at"`
	UpdatedAt            time.Time                      `json:"updated_at"`
}

// MapSkillsPluginsGetOutputFromJSON deserializes JSON data into a SkillsPluginsGetOutput.
func MapSkillsPluginsGetOutputFromJSON(data []byte) (*SkillsPluginsGetOutput, error) {
	var v SkillsPluginsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsGetOutputToJSON serializes a SkillsPluginsGetOutput to JSON.
func MapSkillsPluginsGetOutputToJSON(v *SkillsPluginsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
