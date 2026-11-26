package consumers

import (
	"encoding/json"
	"time"
)

// ConsumersCreateOutput represents the consumers create output type.
type ConsumersCreateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Email - Email address.
	Email string `json:"email"`
	// ImageUrl - Avatar image URL.
	ImageUrl string `json:"image_url"`
	// UserId - User this consumer is linked to. `null` when unlinked.
	UserId *string `json:"user_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapConsumersCreateOutputFromJSON deserializes JSON data into a ConsumersCreateOutput.
func MapConsumersCreateOutputFromJSON(data []byte) (*ConsumersCreateOutput, error) {
	var v ConsumersCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConsumersCreateOutputToJSON serializes a ConsumersCreateOutput to JSON.
func MapConsumersCreateOutputToJSON(v *ConsumersCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ConsumersCreateBody represents the consumers create body type.
type ConsumersCreateBody struct {
	// Name - Display name.
	Name string `json:"name"`
	// Email - Email address. An existing consumer with this email is linked to the instance.
	Email string `json:"email"`
}

// MapConsumersCreateBodyFromJSON deserializes JSON data into a ConsumersCreateBody.
func MapConsumersCreateBodyFromJSON(data []byte) (*ConsumersCreateBody, error) {
	var v ConsumersCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConsumersCreateBodyToJSON serializes a ConsumersCreateBody to JSON.
func MapConsumersCreateBodyToJSON(v *ConsumersCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
