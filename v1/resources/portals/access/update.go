package access

import (
	"encoding/json"
	"time"
)

// PortalsAccessUpdateOutputAccessProviderTemplate represents the portals access update output access provider template type.
type PortalsAccessUpdateOutputAccessProviderTemplate struct {
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

// PortalsAccessUpdateOutputAccessMagicMcpServer represents the portals access update output access magic mcp server type.
type PortalsAccessUpdateOutputAccessMagicMcpServer struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Name        *string `json:"name,omitempty"`
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
}

// PortalsAccessUpdateOutputConsumerGroup represents the portals access update output consumer group type.
type PortalsAccessUpdateOutputConsumerGroup struct {
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

// PortalsAccessUpdateOutput represents the portals access update output type.
type PortalsAccessUpdateOutput struct {
	Object        string                                 `json:"object"`
	Id            string                                 `json:"id"`
	Name          string                                 `json:"name"`
	Description   *string                                `json:"description,omitempty"`
	Readme        *string                                `json:"readme,omitempty"`
	Access        PortalsAccessUpdateOutputAccess        `json:"access"`
	ConsumerGroup PortalsAccessUpdateOutputConsumerGroup `json:"consumer_group"`
	CreatedAt     time.Time                              `json:"created_at"`
	UpdatedAt     time.Time                              `json:"updated_at"`
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
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Readme      *string `json:"readme,omitempty"`
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
