package providerdeployments

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsCreateOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ProviderDeploymentsCreateOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// ProviderDeploymentsCreateOutputLockedVersion represents the provider deployments create output locked version type.
type ProviderDeploymentsCreateOutputLockedVersion struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Version - Release label for the version.
	Version string `json:"version"`
	// ProviderId - Provider this version belongs to.
	ProviderId string `json:"provider_id"`
	// IsCurrent - The version deployments follow unless locked to another.
	IsCurrent bool `json:"is_current"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// SpecificationId - Specification this version belongs to.
	SpecificationId *string `json:"specification_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsCreateOutputDefaultConfig represents the provider deployments create output default config type.
type ProviderDeploymentsCreateOutputDefaultConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsDefault - The deployment falls back to this config when none is specified.
	IsDefault bool `json:"is_default"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider this config belongs to.
	ProviderId string `json:"provider_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsCreateOutput represents the provider deployments create output type.
type ProviderDeploymentsCreateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Archived once the deployment is archived. Archived deployments stay listable.
	Status string `json:"status"`
	// IsDefault - Used when a request doesn't name a deployment.
	IsDefault bool `json:"is_default"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter ProviderDeploymentsCreateOutputToolFilter `json:"tool_filter"`
	// ProviderId - Provider this deployment belongs to.
	ProviderId    string                                        `json:"provider_id"`
	LockedVersion *ProviderDeploymentsCreateOutputLockedVersion `json:"locked_version,omitempty"`
	DefaultConfig *ProviderDeploymentsCreateOutputDefaultConfig `json:"default_config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapProviderDeploymentsCreateOutputFromJSON deserializes JSON data into a ProviderDeploymentsCreateOutput.
func MapProviderDeploymentsCreateOutputFromJSON(data []byte) (*ProviderDeploymentsCreateOutput, error) {
	var v ProviderDeploymentsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsCreateOutputToJSON serializes a ProviderDeploymentsCreateOutput to JSON.
func MapProviderDeploymentsCreateOutputToJSON(v *ProviderDeploymentsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderDeploymentsCreateBodyProviderConfig represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ProviderDeploymentsCreateBodyProviderConfig struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Value - Provider-specific configuration values
	Value *map[string]any `json:"value,omitempty"`
	// ProviderConfigVaultId - Vault to take the values from.
	ProviderConfigVaultId *string `json:"provider_config_vault_id,omitempty"`
	// ProviderId - Explicitly specify provider ID to associate with this config
	ProviderId *string `json:"provider_id,omitempty"`
}

// ProviderDeploymentsCreateBody represents the provider deployments create body type.
type ProviderDeploymentsCreateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the provider deployment. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
	// ProviderId - The provider to deploy
	ProviderId string `json:"provider_id"`
	// LockedProviderVersionId - Pin this deployment to a specific provider version
	LockedProviderVersionId *string `json:"locked_provider_version_id,omitempty"`
	// ProviderConfigId - ID of an existing config to use.
	ProviderConfigId *string                                      `json:"provider_config_id,omitempty"`
	ProviderConfig   *ProviderDeploymentsCreateBodyProviderConfig `json:"provider_config,omitempty"`
}

// MapProviderDeploymentsCreateBodyFromJSON deserializes JSON data into a ProviderDeploymentsCreateBody.
func MapProviderDeploymentsCreateBodyFromJSON(data []byte) (*ProviderDeploymentsCreateBody, error) {
	var v ProviderDeploymentsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsCreateBodyToJSON serializes a ProviderDeploymentsCreateBody to JSON.
func MapProviderDeploymentsCreateBodyToJSON(v *ProviderDeploymentsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
