package consumerprofiles

import (
	"encoding/json"
	"time"
)

// PortalsConsumerProfilesCreateOutputGroupsGroup represents the portals consumer profiles create output groups group type.
type PortalsConsumerProfilesCreateOutputGroupsGroup struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once you delete the group. Archived groups stay listable.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// IsDefault - Whether this is a default group on the portal.
	IsDefault bool `json:"is_default"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// PortalsConsumerProfilesCreateOutputGroups - Group the profile belongs to and how it joined.
type PortalsConsumerProfilesCreateOutputGroups struct {
	// Object - String representing the object’s type.
	Object string                                         `json:"object"`
	Group  PortalsConsumerProfilesCreateOutputGroupsGroup `json:"group"`
	// AssignedVia - How the profile joined the group.
	AssignedVia string `json:"assigned_via"`
	// SsoGroupIds - Identity provider group ids that placed this profile in the group. Empty unless assigned via SSO.
	SsoGroupIds []string `json:"sso_group_ids"`
}

// PortalsConsumerProfilesCreateOutput represents the portals consumer profiles create output type.
type PortalsConsumerProfilesCreateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Email - Email address.
	Email string `json:"email"`
	// ImageUrl - Avatar image URL.
	ImageUrl string `json:"image_url"`
	// ConsumerId - Consumer this profile belongs to.
	ConsumerId string `json:"consumer_id"`
	// UserId - User this profile is linked to. `null` when unlinked.
	UserId *string `json:"user_id,omitempty"`
	// Status - `invited` while the invitation is pending.
	Status string `json:"status"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// Groups - Groups the profile belongs to. `null` when not loaded.
	Groups *[]PortalsConsumerProfilesCreateOutputGroups `json:"groups,omitempty"`
}

// MapPortalsConsumerProfilesCreateOutputFromJSON deserializes JSON data into a PortalsConsumerProfilesCreateOutput.
func MapPortalsConsumerProfilesCreateOutputFromJSON(data []byte) (*PortalsConsumerProfilesCreateOutput, error) {
	var v PortalsConsumerProfilesCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerProfilesCreateOutputToJSON serializes a PortalsConsumerProfilesCreateOutput to JSON.
func MapPortalsConsumerProfilesCreateOutputToJSON(v *PortalsConsumerProfilesCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsConsumerProfilesCreateBody represents the portals consumer profiles create body type.
type PortalsConsumerProfilesCreateBody struct {
	// Email - Email address of the consumer.
	Email string `json:"email"`
	// Name - Display name.
	Name string `json:"name"`
}

// MapPortalsConsumerProfilesCreateBodyFromJSON deserializes JSON data into a PortalsConsumerProfilesCreateBody.
func MapPortalsConsumerProfilesCreateBodyFromJSON(data []byte) (*PortalsConsumerProfilesCreateBody, error) {
	var v PortalsConsumerProfilesCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerProfilesCreateBodyToJSON serializes a PortalsConsumerProfilesCreateBody to JSON.
func MapPortalsConsumerProfilesCreateBodyToJSON(v *PortalsConsumerProfilesCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
