package configurations

import (
	"encoding/json"
	"time"
)

// SkillsConfigurationsCreateOutput represents the skills configurations create output type.
type SkillsConfigurationsCreateOutput struct {
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

// MapSkillsConfigurationsCreateOutputFromJSON deserializes JSON data into a SkillsConfigurationsCreateOutput.
func MapSkillsConfigurationsCreateOutputFromJSON(data []byte) (*SkillsConfigurationsCreateOutput, error) {
	var v SkillsConfigurationsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsConfigurationsCreateOutputToJSON serializes a SkillsConfigurationsCreateOutput to JSON.
func MapSkillsConfigurationsCreateOutputToJSON(v *SkillsConfigurationsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsConfigurationsCreateBody represents the skills configurations create body type.
type SkillsConfigurationsCreateBody struct {
	AllowScripts                *bool     `json:"allow_scripts,omitempty"`
	AllowedFileExtensions       *[]string `json:"allowed_file_extensions,omitempty"`
	AllowNonStandardDirectories *bool     `json:"allow_non_standard_directories,omitempty"`
}

// MapSkillsConfigurationsCreateBodyFromJSON deserializes JSON data into a SkillsConfigurationsCreateBody.
func MapSkillsConfigurationsCreateBodyFromJSON(data []byte) (*SkillsConfigurationsCreateBody, error) {
	var v SkillsConfigurationsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsConfigurationsCreateBodyToJSON serializes a SkillsConfigurationsCreateBody to JSON.
func MapSkillsConfigurationsCreateBodyToJSON(v *SkillsConfigurationsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
