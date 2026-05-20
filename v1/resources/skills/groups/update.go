package groups

import (
	"encoding/json"
	"time"
)

// SkillsGroupsUpdateOutputSkills represents the skills groups update output skills type.
type SkillsGroupsUpdateOutputSkills struct {
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

// SkillsGroupsUpdateOutput represents the skills groups update output type.
type SkillsGroupsUpdateOutput struct {
	Object      string                           `json:"object"`
	Id          string                           `json:"id"`
	Status      string                           `json:"status"`
	Name        string                           `json:"name"`
	Description *string                          `json:"description,omitempty"`
	Metadata    *map[string]any                  `json:"metadata,omitempty"`
	Skills      []SkillsGroupsUpdateOutputSkills `json:"skills"`
	CreatedAt   time.Time                        `json:"created_at"`
	UpdatedAt   time.Time                        `json:"updated_at"`
}

// MapSkillsGroupsUpdateOutputFromJSON deserializes JSON data into a SkillsGroupsUpdateOutput.
func MapSkillsGroupsUpdateOutputFromJSON(data []byte) (*SkillsGroupsUpdateOutput, error) {
	var v SkillsGroupsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsGroupsUpdateOutputToJSON serializes a SkillsGroupsUpdateOutput to JSON.
func MapSkillsGroupsUpdateOutputToJSON(v *SkillsGroupsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsGroupsUpdateBody represents the skills groups update body type.
type SkillsGroupsUpdateBody struct {
	Name        *string         `json:"name,omitempty"`
	Description *string         `json:"description,omitempty"`
	Metadata    *map[string]any `json:"metadata,omitempty"`
	SkillIds    *[]string       `json:"skill_ids,omitempty"`
}

// MapSkillsGroupsUpdateBodyFromJSON deserializes JSON data into a SkillsGroupsUpdateBody.
func MapSkillsGroupsUpdateBodyFromJSON(data []byte) (*SkillsGroupsUpdateBody, error) {
	var v SkillsGroupsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsGroupsUpdateBodyToJSON serializes a SkillsGroupsUpdateBody to JSON.
func MapSkillsGroupsUpdateBodyToJSON(v *SkillsGroupsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
