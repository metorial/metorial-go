package skills

import (
	"encoding/json"
	"time"
)

// SkillsPluginsSkillsGetOutput represents the skills plugins skills get output type.
type SkillsPluginsSkillsGetOutput struct {
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

// MapSkillsPluginsSkillsGetOutputFromJSON deserializes JSON data into a SkillsPluginsSkillsGetOutput.
func MapSkillsPluginsSkillsGetOutputFromJSON(data []byte) (*SkillsPluginsSkillsGetOutput, error) {
	var v SkillsPluginsSkillsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsSkillsGetOutputToJSON serializes a SkillsPluginsSkillsGetOutput to JSON.
func MapSkillsPluginsSkillsGetOutputToJSON(v *SkillsPluginsSkillsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
