package consumerinvites

import (
	"encoding/json"
	"time"
)

// PortalsConsumerInvitesGetOutputConsumerProfile represents the portals consumer invites get output consumer profile type.
type PortalsConsumerInvitesGetOutputConsumerProfile struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

// PortalsConsumerInvitesGetOutputInvitedBy represents the portals consumer invites get output invited by type.
type PortalsConsumerInvitesGetOutputInvitedBy struct {
	Object string  `json:"object"`
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Email  *string `json:"email,omitempty"`
}

// PortalsConsumerInvitesGetOutput represents the portals consumer invites get output type.
type PortalsConsumerInvitesGetOutput struct {
	Object          string                                         `json:"object"`
	Id              string                                         `json:"id"`
	Status          string                                         `json:"status"`
	PortalUrl       *string                                        `json:"portal_url,omitempty"`
	ConsumerProfile PortalsConsumerInvitesGetOutputConsumerProfile `json:"consumer_profile"`
	InvitedBy       PortalsConsumerInvitesGetOutputInvitedBy       `json:"invited_by"`
	Message         *string                                        `json:"message,omitempty"`
	AcceptedAt      *time.Time                                     `json:"accepted_at,omitempty"`
	CreatedAt       time.Time                                      `json:"created_at"`
	UpdatedAt       time.Time                                      `json:"updated_at"`
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
