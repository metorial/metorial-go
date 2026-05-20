package listings

import (
	"encoding/json"
	"time"
)

// PortalsListingsCreateOutputAccessProviderTemplate represents the portals listings create output access provider template type.
type PortalsListingsCreateOutputAccessProviderTemplate struct {
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

// PortalsListingsCreateOutputAccessMagicMcpServer represents the portals listings create output access magic mcp server type.
type PortalsListingsCreateOutputAccessMagicMcpServer struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// PortalsListingsCreateOutputAccessSkill represents the portals listings create output access skill type.
type PortalsListingsCreateOutputAccessSkill struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

// PortalsListingsCreateOutputAccessSkillTemplate represents the portals listings create output access skill template type.
type PortalsListingsCreateOutputAccessSkillTemplate struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Owner       string  `json:"owner"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// PortalsListingsCreateOutputAccessSkillGroup represents the portals listings create output access skill group type.
type PortalsListingsCreateOutputAccessSkillGroup struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// PortalsListingsCreateOutputAccessSkillMarketplace represents the portals listings create output access skill marketplace type.
type PortalsListingsCreateOutputAccessSkillMarketplace struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
}

// PortalsListingsCreateOutputAccess represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type PortalsListingsCreateOutputAccess struct {
	Type             *string                                            `json:"type,omitempty"`
	ProviderTemplate *PortalsListingsCreateOutputAccessProviderTemplate `json:"provider_template,omitempty"`
	MagicMcpServer   *PortalsListingsCreateOutputAccessMagicMcpServer   `json:"magic_mcp_server,omitempty"`
	Skill            *PortalsListingsCreateOutputAccessSkill            `json:"skill,omitempty"`
	SkillTemplate    *PortalsListingsCreateOutputAccessSkillTemplate    `json:"skill_template,omitempty"`
	SkillGroup       *PortalsListingsCreateOutputAccessSkillGroup       `json:"skill_group,omitempty"`
	SkillMarketplace *PortalsListingsCreateOutputAccessSkillMarketplace `json:"skill_marketplace,omitempty"`
}

// PortalsListingsCreateOutputGroups represents the portals listings create output groups type.
type PortalsListingsCreateOutputGroups struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Index       float64 `json:"index"`
}

// PortalsListingsCreateOutput represents the portals listings create output type.
type PortalsListingsCreateOutput struct {
	Object      string                              `json:"object"`
	Id          string                              `json:"id"`
	Name        string                              `json:"name"`
	Description *string                             `json:"description,omitempty"`
	Readme      *string                             `json:"readme,omitempty"`
	Access      PortalsListingsCreateOutputAccess   `json:"access"`
	Groups      []PortalsListingsCreateOutputGroups `json:"groups"`
	CreatedAt   time.Time                           `json:"created_at"`
	UpdatedAt   time.Time                           `json:"updated_at"`
}

// MapPortalsListingsCreateOutputFromJSON deserializes JSON data into a PortalsListingsCreateOutput.
func MapPortalsListingsCreateOutputFromJSON(data []byte) (*PortalsListingsCreateOutput, error) {
	var v PortalsListingsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsListingsCreateOutputToJSON serializes a PortalsListingsCreateOutput to JSON.
func MapPortalsListingsCreateOutputToJSON(v *PortalsListingsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsListingsCreateBodyAccess represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type PortalsListingsCreateBodyAccess struct {
	Type               *string `json:"type,omitempty"`
	ProviderTemplateId *string `json:"provider_template_id,omitempty"`
	MagicMcpServerId   *string `json:"magic_mcp_server_id,omitempty"`
	SkillId            *string `json:"skill_id,omitempty"`
	SkillTemplateId    *string `json:"skill_template_id,omitempty"`
	SkillGroupId       *string `json:"skill_group_id,omitempty"`
	SkillMarketplaceId *string `json:"skill_marketplace_id,omitempty"`
}

// PortalsListingsCreateBody represents the portals listings create body type.
type PortalsListingsCreateBody struct {
	Name        *string                         `json:"name,omitempty"`
	Description *string                         `json:"description,omitempty"`
	Readme      *string                         `json:"readme,omitempty"`
	Access      PortalsListingsCreateBodyAccess `json:"access"`
}

// MapPortalsListingsCreateBodyFromJSON deserializes JSON data into a PortalsListingsCreateBody.
func MapPortalsListingsCreateBodyFromJSON(data []byte) (*PortalsListingsCreateBody, error) {
	var v PortalsListingsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsListingsCreateBodyToJSON serializes a PortalsListingsCreateBody to JSON.
func MapPortalsListingsCreateBodyToJSON(v *PortalsListingsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
