package skills

import (
	"encoding/json"
	"time"
)

// SkillsPluginsSkillsRemoveOutput represents the skills plugins skills remove output type.
type SkillsPluginsSkillsRemoveOutput struct {
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

// MapSkillsPluginsSkillsRemoveOutputFromJSON deserializes JSON data into a SkillsPluginsSkillsRemoveOutput.
func MapSkillsPluginsSkillsRemoveOutputFromJSON(data []byte) (*SkillsPluginsSkillsRemoveOutput, error) {
	var v SkillsPluginsSkillsRemoveOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsSkillsRemoveOutputToJSON serializes a SkillsPluginsSkillsRemoveOutput to JSON.
func MapSkillsPluginsSkillsRemoveOutputToJSON(v *SkillsPluginsSkillsRemoveOutput) ([]byte, error) {
	return json.Marshal(v)
}
