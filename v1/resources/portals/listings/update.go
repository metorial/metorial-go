package listings

import (
	"encoding/json"
	"time"
)

// PortalsListingsUpdateOutputAccessProviderTemplate represents the portals listings update output access provider template type.
type PortalsListingsUpdateOutputAccessProviderTemplate struct {
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

// PortalsListingsUpdateOutputAccessMagicMcpServer represents the portals listings update output access magic mcp server type.
type PortalsListingsUpdateOutputAccessMagicMcpServer struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// PortalsListingsUpdateOutputAccessSkill represents the portals listings update output access skill type.
type PortalsListingsUpdateOutputAccessSkill struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

// PortalsListingsUpdateOutputAccessSkillTemplate represents the portals listings update output access skill template type.
type PortalsListingsUpdateOutputAccessSkillTemplate struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Owner       string  `json:"owner"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// PortalsListingsUpdateOutputAccessSkillGroup represents the portals listings update output access skill group type.
type PortalsListingsUpdateOutputAccessSkillGroup struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// PortalsListingsUpdateOutputAccessSkillMarketplace represents the portals listings update output access skill marketplace type.
type PortalsListingsUpdateOutputAccessSkillMarketplace struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
}

// PortalsListingsUpdateOutputAccess represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type PortalsListingsUpdateOutputAccess struct {
	Type             *string                                            `json:"type,omitempty"`
	ProviderTemplate *PortalsListingsUpdateOutputAccessProviderTemplate `json:"provider_template,omitempty"`
	MagicMcpServer   *PortalsListingsUpdateOutputAccessMagicMcpServer   `json:"magic_mcp_server,omitempty"`
	Skill            *PortalsListingsUpdateOutputAccessSkill            `json:"skill,omitempty"`
	SkillTemplate    *PortalsListingsUpdateOutputAccessSkillTemplate    `json:"skill_template,omitempty"`
	SkillGroup       *PortalsListingsUpdateOutputAccessSkillGroup       `json:"skill_group,omitempty"`
	SkillMarketplace *PortalsListingsUpdateOutputAccessSkillMarketplace `json:"skill_marketplace,omitempty"`
}

// PortalsListingsUpdateOutputGroups represents the portals listings update output groups type.
type PortalsListingsUpdateOutputGroups struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Index       float64 `json:"index"`
}

// PortalsListingsUpdateOutput represents the portals listings update output type.
type PortalsListingsUpdateOutput struct {
	Object      string                              `json:"object"`
	Id          string                              `json:"id"`
	Name        string                              `json:"name"`
	Description *string                             `json:"description,omitempty"`
	Readme      *string                             `json:"readme,omitempty"`
	Access      PortalsListingsUpdateOutputAccess   `json:"access"`
	Groups      []PortalsListingsUpdateOutputGroups `json:"groups"`
	CreatedAt   time.Time                           `json:"created_at"`
	UpdatedAt   time.Time                           `json:"updated_at"`
}

// MapPortalsListingsUpdateOutputFromJSON deserializes JSON data into a PortalsListingsUpdateOutput.
func MapPortalsListingsUpdateOutputFromJSON(data []byte) (*PortalsListingsUpdateOutput, error) {
	var v PortalsListingsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsListingsUpdateOutputToJSON serializes a PortalsListingsUpdateOutput to JSON.
func MapPortalsListingsUpdateOutputToJSON(v *PortalsListingsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsListingsUpdateBody represents the portals listings update body type.
type PortalsListingsUpdateBody struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Readme      *string `json:"readme,omitempty"`
}

// MapPortalsListingsUpdateBodyFromJSON deserializes JSON data into a PortalsListingsUpdateBody.
func MapPortalsListingsUpdateBodyFromJSON(data []byte) (*PortalsListingsUpdateBody, error) {
	var v PortalsListingsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsListingsUpdateBodyToJSON serializes a PortalsListingsUpdateBody to JSON.
func MapPortalsListingsUpdateBodyToJSON(v *PortalsListingsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
