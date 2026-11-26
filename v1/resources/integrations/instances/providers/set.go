package providers

import (
	"encoding/json"
	"time"
)

// IntegrationsInstancesProvidersSetOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstancesProvidersSetOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstancesProvidersSetOutputProvider represents the integrations instances providers set output provider type.
type IntegrationsInstancesProvidersSetOutputProvider struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Slug - URL-safe name for the provider.
	Slug string `json:"slug"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// IntegrationsInstancesProvidersSetOutputIntegrationProviderProviderVersion - Version this snapshot was taken from.
type IntegrationsInstancesProvidersSetOutputIntegrationProviderProviderVersion struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Index - Position of this version, starting at 0.
	Index float64 `json:"index"`
}

// IntegrationsInstancesProvidersSetOutputIntegrationProviderToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstancesProvidersSetOutputIntegrationProviderToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstancesProvidersSetOutputIntegrationProviderConfig represents the integrations instances providers set output integration provider config type.
type IntegrationsInstancesProvidersSetOutputIntegrationProviderConfig struct {
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

// IntegrationsInstancesProvidersSetOutputIntegrationProvider represents the integrations instances providers set output integration provider type.
type IntegrationsInstancesProvidersSetOutputIntegrationProvider struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// ProviderVersion - Version this snapshot was taken from.
	ProviderVersion IntegrationsInstancesProvidersSetOutputIntegrationProviderProviderVersion `json:"provider_version"`
	// Status - `archived` once you archive the integration provider.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter *IntegrationsInstancesProvidersSetOutputIntegrationProviderToolFilter `json:"tool_filter,omitempty"`
	// ProviderId - Provider this snapshot was taken from.
	ProviderId string `json:"provider_id"`
	// DeploymentId - Deployment this snapshot was taken from.
	DeploymentId string `json:"deployment_id"`
	// AuthMethodId - Auth method at snapshot time. `null` when none was set.
	AuthMethodId *string `json:"auth_method_id,omitempty"`
	// AuthCredentialsId - Credentials at snapshot time. `null` when none were set.
	AuthCredentialsId *string                                                           `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsInstancesProvidersSetOutputIntegrationProviderConfig `json:"config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsInstancesProvidersSetOutputConfig represents the integrations instances providers set output config type.
type IntegrationsInstancesProvidersSetOutputConfig struct {
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

// IntegrationsInstancesProvidersSetOutputAuthConfig represents the integrations instances providers set output auth config type.
type IntegrationsInstancesProvidersSetOutputAuthConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsDefault - Used when a request doesn't name an auth config.
	IsDefault bool `json:"is_default"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider this auth config belongs to.
	ProviderId string `json:"provider_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// IntegrationsInstancesProvidersSetOutput represents the integrations instances providers set output type.
type IntegrationsInstancesProvidersSetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the entry is archived.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// IntegrationId - Integration the instance belongs to.
	IntegrationId string `json:"integration_id"`
	// IntegrationInstanceId - Integration instance this entry belongs to.
	IntegrationInstanceId string `json:"integration_instance_id"`
	// ToolFilter - Tool filter configuration
	ToolFilter *IntegrationsInstancesProvidersSetOutputToolFilter `json:"tool_filter,omitempty"`
	// IsOverrideToolFilter - Replaces the integration tool filters instead of combining with them.
	IsOverrideToolFilter bool                                                       `json:"is_override_tool_filter"`
	Provider             IntegrationsInstancesProvidersSetOutputProvider            `json:"provider"`
	IntegrationProvider  IntegrationsInstancesProvidersSetOutputIntegrationProvider `json:"integration_provider"`
	Config               *IntegrationsInstancesProvidersSetOutputConfig             `json:"config,omitempty"`
	AuthConfig           *IntegrationsInstancesProvidersSetOutputAuthConfig         `json:"auth_config,omitempty"`
	// CallbackInstanceId - The active callback instance registered for this integration instance provider, if its integration provider has callbacks enabled.
	CallbackInstanceId *string `json:"callback_instance_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// MapIntegrationsInstancesProvidersSetOutputFromJSON deserializes JSON data into a IntegrationsInstancesProvidersSetOutput.
func MapIntegrationsInstancesProvidersSetOutputFromJSON(data []byte) (*IntegrationsInstancesProvidersSetOutput, error) {
	var v IntegrationsInstancesProvidersSetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstancesProvidersSetOutputToJSON serializes a IntegrationsInstancesProvidersSetOutput to JSON.
func MapIntegrationsInstancesProvidersSetOutputToJSON(v *IntegrationsInstancesProvidersSetOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsInstancesProvidersSetBody represents the integrations instances providers set body type.
type IntegrationsInstancesProvidersSetBody struct {
	// ProviderDeploymentId - Deployment to run with. Omitted uses the integration provider’s deployment.
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Provider config to run with. Omitted keeps the current one, or the shared config when none exists.
	ProviderConfigId *string `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Credentials to run with. Omitted keeps the current ones.
	ProviderAuthConfigId *string `json:"provider_auth_config_id,omitempty"`
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
	// IsOverrideToolFilter - Replaces the integration tool filters instead of combining with them.
	IsOverrideToolFilter *bool `json:"is_override_tool_filter,omitempty"`
}

// MapIntegrationsInstancesProvidersSetBodyFromJSON deserializes JSON data into a IntegrationsInstancesProvidersSetBody.
func MapIntegrationsInstancesProvidersSetBodyFromJSON(data []byte) (*IntegrationsInstancesProvidersSetBody, error) {
	var v IntegrationsInstancesProvidersSetBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstancesProvidersSetBodyToJSON serializes a IntegrationsInstancesProvidersSetBody to JSON.
func MapIntegrationsInstancesProvidersSetBodyToJSON(v *IntegrationsInstancesProvidersSetBody) ([]byte, error) {
	return json.Marshal(v)
}
