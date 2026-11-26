package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/versions/snapshot"
)

// SkillsVersionsSnapshotEndpoint provides access to a skill version is a snapshot of a skill at a point in time.
type SkillsVersionsSnapshotEndpoint struct {
	client *endpoint.Client
}

// NewSkillsVersionsSnapshotEndpoint creates a new SkillsVersionsSnapshotEndpoint.
func NewSkillsVersionsSnapshotEndpoint(client *endpoint.Client) *SkillsVersionsSnapshotEndpoint {
	return &SkillsVersionsSnapshotEndpoint{client: client}
}

// Get retrieves the content snapshot for a skill version.
func (e *SkillsVersionsSnapshotEndpoint) Get(instanceId string, skillId string, skillVersionId string) (*snapshot.SkillsVersionsSnapshotGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skills", skillId, "versions", skillVersionId, "snapshot"},
	}
	var result snapshot.SkillsVersionsSnapshotGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
