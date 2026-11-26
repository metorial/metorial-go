package access

import (
	"encoding/json"
	"time"
)

// PortalsAccessListOutputItemsListing - Shared listing this rule was created from, if any.
type PortalsAccessListOutputItemsListing struct {
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// IsDiscoverable - Whether every consumer on the portal can find the listing in the catalog and request access.
	IsDiscoverable bool `json:"is_discoverable"`
}

// PortalsAccessListOutputItemsAccessProviderTemplate represents the portals access list output items access provider template type.
type PortalsAccessListOutputItemsAccessProviderTemplate struct {
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

// PortalsAccessListOutputItemsAccessMagicMcpServer represents the portals access list output items access magic mcp server type.
type PortalsAccessListOutputItemsAccessMagicMcpServer struct {
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

// PortalsAccessListOutputItemsAccessSkill represents the portals access list output items access skill type.
type PortalsAccessListOutputItemsAccessSkill struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

// PortalsAccessListOutputItemsAccessSkillTemplate represents the portals access list output items access skill template type.
type PortalsAccessListOutputItemsAccessSkillTemplate struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Owner       string  `json:"owner"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// PortalsAccessListOutputItemsAccessSkillGroup represents the portals access list output items access skill group type.
type PortalsAccessListOutputItemsAccessSkillGroup struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// PortalsAccessListOutputItemsAccessSkillMarketplace represents the portals access list output items access skill marketplace type.
type PortalsAccessListOutputItemsAccessSkillMarketplace struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
}

// PortalsAccessListOutputItemsAccessSkillPlugin represents the portals access list output items access skill plugin type.
type PortalsAccessListOutputItemsAccessSkillPlugin struct {
	Object string  `json:"object"`
	Id     string  `json:"id"`
	Status string  `json:"status"`
	Name   *string `json:"name,omitempty"`
}

// PortalsAccessListOutputItemsAccess represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type PortalsAccessListOutputItemsAccess struct {
	Type             *string                                             `json:"type,omitempty"`
	ProviderTemplate *PortalsAccessListOutputItemsAccessProviderTemplate `json:"provider_template,omitempty"`
	MagicMcpServer   *PortalsAccessListOutputItemsAccessMagicMcpServer   `json:"magic_mcp_server,omitempty"`
	Skill            *PortalsAccessListOutputItemsAccessSkill            `json:"skill,omitempty"`
	SkillTemplate    *PortalsAccessListOutputItemsAccessSkillTemplate    `json:"skill_template,omitempty"`
	SkillGroup       *PortalsAccessListOutputItemsAccessSkillGroup       `json:"skill_group,omitempty"`
	SkillMarketplace *PortalsAccessListOutputItemsAccessSkillMarketplace `json:"skill_marketplace,omitempty"`
	SkillPlugin      *PortalsAccessListOutputItemsAccessSkillPlugin      `json:"skill_plugin,omitempty"`
}

// PortalsAccessListOutputItemsConsumerGroup represents the portals access list output items consumer group type.
type PortalsAccessListOutputItemsConsumerGroup struct {
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

// PortalsAccessListOutputItems represents the portals access list output items type.
type PortalsAccessListOutputItems struct {
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
	Listing *PortalsAccessListOutputItemsListing `json:"listing,omitempty"`
	// Access - The integration, skill, or skill set the group can access.
	Access        PortalsAccessListOutputItemsAccess        `json:"access"`
	ConsumerGroup PortalsAccessListOutputItemsConsumerGroup `json:"consumer_group"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// PortalsAccessListOutputPagination represents the portals access list output pagination type.
type PortalsAccessListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// PortalsAccessListOutput represents the portals access list output type.
type PortalsAccessListOutput struct {
	// Object - String representing the object’s type.
	Object     string                            `json:"object"`
	Items      []PortalsAccessListOutputItems    `json:"items"`
	Pagination PortalsAccessListOutputPagination `json:"pagination"`
}

// MapPortalsAccessListOutputFromJSON deserializes JSON data into a PortalsAccessListOutput.
func MapPortalsAccessListOutputFromJSON(data []byte) (*PortalsAccessListOutput, error) {
	var v PortalsAccessListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsAccessListOutputToJSON serializes a PortalsAccessListOutput to JSON.
func MapPortalsAccessListOutputToJSON(v *PortalsAccessListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsAccessListQuery represents the portals access list query type.
type PortalsAccessListQuery struct {
	// Limit - Page size, from 1 to 100. Defaults to 20.
	Limit *float64 `json:"limit,omitempty"`
	// After - Return the page after this point. Pass `pagination.after` if the previous response has it, otherwise the `id` of its last item.
	After *string `json:"after,omitempty"`
	// Before - Return the page before this point. Pass `pagination.before` if the previous response has it, otherwise the `id` of its first item.
	Before *string `json:"before,omitempty"`
	// Cursor - A `cur_` cursor from `pagination.after` or `pagination.before`. Alternative to `after` and `before`.
	Cursor *string `json:"cursor,omitempty"`
	// Order - Sort direction.
	Order *string `json:"order,omitempty"`
	// Search - Search by name or description.
	Search *string `json:"search,omitempty"`
	// ConsumerGroupId - Only return rules for these consumer groups.
	ConsumerGroupId *any `json:"consumer_group_id,omitempty"`
	// ProviderTemplateId - Only return rules for these provider templates.
	ProviderTemplateId *any `json:"provider_template_id,omitempty"`
	// MagicMcpServerId - Only return rules for these magic MCP servers.
	MagicMcpServerId *any `json:"magic_mcp_server_id,omitempty"`
	// SkillId - Only return rules for these skills.
	SkillId *any `json:"skill_id,omitempty"`
	// SkillTemplateId - Only return rules for these skill templates.
	SkillTemplateId *any `json:"skill_template_id,omitempty"`
	// SkillGroupId - Only return rules for these skill groups.
	SkillGroupId *any `json:"skill_group_id,omitempty"`
	// SkillMarketplaceId - Only return rules for these skill marketplaces.
	SkillMarketplaceId *any `json:"skill_marketplace_id,omitempty"`
	// SkillPluginId - Only return rules for these skill plugins.
	SkillPluginId *any `json:"skill_plugin_id,omitempty"`
	// ConsumerAccessListingId - Only return rules attached to these listings.
	ConsumerAccessListingId *any `json:"consumer_access_listing_id,omitempty"`
	// Type - Only return rules of these types.
	Type *any `json:"type,omitempty"`
}

// MapPortalsAccessListQueryFromJSON deserializes JSON data into a PortalsAccessListQuery.
func MapPortalsAccessListQueryFromJSON(data []byte) (*PortalsAccessListQuery, error) {
	var v PortalsAccessListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsAccessListQueryToJSON serializes a PortalsAccessListQuery to JSON.
func MapPortalsAccessListQueryToJSON(v *PortalsAccessListQuery) ([]byte, error) {
	return json.Marshal(v)
}
