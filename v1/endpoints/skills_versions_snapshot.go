package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/versions/snapshot"
)

// SkillsVersionsSnapshotEndpoint provides access to inspect version history and snapshots for a skill.
type SkillsVersionsSnapshotEndpoint struct {
	client *endpoint.Client
}

// NewSkillsVersionsSnapshotEndpoint creates a new SkillsVersionsSnapshotEndpoint.
func NewSkillsVersionsSnapshotEndpoint(client *endpoint.Client) *SkillsVersionsSnapshotEndpoint {
	return &SkillsVersionsSnapshotEndpoint{client: client}
}

// Get retrieves the store-backed snapshot for a specific skill version.
func (e *SkillsVersionsSnapshotEndpoint) Get(skillId string, skillVersionId string) (*snapshot.SkillsVersionsSnapshotGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skills", skillId, "versions", skillVersionId, "snapshot"},
	}
	var result snapshot.SkillsVersionsSnapshotGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
