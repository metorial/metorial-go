package app

import (
	"encoding/json"
	"time"
)

// PortalsAuthAppGetOutput represents the portals auth app get output type.
type PortalsAuthAppGetOutput struct {
	Object string `json:"object"`
	// Id - The Ares app identifier for this portal.
	Id string `json:"id"`
	// Slug - The Ares app slug.
	Slug           *string   `json:"slug,omitempty"`
	EmailWhitelist []string  `json:"email_whitelist"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// MapPortalsAuthAppGetOutputFromJSON deserializes JSON data into a PortalsAuthAppGetOutput.
func MapPortalsAuthAppGetOutputFromJSON(data []byte) (*PortalsAuthAppGetOutput, error) {
	var v PortalsAuthAppGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsAuthAppGetOutputToJSON serializes a PortalsAuthAppGetOutput to JSON.
func MapPortalsAuthAppGetOutputToJSON(v *PortalsAuthAppGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
