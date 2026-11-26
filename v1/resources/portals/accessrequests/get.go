package accessrequests

import (
	"encoding/json"
	"time"
)

// PortalsAccessRequestsGetOutputConsumerProfile - Consumer who requested access.
type PortalsAccessRequestsGetOutputConsumerProfile struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Email - Email address of the consumer.
	Email string `json:"email"`
}

// PortalsAccessRequestsGetOutputTargetProviderTemplate represents the portals access requests get output target provider template type.
type PortalsAccessRequestsGetOutputTargetProviderTemplate struct {
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

// PortalsAccessRequestsGetOutputTargetMagicMcpServer represents the portals access requests get output target magic mcp server type.
type PortalsAccessRequestsGetOutputTargetMagicMcpServer struct {
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

// PortalsAccessRequestsGetOutputTarget represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type PortalsAccessRequestsGetOutputTarget struct {
	Type             *string                                               `json:"type,omitempty"`
	ProviderTemplate *PortalsAccessRequestsGetOutputTargetProviderTemplate `json:"provider_template,omitempty"`
	MagicMcpServer   *PortalsAccessRequestsGetOutputTargetMagicMcpServer   `json:"magic_mcp_server,omitempty"`
}

// PortalsAccessRequestsGetOutput represents the portals access requests get output type.
type PortalsAccessRequestsGetOutput struct {
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
	ConsumerProfile PortalsAccessRequestsGetOutputConsumerProfile `json:"consumer_profile"`
	// Target - The integration the consumer requested access to.
	Target PortalsAccessRequestsGetOutputTarget `json:"target"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ReviewedAt - When the request was reviewed. `null` while pending.
	ReviewedAt *time.Time `json:"reviewed_at,omitempty"`
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
