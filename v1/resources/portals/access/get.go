package access

import (
	"encoding/json"
	"time"
)

// PortalsAccessGetOutputAccessProviderTemplate represents the portals access get output access provider template type.
type PortalsAccessGetOutputAccessProviderTemplate struct {
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

// PortalsAccessGetOutputAccessMagicMcpServer represents the portals access get output access magic mcp server type.
type PortalsAccessGetOutputAccessMagicMcpServer struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Name        *string `json:"name,omitempty"`
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
}

// PortalsAccessGetOutputConsumerGroup represents the portals access get output consumer group type.
type PortalsAccessGetOutputConsumerGroup struct {
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

// PortalsAccessGetOutput represents the portals access get output type.
type PortalsAccessGetOutput struct {
	Object        string                              `json:"object"`
	Id            string                              `json:"id"`
	Name          string                              `json:"name"`
	Description   *string                             `json:"description,omitempty"`
	Readme        *string                             `json:"readme,omitempty"`
	Access        PortalsAccessGetOutputAccess        `json:"access"`
	ConsumerGroup PortalsAccessGetOutputConsumerGroup `json:"consumer_group"`
	CreatedAt     time.Time                           `json:"created_at"`
	UpdatedAt     time.Time                           `json:"updated_at"`
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
