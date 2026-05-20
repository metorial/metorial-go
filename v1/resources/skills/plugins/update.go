package plugins

import (
	"encoding/json"
	"time"
)

// SkillsPluginsUpdateOutputSkills represents the skills plugins update output skills type.
type SkillsPluginsUpdateOutputSkills struct {
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

// SkillsPluginsUpdateOutput represents the skills plugins update output type.
type SkillsPluginsUpdateOutput struct {
	Object               string                            `json:"object"`
	Id                   string                            `json:"id"`
	Status               string                            `json:"status"`
	SyncStatus           string                            `json:"sync_status"`
	ImageUrl             string                            `json:"image_url"`
	Name                 string                            `json:"name"`
	Description          *string                           `json:"description,omitempty"`
	LongDescription      *string                           `json:"long_description,omitempty"`
	Category             *string                           `json:"category,omitempty"`
	Slug                 string                            `json:"slug"`
	SkillConfigurationId *string                           `json:"skill_configuration_id,omitempty"`
	Skills               []SkillsPluginsUpdateOutputSkills `json:"skills"`
	CreatedAt            time.Time                         `json:"created_at"`
	UpdatedAt            time.Time                         `json:"updated_at"`
}

// MapSkillsPluginsUpdateOutputFromJSON deserializes JSON data into a SkillsPluginsUpdateOutput.
func MapSkillsPluginsUpdateOutputFromJSON(data []byte) (*SkillsPluginsUpdateOutput, error) {
	var v SkillsPluginsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsUpdateOutputToJSON serializes a SkillsPluginsUpdateOutput to JSON.
func MapSkillsPluginsUpdateOutputToJSON(v *SkillsPluginsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsPluginsUpdateBody represents the skills plugins update body type.
type SkillsPluginsUpdateBody struct {
	Name                 *string `json:"name,omitempty"`
	Description          *string `json:"description,omitempty"`
	LongDescription      *string `json:"long_description,omitempty"`
	Category             *string `json:"category,omitempty"`
	ImageFileId          *string `json:"image_file_id,omitempty"`
	SkillConfigurationId *string `json:"skill_configuration_id,omitempty"`
}

// MapSkillsPluginsUpdateBodyFromJSON deserializes JSON data into a SkillsPluginsUpdateBody.
func MapSkillsPluginsUpdateBodyFromJSON(data []byte) (*SkillsPluginsUpdateBody, error) {
	var v SkillsPluginsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsUpdateBodyToJSON serializes a SkillsPluginsUpdateBody to JSON.
func MapSkillsPluginsUpdateBodyToJSON(v *SkillsPluginsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
