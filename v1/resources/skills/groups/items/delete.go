package items

import (
	"encoding/json"
	"time"
)

// SkillsGroupsItemsDeleteOutputSkillImage represents the skills groups items delete output skill image type.
type SkillsGroupsItemsDeleteOutputSkillImage struct {
	Object string  `json:"object"`
	Type   string  `json:"type"`
	Icon   *string `json:"icon,omitempty"`
	Color  *string `json:"color,omitempty"`
}

// SkillsGroupsItemsDeleteOutputSkill represents the skills groups items delete output skill type.
type SkillsGroupsItemsDeleteOutputSkill struct {
	Object            string                                  `json:"object"`
	Id                string                                  `json:"id"`
	Status            string                                  `json:"status"`
	Slug              string                                  `json:"slug"`
	Name              string                                  `json:"name"`
	Description       *string                                 `json:"description,omitempty"`
	ImageUrl          string                                  `json:"image_url"`
	Image             SkillsGroupsItemsDeleteOutputSkillImage `json:"image"`
	ClientName        string                                  `json:"client_name"`
	ClientDescription *string                                 `json:"client_description,omitempty"`
	ClientMetadata    *map[string]any                         `json:"client_metadata,omitempty"`
	License           *string                                 `json:"license,omitempty"`
	Compatibility     *string                                 `json:"compatibility,omitempty"`
	Metadata          *map[string]any                         `json:"metadata,omitempty"`
	CreatedAt         time.Time                               `json:"created_at"`
	UpdatedAt         time.Time                               `json:"updated_at"`
}

// SkillsGroupsItemsDeleteOutput represents the skills groups items delete output type.
type SkillsGroupsItemsDeleteOutput struct {
	Object       string                             `json:"object"`
	Id           string                             `json:"id"`
	Status       string                             `json:"status"`
	SkillGroupId string                             `json:"skill_group_id"`
	Skill        SkillsGroupsItemsDeleteOutputSkill `json:"skill"`
	CreatedAt    time.Time                          `json:"created_at"`
}

// MapSkillsGroupsItemsDeleteOutputFromJSON deserializes JSON data into a SkillsGroupsItemsDeleteOutput.
func MapSkillsGroupsItemsDeleteOutputFromJSON(data []byte) (*SkillsGroupsItemsDeleteOutput, error) {
	var v SkillsGroupsItemsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsGroupsItemsDeleteOutputToJSON serializes a SkillsGroupsItemsDeleteOutput to JSON.
func MapSkillsGroupsItemsDeleteOutputToJSON(v *SkillsGroupsItemsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
