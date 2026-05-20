package stores

import (
	"encoding/json"
	"time"
)

// StoresDeleteOutput represents the stores delete output type.
type StoresDeleteOutput struct {
	// Object - String representing the object's type
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Access    string    `json:"access"`
	ItemCount float64   `json:"item_count"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// MapStoresDeleteOutputFromJSON deserializes JSON data into a StoresDeleteOutput.
func MapStoresDeleteOutputFromJSON(data []byte) (*StoresDeleteOutput, error) {
	var v StoresDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapStoresDeleteOutputToJSON serializes a StoresDeleteOutput to JSON.
func MapStoresDeleteOutputToJSON(v *StoresDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
