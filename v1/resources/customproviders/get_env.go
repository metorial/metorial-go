package customproviders

import (
	"encoding/json"
)

// CustomProvidersGetEnvOutput represents the custom providers get env output type.
type CustomProvidersGetEnvOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Env - Key-value pairs representing the custom provider environment variables
	Env *map[string]any `json:"env,omitempty"`
}

// MapCustomProvidersGetEnvOutputFromJSON deserializes JSON data into a CustomProvidersGetEnvOutput.
func MapCustomProvidersGetEnvOutputFromJSON(data []byte) (*CustomProvidersGetEnvOutput, error) {
	var v CustomProvidersGetEnvOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCustomProvidersGetEnvOutputToJSON serializes a CustomProvidersGetEnvOutput to JSON.
func MapCustomProvidersGetEnvOutputToJSON(v *CustomProvidersGetEnvOutput) ([]byte, error) {
	return json.Marshal(v)
}
