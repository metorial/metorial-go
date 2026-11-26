package consumerinvites

import (
	"encoding/json"
	"time"
)

// PortalsConsumerInvitesCreateOutputConsumerProfile - Consumer the invite was sent to.
type PortalsConsumerInvitesCreateOutputConsumerProfile struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Email - Email address of the consumer.
	Email string `json:"email"`
}

// PortalsConsumerInvitesCreateOutputInvitedBy - Organization member who sent the invite.
type PortalsConsumerInvitesCreateOutputInvitedBy struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
}

// PortalsConsumerInvitesCreateOutput represents the portals consumer invites create output type.
type PortalsConsumerInvitesCreateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `accepted` once the consumer joins.
	Status string `json:"status"`
	// PortalUrl - Where the invited consumer signs up. `null` outside portals.
	PortalUrl *string `json:"portal_url,omitempty"`
	// ConsumerProfile - Consumer the invite was sent to.
	ConsumerProfile PortalsConsumerInvitesCreateOutputConsumerProfile `json:"consumer_profile"`
	// InvitedBy - Organization member who sent the invite.
	InvitedBy PortalsConsumerInvitesCreateOutputInvitedBy `json:"invited_by"`
	// Message - Personal message included in the invite.
	Message *string `json:"message,omitempty"`
	// AcceptedAt - When the invite was accepted. `null` while pending.
	AcceptedAt *time.Time `json:"accepted_at,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapPortalsConsumerInvitesCreateOutputFromJSON deserializes JSON data into a PortalsConsumerInvitesCreateOutput.
func MapPortalsConsumerInvitesCreateOutputFromJSON(data []byte) (*PortalsConsumerInvitesCreateOutput, error) {
	var v PortalsConsumerInvitesCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerInvitesCreateOutputToJSON serializes a PortalsConsumerInvitesCreateOutput to JSON.
func MapPortalsConsumerInvitesCreateOutputToJSON(v *PortalsConsumerInvitesCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsConsumerInvitesCreateBody represents the portals consumer invites create body type.
type PortalsConsumerInvitesCreateBody struct {
	// Name - Display name.
	Name string `json:"name"`
	// Email - Email address to invite.
	Email string `json:"email"`
	// Message - Personal message included in the invite.
	Message *string `json:"message,omitempty"`
}

// MapPortalsConsumerInvitesCreateBodyFromJSON deserializes JSON data into a PortalsConsumerInvitesCreateBody.
func MapPortalsConsumerInvitesCreateBodyFromJSON(data []byte) (*PortalsConsumerInvitesCreateBody, error) {
	var v PortalsConsumerInvitesCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerInvitesCreateBodyToJSON serializes a PortalsConsumerInvitesCreateBody to JSON.
func MapPortalsConsumerInvitesCreateBodyToJSON(v *PortalsConsumerInvitesCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
