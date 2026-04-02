package consumers

import (
	"encoding/json"
	"time"
)

// ConsumersGetMemberConsumerOutput represents the consumers get member consumer output type.
type ConsumersGetMemberConsumerOutput struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
	SurfaceIdentifier string `json:"surface_identifier"`
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
