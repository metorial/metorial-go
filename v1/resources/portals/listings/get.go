package listings

import (
	"encoding/json"
	"time"
)

// PortalsListingsGetOutputAccessProviderTemplate represents the portals listings get output access provider template type.
type PortalsListingsGetOutputAccessProviderTemplate struct {
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

// PortalsListingsGetOutputAccessMagicMcpServer represents the portals listings get output access magic mcp server type.
type PortalsListingsGetOutputAccessMagicMcpServer struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// PortalsListingsGetOutputAccessSkill represents the portals listings get output access skill type.
type PortalsListingsGetOutputAccessSkill struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

// PortalsListingsGetOutputAccessSkillTemplate represents the portals listings get output access skill template type.
type PortalsListingsGetOutputAccessSkillTemplate struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Owner       string  `json:"owner"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// PortalsListingsGetOutputAccessSkillGroup represents the portals listings get output access skill group type.
type PortalsListingsGetOutputAccessSkillGroup struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// PortalsListingsGetOutputAccessSkillMarketplace represents the portals listings get output access skill marketplace type.
type PortalsListingsGetOutputAccessSkillMarketplace struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
}

// PortalsListingsGetOutputAccess represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type PortalsListingsGetOutputAccess struct {
	Type             *string                                         `json:"type,omitempty"`
	ProviderTemplate *PortalsListingsGetOutputAccessProviderTemplate `json:"provider_template,omitempty"`
	MagicMcpServer   *PortalsListingsGetOutputAccessMagicMcpServer   `json:"magic_mcp_server,omitempty"`
	Skill            *PortalsListingsGetOutputAccessSkill            `json:"skill,omitempty"`
	SkillTemplate    *PortalsListingsGetOutputAccessSkillTemplate    `json:"skill_template,omitempty"`
	SkillGroup       *PortalsListingsGetOutputAccessSkillGroup       `json:"skill_group,omitempty"`
	SkillMarketplace *PortalsListingsGetOutputAccessSkillMarketplace `json:"skill_marketplace,omitempty"`
}

// PortalsListingsGetOutputGroups represents the portals listings get output groups type.
type PortalsListingsGetOutputGroups struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Index       float64 `json:"index"`
}

// PortalsListingsGetOutput represents the portals listings get output type.
type PortalsListingsGetOutput struct {
	Object      string                           `json:"object"`
	Id          string                           `json:"id"`
	Name        string                           `json:"name"`
	Description *string                          `json:"description,omitempty"`
	Readme      *string                          `json:"readme,omitempty"`
	Access      PortalsListingsGetOutputAccess   `json:"access"`
	Groups      []PortalsListingsGetOutputGroups `json:"groups"`
	CreatedAt   time.Time                        `json:"created_at"`
	UpdatedAt   time.Time                        `json:"updated_at"`
}

// MapPortalsListingsGetOutputFromJSON deserializes JSON data into a PortalsListingsGetOutput.
func MapPortalsListingsGetOutputFromJSON(data []byte) (*PortalsListingsGetOutput, error) {
	var v PortalsListingsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsListingsGetOutputToJSON serializes a PortalsListingsGetOutput to JSON.
func MapPortalsListingsGetOutputToJSON(v *PortalsListingsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
