package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/access"
)

// SkillsAccessEndpoint provides access to inspect who has been granted access to an instance skill.
type SkillsAccessEndpoint struct {
	client *endpoint.Client
}

// NewSkillsAccessEndpoint creates a new SkillsAccessEndpoint.
func NewSkillsAccessEndpoint(client *endpoint.Client) *SkillsAccessEndpoint {
	return &SkillsAccessEndpoint{client: client}
}

// Get returns the access grants of a specific skill.
func (e *SkillsAccessEndpoint) Get(instanceId string, skillId string) (*access.SkillsAccessGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skills", skillId, "access"},
	}
	var result access.SkillsAccessGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
