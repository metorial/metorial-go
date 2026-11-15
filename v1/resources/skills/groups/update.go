package groups

import (
	"encoding/json"
	"github.com/metorial/metorial-go/v1/nullable"
	"time"
)

// SkillsGroupsUpdateOutputImage represents the skills groups update output image type.
type SkillsGroupsUpdateOutputImage struct {
	Object string  `json:"object"`
	Type   string  `json:"type"`
	Icon   *string `json:"icon,omitempty"`
	Color  *string `json:"color,omitempty"`
}

// SkillsGroupsUpdateOutputSkillsImage represents the skills groups update output skills image type.
type SkillsGroupsUpdateOutputSkillsImage struct {
	Object string  `json:"object"`
	Type   string  `json:"type"`
	Icon   *string `json:"icon,omitempty"`
	Color  *string `json:"color,omitempty"`
}

// SkillsGroupsUpdateOutputSkills represents the skills groups update output skills type.
type SkillsGroupsUpdateOutputSkills struct {
	Object            string                              `json:"object"`
	Id                string                              `json:"id"`
	Status            string                              `json:"status"`
	Slug              string                              `json:"slug"`
	Name              string                              `json:"name"`
	Description       *string                             `json:"description,omitempty"`
	ImageUrl          string                              `json:"image_url"`
	Image             SkillsGroupsUpdateOutputSkillsImage `json:"image"`
	ClientName        string                              `json:"client_name"`
	ClientDescription *string                             `json:"client_description,omitempty"`
	ClientMetadata    *map[string]any                     `json:"client_metadata,omitempty"`
	License           *string                             `json:"license,omitempty"`
	Compatibility     *string                             `json:"compatibility,omitempty"`
	Metadata          *map[string]any                     `json:"metadata,omitempty"`
	CreatedAt         time.Time                           `json:"created_at"`
	UpdatedAt         time.Time                           `json:"updated_at"`
}

// SkillsGroupsUpdateOutput represents the skills groups update output type.
type SkillsGroupsUpdateOutput struct {
	Object                       string                           `json:"object"`
	Id                           string                           `json:"id"`
	Status                       string                           `json:"status"`
	Name                         string                           `json:"name"`
	Description                  *string                          `json:"description,omitempty"`
	Metadata                     *map[string]any                  `json:"metadata,omitempty"`
	ImageUrl                     string                           `json:"image_url"`
	Image                        SkillsGroupsUpdateOutputImage    `json:"image"`
	AllowConsumerSkillAssignment bool                             `json:"allow_consumer_skill_assignment"`
	Skills                       []SkillsGroupsUpdateOutputSkills `json:"skills"`
	CreatedAt                    time.Time                        `json:"created_at"`
	UpdatedAt                    time.Time                        `json:"updated_at"`
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

// SkillsGroupsUpdateBodyImageIcon - Use a library icon with a tile color as the skill group image. Pass null to reset to the default image. Cannot be combined with image_file_id.
type SkillsGroupsUpdateBodyImageIcon struct {
	// Icon - Library icon id, e.g. rocket-2-line
	Icon string `json:"icon"`
	// Color - Tile color: gray, blue, green, yellow, orange, red or purple
	Color string `json:"color"`
}

// SkillsGroupsUpdateBody represents the skills groups update body type.
type SkillsGroupsUpdateBody struct {
	Name                         *string         `json:"name,omitempty"`
	Description                  *string         `json:"description,omitempty"`
	Metadata                     *map[string]any `json:"metadata,omitempty"`
	SkillIds                     *[]string       `json:"skill_ids,omitempty"`
	AllowConsumerSkillAssignment *bool           `json:"allow_consumer_skill_assignment,omitempty"`
	ImageFileId                  *string         `json:"image_file_id,omitempty"`
	// ImageIcon - Use a library icon with a tile color as the skill group image. Pass null to reset to the default image. Cannot be combined with image_file_id.
	ImageIcon *nullable.Nullable[SkillsGroupsUpdateBodyImageIcon] `json:"image_icon,omitempty"`
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
