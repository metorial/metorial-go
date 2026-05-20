package plugins

import (
	"encoding/json"
	"time"
)

// SkillsPluginsCreateOutputSkills represents the skills plugins create output skills type.
type SkillsPluginsCreateOutputSkills struct {
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

// SkillsPluginsCreateOutput represents the skills plugins create output type.
type SkillsPluginsCreateOutput struct {
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
	Skills               []SkillsPluginsCreateOutputSkills `json:"skills"`
	CreatedAt            time.Time                         `json:"created_at"`
	UpdatedAt            time.Time                         `json:"updated_at"`
}

// MapSkillsPluginsCreateOutputFromJSON deserializes JSON data into a SkillsPluginsCreateOutput.
func MapSkillsPluginsCreateOutputFromJSON(data []byte) (*SkillsPluginsCreateOutput, error) {
	var v SkillsPluginsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsCreateOutputToJSON serializes a SkillsPluginsCreateOutput to JSON.
func MapSkillsPluginsCreateOutputToJSON(v *SkillsPluginsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsPluginsCreateBody represents the skills plugins create body type.
type SkillsPluginsCreateBody struct {
	Name                 string  `json:"name"`
	Description          *string `json:"description,omitempty"`
	LongDescription      *string `json:"long_description,omitempty"`
	Category             *string `json:"category,omitempty"`
	ImageFileId          *string `json:"image_file_id,omitempty"`
	SkillConfigurationId *string `json:"skill_configuration_id,omitempty"`
}

// MapSkillsPluginsCreateBodyFromJSON deserializes JSON data into a SkillsPluginsCreateBody.
func MapSkillsPluginsCreateBodyFromJSON(data []byte) (*SkillsPluginsCreateBody, error) {
	var v SkillsPluginsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsCreateBodyToJSON serializes a SkillsPluginsCreateBody to JSON.
func MapSkillsPluginsCreateBodyToJSON(v *SkillsPluginsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
