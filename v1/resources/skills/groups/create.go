package groups

import (
	"encoding/json"
	"time"
)

// SkillsGroupsCreateOutputImage - Image shown for the skill.
type SkillsGroupsCreateOutputImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Where the skill image comes from.
	Type string `json:"type"`
	// Icon - Library icon ID. `null` unless `type` is `icon`.
	Icon *string `json:"icon,omitempty"`
	// Color - Tile color. `null` unless `type` is `icon`.
	Color *string `json:"color,omitempty"`
}

// SkillsGroupsCreateOutputSkillsImage - Image shown for the skill.
type SkillsGroupsCreateOutputSkillsImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Where the skill image comes from.
	Type string `json:"type"`
	// Icon - Library icon ID. `null` unless `type` is `icon`.
	Icon *string `json:"icon,omitempty"`
	// Color - Tile color. `null` unless `type` is `icon`.
	Color *string `json:"color,omitempty"`
}

// SkillsGroupsCreateOutputSkills represents the skills groups create output skills type.
type SkillsGroupsCreateOutputSkills struct {
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
	Image SkillsGroupsCreateOutputSkillsImage `json:"image"`
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

// SkillsGroupsCreateOutput represents the skills groups create output type.
type SkillsGroupsCreateOutput struct {
	Object      string          `json:"object"`
	Id          string          `json:"id"`
	Status      string          `json:"status"`
	Name        string          `json:"name"`
	Description *string         `json:"description,omitempty"`
	Metadata    *map[string]any `json:"metadata,omitempty"`
	ImageUrl    string          `json:"image_url"`
	// Image - Image shown for the skill.
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
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the skill group.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// SkillIds - Skills to add to the group, by ID.
	SkillIds *[]string `json:"skill_ids,omitempty"`
	// AllowConsumerSkillAssignment - Allow consumers to assign their skills to the group.
	AllowConsumerSkillAssignment *bool `json:"allow_consumer_skill_assignment,omitempty"`
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
