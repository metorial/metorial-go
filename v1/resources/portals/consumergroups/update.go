package consumergroups

import (
	"encoding/json"
	"time"
)

// PortalsConsumerGroupsUpdateOutput represents the portals consumer groups update output type.
type PortalsConsumerGroupsUpdateOutput struct {
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

// MapPortalsConsumerGroupsUpdateOutputFromJSON deserializes JSON data into a PortalsConsumerGroupsUpdateOutput.
func MapPortalsConsumerGroupsUpdateOutputFromJSON(data []byte) (*PortalsConsumerGroupsUpdateOutput, error) {
	var v PortalsConsumerGroupsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerGroupsUpdateOutputToJSON serializes a PortalsConsumerGroupsUpdateOutput to JSON.
func MapPortalsConsumerGroupsUpdateOutputToJSON(v *PortalsConsumerGroupsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsConsumerGroupsUpdateBody represents the portals consumer groups update body type.
type PortalsConsumerGroupsUpdateBody struct {
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
	SsoGroupIds *[]string `json:"sso_group_ids,omitempty"`
	IsDefault   *bool     `json:"is_default,omitempty"`
}

// MapPortalsConsumerGroupsUpdateBodyFromJSON deserializes JSON data into a PortalsConsumerGroupsUpdateBody.
func MapPortalsConsumerGroupsUpdateBodyFromJSON(data []byte) (*PortalsConsumerGroupsUpdateBody, error) {
	var v PortalsConsumerGroupsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerGroupsUpdateBodyToJSON serializes a PortalsConsumerGroupsUpdateBody to JSON.
func MapPortalsConsumerGroupsUpdateBodyToJSON(v *PortalsConsumerGroupsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
