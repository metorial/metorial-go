package listings

import (
	"encoding/json"
	"time"
)

// PortalsListingsGetOutputAccessProviderTemplate represents the portals listings get output access provider template type.
type PortalsListingsGetOutputAccessProviderTemplate struct {
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

// PortalsListingsGetOutputAccessMagicMcpServer represents the portals listings get output access magic mcp server type.
type PortalsListingsGetOutputAccessMagicMcpServer struct {
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

// PortalsListingsGetOutputAccessSkill represents the portals listings get output access skill type.
type PortalsListingsGetOutputAccessSkill struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

// PortalsListingsGetOutputAccessSkillTemplate represents the portals listings get output access skill template type.
type PortalsListingsGetOutputAccessSkillTemplate struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Owner       string  `json:"owner"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// PortalsListingsGetOutputAccessSkillGroup represents the portals listings get output access skill group type.
type PortalsListingsGetOutputAccessSkillGroup struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// PortalsListingsGetOutputAccessSkillMarketplace represents the portals listings get output access skill marketplace type.
type PortalsListingsGetOutputAccessSkillMarketplace struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
}

// PortalsListingsGetOutputAccessSkillPlugin represents the portals listings get output access skill plugin type.
type PortalsListingsGetOutputAccessSkillPlugin struct {
	Object string  `json:"object"`
	Id     string  `json:"id"`
	Status string  `json:"status"`
	Name   *string `json:"name,omitempty"`
}

// PortalsListingsGetOutputAccess represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type PortalsListingsGetOutputAccess struct {
	Type             *string                                         `json:"type,omitempty"`
	ProviderTemplate *PortalsListingsGetOutputAccessProviderTemplate `json:"provider_template,omitempty"`
	MagicMcpServer   *PortalsListingsGetOutputAccessMagicMcpServer   `json:"magic_mcp_server,omitempty"`
	Skill            *PortalsListingsGetOutputAccessSkill            `json:"skill,omitempty"`
	SkillTemplate    *PortalsListingsGetOutputAccessSkillTemplate    `json:"skill_template,omitempty"`
	SkillGroup       *PortalsListingsGetOutputAccessSkillGroup       `json:"skill_group,omitempty"`
	SkillMarketplace *PortalsListingsGetOutputAccessSkillMarketplace `json:"skill_marketplace,omitempty"`
	SkillPlugin      *PortalsListingsGetOutputAccessSkillPlugin      `json:"skill_plugin,omitempty"`
}

// PortalsListingsGetOutputGroups - Catalog section the listing appears in.
type PortalsListingsGetOutputGroups struct {
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Index - Order of the group in the portal catalog.
	Index float64 `json:"index"`
}

// PortalsListingsGetOutput represents the portals listings get output type.
type PortalsListingsGetOutput struct {
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
	Access PortalsListingsGetOutputAccess `json:"access"`
	// Groups - Catalog sections the listing appears in, in display order.
	Groups []PortalsListingsGetOutputGroups `json:"groups"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapPortalsListingsGetOutputFromJSON deserializes JSON data into a PortalsListingsGetOutput.
func MapPortalsListingsGetOutputFromJSON(data []byte) (*PortalsListingsGetOutput, error) {
	var v PortalsListingsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsListingsGetOutputToJSON serializes a PortalsListingsGetOutput to JSON.
func MapPortalsListingsGetOutputToJSON(v *PortalsListingsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
