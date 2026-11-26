package accessrequests

import (
	"encoding/json"
	"time"
)

// PortalsAccessRequestsListOutputItemsConsumerProfile - Consumer who requested access.
type PortalsAccessRequestsListOutputItemsConsumerProfile struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Email - Email address of the consumer.
	Email string `json:"email"`
}

// PortalsAccessRequestsListOutputItemsTargetProviderTemplate represents the portals access requests list output items target provider template type.
type PortalsAccessRequestsListOutputItemsTargetProviderTemplate struct {
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

// PortalsAccessRequestsListOutputItemsTargetMagicMcpServer represents the portals access requests list output items target magic mcp server type.
type PortalsAccessRequestsListOutputItemsTargetMagicMcpServer struct {
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

// PortalsAccessRequestsListOutputItemsTarget represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type PortalsAccessRequestsListOutputItemsTarget struct {
	Type             *string                                                     `json:"type,omitempty"`
	ProviderTemplate *PortalsAccessRequestsListOutputItemsTargetProviderTemplate `json:"provider_template,omitempty"`
	MagicMcpServer   *PortalsAccessRequestsListOutputItemsTargetMagicMcpServer   `json:"magic_mcp_server,omitempty"`
}

// PortalsAccessRequestsListOutputItems represents the portals access requests list output items type.
type PortalsAccessRequestsListOutputItems struct {
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
	ConsumerProfile PortalsAccessRequestsListOutputItemsConsumerProfile `json:"consumer_profile"`
	// Target - The integration the consumer requested access to.
	Target PortalsAccessRequestsListOutputItemsTarget `json:"target"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ReviewedAt - When the request was reviewed. `null` while pending.
	ReviewedAt *time.Time `json:"reviewed_at,omitempty"`
}

// PortalsAccessRequestsListOutputPagination represents the portals access requests list output pagination type.
type PortalsAccessRequestsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// PortalsAccessRequestsListOutput represents the portals access requests list output type.
type PortalsAccessRequestsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                    `json:"object"`
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
	// Limit - Page size, from 1 to 100. Defaults to 20.
	Limit *float64 `json:"limit,omitempty"`
	// After - Return the page after this point. Pass `pagination.after` if the previous response has it, otherwise the `id` of its last item.
	After *string `json:"after,omitempty"`
	// Before - Return the page before this point. Pass `pagination.before` if the previous response has it, otherwise the `id` of its first item.
	Before *string `json:"before,omitempty"`
	// Cursor - A `cur_` cursor from `pagination.after` or `pagination.before`. Alternative to `after` and `before`.
	Cursor *string `json:"cursor,omitempty"`
	// Order - Sort direction.
	Order *string `json:"order,omitempty"`
	// Status - Only return requests with one of these statuses.
	Status *any `json:"status,omitempty"`
	// ConsumerProfileId - Only return requests from these consumer profiles.
	ConsumerProfileId *any `json:"consumer_profile_id,omitempty"`
	// Search - Search by name or description.
	Search *string `json:"search,omitempty"`
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
