package groups

import (
	"encoding/json"
	"github.com/metorial/metorial-go/v1/nullable"
	"time"
)

// SkillsGroupsUpdateOutputImage - Image shown for the skill.
type SkillsGroupsUpdateOutputImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Where the skill image comes from.
	Type string `json:"type"`
	// Icon - Library icon ID. `null` unless `type` is `icon`.
	Icon *string `json:"icon,omitempty"`
	// Color - Tile color. `null` unless `type` is `icon`.
	Color *string `json:"color,omitempty"`
}

// SkillsGroupsUpdateOutputSkillsImage - Image shown for the skill.
type SkillsGroupsUpdateOutputSkillsImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Where the skill image comes from.
	Type string `json:"type"`
	// Icon - Library icon ID. `null` unless `type` is `icon`.
	Icon *string `json:"icon,omitempty"`
	// Color - Tile color. `null` unless `type` is `icon`.
	Color *string `json:"color,omitempty"`
}

// SkillsGroupsUpdateOutputSkills represents the skills groups update output skills type.
type SkillsGroupsUpdateOutputSkills struct {
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
	Image SkillsGroupsUpdateOutputSkillsImage `json:"image"`
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

// SkillsGroupsUpdateOutput represents the skills groups update output type.
type SkillsGroupsUpdateOutput struct {
	Object      string          `json:"object"`
	Id          string          `json:"id"`
	Status      string          `json:"status"`
	Name        string          `json:"name"`
	Description *string         `json:"description,omitempty"`
	Metadata    *map[string]any `json:"metadata,omitempty"`
	ImageUrl    string          `json:"image_url"`
	// Image - Image shown for the skill.
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
	// Icon - ID of the library icon.
	Icon string `json:"icon"`
	// Color - Tile color. One of `gray`, `blue`, `green`, `yellow`, `orange`, `red`, or `purple`.
	Color string `json:"color"`
}

// SkillsGroupsUpdateBody represents the skills groups update body type.
type SkillsGroupsUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description. `null` clears it.
	Description *nullable.Nullable[string] `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the skill group. Replaces existing metadata. `null` clears it.
	Metadata *nullable.Nullable[map[string]any] `json:"metadata,omitempty"`
	// SkillIds - Skills in the group, by ID. Replaces existing membership.
	SkillIds *[]string `json:"skill_ids,omitempty"`
	// AllowConsumerSkillAssignment - Allow consumers to assign their skills to the group.
	AllowConsumerSkillAssignment *bool `json:"allow_consumer_skill_assignment,omitempty"`
	// ImageFileId - Image file for the skill group, by ID. `null` resets to the default image. Cannot be combined with `image_icon`.
	ImageFileId *nullable.Nullable[string] `json:"image_file_id,omitempty"`
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
