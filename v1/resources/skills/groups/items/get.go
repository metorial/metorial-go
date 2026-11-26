package items

import (
	"encoding/json"
	"time"
)

// SkillsGroupsItemsGetOutputSkillImage - Image shown for the skill.
type SkillsGroupsItemsGetOutputSkillImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Where the skill image comes from.
	Type string `json:"type"`
	// Icon - Library icon ID. `null` unless `type` is `icon`.
	Icon *string `json:"icon,omitempty"`
	// Color - Tile color. `null` unless `type` is `icon`.
	Color *string `json:"color,omitempty"`
}

// SkillsGroupsItemsGetOutputSkill represents the skills groups items get output skill type.
type SkillsGroupsItemsGetOutputSkill struct {
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
	Image SkillsGroupsItemsGetOutputSkillImage `json:"image"`
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
