package profiles

import (
	"encoding/json"
	"time"
)

// ConsumersProfilesGetOutputGroupsGroup represents the consumers profiles get output groups group type.
type ConsumersProfilesGetOutputGroupsGroup struct {
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

// ConsumersProfilesGetOutputGroups - Group the profile belongs to and how it joined.
type ConsumersProfilesGetOutputGroups struct {
	// Object - String representing the object’s type.
	Object string                                `json:"object"`
	Group  ConsumersProfilesGetOutputGroupsGroup `json:"group"`
	// AssignedVia - How the profile joined the group.
	AssignedVia string `json:"assigned_via"`
	// SsoGroupIds - Identity provider group ids that placed this profile in the group. Empty unless assigned via SSO.
	SsoGroupIds []string `json:"sso_group_ids"`
}

// ConsumersProfilesGetOutput represents the consumers profiles get output type.
type ConsumersProfilesGetOutput struct {
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
	Groups *[]ConsumersProfilesGetOutputGroups `json:"groups,omitempty"`
}

// MapConsumersProfilesGetOutputFromJSON deserializes JSON data into a ConsumersProfilesGetOutput.
func MapConsumersProfilesGetOutputFromJSON(data []byte) (*ConsumersProfilesGetOutput, error) {
	var v ConsumersProfilesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConsumersProfilesGetOutputToJSON serializes a ConsumersProfilesGetOutput to JSON.
func MapConsumersProfilesGetOutputToJSON(v *ConsumersProfilesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
