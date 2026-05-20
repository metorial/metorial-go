package configurations

import (
	"encoding/json"
	"time"
)

// SkillsConfigurationsDeleteOutput represents the skills configurations delete output type.
type SkillsConfigurationsDeleteOutput struct {
	// Object - String representing the object's type
	Object                      string     `json:"object"`
	Id                          string     `json:"id"`
	IsDefault                   bool       `json:"is_default"`
	AllowScripts                bool       `json:"allow_scripts"`
	AllowedFileExtensions       []string   `json:"allowed_file_extensions"`
	AllowNonStandardDirectories bool       `json:"allow_non_standard_directories"`
	DeletedAt                   *time.Time `json:"deleted_at,omitempty"`
	CreatedAt                   time.Time  `json:"created_at"`
	UpdatedAt                   time.Time  `json:"updated_at"`
}

// MapSkillsConfigurationsDeleteOutputFromJSON deserializes JSON data into a SkillsConfigurationsDeleteOutput.
func MapSkillsConfigurationsDeleteOutputFromJSON(data []byte) (*SkillsConfigurationsDeleteOutput, error) {
	var v SkillsConfigurationsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsConfigurationsDeleteOutputToJSON serializes a SkillsConfigurationsDeleteOutput to JSON.
func MapSkillsConfigurationsDeleteOutputToJSON(v *SkillsConfigurationsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
