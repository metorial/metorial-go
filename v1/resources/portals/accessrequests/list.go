package accessrequests

import (
	"encoding/json"
	"time"
)

// PortalsAccessRequestsListOutputItemsConsumerProfile represents the portals access requests list output items consumer profile type.
type PortalsAccessRequestsListOutputItemsConsumerProfile struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

// PortalsAccessRequestsListOutputItemsTargetProviderTemplate represents the portals access requests list output items target provider template type.
type PortalsAccessRequestsListOutputItemsTargetProviderTemplate struct {
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

// PortalsAccessRequestsListOutputItemsTargetMagicMcpServer represents the portals access requests list output items target magic mcp server type.
type PortalsAccessRequestsListOutputItemsTargetMagicMcpServer struct {
	Object      string  `json:"object"`
	Id          string  `json:"id"`
	Status      string  `json:"status"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// PortalsAccessRequestsListOutputItemsTarget represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type PortalsAccessRequestsListOutputItemsTarget struct {
	Type             *string                                                     `json:"type,omitempty"`
	ProviderTemplate *PortalsAccessRequestsListOutputItemsTargetProviderTemplate `json:"provider_template,omitempty"`
	MagicMcpServer   *PortalsAccessRequestsListOutputItemsTargetMagicMcpServer   `json:"magic_mcp_server,omitempty"`
}

// PortalsAccessRequestsListOutputItems represents the portals access requests list output items type.
type PortalsAccessRequestsListOutputItems struct {
	Object            string                                              `json:"object"`
	Id                string                                              `json:"id"`
	Status            string                                              `json:"status"`
	Message           *string                                             `json:"message,omitempty"`
	ResolutionMessage *string                                             `json:"resolution_message,omitempty"`
	ConsumerProfile   PortalsAccessRequestsListOutputItemsConsumerProfile `json:"consumer_profile"`
	Target            PortalsAccessRequestsListOutputItemsTarget          `json:"target"`
	CreatedAt         time.Time                                           `json:"created_at"`
	UpdatedAt         time.Time                                           `json:"updated_at"`
	ReviewedAt        *time.Time                                          `json:"reviewed_at,omitempty"`
}

// PortalsAccessRequestsListOutputPagination represents the portals access requests list output pagination type.
type PortalsAccessRequestsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// PortalsAccessRequestsListOutput represents the portals access requests list output type.
type PortalsAccessRequestsListOutput struct {
	Items      []PortalsAccessRequestsListOutputItems    `json:"items"`
	Pagination PortalsAccessRequestsListOutputPagination `json:"pagination"`
}

// MapPortalsAccessRequestsListOutputFromJSON deserializes JSON data into a PortalsAccessRequestsListOutput.
func MapPortalsAccessRequestsListOutputFromJSON(data []byte) (*PortalsAccessRequestsListOutput, error) {
	var v PortalsAccessRequestsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsAccessRequestsListOutputToJSON serializes a PortalsAccessRequestsListOutput to JSON.
func MapPortalsAccessRequestsListOutputToJSON(v *PortalsAccessRequestsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsAccessRequestsListQuery represents the portals access requests list query type.
type PortalsAccessRequestsListQuery struct {
	Limit             *float64 `json:"limit,omitempty"`
	After             *string  `json:"after,omitempty"`
	Before            *string  `json:"before,omitempty"`
	Cursor            *string  `json:"cursor,omitempty"`
	Order             *string  `json:"order,omitempty"`
	Status            *any     `json:"status,omitempty"`
	ConsumerProfileId *any     `json:"consumer_profile_id,omitempty"`
	Search            *string  `json:"search,omitempty"`
}

// MapPortalsAccessRequestsListQueryFromJSON deserializes JSON data into a PortalsAccessRequestsListQuery.
func MapPortalsAccessRequestsListQueryFromJSON(data []byte) (*PortalsAccessRequestsListQuery, error) {
	var v PortalsAccessRequestsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsAccessRequestsListQueryToJSON serializes a PortalsAccessRequestsListQuery to JSON.
func MapPortalsAccessRequestsListQueryToJSON(v *PortalsAccessRequestsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
