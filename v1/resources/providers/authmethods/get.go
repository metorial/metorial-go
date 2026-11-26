package authmethods

import (
	"encoding/json"
	"time"
)

// ProvidersAuthMethodsGetOutputInputSchema represents the providers auth methods get output input schema type.
type ProvidersAuthMethodsGetOutputInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// ProvidersAuthMethodsGetOutputOutputSchema represents the providers auth methods get output output schema type.
type ProvidersAuthMethodsGetOutputOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// ProvidersAuthMethodsGetOutputScopes represents the providers auth methods get output scopes type.
type ProvidersAuthMethodsGetOutputScopes struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Scope - The scope value sent to the provider.
	Scope string `json:"scope"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
}

// ProvidersAuthMethodsGetOutput represents the providers auth methods get output type.
type ProvidersAuthMethodsGetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - How the user authenticates with the provider.
	Type string `json:"type"`
	// Key - Stable key for the auth method.
	Key string `json:"key"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Capabilities - What this auth method supports.
	Capabilities map[string]any                             `json:"capabilities"`
	InputSchema  *ProvidersAuthMethodsGetOutputInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *ProvidersAuthMethodsGetOutputOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]ProvidersAuthMethodsGetOutputScopes `json:"scopes,omitempty"`
	// Adapters - Adapter IDs this auth method can be used with
	Adapters *[]string `json:"adapters,omitempty"`
	// ProviderId - Provider this auth method belongs to.
	ProviderId string `json:"provider_id"`
	// ProviderSpecificationId - Specification this auth method belongs to.
	ProviderSpecificationId string `json:"provider_specification_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapProvidersAuthMethodsGetOutputFromJSON deserializes JSON data into a ProvidersAuthMethodsGetOutput.
func MapProvidersAuthMethodsGetOutputFromJSON(data []byte) (*ProvidersAuthMethodsGetOutput, error) {
	var v ProvidersAuthMethodsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProvidersAuthMethodsGetOutputToJSON serializes a ProvidersAuthMethodsGetOutput to JSON.
func MapProvidersAuthMethodsGetOutputToJSON(v *ProvidersAuthMethodsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
