package magicmcpendpoints

import (
	"encoding/json"
	"time"
)

// MagicMcpEndpointsListOutputItemsSkillSelection represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type MagicMcpEndpointsListOutputItemsSkillSelection struct {
	// Type - Only the listed skills and the skills in the listed groups are exposed.
	Type *string `json:"type,omitempty"`
	// SkillIds - IDs of the selected skills.
	SkillIds *[]string `json:"skill_ids,omitempty"`
	// SkillGroupIds - IDs of the selected skill groups.
	SkillGroupIds *[]string `json:"skill_group_ids,omitempty"`
}

// MagicMcpEndpointsListOutputItemsAccessOwnerOrganizationMember - Organization member associated with this owner.
type MagicMcpEndpointsListOutputItemsAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// MagicMcpEndpointsListOutputItemsAccessOwnerConsumer - Consumer and profile associated with this owner.
type MagicMcpEndpointsListOutputItemsAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// MagicMcpEndpointsListOutputItemsAccessOwner - The actor that owns this resource, if any.
type MagicMcpEndpointsListOutputItemsAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *MagicMcpEndpointsListOutputItemsAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *MagicMcpEndpointsListOutputItemsAccessOwnerConsumer `json:"consumer,omitempty"`
}

// MagicMcpEndpointsListOutputItemsAccess - Resource ownership and access settings.
type MagicMcpEndpointsListOutputItemsAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *MagicMcpEndpointsListOutputItemsAccessOwner `json:"owner,omitempty"`
}

// MagicMcpEndpointsListOutputItems represents the magic mcp endpoints list output items type.
type MagicMcpEndpointsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once you delete the endpoint. Archived endpoints stay listable.
	Status string `json:"status"`
	// Slug - URL-safe name, unique within the instance.
	Slug string `json:"slug"`
	// Url - Connection URL for the endpoint.
	Url string `json:"url"`
	// Servers - Servers the endpoint routes to.
	Servers []map[string]any `json:"servers"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata map[string]any `json:"metadata"`
	// EnableSkillEditing - Whether consumers connecting through this endpoint with a consumer magic MCP token get tools to read and edit the skills they have access to. The tools are hidden while skills are disabled for the instance.
	EnableSkillEditing bool `json:"enable_skill_editing"`
	// EnableSkillResources - Whether consumers connecting through this endpoint with a consumer magic MCP token can read the skills they have access to over MCP, as skill:// resources with skills/list and skills/get. Skills are hidden while skills are disabled for the instance.
	EnableSkillResources bool `json:"enable_skill_resources"`
	// SkillSelection - Which skills the endpoint exposes, both as skill resources and through the skill editing tools.
	SkillSelection MagicMcpEndpointsListOutputItemsSkillSelection `json:"skill_selection"`
	// Access - Resource ownership and access settings.
	Access MagicMcpEndpointsListOutputItemsAccess `json:"access"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpEndpointsListOutputPagination represents the magic mcp endpoints list output pagination type.
type MagicMcpEndpointsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// MagicMcpEndpointsListOutput represents the magic mcp endpoints list output type.
type MagicMcpEndpointsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                `json:"object"`
	Items      []MagicMcpEndpointsListOutputItems    `json:"items"`
	Pagination MagicMcpEndpointsListOutputPagination `json:"pagination"`
}

// MapMagicMcpEndpointsListOutputFromJSON deserializes JSON data into a MagicMcpEndpointsListOutput.
func MapMagicMcpEndpointsListOutputFromJSON(data []byte) (*MagicMcpEndpointsListOutput, error) {
	var v MagicMcpEndpointsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpEndpointsListOutputToJSON serializes a MagicMcpEndpointsListOutput to JSON.
func MapMagicMcpEndpointsListOutputToJSON(v *MagicMcpEndpointsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpEndpointsListQuery represents the magic mcp endpoints list query type.
type MagicMcpEndpointsListQuery struct {
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
	// Ownership - Which resources to return under the resource ownership rules. 'shared' returns only resources that aren't owned, 'usable' returns shared resources plus owned resources the caller can use, 'all' returns every resource the caller can see. Defaults to 'usable' for consumers, and otherwise to 'shared' when ownership is enforced for the project and 'all' when it isn't.
	Ownership *string `json:"ownership,omitempty"`
	// Status - Only return endpoints with one of these statuses.
	Status *any `json:"status,omitempty"`
	// MagicMcpServerId - Only return endpoints routing to one of these servers.
	MagicMcpServerId *any `json:"magic_mcp_server_id,omitempty"`
	// Search - Search by name or description.
	Search *string `json:"search,omitempty"`
}

// MapMagicMcpEndpointsListQueryFromJSON deserializes JSON data into a MagicMcpEndpointsListQuery.
func MapMagicMcpEndpointsListQueryFromJSON(data []byte) (*MagicMcpEndpointsListQuery, error) {
	var v MagicMcpEndpointsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpEndpointsListQueryToJSON serializes a MagicMcpEndpointsListQuery to JSON.
func MapMagicMcpEndpointsListQueryToJSON(v *MagicMcpEndpointsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
