package authcredentials

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsAuthCredentialsGetOutput represents the provider deployments auth credentials get output type.
type ProviderDeploymentsAuthCredentialsGetOutput struct {
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

// MapProviderDeploymentsAuthCredentialsGetOutputFromJSON deserializes JSON data into a ProviderDeploymentsAuthCredentialsGetOutput.
func MapProviderDeploymentsAuthCredentialsGetOutputFromJSON(data []byte) (*ProviderDeploymentsAuthCredentialsGetOutput, error) {
	var v ProviderDeploymentsAuthCredentialsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsAuthCredentialsGetOutputToJSON serializes a ProviderDeploymentsAuthCredentialsGetOutput to JSON.
func MapProviderDeploymentsAuthCredentialsGetOutputToJSON(v *ProviderDeploymentsAuthCredentialsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
