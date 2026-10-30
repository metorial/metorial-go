package magicmcpendpoints

import (
	"encoding/json"
	"time"
)

// MagicMcpEndpointsCreateOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type MagicMcpEndpointsCreateOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// MagicMcpEndpointsCreateOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type MagicMcpEndpointsCreateOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// MagicMcpEndpointsCreateOutputAccessOwner - The actor that owns this resource, if any.
type MagicMcpEndpointsCreateOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *MagicMcpEndpointsCreateOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *MagicMcpEndpointsCreateOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// MagicMcpEndpointsCreateOutputAccess - Resource ownership and access settings.
type MagicMcpEndpointsCreateOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *MagicMcpEndpointsCreateOutputAccessOwner `json:"owner,omitempty"`
}

// MagicMcpEndpointsCreateOutput represents the magic mcp endpoints create output type.
type MagicMcpEndpointsCreateOutput struct {
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
	Access    MagicMcpEndpointsCreateOutputAccess `json:"access"`
	CreatedAt time.Time                           `json:"created_at"`
	UpdatedAt time.Time                           `json:"updated_at"`
}

// MapMagicMcpEndpointsCreateOutputFromJSON deserializes JSON data into a MagicMcpEndpointsCreateOutput.
func MapMagicMcpEndpointsCreateOutputFromJSON(data []byte) (*MagicMcpEndpointsCreateOutput, error) {
	var v MagicMcpEndpointsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpEndpointsCreateOutputToJSON serializes a MagicMcpEndpointsCreateOutput to JSON.
func MapMagicMcpEndpointsCreateOutputToJSON(v *MagicMcpEndpointsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpEndpointsCreateBodyMagicMcpServers represents the magic mcp endpoints create body magic mcp servers type.
type MagicMcpEndpointsCreateBodyMagicMcpServers struct {
	MagicMcpServerId string `json:"magic_mcp_server_id"`
	ToolFilters      *any   `json:"tool_filters,omitempty"`
}

// MagicMcpEndpointsCreateBody represents the magic mcp endpoints create body type.
type MagicMcpEndpointsCreateBody struct {
	Name              *string         `json:"name,omitempty"`
	Description       *string         `json:"description,omitempty"`
	Metadata          *map[string]any `json:"metadata,omitempty"`
	ConsumerProfileId *string         `json:"consumer_profile_id,omitempty"`
	SkillPluginId     *string         `json:"skill_plugin_id,omitempty"`
	// EnableSkillEditing - Expose skill editing tools on this endpoint so the connecting consumer can read and edit the skills they have access to. Only connections authenticated with a magic MCP token bound to a single consumer get the tools; connections using a publishable key do not. Requires skills to be enabled for the instance and consumer skill editing to be allowed for the project. The tools disappear if either is turned off later.
	EnableSkillEditing *bool                                         `json:"enable_skill_editing,omitempty"`
	MagicMcpServers    *[]MagicMcpEndpointsCreateBodyMagicMcpServers `json:"magic_mcp_servers,omitempty"`
}

// MapMagicMcpEndpointsCreateBodyFromJSON deserializes JSON data into a MagicMcpEndpointsCreateBody.
func MapMagicMcpEndpointsCreateBodyFromJSON(data []byte) (*MagicMcpEndpointsCreateBody, error) {
	var v MagicMcpEndpointsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpEndpointsCreateBodyToJSON serializes a MagicMcpEndpointsCreateBody to JSON.
func MapMagicMcpEndpointsCreateBodyToJSON(v *MagicMcpEndpointsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
