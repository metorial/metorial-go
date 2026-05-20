package listings

import (
	"encoding/json"
	"time"
)

// PortalsListingsListOutputItemsAccessProviderTemplate represents the portals listings list output items access provider template type.
type PortalsListingsListOutputItemsAccessProviderTemplate struct {
	Object        string         `json:"object"`
	Id            string         `json:"id"`
	Status        string         `json:"status"`
	Name          string         `json:"name"`
	Description   *string        `json:"description,omitempty"`
	Metadata      map[string]any `json:"metadata"`
	IntegrationId *string        `json:"integration_id,omitempty"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

// PortalsListingsListOutputItemsAccessMagicMcpServer represents the portals listings list output items access magic mcp server type.
type PortalsListingsListOutputItemsAccessMagicMcpServer struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Name        *string `json:"name,omitempty"`
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
}

// PortalsListingsListOutputItemsGroups represents the portals listings list output items groups type.
type PortalsListingsListOutputItemsGroups struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Index       float64 `json:"index"`
}

// PortalsListingsListOutputItems represents the portals listings list output items type.
type PortalsListingsListOutputItems struct {
	Object      string                                 `json:"object"`
	Id          string                                 `json:"id"`
	Name        string                                 `json:"name"`
	Description *string                                `json:"description,omitempty"`
	Readme      *string                                `json:"readme,omitempty"`
	Access      PortalsListingsListOutputItemsAccess   `json:"access"`
	Groups      []PortalsListingsListOutputItemsGroups `json:"groups"`
	CreatedAt   time.Time                              `json:"created_at"`
	UpdatedAt   time.Time                              `json:"updated_at"`
}

// PortalsListingsListOutputPagination represents the portals listings list output pagination type.
type PortalsListingsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// PortalsListingsListOutput represents the portals listings list output type.
type PortalsListingsListOutput struct {
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
	Limit                          *float64 `json:"limit,omitempty"`
	After                          *string  `json:"after,omitempty"`
	Before                         *string  `json:"before,omitempty"`
	Cursor                         *string  `json:"cursor,omitempty"`
	Order                          *string  `json:"order,omitempty"`
	Search                         *string  `json:"search,omitempty"`
	ConsumerSurfaceProviderGroupId *any     `json:"consumer_surface_provider_group_id,omitempty"`
	ProviderTemplateId             *any     `json:"provider_template_id,omitempty"`
	MagicMcpServerId               *any     `json:"magic_mcp_server_id,omitempty"`
	SkillId                        *any     `json:"skill_id,omitempty"`
	SkillTemplateId                *any     `json:"skill_template_id,omitempty"`
	SkillGroupId                   *any     `json:"skill_group_id,omitempty"`
	SkillMarketplaceId             *any     `json:"skill_marketplace_id,omitempty"`
	Type                           *any     `json:"type,omitempty"`
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
