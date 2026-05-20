package listings

import (
	"encoding/json"
	"time"
)

// PortalsListingsDeleteOutputAccessProviderTemplate represents the portals listings delete output access provider template type.
type PortalsListingsDeleteOutputAccessProviderTemplate struct {
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

// PortalsListingsDeleteOutputAccessMagicMcpServer represents the portals listings delete output access magic mcp server type.
type PortalsListingsDeleteOutputAccessMagicMcpServer struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// PortalsListingsDeleteOutputAccessSkill represents the portals listings delete output access skill type.
type PortalsListingsDeleteOutputAccessSkill struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

// PortalsListingsDeleteOutputAccessSkillTemplate represents the portals listings delete output access skill template type.
type PortalsListingsDeleteOutputAccessSkillTemplate struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Owner       string  `json:"owner"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// PortalsListingsDeleteOutputAccessSkillGroup represents the portals listings delete output access skill group type.
type PortalsListingsDeleteOutputAccessSkillGroup struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// PortalsListingsDeleteOutputAccessSkillMarketplace represents the portals listings delete output access skill marketplace type.
type PortalsListingsDeleteOutputAccessSkillMarketplace struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
}

// PortalsListingsDeleteOutputAccess represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type PortalsListingsDeleteOutputAccess struct {
	Type             *string                                            `json:"type,omitempty"`
	ProviderTemplate *PortalsListingsDeleteOutputAccessProviderTemplate `json:"provider_template,omitempty"`
	MagicMcpServer   *PortalsListingsDeleteOutputAccessMagicMcpServer   `json:"magic_mcp_server,omitempty"`
	Skill            *PortalsListingsDeleteOutputAccessSkill            `json:"skill,omitempty"`
	SkillTemplate    *PortalsListingsDeleteOutputAccessSkillTemplate    `json:"skill_template,omitempty"`
	SkillGroup       *PortalsListingsDeleteOutputAccessSkillGroup       `json:"skill_group,omitempty"`
	SkillMarketplace *PortalsListingsDeleteOutputAccessSkillMarketplace `json:"skill_marketplace,omitempty"`
}

// PortalsListingsDeleteOutputGroups represents the portals listings delete output groups type.
type PortalsListingsDeleteOutputGroups struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Index       float64 `json:"index"`
}

// PortalsListingsDeleteOutput represents the portals listings delete output type.
type PortalsListingsDeleteOutput struct {
	Object      string                              `json:"object"`
	Id          string                              `json:"id"`
	Name        string                              `json:"name"`
	Description *string                             `json:"description,omitempty"`
	Readme      *string                             `json:"readme,omitempty"`
	Access      PortalsListingsDeleteOutputAccess   `json:"access"`
	Groups      []PortalsListingsDeleteOutputGroups `json:"groups"`
	CreatedAt   time.Time                           `json:"created_at"`
	UpdatedAt   time.Time                           `json:"updated_at"`
}

// MapPortalsListingsDeleteOutputFromJSON deserializes JSON data into a PortalsListingsDeleteOutput.
func MapPortalsListingsDeleteOutputFromJSON(data []byte) (*PortalsListingsDeleteOutput, error) {
	var v PortalsListingsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsListingsDeleteOutputToJSON serializes a PortalsListingsDeleteOutput to JSON.
func MapPortalsListingsDeleteOutputToJSON(v *PortalsListingsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
