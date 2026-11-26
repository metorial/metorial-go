package access

import (
	"encoding/json"
	"time"
)

// PortalsAccessCreateOutputListing - Shared listing this rule was created from, if any.
type PortalsAccessCreateOutputListing struct {
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// IsDiscoverable - Whether every consumer on the portal can find the listing in the catalog and request access.
	IsDiscoverable bool `json:"is_discoverable"`
}

// PortalsAccessCreateOutputAccessProviderTemplate represents the portals access create output access provider template type.
type PortalsAccessCreateOutputAccessProviderTemplate struct {
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

// PortalsAccessCreateOutputAccessMagicMcpServer represents the portals access create output access magic mcp server type.
type PortalsAccessCreateOutputAccessMagicMcpServer struct {
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

// PortalsAccessCreateOutputAccessSkill represents the portals access create output access skill type.
type PortalsAccessCreateOutputAccessSkill struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

// PortalsAccessCreateOutputAccessSkillTemplate represents the portals access create output access skill template type.
type PortalsAccessCreateOutputAccessSkillTemplate struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Owner       string  `json:"owner"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// PortalsAccessCreateOutputAccessSkillGroup represents the portals access create output access skill group type.
type PortalsAccessCreateOutputAccessSkillGroup struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// PortalsAccessCreateOutputAccessSkillMarketplace represents the portals access create output access skill marketplace type.
type PortalsAccessCreateOutputAccessSkillMarketplace struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
}

// PortalsAccessCreateOutputAccessSkillPlugin represents the portals access create output access skill plugin type.
type PortalsAccessCreateOutputAccessSkillPlugin struct {
	Object string  `json:"object"`
	Id     string  `json:"id"`
	Status string  `json:"status"`
	Name   *string `json:"name,omitempty"`
}

// PortalsAccessCreateOutputAccess represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type PortalsAccessCreateOutputAccess struct {
	Type             *string                                          `json:"type,omitempty"`
	ProviderTemplate *PortalsAccessCreateOutputAccessProviderTemplate `json:"provider_template,omitempty"`
	MagicMcpServer   *PortalsAccessCreateOutputAccessMagicMcpServer   `json:"magic_mcp_server,omitempty"`
	Skill            *PortalsAccessCreateOutputAccessSkill            `json:"skill,omitempty"`
	SkillTemplate    *PortalsAccessCreateOutputAccessSkillTemplate    `json:"skill_template,omitempty"`
	SkillGroup       *PortalsAccessCreateOutputAccessSkillGroup       `json:"skill_group,omitempty"`
	SkillMarketplace *PortalsAccessCreateOutputAccessSkillMarketplace `json:"skill_marketplace,omitempty"`
	SkillPlugin      *PortalsAccessCreateOutputAccessSkillPlugin      `json:"skill_plugin,omitempty"`
}

// PortalsAccessCreateOutputConsumerGroup represents the portals access create output consumer group type.
type PortalsAccessCreateOutputConsumerGroup struct {
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

// PortalsAccessCreateOutput represents the portals access create output type.
type PortalsAccessCreateOutput struct {
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
	Listing *PortalsAccessCreateOutputListing `json:"listing,omitempty"`
	// Access - The integration, skill, or skill set the group can access.
	Access        PortalsAccessCreateOutputAccess        `json:"access"`
	ConsumerGroup PortalsAccessCreateOutputConsumerGroup `json:"consumer_group"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapPortalsAccessCreateOutputFromJSON deserializes JSON data into a PortalsAccessCreateOutput.
func MapPortalsAccessCreateOutputFromJSON(data []byte) (*PortalsAccessCreateOutput, error) {
	var v PortalsAccessCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsAccessCreateOutputToJSON serializes a PortalsAccessCreateOutput to JSON.
func MapPortalsAccessCreateOutputToJSON(v *PortalsAccessCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsAccessCreateBodyAccess represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type PortalsAccessCreateBodyAccess struct {
	Type *string `json:"type,omitempty"`
	// ProviderTemplateId - Provider template the rule grants access to.
	ProviderTemplateId *string `json:"provider_template_id,omitempty"`
	// MagicMcpServerId - Magic MCP server the rule grants access to.
	MagicMcpServerId *string `json:"magic_mcp_server_id,omitempty"`
	// SkillId - Skill the rule grants access to.
	SkillId *string `json:"skill_id,omitempty"`
	// Permission - What the group can do with the skill marketplace.
	Permission *string `json:"permission,omitempty"`
	// SkillTemplateId - Skill template the rule grants access to.
	SkillTemplateId *string `json:"skill_template_id,omitempty"`
	// SkillGroupId - Skill group the rule grants access to.
	SkillGroupId *string `json:"skill_group_id,omitempty"`
	// SkillMarketplaceId - Skill marketplace the plugin is published in.
	SkillMarketplaceId *string `json:"skill_marketplace_id,omitempty"`
	// SkillPluginId - Skill plugin the rule grants access to.
	SkillPluginId *string `json:"skill_plugin_id,omitempty"`
}

// PortalsAccessCreateBody represents the portals access create body type.
type PortalsAccessCreateBody struct {
	// ConsumerGroupId - Consumer group the rule grants access to.
	ConsumerGroupId string `json:"consumer_group_id"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Readme - Longer guide shown to consumers.
	Readme *string `json:"readme,omitempty"`
	// Access - The integration, skill, or skill set the rule grants access to.
	Access PortalsAccessCreateBodyAccess `json:"access"`
}

// MapPortalsAccessCreateBodyFromJSON deserializes JSON data into a PortalsAccessCreateBody.
func MapPortalsAccessCreateBodyFromJSON(data []byte) (*PortalsAccessCreateBody, error) {
	var v PortalsAccessCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsAccessCreateBodyToJSON serializes a PortalsAccessCreateBody to JSON.
func MapPortalsAccessCreateBodyToJSON(v *PortalsAccessCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
