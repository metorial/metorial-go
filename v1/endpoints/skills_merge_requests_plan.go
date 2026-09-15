package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/mergerequests/plan"
)

// SkillsMergeRequestsPlanEndpoint provides access to review, resolve, and apply changes between skills.
type SkillsMergeRequestsPlanEndpoint struct {
	client *endpoint.Client
}

// NewSkillsMergeRequestsPlanEndpoint creates a new SkillsMergeRequestsPlanEndpoint.
func NewSkillsMergeRequestsPlanEndpoint(client *endpoint.Client) *SkillsMergeRequestsPlanEndpoint {
	return &SkillsMergeRequestsPlanEndpoint{client: client}
}

// Get returns the proposed changes and conflicts for a skill merge request.
func (e *SkillsMergeRequestsPlanEndpoint) Get(skillMergeRequestId string) (*plan.SkillsMergeRequestsPlanGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-merge-requests", skillMergeRequestId, "plan"},
	}
	var result plan.SkillsMergeRequestsPlanGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
