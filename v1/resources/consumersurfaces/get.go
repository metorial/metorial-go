package consumersurfaces

import (
	"encoding/json"
	"time"
)

// ConsumerSurfacesGetOutputAuth represents the consumer surfaces get output auth type.
type ConsumerSurfacesGetOutputAuth struct {
	Object                     string  `json:"object"`
	SessionExpiryTimeInSeconds float64 `json:"session_expiry_time_in_seconds"`
}

// ConsumerSurfacesGetOutput represents the consumer surfaces get output type.
type ConsumerSurfacesGetOutput struct {
	Object      string                        `json:"object"`
	Id          string                        `json:"id"`
	Status      string                        `json:"status"`
	Name        string                        `json:"name"`
	Description *string                       `json:"description,omitempty"`
	Auth        ConsumerSurfacesGetOutputAuth `json:"auth"`
	CreatedAt   time.Time                     `json:"created_at"`
	UpdatedAt   time.Time                     `json:"updated_at"`
}

// MapConsumerSurfacesGetOutputFromJSON deserializes JSON data into a ConsumerSurfacesGetOutput.
func MapConsumerSurfacesGetOutputFromJSON(data []byte) (*ConsumerSurfacesGetOutput, error) {
	var v ConsumerSurfacesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConsumerSurfacesGetOutputToJSON serializes a ConsumerSurfacesGetOutput to JSON.
func MapConsumerSurfacesGetOutputToJSON(v *ConsumerSurfacesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
