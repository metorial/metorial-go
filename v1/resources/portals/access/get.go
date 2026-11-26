package access

import (
	"encoding/json"
	"time"
)

// PortalsAccessGetOutputListing - Shared listing this rule was created from, if any.
type PortalsAccessGetOutputListing struct {
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// IsDiscoverable - Whether every consumer on the portal can find the listing in the catalog and request access.
	IsDiscoverable bool `json:"is_discoverable"`
}

// PortalsAccessGetOutputAccessProviderTemplate represents the portals access get output access provider template type.
type PortalsAccessGetOutputAccessProviderTemplate struct {
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

// PortalsAccessGetOutputAccessMagicMcpServer represents the portals access get output access magic mcp server type.
type PortalsAccessGetOutputAccessMagicMcpServer struct {
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

// PortalsAccessGetOutputAccessSkill represents the portals access get output access skill type.
type PortalsAccessGetOutputAccessSkill struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

// PortalsAccessGetOutputAccessSkillTemplate represents the portals access get output access skill template type.
type PortalsAccessGetOutputAccessSkillTemplate struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Owner       string  `json:"owner"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// PortalsAccessGetOutputAccessSkillGroup represents the portals access get output access skill group type.
type PortalsAccessGetOutputAccessSkillGroup struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// PortalsAccessGetOutputAccessSkillMarketplace represents the portals access get output access skill marketplace type.
type PortalsAccessGetOutputAccessSkillMarketplace struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
}

// PortalsAccessGetOutputAccessSkillPlugin represents the portals access get output access skill plugin type.
type PortalsAccessGetOutputAccessSkillPlugin struct {
	Object string  `json:"object"`
	Id     string  `json:"id"`
	Status string  `json:"status"`
	Name   *string `json:"name,omitempty"`
}

// PortalsAccessGetOutputAccess represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type PortalsAccessGetOutputAccess struct {
	Type             *string                                       `json:"type,omitempty"`
	ProviderTemplate *PortalsAccessGetOutputAccessProviderTemplate `json:"provider_template,omitempty"`
	MagicMcpServer   *PortalsAccessGetOutputAccessMagicMcpServer   `json:"magic_mcp_server,omitempty"`
	Skill            *PortalsAccessGetOutputAccessSkill            `json:"skill,omitempty"`
	SkillTemplate    *PortalsAccessGetOutputAccessSkillTemplate    `json:"skill_template,omitempty"`
	SkillGroup       *PortalsAccessGetOutputAccessSkillGroup       `json:"skill_group,omitempty"`
	SkillMarketplace *PortalsAccessGetOutputAccessSkillMarketplace `json:"skill_marketplace,omitempty"`
	SkillPlugin      *PortalsAccessGetOutputAccessSkillPlugin      `json:"skill_plugin,omitempty"`
}

// PortalsAccessGetOutputConsumerGroup represents the portals access get output consumer group type.
type PortalsAccessGetOutputConsumerGroup struct {
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

// PortalsAccessGetOutput represents the portals access get output type.
type PortalsAccessGetOutput struct {
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
	Listing *PortalsAccessGetOutputListing `json:"listing,omitempty"`
	// Access - The integration, skill, or skill set the group can access.
	Access        PortalsAccessGetOutputAccess        `json:"access"`
	ConsumerGroup PortalsAccessGetOutputConsumerGroup `json:"consumer_group"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapPortalsAccessGetOutputFromJSON deserializes JSON data into a PortalsAccessGetOutput.
func MapPortalsAccessGetOutputFromJSON(data []byte) (*PortalsAccessGetOutput, error) {
	var v PortalsAccessGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsAccessGetOutputToJSON serializes a PortalsAccessGetOutput to JSON.
func MapPortalsAccessGetOutputToJSON(v *PortalsAccessGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
