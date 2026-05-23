package consumerprofiles

import (
	"encoding/json"
	"time"
)

// PortalsConsumerProfilesCreateOutputGroupsGroup represents the portals consumer profiles create output groups group type.
type PortalsConsumerProfilesCreateOutputGroupsGroup struct {
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

// PortalsConsumerProfilesCreateOutputGroups represents the portals consumer profiles create output groups type.
type PortalsConsumerProfilesCreateOutputGroups struct {
	Object      string                                         `json:"object"`
	Group       PortalsConsumerProfilesCreateOutputGroupsGroup `json:"group"`
	AssignedVia string                                         `json:"assigned_via"`
}

// PortalsConsumerProfilesCreateOutput represents the portals consumer profiles create output type.
type PortalsConsumerProfilesCreateOutput struct {
	Object     string                                       `json:"object"`
	Id         string                                       `json:"id"`
	Name       string                                       `json:"name"`
	Email      string                                       `json:"email"`
	ImageUrl   string                                       `json:"image_url"`
	ConsumerId string                                       `json:"consumer_id"`
	Status     string                                       `json:"status"`
	CreatedAt  time.Time                                    `json:"created_at"`
	UpdatedAt  time.Time                                    `json:"updated_at"`
	Groups     *[]PortalsConsumerProfilesCreateOutputGroups `json:"groups,omitempty"`
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
	Email string `json:"email"`
	Name  string `json:"name"`
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
