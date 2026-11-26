package stores

import (
	"encoding/json"
	"time"
)

// StoresGetOutput represents the stores get output type.
type StoresGetOutput struct {
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

// MapStoresGetOutputFromJSON deserializes JSON data into a StoresGetOutput.
func MapStoresGetOutputFromJSON(data []byte) (*StoresGetOutput, error) {
	var v StoresGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapStoresGetOutputToJSON serializes a StoresGetOutput to JSON.
func MapStoresGetOutputToJSON(v *StoresGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
