package listings

import (
	"encoding/json"
	"time"
)

// PortalsListingsListOutputItemsAccessProviderTemplate represents the portals listings list output items access provider template type.
type PortalsListingsListOutputItemsAccessProviderTemplate struct {
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

// PortalsListingsListOutputItemsAccessMagicMcpServer represents the portals listings list output items access magic mcp server type.
type PortalsListingsListOutputItemsAccessMagicMcpServer struct {
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

// PortalsListingsListOutputItemsAccessSkill represents the portals listings list output items access skill type.
type PortalsListingsListOutputItemsAccessSkill struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

// PortalsListingsListOutputItemsAccessSkillTemplate represents the portals listings list output items access skill template type.
type PortalsListingsListOutputItemsAccessSkillTemplate struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Owner       string  `json:"owner"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// PortalsListingsListOutputItemsAccessSkillGroup represents the portals listings list output items access skill group type.
type PortalsListingsListOutputItemsAccessSkillGroup struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// PortalsListingsListOutputItemsAccessSkillMarketplace represents the portals listings list output items access skill marketplace type.
type PortalsListingsListOutputItemsAccessSkillMarketplace struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
}

// PortalsListingsListOutputItemsAccessSkillPlugin represents the portals listings list output items access skill plugin type.
type PortalsListingsListOutputItemsAccessSkillPlugin struct {
	Object string  `json:"object"`
	Id     string  `json:"id"`
	Status string  `json:"status"`
	Name   *string `json:"name,omitempty"`
}

// PortalsListingsListOutputItemsAccess represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type PortalsListingsListOutputItemsAccess struct {
	Type             *string                                               `json:"type,omitempty"`
	ProviderTemplate *PortalsListingsListOutputItemsAccessProviderTemplate `json:"provider_template,omitempty"`
	MagicMcpServer   *PortalsListingsListOutputItemsAccessMagicMcpServer   `json:"magic_mcp_server,omitempty"`
	Skill            *PortalsListingsListOutputItemsAccessSkill            `json:"skill,omitempty"`
	SkillTemplate    *PortalsListingsListOutputItemsAccessSkillTemplate    `json:"skill_template,omitempty"`
	SkillGroup       *PortalsListingsListOutputItemsAccessSkillGroup       `json:"skill_group,omitempty"`
	SkillMarketplace *PortalsListingsListOutputItemsAccessSkillMarketplace `json:"skill_marketplace,omitempty"`
	SkillPlugin      *PortalsListingsListOutputItemsAccessSkillPlugin      `json:"skill_plugin,omitempty"`
}

// PortalsListingsListOutputItemsGroups - Catalog section the listing appears in.
type PortalsListingsListOutputItemsGroups struct {
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Index - Order of the group in the portal catalog.
	Index float64 `json:"index"`
}

// PortalsListingsListOutputItems represents the portals listings list output items type.
type PortalsListingsListOutputItems struct {
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
	Access PortalsListingsListOutputItemsAccess `json:"access"`
	// Groups - Catalog sections the listing appears in, in display order.
	Groups []PortalsListingsListOutputItemsGroups `json:"groups"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// PortalsListingsListOutputPagination represents the portals listings list output pagination type.
type PortalsListingsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// PortalsListingsListOutput represents the portals listings list output type.
type PortalsListingsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                              `json:"object"`
	Items      []PortalsListingsListOutputItems    `json:"items"`
	Pagination PortalsListingsListOutputPagination `json:"pagination"`
}

// MapPortalsListingsListOutputFromJSON deserializes JSON data into a PortalsListingsListOutput.
func MapPortalsListingsListOutputFromJSON(data []byte) (*PortalsListingsListOutput, error) {
	var v PortalsListingsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsListingsListOutputToJSON serializes a PortalsListingsListOutput to JSON.
func MapPortalsListingsListOutputToJSON(v *PortalsListingsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsListingsListQuery represents the portals listings list query type.
type PortalsListingsListQuery struct {
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
	// ConsumerSurfaceProviderGroupId - Only return listings in these provider groups.
	ConsumerSurfaceProviderGroupId *any `json:"consumer_surface_provider_group_id,omitempty"`
	// ProviderTemplateId - Only return listings for these provider templates.
	ProviderTemplateId *any `json:"provider_template_id,omitempty"`
	// MagicMcpServerId - Only return listings for these magic MCP servers.
	MagicMcpServerId *any `json:"magic_mcp_server_id,omitempty"`
	// SkillId - Only return listings for these skills.
	SkillId *any `json:"skill_id,omitempty"`
	// SkillTemplateId - Only return listings for these skill templates.
	SkillTemplateId *any `json:"skill_template_id,omitempty"`
	// SkillGroupId - Only return listings for these skill groups.
	SkillGroupId *any `json:"skill_group_id,omitempty"`
	// SkillMarketplaceId - Only return listings for these skill marketplaces.
	SkillMarketplaceId *any `json:"skill_marketplace_id,omitempty"`
	// Type - Only return listings of these types.
	Type *any `json:"type,omitempty"`
	// HasAccess - Only return listings that have at least one consumer access grant (true) or none (false).
	HasAccess *bool `json:"has_access,omitempty"`
}

// MapPortalsListingsListQueryFromJSON deserializes JSON data into a PortalsListingsListQuery.
func MapPortalsListingsListQueryFromJSON(data []byte) (*PortalsListingsListQuery, error) {
	var v PortalsListingsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsListingsListQueryToJSON serializes a PortalsListingsListQuery to JSON.
func MapPortalsListingsListQueryToJSON(v *PortalsListingsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
