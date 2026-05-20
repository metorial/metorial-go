package groups

import (
	"encoding/json"
	"time"
)

// SkillsGroupsGetOutputSkills represents the skills groups get output skills type.
type SkillsGroupsGetOutputSkills struct {
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

// SkillsGroupsGetOutput represents the skills groups get output type.
type SkillsGroupsGetOutput struct {
	Object      string                        `json:"object"`
	Id          string                        `json:"id"`
	Status      string                        `json:"status"`
	Name        string                        `json:"name"`
	Description *string                       `json:"description,omitempty"`
	Metadata    *map[string]any               `json:"metadata,omitempty"`
	Skills      []SkillsGroupsGetOutputSkills `json:"skills"`
	CreatedAt   time.Time                     `json:"created_at"`
	UpdatedAt   time.Time                     `json:"updated_at"`
}

// MapSkillsGroupsGetOutputFromJSON deserializes JSON data into a SkillsGroupsGetOutput.
func MapSkillsGroupsGetOutputFromJSON(data []byte) (*SkillsGroupsGetOutput, error) {
	var v SkillsGroupsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsGroupsGetOutputToJSON serializes a SkillsGroupsGetOutput to JSON.
func MapSkillsGroupsGetOutputToJSON(v *SkillsGroupsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
