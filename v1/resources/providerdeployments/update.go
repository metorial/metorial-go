package providerdeployments

import (
	"encoding/json"
	"github.com/metorial/metorial-go/v1/nullable"
	"time"
)

// ProviderDeploymentsUpdateOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ProviderDeploymentsUpdateOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// ProviderDeploymentsUpdateOutputLockedVersion represents the provider deployments update output locked version type.
type ProviderDeploymentsUpdateOutputLockedVersion struct {
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

// ProviderDeploymentsUpdateOutputDefaultConfig represents the provider deployments update output default config type.
type ProviderDeploymentsUpdateOutputDefaultConfig struct {
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

// ProviderDeploymentsUpdateOutput represents the provider deployments update output type.
type ProviderDeploymentsUpdateOutput struct {
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
	ToolFilter ProviderDeploymentsUpdateOutputToolFilter `json:"tool_filter"`
	// ProviderId - Provider this deployment belongs to.
	ProviderId    string                                        `json:"provider_id"`
	LockedVersion *ProviderDeploymentsUpdateOutputLockedVersion `json:"locked_version,omitempty"`
	DefaultConfig *ProviderDeploymentsUpdateOutputDefaultConfig `json:"default_config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapProviderDeploymentsUpdateOutputFromJSON deserializes JSON data into a ProviderDeploymentsUpdateOutput.
func MapProviderDeploymentsUpdateOutputFromJSON(data []byte) (*ProviderDeploymentsUpdateOutput, error) {
	var v ProviderDeploymentsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsUpdateOutputToJSON serializes a ProviderDeploymentsUpdateOutput to JSON.
func MapProviderDeploymentsUpdateOutputToJSON(v *ProviderDeploymentsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderDeploymentsUpdateBody represents the provider deployments update body type.
type ProviderDeploymentsUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the provider deployment. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
	// LockedProviderVersionId - Pin this deployment to a specific provider version, or null to follow latest
	LockedProviderVersionId *nullable.Nullable[string] `json:"locked_provider_version_id,omitempty"`
}

// MapProviderDeploymentsUpdateBodyFromJSON deserializes JSON data into a ProviderDeploymentsUpdateBody.
func MapProviderDeploymentsUpdateBodyFromJSON(data []byte) (*ProviderDeploymentsUpdateBody, error) {
	var v ProviderDeploymentsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsUpdateBodyToJSON serializes a ProviderDeploymentsUpdateBody to JSON.
func MapProviderDeploymentsUpdateBodyToJSON(v *ProviderDeploymentsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
