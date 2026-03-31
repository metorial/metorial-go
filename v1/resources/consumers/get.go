package consumers

import (
	"encoding/json"
	"time"
)

// ConsumersGetOutput represents the consumers get output type.
type ConsumersGetOutput struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// MapConsumersGetOutputFromJSON deserializes JSON data into a ConsumersGetOutput.
func MapConsumersGetOutputFromJSON(data []byte) (*ConsumersGetOutput, error) {
	var v ConsumersGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConsumersGetOutputToJSON serializes a ConsumersGetOutput to JSON.
func MapConsumersGetOutputToJSON(v *ConsumersGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
