package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/nullable"
	"github.com/metorial/metorial-go/v1/resources/skills/mergerequests"
)

// SkillsMergeRequestsEndpoint provides access to a skill merge request is a proposed change from one skill into another.
type SkillsMergeRequestsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsMergeRequestsEndpoint creates a new SkillsMergeRequestsEndpoint.
func NewSkillsMergeRequestsEndpoint(client *endpoint.Client) *SkillsMergeRequestsEndpoint {
	return &SkillsMergeRequestsEndpoint{client: client}
}

// SkillsMergeRequestsEndpointListParams contains optional query parameters for List.
type SkillsMergeRequestsEndpointListParams struct {
	// Limit - Page size, from 1 to 100. Defaults to 20.
	Limit *float64 `json:"limit,omitempty"`
	// After - Return the page after this point. Pass `pagination.after` if the previous response has it, otherwise the `id` of its last item.
	After *string `json:"after,omitempty"`
	// Before - Return the page before this point. Pass `pagination.before` if the previous response has it, otherwise the `id` of its first item.
	Before *string `json:"before,omitempty"`
	// Cursor - A `cur_` cursor from `pagination.after` or `pagination.before`. Alternative to `after` and `before`.
	Cursor *string `json:"cursor,omitempty"`
	// Order - Sort direction.
	Order *string `json:"order,omitempty"`
	// Id - Only return merge requests with these IDs.
	Id *any `json:"id,omitempty"`
	// SourceSkillId - Only return merge requests from one of these source skills.
	SourceSkillId *any `json:"source_skill_id,omitempty"`
	// TargetSkillId - Only return merge requests into one of these target skills.
	TargetSkillId *any `json:"target_skill_id,omitempty"`
	// Status - Only return merge requests with one of these statuses.
	Status *any `json:"status,omitempty"`
	// CreatedByActorId - Only return merge requests created by one of these actors.
	CreatedByActorId *any `json:"created_by_actor_id,omitempty"`
	// CreatedAt - Only return results whose skill merge request creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
}

// SkillsMergeRequestsEndpointCreateBody contains the request body for Create.
type SkillsMergeRequestsEndpointCreateBody struct {
	// SourceSkillId - Skill to merge from, by ID.
	SourceSkillId string `json:"source_skill_id"`
	// TargetSkillId - Skill to merge into, by ID. Defaults to the upstream skill of the source skill.
	TargetSkillId *string `json:"target_skill_id,omitempty"`
	// Title - Display name.
	Title string `json:"title"`
	// Description - Free-text description. `null` clears it.
	Description *nullable.Nullable[string] `json:"description,omitempty"`
}

// List returns a paginated list of skill merge requests.
func (e *SkillsMergeRequestsEndpoint) List(instanceId string, params *SkillsMergeRequestsEndpointListParams) (*mergerequests.SkillsMergeRequestsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "skill-merge-requests"},
		Query: query,
	}
	var result mergerequests.SkillsMergeRequestsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a merge request from one skill into another.
func (e *SkillsMergeRequestsEndpoint) Create(instanceId string, body *SkillsMergeRequestsEndpointCreateBody) (*mergerequests.SkillsMergeRequestsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-merge-requests"},
		Body: body,
	}
	var result mergerequests.SkillsMergeRequestsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a skill merge request.
func (e *SkillsMergeRequestsEndpoint) Get(instanceId string, skillMergeRequestId string) (*mergerequests.SkillsMergeRequestsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-merge-requests", skillMergeRequestId},
	}
	var result mergerequests.SkillsMergeRequestsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Perform queues application of a resolved skill merge request.
func (e *SkillsMergeRequestsEndpoint) Perform(instanceId string, skillMergeRequestId string) (*mergerequests.SkillsMergeRequestsPerformOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-merge-requests", skillMergeRequestId, "perform"},
	}
	var result mergerequests.SkillsMergeRequestsPerformOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Close closes an open skill merge request without applying it.
func (e *SkillsMergeRequestsEndpoint) Close(instanceId string, skillMergeRequestId string) (*mergerequests.SkillsMergeRequestsCloseOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-merge-requests", skillMergeRequestId, "close"},
	}
	var result mergerequests.SkillsMergeRequestsCloseOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Rollback restores the target skill to its state before a completed merge.
func (e *SkillsMergeRequestsEndpoint) Rollback(instanceId string, skillMergeRequestId string) (*mergerequests.SkillsMergeRequestsRollbackOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-merge-requests", skillMergeRequestId, "rollback"},
	}
	var result mergerequests.SkillsMergeRequestsRollbackOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
