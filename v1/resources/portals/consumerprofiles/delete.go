package consumerprofiles

import (
	"encoding/json"
	"time"
)

// PortalsConsumerProfilesDeleteOutputGroupsGroup represents the portals consumer profiles delete output groups group type.
type PortalsConsumerProfilesDeleteOutputGroupsGroup struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Status      string    `json:"status"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	IsDefault   bool      `json:"is_default"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// PortalsConsumerProfilesDeleteOutputGroups represents the portals consumer profiles delete output groups type.
type PortalsConsumerProfilesDeleteOutputGroups struct {
	Object      string                                         `json:"object"`
	Group       PortalsConsumerProfilesDeleteOutputGroupsGroup `json:"group"`
	AssignedVia string                                         `json:"assigned_via"`
}

// PortalsConsumerProfilesDeleteOutput represents the portals consumer profiles delete output type.
type PortalsConsumerProfilesDeleteOutput struct {
	Object     string                                       `json:"object"`
	Id         string                                       `json:"id"`
	Name       string                                       `json:"name"`
	Email      string                                       `json:"email"`
	ImageUrl   string                                       `json:"image_url"`
	ConsumerId string                                       `json:"consumer_id"`
	UserId     *string                                      `json:"user_id,omitempty"`
	Status     string                                       `json:"status"`
	CreatedAt  time.Time                                    `json:"created_at"`
	UpdatedAt  time.Time                                    `json:"updated_at"`
	Groups     *[]PortalsConsumerProfilesDeleteOutputGroups `json:"groups,omitempty"`
}

// MapPortalsConsumerProfilesDeleteOutputFromJSON deserializes JSON data into a PortalsConsumerProfilesDeleteOutput.
func MapPortalsConsumerProfilesDeleteOutputFromJSON(data []byte) (*PortalsConsumerProfilesDeleteOutput, error) {
	var v PortalsConsumerProfilesDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerProfilesDeleteOutputToJSON serializes a PortalsConsumerProfilesDeleteOutput to JSON.
func MapPortalsConsumerProfilesDeleteOutputToJSON(v *PortalsConsumerProfilesDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
