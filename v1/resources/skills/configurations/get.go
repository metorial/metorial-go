package configurations

import (
	"encoding/json"
	"time"
)

// SkillsConfigurationsGetOutput represents the skills configurations get output type.
type SkillsConfigurationsGetOutput struct {
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

// MapSkillsConfigurationsGetOutputFromJSON deserializes JSON data into a SkillsConfigurationsGetOutput.
func MapSkillsConfigurationsGetOutputFromJSON(data []byte) (*SkillsConfigurationsGetOutput, error) {
	var v SkillsConfigurationsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsConfigurationsGetOutputToJSON serializes a SkillsConfigurationsGetOutput to JSON.
func MapSkillsConfigurationsGetOutputToJSON(v *SkillsConfigurationsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
