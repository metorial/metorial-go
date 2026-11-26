package access

import (
	"encoding/json"
	"time"
)

// PortalsAccessDeleteOutputListing - Shared listing this rule was created from, if any.
type PortalsAccessDeleteOutputListing struct {
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// IsDiscoverable - Whether every consumer on the portal can find the listing in the catalog and request access.
	IsDiscoverable bool `json:"is_discoverable"`
}

// PortalsAccessDeleteOutputAccessProviderTemplate represents the portals access delete output access provider template type.
type PortalsAccessDeleteOutputAccessProviderTemplate struct {
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

// PortalsAccessDeleteOutputAccessMagicMcpServer represents the portals access delete output access magic mcp server type.
type PortalsAccessDeleteOutputAccessMagicMcpServer struct {
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

// PortalsAccessDeleteOutputAccessSkill represents the portals access delete output access skill type.
type PortalsAccessDeleteOutputAccessSkill struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

// PortalsAccessDeleteOutputAccessSkillTemplate represents the portals access delete output access skill template type.
type PortalsAccessDeleteOutputAccessSkillTemplate struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Owner       string  `json:"owner"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// PortalsAccessDeleteOutputAccessSkillGroup represents the portals access delete output access skill group type.
type PortalsAccessDeleteOutputAccessSkillGroup struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// PortalsAccessDeleteOutputAccessSkillMarketplace represents the portals access delete output access skill marketplace type.
type PortalsAccessDeleteOutputAccessSkillMarketplace struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
}

// PortalsAccessDeleteOutputAccessSkillPlugin represents the portals access delete output access skill plugin type.
type PortalsAccessDeleteOutputAccessSkillPlugin struct {
	Object string  `json:"object"`
	Id     string  `json:"id"`
	Status string  `json:"status"`
	Name   *string `json:"name,omitempty"`
}

// PortalsAccessDeleteOutputAccess represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type PortalsAccessDeleteOutputAccess struct {
	Type             *string                                          `json:"type,omitempty"`
	ProviderTemplate *PortalsAccessDeleteOutputAccessProviderTemplate `json:"provider_template,omitempty"`
	MagicMcpServer   *PortalsAccessDeleteOutputAccessMagicMcpServer   `json:"magic_mcp_server,omitempty"`
	Skill            *PortalsAccessDeleteOutputAccessSkill            `json:"skill,omitempty"`
	SkillTemplate    *PortalsAccessDeleteOutputAccessSkillTemplate    `json:"skill_template,omitempty"`
	SkillGroup       *PortalsAccessDeleteOutputAccessSkillGroup       `json:"skill_group,omitempty"`
	SkillMarketplace *PortalsAccessDeleteOutputAccessSkillMarketplace `json:"skill_marketplace,omitempty"`
	SkillPlugin      *PortalsAccessDeleteOutputAccessSkillPlugin      `json:"skill_plugin,omitempty"`
}

// PortalsAccessDeleteOutputConsumerGroup represents the portals access delete output consumer group type.
type PortalsAccessDeleteOutputConsumerGroup struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once you delete the group. Archived groups stay listable.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// IsDefault - Whether this is a default group on the portal.
	IsDefault bool `json:"is_default"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// PortalsAccessDeleteOutput represents the portals access delete output type.
type PortalsAccessDeleteOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// AccessLevel - What the consumer group can do with the resource. Skills and skill groups support read, write and manage (share with others); skill marketplaces support read and manage.
	AccessLevel *string `json:"access_level,omitempty"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Readme - Longer guide shown to consumers.
	Readme *string `json:"readme,omitempty"`
	// Listing - Shared listing this rule was created from, if any.
	Listing *PortalsAccessDeleteOutputListing `json:"listing,omitempty"`
	// Access - The integration, skill, or skill set the group can access.
	Access        PortalsAccessDeleteOutputAccess        `json:"access"`
	ConsumerGroup PortalsAccessDeleteOutputConsumerGroup `json:"consumer_group"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapPortalsAccessDeleteOutputFromJSON deserializes JSON data into a PortalsAccessDeleteOutput.
func MapPortalsAccessDeleteOutputFromJSON(data []byte) (*PortalsAccessDeleteOutput, error) {
	var v PortalsAccessDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsAccessDeleteOutputToJSON serializes a PortalsAccessDeleteOutput to JSON.
func MapPortalsAccessDeleteOutputToJSON(v *PortalsAccessDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
