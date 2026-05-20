package accessrequests

import (
	"encoding/json"
	"time"
)

// PortalsAccessRequestsUpdateOutputConsumerProfile represents the portals access requests update output consumer profile type.
type PortalsAccessRequestsUpdateOutputConsumerProfile struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

// PortalsAccessRequestsUpdateOutputTargetProviderTemplate represents the portals access requests update output target provider template type.
type PortalsAccessRequestsUpdateOutputTargetProviderTemplate struct {
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

// PortalsAccessRequestsUpdateOutputTargetMagicMcpServer represents the portals access requests update output target magic mcp server type.
type PortalsAccessRequestsUpdateOutputTargetMagicMcpServer struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// PortalsAccessRequestsUpdateOutputTarget represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type PortalsAccessRequestsUpdateOutputTarget struct {
	Type             *string                                                  `json:"type,omitempty"`
	ProviderTemplate *PortalsAccessRequestsUpdateOutputTargetProviderTemplate `json:"provider_template,omitempty"`
	MagicMcpServer   *PortalsAccessRequestsUpdateOutputTargetMagicMcpServer   `json:"magic_mcp_server,omitempty"`
}

// PortalsAccessRequestsUpdateOutput represents the portals access requests update output type.
type PortalsAccessRequestsUpdateOutput struct {
	Object            string                                           `json:"object"`
	Id                string                                           `json:"id"`
	Status            string                                           `json:"status"`
	Message           *string                                          `json:"message,omitempty"`
	ResolutionMessage *string                                          `json:"resolution_message,omitempty"`
	ConsumerProfile   PortalsAccessRequestsUpdateOutputConsumerProfile `json:"consumer_profile"`
	Target            PortalsAccessRequestsUpdateOutputTarget          `json:"target"`
	CreatedAt         time.Time                                        `json:"created_at"`
	UpdatedAt         time.Time                                        `json:"updated_at"`
	ReviewedAt        *time.Time                                       `json:"reviewed_at,omitempty"`
}

// MapPortalsAccessRequestsUpdateOutputFromJSON deserializes JSON data into a PortalsAccessRequestsUpdateOutput.
func MapPortalsAccessRequestsUpdateOutputFromJSON(data []byte) (*PortalsAccessRequestsUpdateOutput, error) {
	var v PortalsAccessRequestsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsAccessRequestsUpdateOutputToJSON serializes a PortalsAccessRequestsUpdateOutput to JSON.
func MapPortalsAccessRequestsUpdateOutputToJSON(v *PortalsAccessRequestsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsAccessRequestsUpdateBody represents the portals access requests update body type.
type PortalsAccessRequestsUpdateBody struct {
	Status            string  `json:"status"`
	ResolutionMessage *string `json:"resolution_message,omitempty"`
	ConsumerGroupId   *string `json:"consumer_group_id,omitempty"`
}

// MapPortalsAccessRequestsUpdateBodyFromJSON deserializes JSON data into a PortalsAccessRequestsUpdateBody.
func MapPortalsAccessRequestsUpdateBodyFromJSON(data []byte) (*PortalsAccessRequestsUpdateBody, error) {
	var v PortalsAccessRequestsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsAccessRequestsUpdateBodyToJSON serializes a PortalsAccessRequestsUpdateBody to JSON.
func MapPortalsAccessRequestsUpdateBodyToJSON(v *PortalsAccessRequestsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
