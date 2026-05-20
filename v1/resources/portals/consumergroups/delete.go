package consumergroups

import (
	"encoding/json"
	"time"
)

// PortalsConsumerGroupsDeleteOutput represents the portals consumer groups delete output type.
type PortalsConsumerGroupsDeleteOutput struct {
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

// MapPortalsConsumerGroupsDeleteOutputFromJSON deserializes JSON data into a PortalsConsumerGroupsDeleteOutput.
func MapPortalsConsumerGroupsDeleteOutputFromJSON(data []byte) (*PortalsConsumerGroupsDeleteOutput, error) {
	var v PortalsConsumerGroupsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerGroupsDeleteOutputToJSON serializes a PortalsConsumerGroupsDeleteOutput to JSON.
func MapPortalsConsumerGroupsDeleteOutputToJSON(v *PortalsConsumerGroupsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
