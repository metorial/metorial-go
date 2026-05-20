package items

import (
	"encoding/json"
	"time"
)

// SkillsGroupsItemsCreateOutputSkill represents the skills groups items create output skill type.
type SkillsGroupsItemsCreateOutputSkill struct {
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

// SkillsGroupsItemsCreateOutput represents the skills groups items create output type.
type SkillsGroupsItemsCreateOutput struct {
	Object       string                             `json:"object"`
	Id           string                             `json:"id"`
	Status       string                             `json:"status"`
	SkillGroupId string                             `json:"skill_group_id"`
	Skill        SkillsGroupsItemsCreateOutputSkill `json:"skill"`
	CreatedAt    time.Time                          `json:"created_at"`
}

// MapSkillsGroupsItemsCreateOutputFromJSON deserializes JSON data into a SkillsGroupsItemsCreateOutput.
func MapSkillsGroupsItemsCreateOutputFromJSON(data []byte) (*SkillsGroupsItemsCreateOutput, error) {
	var v SkillsGroupsItemsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsGroupsItemsCreateOutputToJSON serializes a SkillsGroupsItemsCreateOutput to JSON.
func MapSkillsGroupsItemsCreateOutputToJSON(v *SkillsGroupsItemsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsGroupsItemsCreateBody represents the skills groups items create body type.
type SkillsGroupsItemsCreateBody struct {
	SkillId string `json:"skill_id"`
}

// MapSkillsGroupsItemsCreateBodyFromJSON deserializes JSON data into a SkillsGroupsItemsCreateBody.
func MapSkillsGroupsItemsCreateBodyFromJSON(data []byte) (*SkillsGroupsItemsCreateBody, error) {
	var v SkillsGroupsItemsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsGroupsItemsCreateBodyToJSON serializes a SkillsGroupsItemsCreateBody to JSON.
func MapSkillsGroupsItemsCreateBodyToJSON(v *SkillsGroupsItemsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
