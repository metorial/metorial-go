package magicmcpendpoints

import (
	"encoding/json"
	"time"
)

// MagicMcpEndpointsAddServersOutputSkillSelection represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type MagicMcpEndpointsAddServersOutputSkillSelection struct {
	// Type - Only the listed skills and the skills in the listed groups are exposed.
	Type *string `json:"type,omitempty"`
	// SkillIds - IDs of the selected skills.
	SkillIds *[]string `json:"skill_ids,omitempty"`
	// SkillGroupIds - IDs of the selected skill groups.
	SkillGroupIds *[]string `json:"skill_group_ids,omitempty"`
}

// MagicMcpEndpointsAddServersOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type MagicMcpEndpointsAddServersOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// MagicMcpEndpointsAddServersOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type MagicMcpEndpointsAddServersOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// MagicMcpEndpointsAddServersOutputAccessOwner - The actor that owns this resource, if any.
type MagicMcpEndpointsAddServersOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *MagicMcpEndpointsAddServersOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *MagicMcpEndpointsAddServersOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// MagicMcpEndpointsAddServersOutputAccess - Resource ownership and access settings.
type MagicMcpEndpointsAddServersOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *MagicMcpEndpointsAddServersOutputAccessOwner `json:"owner,omitempty"`
}

// MagicMcpEndpointsAddServersOutput represents the magic mcp endpoints add servers output type.
type MagicMcpEndpointsAddServersOutput struct {
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
	SkillSelection MagicMcpEndpointsAddServersOutputSkillSelection `json:"skill_selection"`
	// Access - Resource ownership and access settings.
	Access MagicMcpEndpointsAddServersOutputAccess `json:"access"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapMagicMcpEndpointsAddServersOutputFromJSON deserializes JSON data into a MagicMcpEndpointsAddServersOutput.
func MapMagicMcpEndpointsAddServersOutputFromJSON(data []byte) (*MagicMcpEndpointsAddServersOutput, error) {
	var v MagicMcpEndpointsAddServersOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpEndpointsAddServersOutputToJSON serializes a MagicMcpEndpointsAddServersOutput to JSON.
func MapMagicMcpEndpointsAddServersOutputToJSON(v *MagicMcpEndpointsAddServersOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpEndpointsAddServersBodyMagicMcpServers represents the magic mcp endpoints add servers body magic mcp servers type.
type MagicMcpEndpointsAddServersBodyMagicMcpServers struct {
	// MagicMcpServerId - Server to route to.
	MagicMcpServerId string `json:"magic_mcp_server_id"`
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
}

// MagicMcpEndpointsAddServersBody represents the magic mcp endpoints add servers body type.
type MagicMcpEndpointsAddServersBody struct {
	// MagicMcpServers - Servers to add to the endpoint.
	MagicMcpServers *[]MagicMcpEndpointsAddServersBodyMagicMcpServers `json:"magic_mcp_servers,omitempty"`
}

// MapMagicMcpEndpointsAddServersBodyFromJSON deserializes JSON data into a MagicMcpEndpointsAddServersBody.
func MapMagicMcpEndpointsAddServersBodyFromJSON(data []byte) (*MagicMcpEndpointsAddServersBody, error) {
	var v MagicMcpEndpointsAddServersBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpEndpointsAddServersBodyToJSON serializes a MagicMcpEndpointsAddServersBody to JSON.
func MapMagicMcpEndpointsAddServersBodyToJSON(v *MagicMcpEndpointsAddServersBody) ([]byte, error) {
	return json.Marshal(v)
}
