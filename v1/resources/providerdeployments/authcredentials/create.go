package authcredentials

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsAuthCredentialsCreateOutput represents the provider deployments auth credentials create output type.
type ProviderDeploymentsAuthCredentialsCreateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - Must be `oauth`.
	Type string `json:"type"`
	// Status - Archived once the credentials are archived. Archived credentials stay listable.
	Status string `json:"status"`
	// IsDefault - Used when a request doesn't name credentials.
	IsDefault bool `json:"is_default"`
	// IsManaged - Whether Metorial or the installation administrator manages these OAuth credentials.
	IsManaged bool `json:"is_managed"`
	// ManagedBy - Who manages these credentials: Metorial, or the administrator of this self-hosted installation. Null for credentials created in the project.
	ManagedBy *string `json:"managed_by,omitempty"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Scopes - OAuth scopes requested by this credential
	Scopes *[]string `json:"scopes,omitempty"`
	// ProviderId - Provider these credentials belong to.
	ProviderId string `json:"provider_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapProviderDeploymentsAuthCredentialsCreateOutputFromJSON deserializes JSON data into a ProviderDeploymentsAuthCredentialsCreateOutput.
func MapProviderDeploymentsAuthCredentialsCreateOutputFromJSON(data []byte) (*ProviderDeploymentsAuthCredentialsCreateOutput, error) {
	var v ProviderDeploymentsAuthCredentialsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsAuthCredentialsCreateOutputToJSON serializes a ProviderDeploymentsAuthCredentialsCreateOutput to JSON.
func MapProviderDeploymentsAuthCredentialsCreateOutputToJSON(v *ProviderDeploymentsAuthCredentialsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderDeploymentsAuthCredentialsCreateBodyConfig - OAuth app credentials to store.
type ProviderDeploymentsAuthCredentialsCreateBodyConfig struct {
	// Type - Must be `oauth`.
	Type *string `json:"type,omitempty"`
	// ClientId - Client ID of your OAuth app.
	ClientId string `json:"client_id"`
	// ClientSecret - Client secret of your OAuth app.
	ClientSecret string `json:"client_secret"`
	// Scopes - OAuth scopes to request
	Scopes []string `json:"scopes"`
}

// ProviderDeploymentsAuthCredentialsCreateBody represents the provider deployments auth credentials create body type.
type ProviderDeploymentsAuthCredentialsCreateBody struct {
	// ProviderId - Provider to create credentials for.
	ProviderId string `json:"provider_id"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the provider auth credentials. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Config - OAuth app credentials to store.
	Config ProviderDeploymentsAuthCredentialsCreateBodyConfig `json:"config"`
}

// MapProviderDeploymentsAuthCredentialsCreateBodyFromJSON deserializes JSON data into a ProviderDeploymentsAuthCredentialsCreateBody.
func MapProviderDeploymentsAuthCredentialsCreateBodyFromJSON(data []byte) (*ProviderDeploymentsAuthCredentialsCreateBody, error) {
	var v ProviderDeploymentsAuthCredentialsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsAuthCredentialsCreateBodyToJSON serializes a ProviderDeploymentsAuthCredentialsCreateBody to JSON.
func MapProviderDeploymentsAuthCredentialsCreateBodyToJSON(v *ProviderDeploymentsAuthCredentialsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
