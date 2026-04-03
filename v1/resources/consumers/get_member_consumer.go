package consumers

import (
	"encoding/json"
	"time"
)

// ConsumersGetMemberConsumerOutputProfileGroupsGroup represents the consumers get member consumer output profile groups group type.
type ConsumersGetMemberConsumerOutputProfileGroupsGroup struct {
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

// ConsumersGetMemberConsumerOutputProfileGroups represents the consumers get member consumer output profile groups type.
type ConsumersGetMemberConsumerOutputProfileGroups struct {
	Object      string                                             `json:"object"`
	Group       ConsumersGetMemberConsumerOutputProfileGroupsGroup `json:"group"`
	AssignedVia string                                             `json:"assigned_via"`
}

// ConsumersGetMemberConsumerOutputProfile represents the consumers get member consumer output profile type.
type ConsumersGetMemberConsumerOutputProfile struct {
	Object     string                                           `json:"object"`
	Id         string                                           `json:"id"`
	Name       string                                           `json:"name"`
	Email      string                                           `json:"email"`
	ImageUrl   string                                           `json:"image_url"`
	Groups     *[]ConsumersGetMemberConsumerOutputProfileGroups `json:"groups,omitempty"`
	ConsumerId string                                           `json:"consumer_id"`
	CreatedAt  time.Time                                        `json:"created_at"`
	UpdatedAt  time.Time                                        `json:"updated_at"`
}

// ConsumersGetMemberConsumerOutput represents the consumers get member consumer output type.
type ConsumersGetMemberConsumerOutput struct {
	Object    string                                  `json:"object"`
	Id        string                                  `json:"id"`
	Name      string                                  `json:"name"`
	Email     string                                  `json:"email"`
	CreatedAt time.Time                               `json:"created_at"`
	UpdatedAt time.Time                               `json:"updated_at"`
	Profile   ConsumersGetMemberConsumerOutputProfile `json:"profile"`
}

// MapConsumersGetMemberConsumerOutputFromJSON deserializes JSON data into a ConsumersGetMemberConsumerOutput.
func MapConsumersGetMemberConsumerOutputFromJSON(data []byte) (*ConsumersGetMemberConsumerOutput, error) {
	var v ConsumersGetMemberConsumerOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConsumersGetMemberConsumerOutputToJSON serializes a ConsumersGetMemberConsumerOutput to JSON.
func MapConsumersGetMemberConsumerOutputToJSON(v *ConsumersGetMemberConsumerOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ConsumersGetMemberConsumerBody represents the consumers get member consumer body type.
type ConsumersGetMemberConsumerBody struct {
	SurfaceIdentifier *string `json:"surface_identifier,omitempty"`
}

// MapConsumersGetMemberConsumerBodyFromJSON deserializes JSON data into a ConsumersGetMemberConsumerBody.
func MapConsumersGetMemberConsumerBodyFromJSON(data []byte) (*ConsumersGetMemberConsumerBody, error) {
	var v ConsumersGetMemberConsumerBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConsumersGetMemberConsumerBodyToJSON serializes a ConsumersGetMemberConsumerBody to JSON.
func MapConsumersGetMemberConsumerBodyToJSON(v *ConsumersGetMemberConsumerBody) ([]byte, error) {
	return json.Marshal(v)
}
