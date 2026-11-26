package consumerprofiles

import (
	"encoding/json"
	"time"
)

// PortalsConsumerProfilesGetOutputGroupsGroup represents the portals consumer profiles get output groups group type.
type PortalsConsumerProfilesGetOutputGroupsGroup struct {
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

// PortalsConsumerProfilesGetOutputGroups - Group the profile belongs to and how it joined.
type PortalsConsumerProfilesGetOutputGroups struct {
	// Object - String representing the object’s type.
	Object string                                      `json:"object"`
	Group  PortalsConsumerProfilesGetOutputGroupsGroup `json:"group"`
	// AssignedVia - How the profile joined the group.
	AssignedVia string `json:"assigned_via"`
	// SsoGroupIds - Identity provider group ids that placed this profile in the group. Empty unless assigned via SSO.
	SsoGroupIds []string `json:"sso_group_ids"`
}

// PortalsConsumerProfilesGetOutput represents the portals consumer profiles get output type.
type PortalsConsumerProfilesGetOutput struct {
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
	Groups *[]PortalsConsumerProfilesGetOutputGroups `json:"groups,omitempty"`
}

// MapPortalsConsumerProfilesGetOutputFromJSON deserializes JSON data into a PortalsConsumerProfilesGetOutput.
func MapPortalsConsumerProfilesGetOutputFromJSON(data []byte) (*PortalsConsumerProfilesGetOutput, error) {
	var v PortalsConsumerProfilesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerProfilesGetOutputToJSON serializes a PortalsConsumerProfilesGetOutput to JSON.
func MapPortalsConsumerProfilesGetOutputToJSON(v *PortalsConsumerProfilesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
