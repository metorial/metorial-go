package consumerinvites

import (
	"encoding/json"
	"time"
)

// PortalsConsumerInvitesListOutputItemsConsumerProfile represents the portals consumer invites list output items consumer profile type.
type PortalsConsumerInvitesListOutputItemsConsumerProfile struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

// PortalsConsumerInvitesListOutputItemsInvitedBy represents the portals consumer invites list output items invited by type.
type PortalsConsumerInvitesListOutputItemsInvitedBy struct {
	Object string  `json:"object"`
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Email  *string `json:"email,omitempty"`
}

// PortalsConsumerInvitesListOutputItems represents the portals consumer invites list output items type.
type PortalsConsumerInvitesListOutputItems struct {
	Object          string                                               `json:"object"`
	Id              string                                               `json:"id"`
	Status          string                                               `json:"status"`
	PortalUrl       *string                                              `json:"portal_url,omitempty"`
	ConsumerProfile PortalsConsumerInvitesListOutputItemsConsumerProfile `json:"consumer_profile"`
	InvitedBy       PortalsConsumerInvitesListOutputItemsInvitedBy       `json:"invited_by"`
	Message         *string                                              `json:"message,omitempty"`
	AcceptedAt      *time.Time                                           `json:"accepted_at,omitempty"`
	CreatedAt       time.Time                                            `json:"created_at"`
	UpdatedAt       time.Time                                            `json:"updated_at"`
}

// PortalsConsumerInvitesListOutputPagination represents the portals consumer invites list output pagination type.
type PortalsConsumerInvitesListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// PortalsConsumerInvitesListOutput represents the portals consumer invites list output type.
type PortalsConsumerInvitesListOutput struct {
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
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	Search *string  `json:"search,omitempty"`
	Status *any     `json:"status,omitempty"`
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
