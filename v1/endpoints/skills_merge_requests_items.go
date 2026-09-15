package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/mergerequests/items"
)

// SkillsMergeRequestsItemsEndpoint provides access to review, resolve, and apply changes between skills.
type SkillsMergeRequestsItemsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsMergeRequestsItemsEndpoint creates a new SkillsMergeRequestsItemsEndpoint.
func NewSkillsMergeRequestsItemsEndpoint(client *endpoint.Client) *SkillsMergeRequestsItemsEndpoint {
	return &SkillsMergeRequestsItemsEndpoint{client: client}
}

// SkillsMergeRequestsItemsEndpointResolveBody contains the request body for Resolve.
type SkillsMergeRequestsItemsEndpointResolveBody struct {
	ResolutionType string          `json:"resolution_type"`
	Resolution     *map[string]any `json:"resolution,omitempty"`
}

// SkillsMergeRequestsItemsEndpointBulkResolveBody contains the request body for BulkResolve.
type SkillsMergeRequestsItemsEndpointBulkResolveBody struct {
	Items []map[string]any `json:"items"`
}

// Resolve saves a resolution for one proposed skill change.
func (e *SkillsMergeRequestsItemsEndpoint) Resolve(skillMergeRequestId string, itemId string, body *SkillsMergeRequestsItemsEndpointResolveBody) (*items.SkillsMergeRequestsItemsResolveOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-merge-requests", skillMergeRequestId, "items", itemId},
		Body: body,
	}
	var result items.SkillsMergeRequestsItemsResolveOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BulkResolve saves resolutions for multiple proposed skill changes.
func (e *SkillsMergeRequestsItemsEndpoint) BulkResolve(skillMergeRequestId string, body *SkillsMergeRequestsItemsEndpointBulkResolveBody) (*items.SkillsMergeRequestsItemsBulkResolveOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-merge-requests", skillMergeRequestId, "items"},
		Body: body,
	}
	var result items.SkillsMergeRequestsItemsBulkResolveOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
