package plugins

import (
	"encoding/json"
	"time"
)

// SkillsPluginsListOutputItemsSourceSkillGroup - The skill group the plugin's skills are synced from.
type SkillsPluginsListOutputItemsSourceSkillGroup struct {
	// Id - The ID of the skill group.
	Id string `json:"id"`
	// Name - The name of the skill group.
	Name string `json:"name"`
}

// SkillsPluginsListOutputItemsSourceUnsyncedSkills represents the skills plugins list output items source unsynced skills type.
type SkillsPluginsListOutputItemsSourceUnsyncedSkills struct {
	// SkillId - The ID of the skill that was not synced.
	SkillId string `json:"skill_id"`
	// Reason - Why the skill was not synced.
	Reason string `json:"reason"`
	// Message - A human-readable explanation.
	Message string `json:"message"`
}

// SkillsPluginsListOutputItemsSource represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SkillsPluginsListOutputItemsSource struct {
	// Type - The plugin's skills are synced from a skill group and cannot be edited on the plugin directly.
	Type *string `json:"type,omitempty"`
	// SkillGroup - The skill group the plugin's skills are synced from.
	SkillGroup *SkillsPluginsListOutputItemsSourceSkillGroup `json:"skill_group,omitempty"`
	// UnsyncedSkills - Group skills that are not in the plugin, for example because a plugin or marketplace limit was reached.
	UnsyncedSkills *[]SkillsPluginsListOutputItemsSourceUnsyncedSkills `json:"unsynced_skills,omitempty"`
	// LastSyncedAt - When the plugin was last reconciled with the skill group.
	LastSyncedAt *time.Time `json:"last_synced_at,omitempty"`
}

// SkillsPluginsListOutputItemsSkillsSkillImage represents the skills plugins list output items skills skill image type.
type SkillsPluginsListOutputItemsSkillsSkillImage struct {
	Object string  `json:"object"`
	Type   string  `json:"type"`
	Icon   *string `json:"icon,omitempty"`
	Color  *string `json:"color,omitempty"`
}

// SkillsPluginsListOutputItemsSkillsSkill represents the skills plugins list output items skills skill type.
type SkillsPluginsListOutputItemsSkillsSkill struct {
	Object            string                                       `json:"object"`
	Id                string                                       `json:"id"`
	Status            string                                       `json:"status"`
	Slug              string                                       `json:"slug"`
	Name              string                                       `json:"name"`
	Description       *string                                      `json:"description,omitempty"`
	ImageUrl          string                                       `json:"image_url"`
	Image             SkillsPluginsListOutputItemsSkillsSkillImage `json:"image"`
	ClientName        string                                       `json:"client_name"`
	ClientDescription *string                                      `json:"client_description,omitempty"`
	ClientMetadata    *map[string]any                              `json:"client_metadata,omitempty"`
	License           *string                                      `json:"license,omitempty"`
	Compatibility     *string                                      `json:"compatibility,omitempty"`
	Metadata          *map[string]any                              `json:"metadata,omitempty"`
	CreatedAt         time.Time                                    `json:"created_at"`
	UpdatedAt         time.Time                                    `json:"updated_at"`
}

// SkillsPluginsListOutputItemsSkills represents the skills plugins list output items skills type.
type SkillsPluginsListOutputItemsSkills struct {
	Object               string                                  `json:"object"`
	Id                   string                                  `json:"id"`
	Identifier           string                                  `json:"identifier"`
	Status               string                                  `json:"status"`
	ClientName           *string                                 `json:"client_name,omitempty"`
	ClientDescription    *string                                 `json:"client_description,omitempty"`
	ClientMetadata       *map[string]any                         `json:"client_metadata,omitempty"`
	License              *string                                 `json:"license,omitempty"`
	Compatibility        *string                                 `json:"compatibility,omitempty"`
	SkillConfigurationId *string                                 `json:"skill_configuration_id,omitempty"`
	SkillId              string                                  `json:"skill_id"`
	Skill                SkillsPluginsListOutputItemsSkillsSkill `json:"skill"`
	CreatedAt            time.Time                               `json:"created_at"`
	UpdatedAt            time.Time                               `json:"updated_at"`
}

// SkillsPluginsListOutputItems represents the skills plugins list output items type.
type SkillsPluginsListOutputItems struct {
	Object               string  `json:"object"`
	Id                   string  `json:"id"`
	Status               string  `json:"status"`
	SyncStatus           string  `json:"sync_status"`
	ImageUrl             string  `json:"image_url"`
	Name                 string  `json:"name"`
	Description          *string `json:"description,omitempty"`
	LongDescription      *string `json:"long_description,omitempty"`
	Category             *string `json:"category,omitempty"`
	Slug                 string  `json:"slug"`
	SkillConfigurationId *string `json:"skill_configuration_id,omitempty"`
	// Source - Where the plugin's skills come from.
	Source    SkillsPluginsListOutputItemsSource   `json:"source"`
	Skills    []SkillsPluginsListOutputItemsSkills `json:"skills"`
	CreatedAt time.Time                            `json:"created_at"`
	UpdatedAt time.Time                            `json:"updated_at"`
}

// SkillsPluginsListOutputPagination represents the skills plugins list output pagination type.
type SkillsPluginsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SkillsPluginsListOutput represents the skills plugins list output type.
type SkillsPluginsListOutput struct {
	Items      []SkillsPluginsListOutputItems    `json:"items"`
	Pagination SkillsPluginsListOutputPagination `json:"pagination"`
}

// MapSkillsPluginsListOutputFromJSON deserializes JSON data into a SkillsPluginsListOutput.
func MapSkillsPluginsListOutputFromJSON(data []byte) (*SkillsPluginsListOutput, error) {
	var v SkillsPluginsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsListOutputToJSON serializes a SkillsPluginsListOutput to JSON.
func MapSkillsPluginsListOutputToJSON(v *SkillsPluginsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsPluginsListQueryCreatedAt - Filter skill plugin creation time by date range
type SkillsPluginsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for skill plugin creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for skill plugin creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsPluginsListQueryUpdatedAt - Filter skill plugin last update time by date range
type SkillsPluginsListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for skill plugin last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for skill plugin last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsPluginsListQuery represents the skills plugins list query type.
type SkillsPluginsListQuery struct {
	Limit                *float64 `json:"limit,omitempty"`
	After                *string  `json:"after,omitempty"`
	Before               *string  `json:"before,omitempty"`
	Cursor               *string  `json:"cursor,omitempty"`
	Order                *string  `json:"order,omitempty"`
	Id                   *any     `json:"id,omitempty"`
	SkillMarketplaceId   *any     `json:"skill_marketplace_id,omitempty"`
	Status               *any     `json:"status,omitempty"`
	Category             *string  `json:"category,omitempty"`
	Search               *string  `json:"search,omitempty"`
	SkillConfigurationId *any     `json:"skill_configuration_id,omitempty"`
	// SkillGroupId - Only return plugins whose skills are synced from these skill groups.
	SkillGroupId *any `json:"skill_group_id,omitempty"`
	// CreatedAt - Filter skill plugin creation time by date range
	CreatedAt *SkillsPluginsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter skill plugin last update time by date range
	UpdatedAt *SkillsPluginsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapSkillsPluginsListQueryFromJSON deserializes JSON data into a SkillsPluginsListQuery.
func MapSkillsPluginsListQueryFromJSON(data []byte) (*SkillsPluginsListQuery, error) {
	var v SkillsPluginsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsListQueryToJSON serializes a SkillsPluginsListQuery to JSON.
func MapSkillsPluginsListQueryToJSON(v *SkillsPluginsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
