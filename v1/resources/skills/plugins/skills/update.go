package skills

import (
	"encoding/json"
	"time"
)

// SkillsPluginsSkillsUpdateOutputSkill represents the skills plugins skills update output skill type.
type SkillsPluginsSkillsUpdateOutputSkill struct {
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

// SkillsPluginsSkillsUpdateOutput represents the skills plugins skills update output type.
type SkillsPluginsSkillsUpdateOutput struct {
	Object               string                               `json:"object"`
	Id                   string                               `json:"id"`
	Identifier           string                               `json:"identifier"`
	Status               string                               `json:"status"`
	ClientName           *string                              `json:"client_name,omitempty"`
	ClientDescription    *string                              `json:"client_description,omitempty"`
	ClientMetadata       *map[string]any                      `json:"client_metadata,omitempty"`
	License              *string                              `json:"license,omitempty"`
	Compatibility        *string                              `json:"compatibility,omitempty"`
	SkillConfigurationId *string                              `json:"skill_configuration_id,omitempty"`
	SkillId              string                               `json:"skill_id"`
	Skill                SkillsPluginsSkillsUpdateOutputSkill `json:"skill"`
	CreatedAt            time.Time                            `json:"created_at"`
	UpdatedAt            time.Time                            `json:"updated_at"`
}

// MapSkillsPluginsSkillsUpdateOutputFromJSON deserializes JSON data into a SkillsPluginsSkillsUpdateOutput.
func MapSkillsPluginsSkillsUpdateOutputFromJSON(data []byte) (*SkillsPluginsSkillsUpdateOutput, error) {
	var v SkillsPluginsSkillsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsSkillsUpdateOutputToJSON serializes a SkillsPluginsSkillsUpdateOutput to JSON.
func MapSkillsPluginsSkillsUpdateOutputToJSON(v *SkillsPluginsSkillsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsPluginsSkillsUpdateBody represents the skills plugins skills update body type.
type SkillsPluginsSkillsUpdateBody struct {
	ClientName           *string         `json:"client_name,omitempty"`
	ClientDescription    *string         `json:"client_description,omitempty"`
	ClientMetadata       *map[string]any `json:"client_metadata,omitempty"`
	License              *string         `json:"license,omitempty"`
	Compatibility        *string         `json:"compatibility,omitempty"`
	SkillConfigurationId *string         `json:"skill_configuration_id,omitempty"`
}

// MapSkillsPluginsSkillsUpdateBodyFromJSON deserializes JSON data into a SkillsPluginsSkillsUpdateBody.
func MapSkillsPluginsSkillsUpdateBodyFromJSON(data []byte) (*SkillsPluginsSkillsUpdateBody, error) {
	var v SkillsPluginsSkillsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsSkillsUpdateBodyToJSON serializes a SkillsPluginsSkillsUpdateBody to JSON.
func MapSkillsPluginsSkillsUpdateBodyToJSON(v *SkillsPluginsSkillsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
