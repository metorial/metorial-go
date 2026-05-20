package stores

import (
	"encoding/json"
	"time"
)

// StoresGetOutput represents the stores get output type.
type StoresGetOutput struct {
	// Object - String representing the object's type
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Access    string    `json:"access"`
	ItemCount float64   `json:"item_count"`
	CreatedAt time.Time `json:"created_at"`
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
