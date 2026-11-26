package providerdeployments

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsDeleteOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ProviderDeploymentsDeleteOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// ProviderDeploymentsDeleteOutputLockedVersion represents the provider deployments delete output locked version type.
type ProviderDeploymentsDeleteOutputLockedVersion struct {
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

// ProviderDeploymentsDeleteOutputDefaultConfig represents the provider deployments delete output default config type.
type ProviderDeploymentsDeleteOutputDefaultConfig struct {
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

// ProviderDeploymentsDeleteOutput represents the provider deployments delete output type.
type ProviderDeploymentsDeleteOutput struct {
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
	ToolFilter ProviderDeploymentsDeleteOutputToolFilter `json:"tool_filter"`
	// ProviderId - Provider this deployment belongs to.
	ProviderId    string                                        `json:"provider_id"`
	LockedVersion *ProviderDeploymentsDeleteOutputLockedVersion `json:"locked_version,omitempty"`
	DefaultConfig *ProviderDeploymentsDeleteOutputDefaultConfig `json:"default_config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapProviderDeploymentsDeleteOutputFromJSON deserializes JSON data into a ProviderDeploymentsDeleteOutput.
func MapProviderDeploymentsDeleteOutputFromJSON(data []byte) (*ProviderDeploymentsDeleteOutput, error) {
	var v ProviderDeploymentsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsDeleteOutputToJSON serializes a ProviderDeploymentsDeleteOutput to JSON.
func MapProviderDeploymentsDeleteOutputToJSON(v *ProviderDeploymentsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
