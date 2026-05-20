package access

import (
	"encoding/json"
	"time"
)

// PortalsAccessListOutputItemsAccessProviderTemplate represents the portals access list output items access provider template type.
type PortalsAccessListOutputItemsAccessProviderTemplate struct {
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

// PortalsAccessListOutputItemsAccessMagicMcpServer represents the portals access list output items access magic mcp server type.
type PortalsAccessListOutputItemsAccessMagicMcpServer struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Name        *string `json:"name,omitempty"`
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
}

// PortalsAccessListOutputItemsConsumerGroup represents the portals access list output items consumer group type.
type PortalsAccessListOutputItemsConsumerGroup struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Status      string    `json:"status"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	IsDefault   bool      `json:"is_default"`
	SsoGroupIds []string  `json:"sso_group_ids"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// PortalsAccessListOutputItems represents the portals access list output items type.
type PortalsAccessListOutputItems struct {
	Object        string                                    `json:"object"`
	Id            string                                    `json:"id"`
	Name          string                                    `json:"name"`
	Description   *string                                   `json:"description,omitempty"`
	Readme        *string                                   `json:"readme,omitempty"`
	Access        PortalsAccessListOutputItemsAccess        `json:"access"`
	ConsumerGroup PortalsAccessListOutputItemsConsumerGroup `json:"consumer_group"`
	CreatedAt     time.Time                                 `json:"created_at"`
	UpdatedAt     time.Time                                 `json:"updated_at"`
}

// PortalsAccessListOutputPagination represents the portals access list output pagination type.
type PortalsAccessListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// PortalsAccessListOutput represents the portals access list output type.
type PortalsAccessListOutput struct {
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
	Limit                   *float64 `json:"limit,omitempty"`
	After                   *string  `json:"after,omitempty"`
	Before                  *string  `json:"before,omitempty"`
	Cursor                  *string  `json:"cursor,omitempty"`
	Order                   *string  `json:"order,omitempty"`
	Search                  *string  `json:"search,omitempty"`
	ConsumerGroupId         *any     `json:"consumer_group_id,omitempty"`
	ProviderTemplateId      *any     `json:"provider_template_id,omitempty"`
	MagicMcpServerId        *any     `json:"magic_mcp_server_id,omitempty"`
	SkillId                 *any     `json:"skill_id,omitempty"`
	SkillTemplateId         *any     `json:"skill_template_id,omitempty"`
	SkillGroupId            *any     `json:"skill_group_id,omitempty"`
	SkillMarketplaceId      *any     `json:"skill_marketplace_id,omitempty"`
	ConsumerAccessListingId *any     `json:"consumer_access_listing_id,omitempty"`
	Type                    *any     `json:"type,omitempty"`
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
