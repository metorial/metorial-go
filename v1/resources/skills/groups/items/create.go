package items

import (
	"encoding/json"
	"time"
)

// SkillsGroupsItemsCreateOutputSkillImage - Image shown for the skill.
type SkillsGroupsItemsCreateOutputSkillImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Where the skill image comes from.
	Type string `json:"type"`
	// Icon - Library icon ID. `null` unless `type` is `icon`.
	Icon *string `json:"icon,omitempty"`
	// Color - Tile color. `null` unless `type` is `icon`.
	Color *string `json:"color,omitempty"`
}

// SkillsGroupsItemsCreateOutputSkill represents the skills groups items create output skill type.
type SkillsGroupsItemsCreateOutputSkill struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Whether the skill is available, archived, or deleted.
	Status string `json:"status"`
	// Slug - URL-safe name, unique within the instance.
	Slug string `json:"slug"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// ImageUrl - Public URL of the skill image.
	ImageUrl string `json:"image_url"`
	// Image - Image shown for the skill.
	Image SkillsGroupsItemsCreateOutputSkillImage `json:"image"`
	// ClientName - Client-facing skill name.
	ClientName string `json:"client_name"`
	// ClientDescription - Client-facing description.
	ClientDescription *string `json:"client_description,omitempty"`
	// ClientMetadata - Custom data for the client-facing skill.
	ClientMetadata *map[string]any `json:"client_metadata,omitempty"`
	// License - License terms for the skill.
	License *string `json:"license,omitempty"`
	// Compatibility - Compatibility notes for the skill.
	Compatibility *string `json:"compatibility,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
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
	// SkillId - Skill to add to the group, by ID.
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
