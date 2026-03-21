package token

import (
	"encoding/json"
)

// TokenGetOutput represents the token get output type.
type TokenGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Type - The token's type
	Type string `json:"type"`
}

// MapTokenGetOutputFromJSON deserializes JSON data into a TokenGetOutput.
func MapTokenGetOutputFromJSON(data []byte) (*TokenGetOutput, error) {
	var v TokenGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapTokenGetOutputToJSON serializes a TokenGetOutput to JSON.
func MapTokenGetOutputToJSON(v *TokenGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
