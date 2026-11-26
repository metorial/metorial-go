package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/mergerequests/items"
)

// SkillsMergeRequestsItemsEndpoint provides access to a skill merge request is a proposed change from one skill into another.
type SkillsMergeRequestsItemsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsMergeRequestsItemsEndpoint creates a new SkillsMergeRequestsItemsEndpoint.
func NewSkillsMergeRequestsItemsEndpoint(client *endpoint.Client) *SkillsMergeRequestsItemsEndpoint {
	return &SkillsMergeRequestsItemsEndpoint{client: client}
}

// SkillsMergeRequestsItemsEndpointResolveBody contains the request body for Resolve.
type SkillsMergeRequestsItemsEndpointResolveBody struct {
	// ResolutionType - How to resolve the proposed change.
	ResolutionType string `json:"resolution_type"`
	// Resolution - Custom content for the resolution. Required for `edit_document` and `replace_file`.
	Resolution *map[string]any `json:"resolution,omitempty"`
}

// SkillsMergeRequestsItemsEndpointBulkResolveBody contains the request body for BulkResolve.
type SkillsMergeRequestsItemsEndpointBulkResolveBody struct {
	// Items - Resolutions to save.
	Items []map[string]any `json:"items"`
}

// Resolve saves a resolution for one proposed skill change. The merge request must be open.
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

// BulkResolve saves resolutions for multiple proposed skill changes. The merge request must be open.
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
