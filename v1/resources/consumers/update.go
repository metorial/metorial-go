package consumers

import (
	"encoding/json"
	"time"
)

// ConsumersUpdateOutput represents the consumers update output type.
type ConsumersUpdateOutput struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// MapConsumersUpdateOutputFromJSON deserializes JSON data into a ConsumersUpdateOutput.
func MapConsumersUpdateOutputFromJSON(data []byte) (*ConsumersUpdateOutput, error) {
	var v ConsumersUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConsumersUpdateOutputToJSON serializes a ConsumersUpdateOutput to JSON.
func MapConsumersUpdateOutputToJSON(v *ConsumersUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ConsumersUpdateBody represents the consumers update body type.
type ConsumersUpdateBody struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

// MapConsumersUpdateBodyFromJSON deserializes JSON data into a ConsumersUpdateBody.
func MapConsumersUpdateBodyFromJSON(data []byte) (*ConsumersUpdateBody, error) {
	var v ConsumersUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConsumersUpdateBodyToJSON serializes a ConsumersUpdateBody to JSON.
func MapConsumersUpdateBodyToJSON(v *ConsumersUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
