package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/plugins/skillgroupmatches"
)

// SkillsPluginsSkillGroupMatchesEndpoint provides access to a skill plugin is a set of skills distributed through a skill marketplace.
type SkillsPluginsSkillGroupMatchesEndpoint struct {
	client *endpoint.Client
}

// NewSkillsPluginsSkillGroupMatchesEndpoint creates a new SkillsPluginsSkillGroupMatchesEndpoint.
func NewSkillsPluginsSkillGroupMatchesEndpoint(client *endpoint.Client) *SkillsPluginsSkillGroupMatchesEndpoint {
	return &SkillsPluginsSkillGroupMatchesEndpoint{client: client}
}

// SkillsPluginsSkillGroupMatchesEndpointListParams contains optional query parameters for List.
type SkillsPluginsSkillGroupMatchesEndpointListParams struct {
	// SkillGroupId - Only compare with these skill groups, up to 5.
	SkillGroupId *any `json:"skill_group_id,omitempty"`
}

// List compares a skill plugin with skill groups and returns the skills attaching each group would add or remove. Without a filter, returns up to 5 groups that share the most skills with the plugin.
func (e *SkillsPluginsSkillGroupMatchesEndpoint) List(skillPluginId string, params *SkillsPluginsSkillGroupMatchesEndpointListParams) (*skillgroupmatches.SkillsPluginsSkillGroupMatchesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"skill-plugins", skillPluginId, "skill-group-matches"},
		Query: query,
	}
	var result skillgroupmatches.SkillsPluginsSkillGroupMatchesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
