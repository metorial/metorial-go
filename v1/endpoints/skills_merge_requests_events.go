package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/mergerequests/events"
)

// SkillsMergeRequestsEventsEndpoint provides access to inspect the activity history of skill merge requests.
type SkillsMergeRequestsEventsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsMergeRequestsEventsEndpoint creates a new SkillsMergeRequestsEventsEndpoint.
func NewSkillsMergeRequestsEventsEndpoint(client *endpoint.Client) *SkillsMergeRequestsEventsEndpoint {
	return &SkillsMergeRequestsEventsEndpoint{client: client}
}

// SkillsMergeRequestsEventsEndpointListParams contains optional query parameters for List.
type SkillsMergeRequestsEventsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	Type   *any     `json:"type,omitempty"`
	// CreatedAt - Filter skill merge request event creation time by date range
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
