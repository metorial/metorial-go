package groups

import (
	"encoding/json"
	"time"
)

// SkillsGroupsCreateOutputImage represents the skills groups create output image type.
type SkillsGroupsCreateOutputImage struct {
	Object string  `json:"object"`
	Type   string  `json:"type"`
	Icon   *string `json:"icon,omitempty"`
	Color  *string `json:"color,omitempty"`
}

// SkillsGroupsCreateOutputSkillsImage represents the skills groups create output skills image type.
type SkillsGroupsCreateOutputSkillsImage struct {
	Object string  `json:"object"`
	Type   string  `json:"type"`
	Icon   *string `json:"icon,omitempty"`
	Color  *string `json:"color,omitempty"`
}

// SkillsGroupsCreateOutputSkills represents the skills groups create output skills type.
type SkillsGroupsCreateOutputSkills struct {
	Object            string                              `json:"object"`
	Id                string                              `json:"id"`
	Status            string                              `json:"status"`
	Slug              string                              `json:"slug"`
	Name              string                              `json:"name"`
	Description       *string                             `json:"description,omitempty"`
	ImageUrl          string                              `json:"image_url"`
	Image             SkillsGroupsCreateOutputSkillsImage `json:"image"`
	ClientName        string                              `json:"client_name"`
	ClientDescription *string                             `json:"client_description,omitempty"`
	ClientMetadata    *map[string]any                     `json:"client_metadata,omitempty"`
	License           *string                             `json:"license,omitempty"`
	Compatibility     *string                             `json:"compatibility,omitempty"`
	Metadata          *map[string]any                     `json:"metadata,omitempty"`
	CreatedAt         time.Time                           `json:"created_at"`
	UpdatedAt         time.Time                           `json:"updated_at"`
}

// SkillsGroupsCreateOutput represents the skills groups create output type.
type SkillsGroupsCreateOutput struct {
	Object                       string                           `json:"object"`
	Id                           string                           `json:"id"`
	Status                       string                           `json:"status"`
	Name                         string                           `json:"name"`
	Description                  *string                          `json:"description,omitempty"`
	Metadata                     *map[string]any                  `json:"metadata,omitempty"`
	ImageUrl                     string                           `json:"image_url"`
	Image                        SkillsGroupsCreateOutputImage    `json:"image"`
	AllowConsumerSkillAssignment bool                             `json:"allow_consumer_skill_assignment"`
	Skills                       []SkillsGroupsCreateOutputSkills `json:"skills"`
	CreatedAt                    time.Time                        `json:"created_at"`
	UpdatedAt                    time.Time                        `json:"updated_at"`
}

// MapSkillsGroupsCreateOutputFromJSON deserializes JSON data into a SkillsGroupsCreateOutput.
func MapSkillsGroupsCreateOutputFromJSON(data []byte) (*SkillsGroupsCreateOutput, error) {
	var v SkillsGroupsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsGroupsCreateOutputToJSON serializes a SkillsGroupsCreateOutput to JSON.
func MapSkillsGroupsCreateOutputToJSON(v *SkillsGroupsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsGroupsCreateBody represents the skills groups create body type.
type SkillsGroupsCreateBody struct {
	Name                         string          `json:"name"`
	Description                  *string         `json:"description,omitempty"`
	Metadata                     *map[string]any `json:"metadata,omitempty"`
	SkillIds                     *[]string       `json:"skill_ids,omitempty"`
	AllowConsumerSkillAssignment *bool           `json:"allow_consumer_skill_assignment,omitempty"`
}

// MapSkillsGroupsCreateBodyFromJSON deserializes JSON data into a SkillsGroupsCreateBody.
func MapSkillsGroupsCreateBodyFromJSON(data []byte) (*SkillsGroupsCreateBody, error) {
	var v SkillsGroupsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsGroupsCreateBodyToJSON serializes a SkillsGroupsCreateBody to JSON.
func MapSkillsGroupsCreateBodyToJSON(v *SkillsGroupsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
