package stores

import (
	"encoding/json"
	"time"
)

// StoresCreateOutput represents the stores create output type.
type StoresCreateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Access - Who else can access the store: `private`, `public_read`, or `public_write`.
	Access string `json:"access"`
	// ItemCount - Number of items in the store.
	ItemCount float64 `json:"item_count"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
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
	// Name - Display name for the store.
	Name string `json:"name"`
	// Access - Who else can access the store: `private`, `public_read`, or `public_write`.
	Access *string `json:"access,omitempty"`
	// TemplateId - Template to create the store from, by ID. Cannot be combined with `parent_id`.
	TemplateId *string `json:"template_id,omitempty"`
	// ParentId - Store to copy, by ID. Cannot be combined with `template_id`.
	ParentId *string `json:"parent_id,omitempty"`
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
