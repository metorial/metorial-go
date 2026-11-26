package authcredentials

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsAuthCredentialsUpdateOutput represents the provider deployments auth credentials update output type.
type ProviderDeploymentsAuthCredentialsUpdateOutput struct {
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

// MapProviderDeploymentsAuthCredentialsUpdateOutputFromJSON deserializes JSON data into a ProviderDeploymentsAuthCredentialsUpdateOutput.
func MapProviderDeploymentsAuthCredentialsUpdateOutputFromJSON(data []byte) (*ProviderDeploymentsAuthCredentialsUpdateOutput, error) {
	var v ProviderDeploymentsAuthCredentialsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsAuthCredentialsUpdateOutputToJSON serializes a ProviderDeploymentsAuthCredentialsUpdateOutput to JSON.
func MapProviderDeploymentsAuthCredentialsUpdateOutputToJSON(v *ProviderDeploymentsAuthCredentialsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderDeploymentsAuthCredentialsUpdateBody represents the provider deployments auth credentials update body type.
type ProviderDeploymentsAuthCredentialsUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the provider auth credentials. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ClientId - Client ID of your OAuth app.
	ClientId *string `json:"client_id,omitempty"`
	// ClientSecret - Client secret of your OAuth app.
	ClientSecret *string `json:"client_secret,omitempty"`
	// Scopes - OAuth scopes to request.
	Scopes *[]string `json:"scopes,omitempty"`
}

// MapProviderDeploymentsAuthCredentialsUpdateBodyFromJSON deserializes JSON data into a ProviderDeploymentsAuthCredentialsUpdateBody.
func MapProviderDeploymentsAuthCredentialsUpdateBodyFromJSON(data []byte) (*ProviderDeploymentsAuthCredentialsUpdateBody, error) {
	var v ProviderDeploymentsAuthCredentialsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsAuthCredentialsUpdateBodyToJSON serializes a ProviderDeploymentsAuthCredentialsUpdateBody to JSON.
func MapProviderDeploymentsAuthCredentialsUpdateBodyToJSON(v *ProviderDeploymentsAuthCredentialsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
