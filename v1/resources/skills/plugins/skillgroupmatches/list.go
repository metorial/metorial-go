package skillgroupmatches

import (
	"encoding/json"
)

// SkillsPluginsSkillGroupMatchesListOutputItemsSkillGroup - Skill group compared with the plugin.
type SkillsPluginsSkillGroupMatchesListOutputItemsSkillGroup struct {
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
}

// SkillsPluginsSkillGroupMatchesListOutputItems represents the skills plugins skill group matches list output items type.
type SkillsPluginsSkillGroupMatchesListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// SkillPluginId - Skill plugin the skill group is compared with.
	SkillPluginId string `json:"skill_plugin_id"`
	// SkillGroup - Skill group compared with the plugin.
	SkillGroup SkillsPluginsSkillGroupMatchesListOutputItemsSkillGroup `json:"skill_group"`
	// SharedSkillCount - Number of skills that are in both the plugin and the skill group.
	SharedSkillCount float64 `json:"shared_skill_count"`
	// SkillIdsToAdd - Skills in the group that the plugin lacks. Attaching the group adds them to the plugin.
	SkillIdsToAdd []string `json:"skill_ids_to_add"`
	// SkillIdsToRemove - Skills in the plugin that the group lacks. Attaching the group removes them from the plugin.
	SkillIdsToRemove []string `json:"skill_ids_to_remove"`
}

// SkillsPluginsSkillGroupMatchesListOutputPagination represents the skills plugins skill group matches list output pagination type.
type SkillsPluginsSkillGroupMatchesListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SkillsPluginsSkillGroupMatchesListOutput represents the skills plugins skill group matches list output type.
type SkillsPluginsSkillGroupMatchesListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                             `json:"object"`
	Items      []SkillsPluginsSkillGroupMatchesListOutputItems    `json:"items"`
	Pagination SkillsPluginsSkillGroupMatchesListOutputPagination `json:"pagination"`
}

// MapSkillsPluginsSkillGroupMatchesListOutputFromJSON deserializes JSON data into a SkillsPluginsSkillGroupMatchesListOutput.
func MapSkillsPluginsSkillGroupMatchesListOutputFromJSON(data []byte) (*SkillsPluginsSkillGroupMatchesListOutput, error) {
	var v SkillsPluginsSkillGroupMatchesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsSkillGroupMatchesListOutputToJSON serializes a SkillsPluginsSkillGroupMatchesListOutput to JSON.
func MapSkillsPluginsSkillGroupMatchesListOutputToJSON(v *SkillsPluginsSkillGroupMatchesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsPluginsSkillGroupMatchesListQuery represents the skills plugins skill group matches list query type.
type SkillsPluginsSkillGroupMatchesListQuery struct {
	// SkillGroupId - Only compare with these skill groups, up to 5.
	SkillGroupId *any `json:"skill_group_id,omitempty"`
}

// MapSkillsPluginsSkillGroupMatchesListQueryFromJSON deserializes JSON data into a SkillsPluginsSkillGroupMatchesListQuery.
func MapSkillsPluginsSkillGroupMatchesListQueryFromJSON(data []byte) (*SkillsPluginsSkillGroupMatchesListQuery, error) {
	var v SkillsPluginsSkillGroupMatchesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsSkillGroupMatchesListQueryToJSON serializes a SkillsPluginsSkillGroupMatchesListQuery to JSON.
func MapSkillsPluginsSkillGroupMatchesListQueryToJSON(v *SkillsPluginsSkillGroupMatchesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
