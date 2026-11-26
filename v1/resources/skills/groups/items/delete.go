package items

import (
	"encoding/json"
	"time"
)

// SkillsGroupsItemsDeleteOutputSkillImage - Image shown for the skill.
type SkillsGroupsItemsDeleteOutputSkillImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Where the skill image comes from.
	Type string `json:"type"`
	// Icon - Library icon ID. `null` unless `type` is `icon`.
	Icon *string `json:"icon,omitempty"`
	// Color - Tile color. `null` unless `type` is `icon`.
	Color *string `json:"color,omitempty"`
}

// SkillsGroupsItemsDeleteOutputSkill represents the skills groups items delete output skill type.
type SkillsGroupsItemsDeleteOutputSkill struct {
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
	Image SkillsGroupsItemsDeleteOutputSkillImage `json:"image"`
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
