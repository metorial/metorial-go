package access

import (
	"encoding/json"
	"time"
)

// PortalsAccessCreateOutputAccessProviderTemplate represents the portals access create output access provider template type.
type PortalsAccessCreateOutputAccessProviderTemplate struct {
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

// PortalsAccessCreateOutputAccessMagicMcpServer represents the portals access create output access magic mcp server type.
type PortalsAccessCreateOutputAccessMagicMcpServer struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Name        *string `json:"name,omitempty"`
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
}

// PortalsAccessCreateOutputConsumerGroup represents the portals access create output consumer group type.
type PortalsAccessCreateOutputConsumerGroup struct {
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

// PortalsAccessCreateOutput represents the portals access create output type.
type PortalsAccessCreateOutput struct {
	Object        string                                 `json:"object"`
	Id            string                                 `json:"id"`
	Name          string                                 `json:"name"`
	Description   *string                                `json:"description,omitempty"`
	Readme        *string                                `json:"readme,omitempty"`
	Access        PortalsAccessCreateOutputAccess        `json:"access"`
	ConsumerGroup PortalsAccessCreateOutputConsumerGroup `json:"consumer_group"`
	CreatedAt     time.Time                              `json:"created_at"`
	UpdatedAt     time.Time                              `json:"updated_at"`
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
	Type               *string `json:"type,omitempty"`
	ProviderTemplateId *string `json:"provider_template_id,omitempty"`
	MagicMcpServerId   *string `json:"magic_mcp_server_id,omitempty"`
	SkillId            *string `json:"skill_id,omitempty"`
	SkillTemplateId    *string `json:"skill_template_id,omitempty"`
	SkillGroupId       *string `json:"skill_group_id,omitempty"`
	SkillMarketplaceId *string `json:"skill_marketplace_id,omitempty"`
}

// PortalsAccessCreateBody represents the portals access create body type.
type PortalsAccessCreateBody struct {
	ConsumerGroupId string                        `json:"consumer_group_id"`
	Name            *string                       `json:"name,omitempty"`
	Description     *string                       `json:"description,omitempty"`
	Readme          *string                       `json:"readme,omitempty"`
	Access          PortalsAccessCreateBodyAccess `json:"access"`
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
