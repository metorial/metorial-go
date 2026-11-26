package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/mergerequests/events"
)

// SkillsMergeRequestsEventsEndpoint provides access to a skill merge request event is one entry in a merge request activity history.
type SkillsMergeRequestsEventsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsMergeRequestsEventsEndpoint creates a new SkillsMergeRequestsEventsEndpoint.
func NewSkillsMergeRequestsEventsEndpoint(client *endpoint.Client) *SkillsMergeRequestsEventsEndpoint {
	return &SkillsMergeRequestsEventsEndpoint{client: client}
}

// SkillsMergeRequestsEventsEndpointListParams contains optional query parameters for List.
type SkillsMergeRequestsEventsEndpointListParams struct {
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
	// Type - Only return events of one of these types.
	Type *any `json:"type,omitempty"`
	// CreatedAt - Only return results whose skill merge request event creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
}

// List returns a paginated activity history for a skill merge request.
func (e *SkillsMergeRequestsEventsEndpoint) List(skillMergeRequestId string, params *SkillsMergeRequestsEventsEndpointListParams) (*events.SkillsMergeRequestsEventsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"skill-merge-requests", skillMergeRequestId, "events"},
		Query: query,
	}
	var result events.SkillsMergeRequestsEventsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves one event from a skill merge request activity history.
func (e *SkillsMergeRequestsEventsEndpoint) Get(skillMergeRequestId string, eventId string) (*events.SkillsMergeRequestsEventsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-merge-requests", skillMergeRequestId, "events", eventId},
	}
	var result events.SkillsMergeRequestsEventsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
