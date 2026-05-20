package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/participants"
)

// SkillsParticipantsEndpoint provides access to inspect participants associated with an instance skill.
type SkillsParticipantsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsParticipantsEndpoint creates a new SkillsParticipantsEndpoint.
func NewSkillsParticipantsEndpoint(client *endpoint.Client) *SkillsParticipantsEndpoint {
	return &SkillsParticipantsEndpoint{client: client}
}

// SkillsParticipantsEndpointListParams contains optional query parameters for List.
type SkillsParticipantsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
}

// List returns a paginated list of participants for a specific skill.
func (e *SkillsParticipantsEndpoint) List(instanceId string, skillId string, params *SkillsParticipantsEndpointListParams) (*participants.SkillsParticipantsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "skills", skillId, "participants"},
		Query: query,
	}
	var result participants.SkillsParticipantsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific participant within a skill.
func (e *SkillsParticipantsEndpoint) Get(instanceId string, skillId string, skillParticipantId string) (*participants.SkillsParticipantsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skills", skillId, "participants", skillParticipantId},
	}
	var result participants.SkillsParticipantsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
