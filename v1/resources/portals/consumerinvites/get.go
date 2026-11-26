package consumerinvites

import (
	"encoding/json"
	"time"
)

// PortalsConsumerInvitesGetOutputConsumerProfile - Consumer the invite was sent to.
type PortalsConsumerInvitesGetOutputConsumerProfile struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Email - Email address of the consumer.
	Email string `json:"email"`
}

// PortalsConsumerInvitesGetOutputInvitedBy - Organization member who sent the invite.
type PortalsConsumerInvitesGetOutputInvitedBy struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
}

// PortalsConsumerInvitesGetOutput represents the portals consumer invites get output type.
type PortalsConsumerInvitesGetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `accepted` once the consumer joins.
	Status string `json:"status"`
	// PortalUrl - Where the invited consumer signs up. `null` outside portals.
	PortalUrl *string `json:"portal_url,omitempty"`
	// ConsumerProfile - Consumer the invite was sent to.
	ConsumerProfile PortalsConsumerInvitesGetOutputConsumerProfile `json:"consumer_profile"`
	// InvitedBy - Organization member who sent the invite.
	InvitedBy PortalsConsumerInvitesGetOutputInvitedBy `json:"invited_by"`
	// Message - Personal message included in the invite.
	Message *string `json:"message,omitempty"`
	// AcceptedAt - When the invite was accepted. `null` while pending.
	AcceptedAt *time.Time `json:"accepted_at,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapPortalsConsumerInvitesGetOutputFromJSON deserializes JSON data into a PortalsConsumerInvitesGetOutput.
func MapPortalsConsumerInvitesGetOutputFromJSON(data []byte) (*PortalsConsumerInvitesGetOutput, error) {
	var v PortalsConsumerInvitesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerInvitesGetOutputToJSON serializes a PortalsConsumerInvitesGetOutput to JSON.
func MapPortalsConsumerInvitesGetOutputToJSON(v *PortalsConsumerInvitesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
