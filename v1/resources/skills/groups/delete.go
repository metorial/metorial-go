package groups

import (
	"encoding/json"
	"time"
)

// SkillsGroupsDeleteOutputImage - Image shown for the skill.
type SkillsGroupsDeleteOutputImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Where the skill image comes from.
	Type string `json:"type"`
	// Icon - Library icon ID. `null` unless `type` is `icon`.
	Icon *string `json:"icon,omitempty"`
	// Color - Tile color. `null` unless `type` is `icon`.
	Color *string `json:"color,omitempty"`
}

// SkillsGroupsDeleteOutputSkillsImage - Image shown for the skill.
type SkillsGroupsDeleteOutputSkillsImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Where the skill image comes from.
	Type string `json:"type"`
	// Icon - Library icon ID. `null` unless `type` is `icon`.
	Icon *string `json:"icon,omitempty"`
	// Color - Tile color. `null` unless `type` is `icon`.
	Color *string `json:"color,omitempty"`
}

// SkillsGroupsDeleteOutputSkills represents the skills groups delete output skills type.
type SkillsGroupsDeleteOutputSkills struct {
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
	Image SkillsGroupsDeleteOutputSkillsImage `json:"image"`
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

// SkillsGroupsDeleteOutput represents the skills groups delete output type.
type SkillsGroupsDeleteOutput struct {
	Object      string          `json:"object"`
	Id          string          `json:"id"`
	Status      string          `json:"status"`
	Name        string          `json:"name"`
	Description *string         `json:"description,omitempty"`
	Metadata    *map[string]any `json:"metadata,omitempty"`
	ImageUrl    string          `json:"image_url"`
	// Image - Image shown for the skill.
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
