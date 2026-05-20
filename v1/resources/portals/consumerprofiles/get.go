package consumerprofiles

import (
	"encoding/json"
	"time"
)

// PortalsConsumerProfilesGetOutputGroupsGroup represents the portals consumer profiles get output groups group type.
type PortalsConsumerProfilesGetOutputGroupsGroup struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Status      string    `json:"status"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	IsDefault   bool      `json:"is_default"`
	SsoGroupIds []string  `json:"sso_group_ids"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// PortalsConsumerProfilesGetOutputGroups represents the portals consumer profiles get output groups type.
type PortalsConsumerProfilesGetOutputGroups struct {
	Object      string                                      `json:"object"`
	Group       PortalsConsumerProfilesGetOutputGroupsGroup `json:"group"`
	AssignedVia string                                      `json:"assigned_via"`
}

// PortalsConsumerProfilesGetOutput represents the portals consumer profiles get output type.
type PortalsConsumerProfilesGetOutput struct {
	Object     string                                    `json:"object"`
	Id         string                                    `json:"id"`
	Name       string                                    `json:"name"`
	Email      string                                    `json:"email"`
	ImageUrl   string                                    `json:"image_url"`
	ConsumerId string                                    `json:"consumer_id"`
	Status     string                                    `json:"status"`
	CreatedAt  time.Time                                 `json:"created_at"`
	UpdatedAt  time.Time                                 `json:"updated_at"`
	Groups     *[]PortalsConsumerProfilesGetOutputGroups `json:"groups,omitempty"`
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
