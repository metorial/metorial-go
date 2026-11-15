package groups

import (
	"encoding/json"
	"time"
)

// SkillsGroupsDeleteOutputImage represents the skills groups delete output image type.
type SkillsGroupsDeleteOutputImage struct {
	Object string  `json:"object"`
	Type   string  `json:"type"`
	Icon   *string `json:"icon,omitempty"`
	Color  *string `json:"color,omitempty"`
}

// SkillsGroupsDeleteOutputSkillsImage represents the skills groups delete output skills image type.
type SkillsGroupsDeleteOutputSkillsImage struct {
	Object string  `json:"object"`
	Type   string  `json:"type"`
	Icon   *string `json:"icon,omitempty"`
	Color  *string `json:"color,omitempty"`
}

// SkillsGroupsDeleteOutputSkills represents the skills groups delete output skills type.
type SkillsGroupsDeleteOutputSkills struct {
	Object            string                              `json:"object"`
	Id                string                              `json:"id"`
	Status            string                              `json:"status"`
	Slug              string                              `json:"slug"`
	Name              string                              `json:"name"`
	Description       *string                             `json:"description,omitempty"`
	ImageUrl          string                              `json:"image_url"`
	Image             SkillsGroupsDeleteOutputSkillsImage `json:"image"`
	ClientName        string                              `json:"client_name"`
	ClientDescription *string                             `json:"client_description,omitempty"`
	ClientMetadata    *map[string]any                     `json:"client_metadata,omitempty"`
	License           *string                             `json:"license,omitempty"`
	Compatibility     *string                             `json:"compatibility,omitempty"`
	Metadata          *map[string]any                     `json:"metadata,omitempty"`
	CreatedAt         time.Time                           `json:"created_at"`
	UpdatedAt         time.Time                           `json:"updated_at"`
}

// SkillsGroupsDeleteOutput represents the skills groups delete output type.
type SkillsGroupsDeleteOutput struct {
	Object                       string                           `json:"object"`
	Id                           string                           `json:"id"`
	Status                       string                           `json:"status"`
	Name                         string                           `json:"name"`
	Description                  *string                          `json:"description,omitempty"`
	Metadata                     *map[string]any                  `json:"metadata,omitempty"`
	ImageUrl                     string                           `json:"image_url"`
	Image                        SkillsGroupsDeleteOutputImage    `json:"image"`
	AllowConsumerSkillAssignment bool                             `json:"allow_consumer_skill_assignment"`
	Skills                       []SkillsGroupsDeleteOutputSkills `json:"skills"`
	CreatedAt                    time.Time                        `json:"created_at"`
	UpdatedAt                    time.Time                        `json:"updated_at"`
}

// MapSkillsGroupsDeleteOutputFromJSON deserializes JSON data into a SkillsGroupsDeleteOutput.
func MapSkillsGroupsDeleteOutputFromJSON(data []byte) (*SkillsGroupsDeleteOutput, error) {
	var v SkillsGroupsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsGroupsDeleteOutputToJSON serializes a SkillsGroupsDeleteOutput to JSON.
func MapSkillsGroupsDeleteOutputToJSON(v *SkillsGroupsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
