package listings

import (
	"encoding/json"
	"time"
)

// PortalsListingsCreateOutputAccessProviderTemplate represents the portals listings create output access provider template type.
type PortalsListingsCreateOutputAccessProviderTemplate struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once you archive the template.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata map[string]any `json:"metadata"`
	// IntegrationId - Integration the template was created from. `null` when it wasn’t created from one.
	IntegrationId *string `json:"integration_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// PortalsListingsCreateOutputAccessMagicMcpServer represents the portals listings create output access magic mcp server type.
type PortalsListingsCreateOutputAccessMagicMcpServer struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the server is deleted.
	Status string `json:"status"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
}

// PortalsListingsCreateOutputAccessSkill represents the portals listings create output access skill type.
type PortalsListingsCreateOutputAccessSkill struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

// PortalsListingsCreateOutputAccessSkillTemplate represents the portals listings create output access skill template type.
type PortalsListingsCreateOutputAccessSkillTemplate struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Owner       string  `json:"owner"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// PortalsListingsCreateOutputAccessSkillGroup represents the portals listings create output access skill group type.
type PortalsListingsCreateOutputAccessSkillGroup struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// PortalsListingsCreateOutputAccessSkillMarketplace represents the portals listings create output access skill marketplace type.
type PortalsListingsCreateOutputAccessSkillMarketplace struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
}

// PortalsListingsCreateOutputAccessSkillPlugin represents the portals listings create output access skill plugin type.
type PortalsListingsCreateOutputAccessSkillPlugin struct {
	Object string  `json:"object"`
	Id     string  `json:"id"`
	Status string  `json:"status"`
	Name   *string `json:"name,omitempty"`
}

// PortalsListingsCreateOutputAccess represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type PortalsListingsCreateOutputAccess struct {
	Type             *string                                            `json:"type,omitempty"`
	ProviderTemplate *PortalsListingsCreateOutputAccessProviderTemplate `json:"provider_template,omitempty"`
	MagicMcpServer   *PortalsListingsCreateOutputAccessMagicMcpServer   `json:"magic_mcp_server,omitempty"`
	Skill            *PortalsListingsCreateOutputAccessSkill            `json:"skill,omitempty"`
	SkillTemplate    *PortalsListingsCreateOutputAccessSkillTemplate    `json:"skill_template,omitempty"`
	SkillGroup       *PortalsListingsCreateOutputAccessSkillGroup       `json:"skill_group,omitempty"`
	SkillMarketplace *PortalsListingsCreateOutputAccessSkillMarketplace `json:"skill_marketplace,omitempty"`
	SkillPlugin      *PortalsListingsCreateOutputAccessSkillPlugin      `json:"skill_plugin,omitempty"`
}

// PortalsListingsCreateOutputGroups - Catalog section the listing appears in.
type PortalsListingsCreateOutputGroups struct {
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Index - Order of the group in the portal catalog.
	Index float64 `json:"index"`
}

// PortalsListingsCreateOutput represents the portals listings create output type.
type PortalsListingsCreateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Readme - Longer guide shown to consumers.
	Readme *string `json:"readme,omitempty"`
	// IsDiscoverable - Whether every consumer on the portal can find this listing in the catalog and request access. When false, only consumers who already have access through one of their groups see it.
	IsDiscoverable bool `json:"is_discoverable"`
	// Access - The integration, skill, or skill set the listing is for.
	Access PortalsListingsCreateOutputAccess `json:"access"`
	// Groups - Catalog sections the listing appears in, in display order.
	Groups []PortalsListingsCreateOutputGroups `json:"groups"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapPortalsListingsCreateOutputFromJSON deserializes JSON data into a PortalsListingsCreateOutput.
func MapPortalsListingsCreateOutputFromJSON(data []byte) (*PortalsListingsCreateOutput, error) {
	var v PortalsListingsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsListingsCreateOutputToJSON serializes a PortalsListingsCreateOutput to JSON.
func MapPortalsListingsCreateOutputToJSON(v *PortalsListingsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsListingsCreateBodyAccess represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type PortalsListingsCreateBodyAccess struct {
	Type *string `json:"type,omitempty"`
	// ProviderTemplateId - Provider template the listing is for.
	ProviderTemplateId *string `json:"provider_template_id,omitempty"`
	// MagicMcpServerId - Magic MCP server the listing is for.
	MagicMcpServerId *string `json:"magic_mcp_server_id,omitempty"`
	// SkillId - Skill the listing is for.
	SkillId *string `json:"skill_id,omitempty"`
	// SkillTemplateId - Skill template the listing is for.
	SkillTemplateId *string `json:"skill_template_id,omitempty"`
	// SkillGroupId - Skill group the listing is for.
	SkillGroupId *string `json:"skill_group_id,omitempty"`
	// SkillMarketplaceId - Skill marketplace the listing is for.
	SkillMarketplaceId *string `json:"skill_marketplace_id,omitempty"`
}

// PortalsListingsCreateBody represents the portals listings create body type.
type PortalsListingsCreateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Readme - Longer guide shown to consumers.
	Readme *string `json:"readme,omitempty"`
	// IsDiscoverable - Whether every consumer on the portal can find this listing in the catalog and request access. When false, only consumers who already have access through one of their groups see it. Defaults to true.
	IsDiscoverable *bool `json:"is_discoverable,omitempty"`
	// Access - The integration, skill, or skill set the listing is for.
	Access PortalsListingsCreateBodyAccess `json:"access"`
}

// MapPortalsListingsCreateBodyFromJSON deserializes JSON data into a PortalsListingsCreateBody.
func MapPortalsListingsCreateBodyFromJSON(data []byte) (*PortalsListingsCreateBody, error) {
	var v PortalsListingsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsListingsCreateBodyToJSON serializes a PortalsListingsCreateBody to JSON.
func MapPortalsListingsCreateBodyToJSON(v *PortalsListingsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
