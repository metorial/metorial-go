package access

import (
	"encoding/json"
	"time"
)

// PortalsAccessUpdateOutputListing - Shared listing this rule was created from, if any.
type PortalsAccessUpdateOutputListing struct {
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// IsDiscoverable - Whether every consumer on the portal can find the listing in the catalog and request access.
	IsDiscoverable bool `json:"is_discoverable"`
}

// PortalsAccessUpdateOutputAccessProviderTemplate represents the portals access update output access provider template type.
type PortalsAccessUpdateOutputAccessProviderTemplate struct {
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

// PortalsAccessUpdateOutputAccessMagicMcpServer represents the portals access update output access magic mcp server type.
type PortalsAccessUpdateOutputAccessMagicMcpServer struct {
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

// PortalsAccessUpdateOutputAccessSkill represents the portals access update output access skill type.
type PortalsAccessUpdateOutputAccessSkill struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

// PortalsAccessUpdateOutputAccessSkillTemplate represents the portals access update output access skill template type.
type PortalsAccessUpdateOutputAccessSkillTemplate struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Owner       string  `json:"owner"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// PortalsAccessUpdateOutputAccessSkillGroup represents the portals access update output access skill group type.
type PortalsAccessUpdateOutputAccessSkillGroup struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// PortalsAccessUpdateOutputAccessSkillMarketplace represents the portals access update output access skill marketplace type.
type PortalsAccessUpdateOutputAccessSkillMarketplace struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
}

// PortalsAccessUpdateOutputAccessSkillPlugin represents the portals access update output access skill plugin type.
type PortalsAccessUpdateOutputAccessSkillPlugin struct {
	Object string  `json:"object"`
	Id     string  `json:"id"`
	Status string  `json:"status"`
	Name   *string `json:"name,omitempty"`
}

// PortalsAccessUpdateOutputAccess represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type PortalsAccessUpdateOutputAccess struct {
	Type             *string                                          `json:"type,omitempty"`
	ProviderTemplate *PortalsAccessUpdateOutputAccessProviderTemplate `json:"provider_template,omitempty"`
	MagicMcpServer   *PortalsAccessUpdateOutputAccessMagicMcpServer   `json:"magic_mcp_server,omitempty"`
	Skill            *PortalsAccessUpdateOutputAccessSkill            `json:"skill,omitempty"`
	SkillTemplate    *PortalsAccessUpdateOutputAccessSkillTemplate    `json:"skill_template,omitempty"`
	SkillGroup       *PortalsAccessUpdateOutputAccessSkillGroup       `json:"skill_group,omitempty"`
	SkillMarketplace *PortalsAccessUpdateOutputAccessSkillMarketplace `json:"skill_marketplace,omitempty"`
	SkillPlugin      *PortalsAccessUpdateOutputAccessSkillPlugin      `json:"skill_plugin,omitempty"`
}

// PortalsAccessUpdateOutputConsumerGroup represents the portals access update output consumer group type.
type PortalsAccessUpdateOutputConsumerGroup struct {
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

// PortalsAccessUpdateOutput represents the portals access update output type.
type PortalsAccessUpdateOutput struct {
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
	Listing *PortalsAccessUpdateOutputListing `json:"listing,omitempty"`
	// Access - The integration, skill, or skill set the group can access.
	Access        PortalsAccessUpdateOutputAccess        `json:"access"`
	ConsumerGroup PortalsAccessUpdateOutputConsumerGroup `json:"consumer_group"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapPortalsAccessUpdateOutputFromJSON deserializes JSON data into a PortalsAccessUpdateOutput.
func MapPortalsAccessUpdateOutputFromJSON(data []byte) (*PortalsAccessUpdateOutput, error) {
	var v PortalsAccessUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsAccessUpdateOutputToJSON serializes a PortalsAccessUpdateOutput to JSON.
func MapPortalsAccessUpdateOutputToJSON(v *PortalsAccessUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsAccessUpdateBody represents the portals access update body type.
type PortalsAccessUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Readme - Longer guide shown to consumers.
	Readme *string `json:"readme,omitempty"`
	// Permission - What the group can do with the skill or skill group.
	Permission *string `json:"permission,omitempty"`
}

// MapPortalsAccessUpdateBodyFromJSON deserializes JSON data into a PortalsAccessUpdateBody.
func MapPortalsAccessUpdateBodyFromJSON(data []byte) (*PortalsAccessUpdateBody, error) {
	var v PortalsAccessUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsAccessUpdateBodyToJSON serializes a PortalsAccessUpdateBody to JSON.
func MapPortalsAccessUpdateBodyToJSON(v *PortalsAccessUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
