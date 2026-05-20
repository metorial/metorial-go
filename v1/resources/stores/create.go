package stores

import (
	"encoding/json"
	"time"
)

// StoresCreateOutput represents the stores create output type.
type StoresCreateOutput struct {
	// Object - String representing the object's type
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Access    string    `json:"access"`
	ItemCount float64   `json:"item_count"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// MapStoresCreateOutputFromJSON deserializes JSON data into a StoresCreateOutput.
func MapStoresCreateOutputFromJSON(data []byte) (*StoresCreateOutput, error) {
	var v StoresCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapStoresCreateOutputToJSON serializes a StoresCreateOutput to JSON.
func MapStoresCreateOutputToJSON(v *StoresCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// StoresCreateBody represents the stores create body type.
type StoresCreateBody struct {
	Name       string  `json:"name"`
	Access     *string `json:"access,omitempty"`
	TemplateId *string `json:"template_id,omitempty"`
	ParentId   *string `json:"parent_id,omitempty"`
}

// MapStoresCreateBodyFromJSON deserializes JSON data into a StoresCreateBody.
func MapStoresCreateBodyFromJSON(data []byte) (*StoresCreateBody, error) {
	var v StoresCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapStoresCreateBodyToJSON serializes a StoresCreateBody to JSON.
func MapStoresCreateBodyToJSON(v *StoresCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
