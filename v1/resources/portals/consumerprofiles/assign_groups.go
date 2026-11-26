package consumerprofiles

import (
	"encoding/json"
	"time"
)

// PortalsConsumerProfilesAssignGroupsOutputGroupsGroup represents the portals consumer profiles assign groups output groups group type.
type PortalsConsumerProfilesAssignGroupsOutputGroupsGroup struct {
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

// PortalsConsumerProfilesAssignGroupsOutputGroups - Group the profile belongs to and how it joined.
type PortalsConsumerProfilesAssignGroupsOutputGroups struct {
	// Object - String representing the object’s type.
	Object string                                               `json:"object"`
	Group  PortalsConsumerProfilesAssignGroupsOutputGroupsGroup `json:"group"`
	// AssignedVia - How the profile joined the group.
	AssignedVia string `json:"assigned_via"`
	// SsoGroupIds - Identity provider group ids that placed this profile in the group. Empty unless assigned via SSO.
	SsoGroupIds []string `json:"sso_group_ids"`
}

// PortalsConsumerProfilesAssignGroupsOutput represents the portals consumer profiles assign groups output type.
type PortalsConsumerProfilesAssignGroupsOutput struct {
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
	Groups *[]PortalsConsumerProfilesAssignGroupsOutputGroups `json:"groups,omitempty"`
}

// MapPortalsConsumerProfilesAssignGroupsOutputFromJSON deserializes JSON data into a PortalsConsumerProfilesAssignGroupsOutput.
func MapPortalsConsumerProfilesAssignGroupsOutputFromJSON(data []byte) (*PortalsConsumerProfilesAssignGroupsOutput, error) {
	var v PortalsConsumerProfilesAssignGroupsOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerProfilesAssignGroupsOutputToJSON serializes a PortalsConsumerProfilesAssignGroupsOutput to JSON.
func MapPortalsConsumerProfilesAssignGroupsOutputToJSON(v *PortalsConsumerProfilesAssignGroupsOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsConsumerProfilesAssignGroupsBody represents the portals consumer profiles assign groups body type.
type PortalsConsumerProfilesAssignGroupsBody struct {
	// GroupIds - Groups to assign the profile to.
	GroupIds []string `json:"group_ids"`
}

// MapPortalsConsumerProfilesAssignGroupsBodyFromJSON deserializes JSON data into a PortalsConsumerProfilesAssignGroupsBody.
func MapPortalsConsumerProfilesAssignGroupsBodyFromJSON(data []byte) (*PortalsConsumerProfilesAssignGroupsBody, error) {
	var v PortalsConsumerProfilesAssignGroupsBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerProfilesAssignGroupsBodyToJSON serializes a PortalsConsumerProfilesAssignGroupsBody to JSON.
func MapPortalsConsumerProfilesAssignGroupsBodyToJSON(v *PortalsConsumerProfilesAssignGroupsBody) ([]byte, error) {
	return json.Marshal(v)
}
