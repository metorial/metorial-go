package consumerprofiles

import (
	"encoding/json"
	"time"
)

// PortalsConsumerProfilesUnassignGroupsOutputGroupsGroup represents the portals consumer profiles unassign groups output groups group type.
type PortalsConsumerProfilesUnassignGroupsOutputGroupsGroup struct {
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

// PortalsConsumerProfilesUnassignGroupsOutputGroups - Group the profile belongs to and how it joined.
type PortalsConsumerProfilesUnassignGroupsOutputGroups struct {
	// Object - String representing the object’s type.
	Object string                                                 `json:"object"`
	Group  PortalsConsumerProfilesUnassignGroupsOutputGroupsGroup `json:"group"`
	// AssignedVia - How the profile joined the group.
	AssignedVia string `json:"assigned_via"`
	// SsoGroupIds - Identity provider group ids that placed this profile in the group. Empty unless assigned via SSO.
	SsoGroupIds []string `json:"sso_group_ids"`
}

// PortalsConsumerProfilesUnassignGroupsOutput represents the portals consumer profiles unassign groups output type.
type PortalsConsumerProfilesUnassignGroupsOutput struct {
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
	Groups *[]PortalsConsumerProfilesUnassignGroupsOutputGroups `json:"groups,omitempty"`
}

// MapPortalsConsumerProfilesUnassignGroupsOutputFromJSON deserializes JSON data into a PortalsConsumerProfilesUnassignGroupsOutput.
func MapPortalsConsumerProfilesUnassignGroupsOutputFromJSON(data []byte) (*PortalsConsumerProfilesUnassignGroupsOutput, error) {
	var v PortalsConsumerProfilesUnassignGroupsOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerProfilesUnassignGroupsOutputToJSON serializes a PortalsConsumerProfilesUnassignGroupsOutput to JSON.
func MapPortalsConsumerProfilesUnassignGroupsOutputToJSON(v *PortalsConsumerProfilesUnassignGroupsOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsConsumerProfilesUnassignGroupsBody represents the portals consumer profiles unassign groups body type.
type PortalsConsumerProfilesUnassignGroupsBody struct {
	// GroupIds - Groups to remove the profile from.
	GroupIds []string `json:"group_ids"`
}

// MapPortalsConsumerProfilesUnassignGroupsBodyFromJSON deserializes JSON data into a PortalsConsumerProfilesUnassignGroupsBody.
func MapPortalsConsumerProfilesUnassignGroupsBodyFromJSON(data []byte) (*PortalsConsumerProfilesUnassignGroupsBody, error) {
	var v PortalsConsumerProfilesUnassignGroupsBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerProfilesUnassignGroupsBodyToJSON serializes a PortalsConsumerProfilesUnassignGroupsBody to JSON.
func MapPortalsConsumerProfilesUnassignGroupsBodyToJSON(v *PortalsConsumerProfilesUnassignGroupsBody) ([]byte, error) {
	return json.Marshal(v)
}
