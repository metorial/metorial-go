package listings

import (
	"encoding/json"
	"time"
)

// PortalsListingsUpdateOutputAccessProviderTemplate represents the portals listings update output access provider template type.
type PortalsListingsUpdateOutputAccessProviderTemplate struct {
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

// PortalsListingsUpdateOutputAccessMagicMcpServer represents the portals listings update output access magic mcp server type.
type PortalsListingsUpdateOutputAccessMagicMcpServer struct {
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

// PortalsListingsUpdateOutputAccessSkill represents the portals listings update output access skill type.
type PortalsListingsUpdateOutputAccessSkill struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

// PortalsListingsUpdateOutputAccessSkillTemplate represents the portals listings update output access skill template type.
type PortalsListingsUpdateOutputAccessSkillTemplate struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Owner       string  `json:"owner"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// PortalsListingsUpdateOutputAccessSkillGroup represents the portals listings update output access skill group type.
type PortalsListingsUpdateOutputAccessSkillGroup struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// PortalsListingsUpdateOutputAccessSkillMarketplace represents the portals listings update output access skill marketplace type.
type PortalsListingsUpdateOutputAccessSkillMarketplace struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
}

// PortalsListingsUpdateOutputAccessSkillPlugin represents the portals listings update output access skill plugin type.
type PortalsListingsUpdateOutputAccessSkillPlugin struct {
	Object string  `json:"object"`
	Id     string  `json:"id"`
	Status string  `json:"status"`
	Name   *string `json:"name,omitempty"`
}

// PortalsListingsUpdateOutputAccess represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type PortalsListingsUpdateOutputAccess struct {
	Type             *string                                            `json:"type,omitempty"`
	ProviderTemplate *PortalsListingsUpdateOutputAccessProviderTemplate `json:"provider_template,omitempty"`
	MagicMcpServer   *PortalsListingsUpdateOutputAccessMagicMcpServer   `json:"magic_mcp_server,omitempty"`
	Skill            *PortalsListingsUpdateOutputAccessSkill            `json:"skill,omitempty"`
	SkillTemplate    *PortalsListingsUpdateOutputAccessSkillTemplate    `json:"skill_template,omitempty"`
	SkillGroup       *PortalsListingsUpdateOutputAccessSkillGroup       `json:"skill_group,omitempty"`
	SkillMarketplace *PortalsListingsUpdateOutputAccessSkillMarketplace `json:"skill_marketplace,omitempty"`
	SkillPlugin      *PortalsListingsUpdateOutputAccessSkillPlugin      `json:"skill_plugin,omitempty"`
}

// PortalsListingsUpdateOutputGroups - Catalog section the listing appears in.
type PortalsListingsUpdateOutputGroups struct {
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Index - Order of the group in the portal catalog.
	Index float64 `json:"index"`
}

// PortalsListingsUpdateOutput represents the portals listings update output type.
type PortalsListingsUpdateOutput struct {
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
	Access PortalsListingsUpdateOutputAccess `json:"access"`
	// Groups - Catalog sections the listing appears in, in display order.
	Groups []PortalsListingsUpdateOutputGroups `json:"groups"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapPortalsListingsUpdateOutputFromJSON deserializes JSON data into a PortalsListingsUpdateOutput.
func MapPortalsListingsUpdateOutputFromJSON(data []byte) (*PortalsListingsUpdateOutput, error) {
	var v PortalsListingsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsListingsUpdateOutputToJSON serializes a PortalsListingsUpdateOutput to JSON.
func MapPortalsListingsUpdateOutputToJSON(v *PortalsListingsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsListingsUpdateBody represents the portals listings update body type.
type PortalsListingsUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Readme - Longer guide shown to consumers.
	Readme *string `json:"readme,omitempty"`
	// IsDiscoverable - Whether every consumer on the portal can find this listing in the catalog and request access. When false, only consumers who already have access through one of their groups see it.
	IsDiscoverable *bool `json:"is_discoverable,omitempty"`
}

// MapPortalsListingsUpdateBodyFromJSON deserializes JSON data into a PortalsListingsUpdateBody.
func MapPortalsListingsUpdateBodyFromJSON(data []byte) (*PortalsListingsUpdateBody, error) {
	var v PortalsListingsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsListingsUpdateBodyToJSON serializes a PortalsListingsUpdateBody to JSON.
func MapPortalsListingsUpdateBodyToJSON(v *PortalsListingsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
