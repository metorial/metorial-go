package magicmcpendpoints

import (
	"encoding/json"
	"time"
)

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
	Object      string           `json:"object"`
	Id          string           `json:"id"`
	Status      string           `json:"status"`
	Slug        string           `json:"slug"`
	Url         string           `json:"url"`
	Servers     []map[string]any `json:"servers"`
	Name        *string          `json:"name,omitempty"`
	Description *string          `json:"description,omitempty"`
	Metadata    map[string]any   `json:"metadata"`
	// EnableSkillEditing - Whether consumers connecting through this endpoint with a consumer magic MCP token get tools to read and edit the skills they have access to. The tools are hidden while skills are disabled for the instance.
	EnableSkillEditing bool `json:"enable_skill_editing"`
	// Access - Resource ownership and access settings.
	Access    MagicMcpEndpointsListOutputItemsAccess `json:"access"`
	CreatedAt time.Time                              `json:"created_at"`
	UpdatedAt time.Time                              `json:"updated_at"`
}

// MagicMcpEndpointsListOutputPagination represents the magic mcp endpoints list output pagination type.
type MagicMcpEndpointsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// MagicMcpEndpointsListOutput represents the magic mcp endpoints list output type.
type MagicMcpEndpointsListOutput struct {
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
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Ownership - Which resources to return under the resource ownership rules. 'shared' returns only resources that aren't owned, 'usable' returns shared resources plus owned resources the caller can use, 'all' returns every resource the caller can see. Defaults to 'shared' when ownership is enforced for the project and 'all' otherwise.
	Ownership        *string `json:"ownership,omitempty"`
	Status           *any    `json:"status,omitempty"`
	MagicMcpServerId *any    `json:"magic_mcp_server_id,omitempty"`
	Search           *string `json:"search,omitempty"`
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
