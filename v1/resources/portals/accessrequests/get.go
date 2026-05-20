package accessrequests

import (
	"encoding/json"
	"time"
)

// PortalsAccessRequestsGetOutputConsumerProfile represents the portals access requests get output consumer profile type.
type PortalsAccessRequestsGetOutputConsumerProfile struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

// PortalsAccessRequestsGetOutputTargetProviderTemplate represents the portals access requests get output target provider template type.
type PortalsAccessRequestsGetOutputTargetProviderTemplate struct {
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

// PortalsAccessRequestsGetOutputTargetMagicMcpServer represents the portals access requests get output target magic mcp server type.
type PortalsAccessRequestsGetOutputTargetMagicMcpServer struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// PortalsAccessRequestsGetOutputTarget represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type PortalsAccessRequestsGetOutputTarget struct {
	Type             *string                                               `json:"type,omitempty"`
	ProviderTemplate *PortalsAccessRequestsGetOutputTargetProviderTemplate `json:"provider_template,omitempty"`
	MagicMcpServer   *PortalsAccessRequestsGetOutputTargetMagicMcpServer   `json:"magic_mcp_server,omitempty"`
}

// PortalsAccessRequestsGetOutput represents the portals access requests get output type.
type PortalsAccessRequestsGetOutput struct {
	Object            string                                        `json:"object"`
	Id                string                                        `json:"id"`
	Status            string                                        `json:"status"`
	Message           *string                                       `json:"message,omitempty"`
	ResolutionMessage *string                                       `json:"resolution_message,omitempty"`
	ConsumerProfile   PortalsAccessRequestsGetOutputConsumerProfile `json:"consumer_profile"`
	Target            PortalsAccessRequestsGetOutputTarget          `json:"target"`
	CreatedAt         time.Time                                     `json:"created_at"`
	UpdatedAt         time.Time                                     `json:"updated_at"`
	ReviewedAt        *time.Time                                    `json:"reviewed_at,omitempty"`
}

// MapPortalsAccessRequestsGetOutputFromJSON deserializes JSON data into a PortalsAccessRequestsGetOutput.
func MapPortalsAccessRequestsGetOutputFromJSON(data []byte) (*PortalsAccessRequestsGetOutput, error) {
	var v PortalsAccessRequestsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsAccessRequestsGetOutputToJSON serializes a PortalsAccessRequestsGetOutput to JSON.
func MapPortalsAccessRequestsGetOutputToJSON(v *PortalsAccessRequestsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
