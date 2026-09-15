package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/forksyncs"
)

// SkillsForkSyncsEndpoint provides access to synchronize changes from an upstream skill into a fork.
type SkillsForkSyncsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsForkSyncsEndpoint creates a new SkillsForkSyncsEndpoint.
func NewSkillsForkSyncsEndpoint(client *endpoint.Client) *SkillsForkSyncsEndpoint {
	return &SkillsForkSyncsEndpoint{client: client}
}

// SkillsForkSyncsEndpointCreateBody contains the request body for Create.
type SkillsForkSyncsEndpointCreateBody struct {
	SkillId string `json:"skill_id"`
}

// Create queues synchronization of upstream changes into a forked skill.
func (e *SkillsForkSyncsEndpoint) Create(body *SkillsForkSyncsEndpointCreateBody) (*forksyncs.SkillsForkSyncsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-fork-syncs"},
		Body: body,
	}
	var result forksyncs.SkillsForkSyncsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves the state of a fork synchronization.
func (e *SkillsForkSyncsEndpoint) Get(skillForkSyncId string) (*forksyncs.SkillsForkSyncsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-fork-syncs", skillForkSyncId},
	}
	var result forksyncs.SkillsForkSyncsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
