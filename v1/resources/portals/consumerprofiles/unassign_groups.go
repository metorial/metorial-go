package consumerprofiles

import (
	"encoding/json"
	"time"
)

// PortalsConsumerProfilesUnassignGroupsOutputGroupsGroup represents the portals consumer profiles unassign groups output groups group type.
type PortalsConsumerProfilesUnassignGroupsOutputGroupsGroup struct {
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

// PortalsConsumerProfilesUnassignGroupsOutputGroups represents the portals consumer profiles unassign groups output groups type.
type PortalsConsumerProfilesUnassignGroupsOutputGroups struct {
	Object      string                                                 `json:"object"`
	Group       PortalsConsumerProfilesUnassignGroupsOutputGroupsGroup `json:"group"`
	AssignedVia string                                                 `json:"assigned_via"`
}

// PortalsConsumerProfilesUnassignGroupsOutput represents the portals consumer profiles unassign groups output type.
type PortalsConsumerProfilesUnassignGroupsOutput struct {
	Object     string                                               `json:"object"`
	Id         string                                               `json:"id"`
	Name       string                                               `json:"name"`
	Email      string                                               `json:"email"`
	ImageUrl   string                                               `json:"image_url"`
	ConsumerId string                                               `json:"consumer_id"`
	Status     string                                               `json:"status"`
	CreatedAt  time.Time                                            `json:"created_at"`
	UpdatedAt  time.Time                                            `json:"updated_at"`
	Groups     *[]PortalsConsumerProfilesUnassignGroupsOutputGroups `json:"groups,omitempty"`
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
