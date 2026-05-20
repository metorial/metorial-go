package stores

import (
	"encoding/json"
	"time"
)

// StoresUpdateOutput represents the stores update output type.
type StoresUpdateOutput struct {
	// Object - String representing the object's type
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Access    string    `json:"access"`
	ItemCount float64   `json:"item_count"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// MapStoresUpdateOutputFromJSON deserializes JSON data into a StoresUpdateOutput.
func MapStoresUpdateOutputFromJSON(data []byte) (*StoresUpdateOutput, error) {
	var v StoresUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapStoresUpdateOutputToJSON serializes a StoresUpdateOutput to JSON.
func MapStoresUpdateOutputToJSON(v *StoresUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// StoresUpdateBody represents the stores update body type.
type StoresUpdateBody struct {
	Name   *string `json:"name,omitempty"`
	Access *string `json:"access,omitempty"`
}

// MapStoresUpdateBodyFromJSON deserializes JSON data into a StoresUpdateBody.
func MapStoresUpdateBodyFromJSON(data []byte) (*StoresUpdateBody, error) {
	var v StoresUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapStoresUpdateBodyToJSON serializes a StoresUpdateBody to JSON.
func MapStoresUpdateBodyToJSON(v *StoresUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
