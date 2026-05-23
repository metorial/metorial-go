package versions

import (
	"encoding/json"
)

// CustomProvidersVersionsGetEnvOutput represents the custom providers versions get env output type.
type CustomProvidersVersionsGetEnvOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Env - Key-value pairs representing the custom provider environment variables
	Env *map[string]any `json:"env,omitempty"`
}

// MapCustomProvidersVersionsGetEnvOutputFromJSON deserializes JSON data into a CustomProvidersVersionsGetEnvOutput.
func MapCustomProvidersVersionsGetEnvOutputFromJSON(data []byte) (*CustomProvidersVersionsGetEnvOutput, error) {
	var v CustomProvidersVersionsGetEnvOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCustomProvidersVersionsGetEnvOutputToJSON serializes a CustomProvidersVersionsGetEnvOutput to JSON.
func MapCustomProvidersVersionsGetEnvOutputToJSON(v *CustomProvidersVersionsGetEnvOutput) ([]byte, error) {
	return json.Marshal(v)
}
