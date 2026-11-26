package accessrequests

import (
	"encoding/json"
	"time"
)

// PortalsAccessRequestsUpdateOutputConsumerProfile - Consumer who requested access.
type PortalsAccessRequestsUpdateOutputConsumerProfile struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Email - Email address of the consumer.
	Email string `json:"email"`
}

// PortalsAccessRequestsUpdateOutputTargetProviderTemplate represents the portals access requests update output target provider template type.
type PortalsAccessRequestsUpdateOutputTargetProviderTemplate struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once you archive the template.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata map[string]any `json:"metadata"`
	// IntegrationId - Integration the template was created from. `null` when it wasn’t created from one.
	IntegrationId *string `json:"integration_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// PortalsAccessRequestsUpdateOutputTargetMagicMcpServer represents the portals access requests update output target magic mcp server type.
type PortalsAccessRequestsUpdateOutputTargetMagicMcpServer struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the server is deleted.
	Status string `json:"status"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
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
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Whether the request still waits for review.
	Status string `json:"status"`
	// Message - Message the consumer sent with the request.
	Message *string `json:"message,omitempty"`
	// ResolutionMessage - Message shown to the consumer about the decision.
	ResolutionMessage *string `json:"resolution_message,omitempty"`
	// ConsumerProfile - Consumer who requested access.
	ConsumerProfile PortalsAccessRequestsUpdateOutputConsumerProfile `json:"consumer_profile"`
	// Target - The integration the consumer requested access to.
	Target PortalsAccessRequestsUpdateOutputTarget `json:"target"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ReviewedAt - When the request was reviewed. `null` while pending.
	ReviewedAt *time.Time `json:"reviewed_at,omitempty"`
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
	// Status - Whether to approve or reject the request.
	Status string `json:"status"`
	// ResolutionMessage - Message shown to the consumer about the decision.
	ResolutionMessage *string `json:"resolution_message,omitempty"`
	// ConsumerGroupId - Consumer group the approval grants access through.
	ConsumerGroupId *string `json:"consumer_group_id,omitempty"`
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
