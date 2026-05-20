package consumerinvites

import (
	"encoding/json"
	"time"
)

// PortalsConsumerInvitesCreateOutputConsumerProfile represents the portals consumer invites create output consumer profile type.
type PortalsConsumerInvitesCreateOutputConsumerProfile struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

// PortalsConsumerInvitesCreateOutputInvitedBy represents the portals consumer invites create output invited by type.
type PortalsConsumerInvitesCreateOutputInvitedBy struct {
	Object string  `json:"object"`
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Email  *string `json:"email,omitempty"`
}

// PortalsConsumerInvitesCreateOutput represents the portals consumer invites create output type.
type PortalsConsumerInvitesCreateOutput struct {
	Object          string                                            `json:"object"`
	Id              string                                            `json:"id"`
	Status          string                                            `json:"status"`
	PortalUrl       *string                                           `json:"portal_url,omitempty"`
	ConsumerProfile PortalsConsumerInvitesCreateOutputConsumerProfile `json:"consumer_profile"`
	InvitedBy       PortalsConsumerInvitesCreateOutputInvitedBy       `json:"invited_by"`
	Message         *string                                           `json:"message,omitempty"`
	AcceptedAt      *time.Time                                        `json:"accepted_at,omitempty"`
	CreatedAt       time.Time                                         `json:"created_at"`
	UpdatedAt       time.Time                                         `json:"updated_at"`
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
	Name    string  `json:"name"`
	Email   string  `json:"email"`
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
