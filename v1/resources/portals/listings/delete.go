package listings

import (
	"encoding/json"
	"time"
)

// PortalsListingsDeleteOutputAccessProviderTemplate represents the portals listings delete output access provider template type.
type PortalsListingsDeleteOutputAccessProviderTemplate struct {
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

// PortalsListingsDeleteOutputAccessMagicMcpServer represents the portals listings delete output access magic mcp server type.
type PortalsListingsDeleteOutputAccessMagicMcpServer struct {
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

// PortalsListingsDeleteOutputAccessSkill represents the portals listings delete output access skill type.
type PortalsListingsDeleteOutputAccessSkill struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

// PortalsListingsDeleteOutputAccessSkillTemplate represents the portals listings delete output access skill template type.
type PortalsListingsDeleteOutputAccessSkillTemplate struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Owner       string  `json:"owner"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// PortalsListingsDeleteOutputAccessSkillGroup represents the portals listings delete output access skill group type.
type PortalsListingsDeleteOutputAccessSkillGroup struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// PortalsListingsDeleteOutputAccessSkillMarketplace represents the portals listings delete output access skill marketplace type.
type PortalsListingsDeleteOutputAccessSkillMarketplace struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
}

// PortalsListingsDeleteOutputAccessSkillPlugin represents the portals listings delete output access skill plugin type.
type PortalsListingsDeleteOutputAccessSkillPlugin struct {
	Object string  `json:"object"`
	Id     string  `json:"id"`
	Status string  `json:"status"`
	Name   *string `json:"name,omitempty"`
}

// PortalsListingsDeleteOutputAccess represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type PortalsListingsDeleteOutputAccess struct {
	Type             *string                                            `json:"type,omitempty"`
	ProviderTemplate *PortalsListingsDeleteOutputAccessProviderTemplate `json:"provider_template,omitempty"`
	MagicMcpServer   *PortalsListingsDeleteOutputAccessMagicMcpServer   `json:"magic_mcp_server,omitempty"`
	Skill            *PortalsListingsDeleteOutputAccessSkill            `json:"skill,omitempty"`
	SkillTemplate    *PortalsListingsDeleteOutputAccessSkillTemplate    `json:"skill_template,omitempty"`
	SkillGroup       *PortalsListingsDeleteOutputAccessSkillGroup       `json:"skill_group,omitempty"`
	SkillMarketplace *PortalsListingsDeleteOutputAccessSkillMarketplace `json:"skill_marketplace,omitempty"`
	SkillPlugin      *PortalsListingsDeleteOutputAccessSkillPlugin      `json:"skill_plugin,omitempty"`
}

// PortalsListingsDeleteOutputGroups - Catalog section the listing appears in.
type PortalsListingsDeleteOutputGroups struct {
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Index - Order of the group in the portal catalog.
	Index float64 `json:"index"`
}

// PortalsListingsDeleteOutput represents the portals listings delete output type.
type PortalsListingsDeleteOutput struct {
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
	Access PortalsListingsDeleteOutputAccess `json:"access"`
	// Groups - Catalog sections the listing appears in, in display order.
	Groups []PortalsListingsDeleteOutputGroups `json:"groups"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapPortalsListingsDeleteOutputFromJSON deserializes JSON data into a PortalsListingsDeleteOutput.
func MapPortalsListingsDeleteOutputFromJSON(data []byte) (*PortalsListingsDeleteOutput, error) {
	var v PortalsListingsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsListingsDeleteOutputToJSON serializes a PortalsListingsDeleteOutput to JSON.
func MapPortalsListingsDeleteOutputToJSON(v *PortalsListingsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
