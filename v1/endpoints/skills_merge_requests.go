package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/mergerequests"
)

// SkillsMergeRequestsEndpoint provides access to review, resolve, and apply changes between skills.
type SkillsMergeRequestsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsMergeRequestsEndpoint creates a new SkillsMergeRequestsEndpoint.
func NewSkillsMergeRequestsEndpoint(client *endpoint.Client) *SkillsMergeRequestsEndpoint {
	return &SkillsMergeRequestsEndpoint{client: client}
}

// SkillsMergeRequestsEndpointListParams contains optional query parameters for List.
type SkillsMergeRequestsEndpointListParams struct {
	Limit            *float64 `json:"limit,omitempty"`
	After            *string  `json:"after,omitempty"`
	Before           *string  `json:"before,omitempty"`
	Cursor           *string  `json:"cursor,omitempty"`
	Order            *string  `json:"order,omitempty"`
	Id               *any     `json:"id,omitempty"`
	SourceSkillId    *any     `json:"source_skill_id,omitempty"`
	TargetSkillId    *any     `json:"target_skill_id,omitempty"`
	Status           *any     `json:"status,omitempty"`
	CreatedByActorId *any     `json:"created_by_actor_id,omitempty"`
	// CreatedAt - Filter skill merge request creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
}

// SkillsMergeRequestsEndpointCreateBody contains the request body for Create.
type SkillsMergeRequestsEndpointCreateBody struct {
	SourceSkillId string  `json:"source_skill_id"`
	TargetSkillId *string `json:"target_skill_id,omitempty"`
	Title         string  `json:"title"`
	Description   *string `json:"description,omitempty"`
}

// List returns a paginated list of skill merge requests.
func (e *SkillsMergeRequestsEndpoint) List(params *SkillsMergeRequestsEndpointListParams) (*mergerequests.SkillsMergeRequestsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"skill-merge-requests"},
		Query: query,
	}
	var result mergerequests.SkillsMergeRequestsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a merge request from one skill into another.
func (e *SkillsMergeRequestsEndpoint) Create(body *SkillsMergeRequestsEndpointCreateBody) (*mergerequests.SkillsMergeRequestsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-merge-requests"},
		Body: body,
	}
	var result mergerequests.SkillsMergeRequestsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a skill merge request.
func (e *SkillsMergeRequestsEndpoint) Get(skillMergeRequestId string) (*mergerequests.SkillsMergeRequestsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-merge-requests", skillMergeRequestId},
	}
	var result mergerequests.SkillsMergeRequestsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Perform queues application of a resolved skill merge request.
func (e *SkillsMergeRequestsEndpoint) Perform(skillMergeRequestId string) (*mergerequests.SkillsMergeRequestsPerformOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-merge-requests", skillMergeRequestId, "perform"},
	}
	var result mergerequests.SkillsMergeRequestsPerformOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Close closes an open skill merge request without applying it.
func (e *SkillsMergeRequestsEndpoint) Close(skillMergeRequestId string) (*mergerequests.SkillsMergeRequestsCloseOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-merge-requests", skillMergeRequestId, "close"},
	}
	var result mergerequests.SkillsMergeRequestsCloseOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Rollback restores the target skill to its state before a completed merge.
func (e *SkillsMergeRequestsEndpoint) Rollback(skillMergeRequestId string) (*mergerequests.SkillsMergeRequestsRollbackOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-merge-requests", skillMergeRequestId, "rollback"},
	}
	var result mergerequests.SkillsMergeRequestsRollbackOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
