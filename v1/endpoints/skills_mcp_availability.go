package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/mcpavailability"
)

// SkillsMcpAvailabilityEndpoint provides access to read and edit the files of a skill by path.
type SkillsMcpAvailabilityEndpoint struct {
	client *endpoint.Client
}

// NewSkillsMcpAvailabilityEndpoint creates a new SkillsMcpAvailabilityEndpoint.
func NewSkillsMcpAvailabilityEndpoint(client *endpoint.Client) *SkillsMcpAvailabilityEndpoint {
	return &SkillsMcpAvailabilityEndpoint{client: client}
}

// Get returns whether apps connected over MCP can see the skill. Skills with too many files or too large a total size are left out of MCP connections and keep working everywhere else.
func (e *SkillsMcpAvailabilityEndpoint) Get(skillId string) (*mcpavailability.SkillsMcpAvailabilityGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skills", skillId, "mcp-availability"},
	}
	var result mcpavailability.SkillsMcpAvailabilityGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
