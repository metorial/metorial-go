package skills

import (
	"encoding/json"
	"time"
)

// SkillsPluginsSkillsAddOutput represents the skills plugins skills add output type.
type SkillsPluginsSkillsAddOutput struct {
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

// MapSkillsPluginsSkillsAddOutputFromJSON deserializes JSON data into a SkillsPluginsSkillsAddOutput.
func MapSkillsPluginsSkillsAddOutputFromJSON(data []byte) (*SkillsPluginsSkillsAddOutput, error) {
	var v SkillsPluginsSkillsAddOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsSkillsAddOutputToJSON serializes a SkillsPluginsSkillsAddOutput to JSON.
func MapSkillsPluginsSkillsAddOutputToJSON(v *SkillsPluginsSkillsAddOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsPluginsSkillsAddBody represents the skills plugins skills add body type.
type SkillsPluginsSkillsAddBody struct {
	SkillId              string          `json:"skill_id"`
	Identifier           *string         `json:"identifier,omitempty"`
	ClientName           *string         `json:"client_name,omitempty"`
	ClientDescription    *string         `json:"client_description,omitempty"`
	ClientMetadata       *map[string]any `json:"client_metadata,omitempty"`
	License              *string         `json:"license,omitempty"`
	Compatibility        *string         `json:"compatibility,omitempty"`
	SkillConfigurationId *string         `json:"skill_configuration_id,omitempty"`
}

// MapSkillsPluginsSkillsAddBodyFromJSON deserializes JSON data into a SkillsPluginsSkillsAddBody.
func MapSkillsPluginsSkillsAddBodyFromJSON(data []byte) (*SkillsPluginsSkillsAddBody, error) {
	var v SkillsPluginsSkillsAddBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsSkillsAddBodyToJSON serializes a SkillsPluginsSkillsAddBody to JSON.
func MapSkillsPluginsSkillsAddBodyToJSON(v *SkillsPluginsSkillsAddBody) ([]byte, error) {
	return json.Marshal(v)
}
