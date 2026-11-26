package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/nullable"
	"github.com/metorial/metorial-go/v1/resources/skills/mergerequests/comments"
)

// SkillsMergeRequestsCommentsEndpoint provides access to a skill merge request comment is a comment on a merge request or one of its proposed changes.
type SkillsMergeRequestsCommentsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsMergeRequestsCommentsEndpoint creates a new SkillsMergeRequestsCommentsEndpoint.
func NewSkillsMergeRequestsCommentsEndpoint(client *endpoint.Client) *SkillsMergeRequestsCommentsEndpoint {
	return &SkillsMergeRequestsCommentsEndpoint{client: client}
}

// SkillsMergeRequestsCommentsEndpointListParams contains optional query parameters for List.
type SkillsMergeRequestsCommentsEndpointListParams struct {
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
	// ItemId - Only return comments on this proposed change.
	ItemId *string `json:"item_id,omitempty"`
}

// SkillsMergeRequestsCommentsEndpointCreateBody contains the request body for Create.
type SkillsMergeRequestsCommentsEndpointCreateBody struct {
	// ItemId - Proposed change to comment on. Omit to comment on the merge request.
	ItemId *string `json:"item_id,omitempty"`
	// InReplyToCommentId - Comment this replies to, by ID.
	InReplyToCommentId *string `json:"in_reply_to_comment_id,omitempty"`
	// Body - Comment text.
	Body string `json:"body"`
	// Path - File path the comment refers to. `null` for general comments.
	Path *nullable.Nullable[string] `json:"path,omitempty"`
}

// SkillsMergeRequestsCommentsEndpointUpdateBody contains the request body for Update.
type SkillsMergeRequestsCommentsEndpointUpdateBody struct {
	// Body - Comment text.
	Body string `json:"body"`
}

// List returns a paginated list of skill merge request comments.
func (e *SkillsMergeRequestsCommentsEndpoint) List(instanceId string, skillMergeRequestId string, params *SkillsMergeRequestsCommentsEndpointListParams) (*comments.SkillsMergeRequestsCommentsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "skill-merge-requests", skillMergeRequestId, "comments"},
		Query: query,
	}
	var result comments.SkillsMergeRequestsCommentsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a comment on a skill merge request or one of its items.
func (e *SkillsMergeRequestsCommentsEndpoint) Create(instanceId string, skillMergeRequestId string, body *SkillsMergeRequestsCommentsEndpointCreateBody) (*comments.SkillsMergeRequestsCommentsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-merge-requests", skillMergeRequestId, "comments"},
		Body: body,
	}
	var result comments.SkillsMergeRequestsCommentsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a comment on a skill merge request.
func (e *SkillsMergeRequestsCommentsEndpoint) Get(instanceId string, skillMergeRequestId string, commentId string) (*comments.SkillsMergeRequestsCommentsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-merge-requests", skillMergeRequestId, "comments", commentId},
	}
	var result comments.SkillsMergeRequestsCommentsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a comment authored by the current actor.
func (e *SkillsMergeRequestsCommentsEndpoint) Update(instanceId string, skillMergeRequestId string, commentId string, body *SkillsMergeRequestsCommentsEndpointUpdateBody) (*comments.SkillsMergeRequestsCommentsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-merge-requests", skillMergeRequestId, "comments", commentId},
		Body: body,
	}
	var result comments.SkillsMergeRequestsCommentsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a comment authored by the current actor.
func (e *SkillsMergeRequestsCommentsEndpoint) Delete(instanceId string, skillMergeRequestId string, commentId string) (*comments.SkillsMergeRequestsCommentsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-merge-requests", skillMergeRequestId, "comments", commentId},
	}
	var result comments.SkillsMergeRequestsCommentsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
