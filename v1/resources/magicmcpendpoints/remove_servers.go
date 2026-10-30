package magicmcpendpoints

import (
	"encoding/json"
	"time"
)

// MagicMcpEndpointsRemoveServersOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type MagicMcpEndpointsRemoveServersOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// MagicMcpEndpointsRemoveServersOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type MagicMcpEndpointsRemoveServersOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// MagicMcpEndpointsRemoveServersOutputAccessOwner - The actor that owns this resource, if any.
type MagicMcpEndpointsRemoveServersOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *MagicMcpEndpointsRemoveServersOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *MagicMcpEndpointsRemoveServersOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// MagicMcpEndpointsRemoveServersOutputAccess - Resource ownership and access settings.
type MagicMcpEndpointsRemoveServersOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *MagicMcpEndpointsRemoveServersOutputAccessOwner `json:"owner,omitempty"`
}

// MagicMcpEndpointsRemoveServersOutput represents the magic mcp endpoints remove servers output type.
type MagicMcpEndpointsRemoveServersOutput struct {
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
	Access    MagicMcpEndpointsRemoveServersOutputAccess `json:"access"`
	CreatedAt time.Time                                  `json:"created_at"`
	UpdatedAt time.Time                                  `json:"updated_at"`
}

// MapMagicMcpEndpointsRemoveServersOutputFromJSON deserializes JSON data into a MagicMcpEndpointsRemoveServersOutput.
func MapMagicMcpEndpointsRemoveServersOutputFromJSON(data []byte) (*MagicMcpEndpointsRemoveServersOutput, error) {
	var v MagicMcpEndpointsRemoveServersOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpEndpointsRemoveServersOutputToJSON serializes a MagicMcpEndpointsRemoveServersOutput to JSON.
func MapMagicMcpEndpointsRemoveServersOutputToJSON(v *MagicMcpEndpointsRemoveServersOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpEndpointsRemoveServersBody represents the magic mcp endpoints remove servers body type.
type MagicMcpEndpointsRemoveServersBody struct {
	MagicMcpServerIds []string `json:"magic_mcp_server_ids"`
}

// MapMagicMcpEndpointsRemoveServersBodyFromJSON deserializes JSON data into a MagicMcpEndpointsRemoveServersBody.
func MapMagicMcpEndpointsRemoveServersBodyFromJSON(data []byte) (*MagicMcpEndpointsRemoveServersBody, error) {
	var v MagicMcpEndpointsRemoveServersBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpEndpointsRemoveServersBodyToJSON serializes a MagicMcpEndpointsRemoveServersBody to JSON.
func MapMagicMcpEndpointsRemoveServersBodyToJSON(v *MagicMcpEndpointsRemoveServersBody) ([]byte, error) {
	return json.Marshal(v)
}
