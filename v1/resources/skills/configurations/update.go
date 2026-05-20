package configurations

import (
	"encoding/json"
	"time"
)

// SkillsConfigurationsUpdateOutput represents the skills configurations update output type.
type SkillsConfigurationsUpdateOutput struct {
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

// MapSkillsConfigurationsUpdateOutputFromJSON deserializes JSON data into a SkillsConfigurationsUpdateOutput.
func MapSkillsConfigurationsUpdateOutputFromJSON(data []byte) (*SkillsConfigurationsUpdateOutput, error) {
	var v SkillsConfigurationsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsConfigurationsUpdateOutputToJSON serializes a SkillsConfigurationsUpdateOutput to JSON.
func MapSkillsConfigurationsUpdateOutputToJSON(v *SkillsConfigurationsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsConfigurationsUpdateBody represents the skills configurations update body type.
type SkillsConfigurationsUpdateBody struct {
	AllowScripts                *bool     `json:"allow_scripts,omitempty"`
	AllowedFileExtensions       *[]string `json:"allowed_file_extensions,omitempty"`
	AllowNonStandardDirectories *bool     `json:"allow_non_standard_directories,omitempty"`
}

// MapSkillsConfigurationsUpdateBodyFromJSON deserializes JSON data into a SkillsConfigurationsUpdateBody.
func MapSkillsConfigurationsUpdateBodyFromJSON(data []byte) (*SkillsConfigurationsUpdateBody, error) {
	var v SkillsConfigurationsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsConfigurationsUpdateBodyToJSON serializes a SkillsConfigurationsUpdateBody to JSON.
func MapSkillsConfigurationsUpdateBodyToJSON(v *SkillsConfigurationsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
