package consumerprofiles

import (
	"encoding/json"
	"time"
)

// PortalsConsumerProfilesAssignGroupsOutputGroupsGroup represents the portals consumer profiles assign groups output groups group type.
type PortalsConsumerProfilesAssignGroupsOutputGroupsGroup struct {
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

// PortalsConsumerProfilesAssignGroupsOutputGroups represents the portals consumer profiles assign groups output groups type.
type PortalsConsumerProfilesAssignGroupsOutputGroups struct {
	Object      string                                               `json:"object"`
	Group       PortalsConsumerProfilesAssignGroupsOutputGroupsGroup `json:"group"`
	AssignedVia string                                               `json:"assigned_via"`
}

// PortalsConsumerProfilesAssignGroupsOutput represents the portals consumer profiles assign groups output type.
type PortalsConsumerProfilesAssignGroupsOutput struct {
	Object     string                                             `json:"object"`
	Id         string                                             `json:"id"`
	Name       string                                             `json:"name"`
	Email      string                                             `json:"email"`
	ImageUrl   string                                             `json:"image_url"`
	ConsumerId string                                             `json:"consumer_id"`
	Status     string                                             `json:"status"`
	CreatedAt  time.Time                                          `json:"created_at"`
	UpdatedAt  time.Time                                          `json:"updated_at"`
	Groups     *[]PortalsConsumerProfilesAssignGroupsOutputGroups `json:"groups,omitempty"`
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
