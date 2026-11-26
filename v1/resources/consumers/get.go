package consumers

import (
	"encoding/json"
	"time"
)

// ConsumersGetOutput represents the consumers get output type.
type ConsumersGetOutput struct {
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
