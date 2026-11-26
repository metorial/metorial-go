package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/participants"
)

// SkillsParticipantsEndpoint provides access to a skill participant is a member or consumer participating in a skill.
type SkillsParticipantsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsParticipantsEndpoint creates a new SkillsParticipantsEndpoint.
func NewSkillsParticipantsEndpoint(client *endpoint.Client) *SkillsParticipantsEndpoint {
	return &SkillsParticipantsEndpoint{client: client}
}

// SkillsParticipantsEndpointListParams contains optional query parameters for List.
type SkillsParticipantsEndpointListParams struct {
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
}

// List returns a paginated list of skill participants.
func (e *SkillsParticipantsEndpoint) List(skillId string, params *SkillsParticipantsEndpointListParams) (*participants.SkillsParticipantsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"skills", skillId, "participants"},
		Query: query,
	}
	var result participants.SkillsParticipantsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a skill participant.
func (e *SkillsParticipantsEndpoint) Get(skillId string, skillParticipantId string) (*participants.SkillsParticipantsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skills", skillId, "participants", skillParticipantId},
	}
	var result participants.SkillsParticipantsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
