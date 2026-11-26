package consumerinvites

import (
	"encoding/json"
	"time"
)

// PortalsConsumerInvitesListOutputItemsConsumerProfile - Consumer the invite was sent to.
type PortalsConsumerInvitesListOutputItemsConsumerProfile struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Email - Email address of the consumer.
	Email string `json:"email"`
}

// PortalsConsumerInvitesListOutputItemsInvitedBy - Organization member who sent the invite.
type PortalsConsumerInvitesListOutputItemsInvitedBy struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
}

// PortalsConsumerInvitesListOutputItems represents the portals consumer invites list output items type.
type PortalsConsumerInvitesListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `accepted` once the consumer joins.
	Status string `json:"status"`
	// PortalUrl - Where the invited consumer signs up. `null` outside portals.
	PortalUrl *string `json:"portal_url,omitempty"`
	// ConsumerProfile - Consumer the invite was sent to.
	ConsumerProfile PortalsConsumerInvitesListOutputItemsConsumerProfile `json:"consumer_profile"`
	// InvitedBy - Organization member who sent the invite.
	InvitedBy PortalsConsumerInvitesListOutputItemsInvitedBy `json:"invited_by"`
	// Message - Personal message included in the invite.
	Message *string `json:"message,omitempty"`
	// AcceptedAt - When the invite was accepted. `null` while pending.
	AcceptedAt *time.Time `json:"accepted_at,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// PortalsConsumerInvitesListOutputPagination represents the portals consumer invites list output pagination type.
type PortalsConsumerInvitesListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// PortalsConsumerInvitesListOutput represents the portals consumer invites list output type.
type PortalsConsumerInvitesListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                     `json:"object"`
	Items      []PortalsConsumerInvitesListOutputItems    `json:"items"`
	Pagination PortalsConsumerInvitesListOutputPagination `json:"pagination"`
}

// MapPortalsConsumerInvitesListOutputFromJSON deserializes JSON data into a PortalsConsumerInvitesListOutput.
func MapPortalsConsumerInvitesListOutputFromJSON(data []byte) (*PortalsConsumerInvitesListOutput, error) {
	var v PortalsConsumerInvitesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerInvitesListOutputToJSON serializes a PortalsConsumerInvitesListOutput to JSON.
func MapPortalsConsumerInvitesListOutputToJSON(v *PortalsConsumerInvitesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsConsumerInvitesListQuery represents the portals consumer invites list query type.
type PortalsConsumerInvitesListQuery struct {
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
	// Search - Search by name or email.
	Search *string `json:"search,omitempty"`
	// Email - Only return invites sent to one of these emails.
	Email *any `json:"email,omitempty"`
	// Status - Only return invites with one of these statuses.
	Status *any `json:"status,omitempty"`
}

// MapPortalsConsumerInvitesListQueryFromJSON deserializes JSON data into a PortalsConsumerInvitesListQuery.
func MapPortalsConsumerInvitesListQueryFromJSON(data []byte) (*PortalsConsumerInvitesListQuery, error) {
	var v PortalsConsumerInvitesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerInvitesListQueryToJSON serializes a PortalsConsumerInvitesListQuery to JSON.
func MapPortalsConsumerInvitesListQueryToJSON(v *PortalsConsumerInvitesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
