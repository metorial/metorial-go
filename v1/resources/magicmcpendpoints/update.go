package magicmcpendpoints

import (
	"encoding/json"
	"time"
)

// MagicMcpEndpointsUpdateOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type MagicMcpEndpointsUpdateOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// MagicMcpEndpointsUpdateOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type MagicMcpEndpointsUpdateOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// MagicMcpEndpointsUpdateOutputAccessOwner - The actor that owns this resource, if any.
type MagicMcpEndpointsUpdateOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *MagicMcpEndpointsUpdateOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *MagicMcpEndpointsUpdateOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// MagicMcpEndpointsUpdateOutputAccess - Resource ownership and access settings.
type MagicMcpEndpointsUpdateOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *MagicMcpEndpointsUpdateOutputAccessOwner `json:"owner,omitempty"`
}

// MagicMcpEndpointsUpdateOutput represents the magic mcp endpoints update output type.
type MagicMcpEndpointsUpdateOutput struct {
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
	Access    MagicMcpEndpointsUpdateOutputAccess `json:"access"`
	CreatedAt time.Time                           `json:"created_at"`
	UpdatedAt time.Time                           `json:"updated_at"`
}

// MapMagicMcpEndpointsUpdateOutputFromJSON deserializes JSON data into a MagicMcpEndpointsUpdateOutput.
func MapMagicMcpEndpointsUpdateOutputFromJSON(data []byte) (*MagicMcpEndpointsUpdateOutput, error) {
	var v MagicMcpEndpointsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpEndpointsUpdateOutputToJSON serializes a MagicMcpEndpointsUpdateOutput to JSON.
func MapMagicMcpEndpointsUpdateOutputToJSON(v *MagicMcpEndpointsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpEndpointsUpdateBody represents the magic mcp endpoints update body type.
type MagicMcpEndpointsUpdateBody struct {
	Name        *string         `json:"name,omitempty"`
	Description *string         `json:"description,omitempty"`
	Metadata    *map[string]any `json:"metadata,omitempty"`
	// EnableSkillEditing - Expose skill editing tools on this endpoint so the connecting consumer can read and edit the skills they have access to. Only connections authenticated with a magic MCP token bound to a single consumer get the tools; connections using a publishable key do not. Requires skills to be enabled for the instance and consumer skill editing to be allowed for the project. The tools disappear if either is turned off later.
	EnableSkillEditing *bool `json:"enable_skill_editing,omitempty"`
}

// MapMagicMcpEndpointsUpdateBodyFromJSON deserializes JSON data into a MagicMcpEndpointsUpdateBody.
func MapMagicMcpEndpointsUpdateBodyFromJSON(data []byte) (*MagicMcpEndpointsUpdateBody, error) {
	var v MagicMcpEndpointsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpEndpointsUpdateBodyToJSON serializes a MagicMcpEndpointsUpdateBody to JSON.
func MapMagicMcpEndpointsUpdateBodyToJSON(v *MagicMcpEndpointsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
