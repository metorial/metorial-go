package access

import (
	"encoding/json"
	"time"
)

// PortalsAccessDeleteOutputAccessProviderTemplate represents the portals access delete output access provider template type.
type PortalsAccessDeleteOutputAccessProviderTemplate struct {
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

// PortalsAccessDeleteOutputAccessMagicMcpServer represents the portals access delete output access magic mcp server type.
type PortalsAccessDeleteOutputAccessMagicMcpServer struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Name        *string `json:"name,omitempty"`
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
}

// PortalsAccessDeleteOutputConsumerGroup represents the portals access delete output consumer group type.
type PortalsAccessDeleteOutputConsumerGroup struct {
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

// PortalsAccessDeleteOutput represents the portals access delete output type.
type PortalsAccessDeleteOutput struct {
	Object        string                                 `json:"object"`
	Id            string                                 `json:"id"`
	Name          string                                 `json:"name"`
	Description   *string                                `json:"description,omitempty"`
	Readme        *string                                `json:"readme,omitempty"`
	Access        PortalsAccessDeleteOutputAccess        `json:"access"`
	ConsumerGroup PortalsAccessDeleteOutputConsumerGroup `json:"consumer_group"`
	CreatedAt     time.Time                              `json:"created_at"`
	UpdatedAt     time.Time                              `json:"updated_at"`
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
