package items

import (
	"encoding/json"
	"time"
)

// SkillsGroupsItemsGetOutputSkill represents the skills groups items get output skill type.
type SkillsGroupsItemsGetOutputSkill struct {
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

// SkillsGroupsItemsGetOutput represents the skills groups items get output type.
type SkillsGroupsItemsGetOutput struct {
	Object       string                          `json:"object"`
	Id           string                          `json:"id"`
	Status       string                          `json:"status"`
	SkillGroupId string                          `json:"skill_group_id"`
	Skill        SkillsGroupsItemsGetOutputSkill `json:"skill"`
	CreatedAt    time.Time                       `json:"created_at"`
}

// MapSkillsGroupsItemsGetOutputFromJSON deserializes JSON data into a SkillsGroupsItemsGetOutput.
func MapSkillsGroupsItemsGetOutputFromJSON(data []byte) (*SkillsGroupsItemsGetOutput, error) {
	var v SkillsGroupsItemsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsGroupsItemsGetOutputToJSON serializes a SkillsGroupsItemsGetOutput to JSON.
func MapSkillsGroupsItemsGetOutputToJSON(v *SkillsGroupsItemsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
