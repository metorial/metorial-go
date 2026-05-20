package plugins

import (
	"encoding/json"
	"time"
)

// SkillsPluginsSyncOutputSkills represents the skills plugins sync output skills type.
type SkillsPluginsSyncOutputSkills struct {
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

// SkillsPluginsSyncOutput represents the skills plugins sync output type.
type SkillsPluginsSyncOutput struct {
	Object               string                          `json:"object"`
	Id                   string                          `json:"id"`
	Status               string                          `json:"status"`
	SyncStatus           string                          `json:"sync_status"`
	ImageUrl             string                          `json:"image_url"`
	Name                 string                          `json:"name"`
	Description          *string                         `json:"description,omitempty"`
	LongDescription      *string                         `json:"long_description,omitempty"`
	Category             *string                         `json:"category,omitempty"`
	Slug                 string                          `json:"slug"`
	SkillConfigurationId *string                         `json:"skill_configuration_id,omitempty"`
	Skills               []SkillsPluginsSyncOutputSkills `json:"skills"`
	CreatedAt            time.Time                       `json:"created_at"`
	UpdatedAt            time.Time                       `json:"updated_at"`
}

// MapSkillsPluginsSyncOutputFromJSON deserializes JSON data into a SkillsPluginsSyncOutput.
func MapSkillsPluginsSyncOutputFromJSON(data []byte) (*SkillsPluginsSyncOutput, error) {
	var v SkillsPluginsSyncOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsSyncOutputToJSON serializes a SkillsPluginsSyncOutput to JSON.
func MapSkillsPluginsSyncOutputToJSON(v *SkillsPluginsSyncOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsPluginsSyncBody represents the skills plugins sync body type.
type SkillsPluginsSyncBody struct{}

// MapSkillsPluginsSyncBodyFromJSON deserializes JSON data into a SkillsPluginsSyncBody.
func MapSkillsPluginsSyncBodyFromJSON(data []byte) (*SkillsPluginsSyncBody, error) {
	var v SkillsPluginsSyncBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsSyncBodyToJSON serializes a SkillsPluginsSyncBody to JSON.
func MapSkillsPluginsSyncBodyToJSON(v *SkillsPluginsSyncBody) ([]byte, error) {
	return json.Marshal(v)
}
