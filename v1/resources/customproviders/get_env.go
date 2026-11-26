package customproviders

import (
	"encoding/json"
)

// CustomProvidersGetEnvOutput represents the custom providers get env output type.
type CustomProvidersGetEnvOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Env - Environment variables available to the function. `null` unless the provider is a `function` provider.
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
