package consumergroups

import (
	"encoding/json"
	"time"
)

// PortalsConsumerGroupsCreateOutput represents the portals consumer groups create output type.
type PortalsConsumerGroupsCreateOutput struct {
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

// MapPortalsConsumerGroupsCreateOutputFromJSON deserializes JSON data into a PortalsConsumerGroupsCreateOutput.
func MapPortalsConsumerGroupsCreateOutputFromJSON(data []byte) (*PortalsConsumerGroupsCreateOutput, error) {
	var v PortalsConsumerGroupsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerGroupsCreateOutputToJSON serializes a PortalsConsumerGroupsCreateOutput to JSON.
func MapPortalsConsumerGroupsCreateOutputToJSON(v *PortalsConsumerGroupsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsConsumerGroupsCreateBody represents the portals consumer groups create body type.
type PortalsConsumerGroupsCreateBody struct {
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	SsoGroupIds *[]string `json:"sso_group_ids,omitempty"`
	IsDefault   *bool     `json:"is_default,omitempty"`
}

// MapPortalsConsumerGroupsCreateBodyFromJSON deserializes JSON data into a PortalsConsumerGroupsCreateBody.
func MapPortalsConsumerGroupsCreateBodyFromJSON(data []byte) (*PortalsConsumerGroupsCreateBody, error) {
	var v PortalsConsumerGroupsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerGroupsCreateBodyToJSON serializes a PortalsConsumerGroupsCreateBody to JSON.
func MapPortalsConsumerGroupsCreateBodyToJSON(v *PortalsConsumerGroupsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
