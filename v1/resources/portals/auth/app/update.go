package app

import (
	"encoding/json"
	"time"
)

// PortalsAuthAppUpdateOutput represents the portals auth app update output type.
type PortalsAuthAppUpdateOutput struct {
	Object string `json:"object"`
	// Id - The Ares app identifier for this portal.
	Id string `json:"id"`
	// Slug - The Ares app slug.
	Slug           *string   `json:"slug,omitempty"`
	EmailWhitelist []string  `json:"email_whitelist"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// MapPortalsAuthAppUpdateOutputFromJSON deserializes JSON data into a PortalsAuthAppUpdateOutput.
func MapPortalsAuthAppUpdateOutputFromJSON(data []byte) (*PortalsAuthAppUpdateOutput, error) {
	var v PortalsAuthAppUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsAuthAppUpdateOutputToJSON serializes a PortalsAuthAppUpdateOutput to JSON.
func MapPortalsAuthAppUpdateOutputToJSON(v *PortalsAuthAppUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsAuthAppUpdateBody represents the portals auth app update body type.
type PortalsAuthAppUpdateBody struct {
	EmailWhitelist *[]string `json:"email_whitelist,omitempty"`
}

// MapPortalsAuthAppUpdateBodyFromJSON deserializes JSON data into a PortalsAuthAppUpdateBody.
func MapPortalsAuthAppUpdateBodyFromJSON(data []byte) (*PortalsAuthAppUpdateBody, error) {
	var v PortalsAuthAppUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsAuthAppUpdateBodyToJSON serializes a PortalsAuthAppUpdateBody to JSON.
func MapPortalsAuthAppUpdateBodyToJSON(v *PortalsAuthAppUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
