package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/mergerequests/comments"
)

// SkillsMergeRequestsCommentsEndpoint provides access to discuss skill merge requests and individual proposed changes.
type SkillsMergeRequestsCommentsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsMergeRequestsCommentsEndpoint creates a new SkillsMergeRequestsCommentsEndpoint.
func NewSkillsMergeRequestsCommentsEndpoint(client *endpoint.Client) *SkillsMergeRequestsCommentsEndpoint {
	return &SkillsMergeRequestsCommentsEndpoint{client: client}
}

// SkillsMergeRequestsCommentsEndpointListParams contains optional query parameters for List.
type SkillsMergeRequestsCommentsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	ItemId *string  `json:"item_id,omitempty"`
}

// SkillsMergeRequestsCommentsEndpointCreateBody contains the request body for Create.
type SkillsMergeRequestsCommentsEndpointCreateBody struct {
	ItemId             *string `json:"item_id,omitempty"`
	InReplyToCommentId *string `json:"in_reply_to_comment_id,omitempty"`
	Body               string  `json:"body"`
	Path               *string `json:"path,omitempty"`
}

// SkillsMergeRequestsCommentsEndpointUpdateBody contains the request body for Update.
type SkillsMergeRequestsCommentsEndpointUpdateBody struct {
	Body string `json:"body"`
}

// List lists comments on a skill merge request or one of its items.
func (e *SkillsMergeRequestsCommentsEndpoint) List(skillMergeRequestId string, params *SkillsMergeRequestsCommentsEndpointListParams) (*comments.SkillsMergeRequestsCommentsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"skill-merge-requests", skillMergeRequestId, "comments"},
		Query: query,
	}
	var result comments.SkillsMergeRequestsCommentsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create adds a comment to a skill merge request or one of its items.
func (e *SkillsMergeRequestsCommentsEndpoint) Create(skillMergeRequestId string, body *SkillsMergeRequestsCommentsEndpointCreateBody) (*comments.SkillsMergeRequestsCommentsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-merge-requests", skillMergeRequestId, "comments"},
		Body: body,
	}
	var result comments.SkillsMergeRequestsCommentsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a comment on a skill merge request.
func (e *SkillsMergeRequestsCommentsEndpoint) Get(skillMergeRequestId string, commentId string) (*comments.SkillsMergeRequestsCommentsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-merge-requests", skillMergeRequestId, "comments", commentId},
	}
	var result comments.SkillsMergeRequestsCommentsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a comment authored by the current actor.
func (e *SkillsMergeRequestsCommentsEndpoint) Update(skillMergeRequestId string, commentId string, body *SkillsMergeRequestsCommentsEndpointUpdateBody) (*comments.SkillsMergeRequestsCommentsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-merge-requests", skillMergeRequestId, "comments", commentId},
		Body: body,
	}
	var result comments.SkillsMergeRequestsCommentsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a comment authored by the current actor.
func (e *SkillsMergeRequestsCommentsEndpoint) Delete(skillMergeRequestId string, commentId string) (*comments.SkillsMergeRequestsCommentsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-merge-requests", skillMergeRequestId, "comments", commentId},
	}
	var result comments.SkillsMergeRequestsCommentsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
