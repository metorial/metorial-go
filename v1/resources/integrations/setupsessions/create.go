package setupsessions

import (
	"encoding/json"
	"time"
)

// IntegrationsSetupSessionsCreateOutputIntegrationInstanceImplementation represents the integrations setup sessions create output integration instance implementation type.
type IntegrationsSetupSessionsCreateOutputIntegrationInstanceImplementation struct {
	Type             string `json:"type"`
	MagicMcpServerId string `json:"magic_mcp_server_id"`
}

// IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersProvider represents the integrations setup sessions create output integration instance providers provider type.
type IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersProvider struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique provider identifier
	Id string `json:"id"`
	// Name - Display name of the provider
	Name string `json:"name"`
	// Description - Brief description of the provider
	Description *string `json:"description,omitempty"`
	// Slug - URL-friendly identifier
	Slug string `json:"slug"`
	// CreatedAt - Timestamp when the provider was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the provider was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersIntegrationProviderProviderVersion represents the integrations setup sessions create output integration instance providers integration provider provider version type.
type IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersIntegrationProviderProviderVersion struct {
	Object string  `json:"object"`
	Id     string  `json:"id"`
	Index  float64 `json:"index"`
}

// IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersIntegrationProviderToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersIntegrationProviderToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersIntegrationProviderConfig represents the integrations setup sessions create output integration instance providers integration provider config type.
type IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersIntegrationProviderConfig struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Config ID
	Id string `json:"id"`
	// IsDefault - Whether this is the default config
	IsDefault bool `json:"is_default"`
	// Name - Config name
	Name *string `json:"name,omitempty"`
	// Description - Description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersIntegrationProvider represents the integrations setup sessions create output integration instance providers integration provider type.
type IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersIntegrationProvider struct {
	Object          string                                                                                              `json:"object"`
	Id              string                                                                                              `json:"id"`
	ProviderVersion IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersIntegrationProviderProviderVersion `json:"provider_version"`
	Status          string                                                                                              `json:"status"`
	Name            string                                                                                              `json:"name"`
	Description     *string                                                                                             `json:"description,omitempty"`
	Metadata        *map[string]any                                                                                     `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter        *IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersIntegrationProviderToolFilter `json:"tool_filter,omitempty"`
	ProviderId        string                                                                                          `json:"provider_id"`
	DeploymentId      string                                                                                          `json:"deployment_id"`
	AuthMethodId      *string                                                                                         `json:"auth_method_id,omitempty"`
	AuthCredentialsId *string                                                                                         `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersIntegrationProviderConfig     `json:"config,omitempty"`
	CreatedAt         time.Time                                                                                       `json:"created_at"`
	UpdatedAt         time.Time                                                                                       `json:"updated_at"`
	ArchivedAt        *time.Time                                                                                      `json:"archived_at,omitempty"`
}

// IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersConfig represents the integrations setup sessions create output integration instance providers config type.
type IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersConfig struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Config ID
	Id string `json:"id"`
	// IsDefault - Whether this is the default config
	IsDefault bool `json:"is_default"`
	// Name - Config name
	Name *string `json:"name,omitempty"`
	// Description - Description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersAuthConfig represents the integrations setup sessions create output integration instance providers auth config type.
type IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersAuthConfig struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Config ID
	Id string `json:"id"`
	// IsDefault - Whether this is the default config
	IsDefault bool `json:"is_default"`
	// Name - Config name
	Name *string `json:"name,omitempty"`
	// Description - Description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// IntegrationsSetupSessionsCreateOutputIntegrationInstanceProviders represents the integrations setup sessions create output integration instance providers type.
type IntegrationsSetupSessionsCreateOutputIntegrationInstanceProviders struct {
	Object                string          `json:"object"`
	Id                    string          `json:"id"`
	Status                string          `json:"status"`
	Name                  string          `json:"name"`
	Description           *string         `json:"description,omitempty"`
	Metadata              *map[string]any `json:"metadata,omitempty"`
	IntegrationId         string          `json:"integration_id"`
	IntegrationInstanceId string          `json:"integration_instance_id"`
	// ToolFilter - Tool filter configuration
	ToolFilter           *IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersToolFilter         `json:"tool_filter,omitempty"`
	IsOverrideToolFilter bool                                                                                 `json:"is_override_tool_filter"`
	Provider             IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersProvider            `json:"provider"`
	IntegrationProvider  IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersIntegrationProvider `json:"integration_provider"`
	Config               *IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersConfig             `json:"config,omitempty"`
	AuthConfig           *IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersAuthConfig         `json:"auth_config,omitempty"`
	CreatedAt            time.Time                                                                            `json:"created_at"`
	UpdatedAt            time.Time                                                                            `json:"updated_at"`
	ArchivedAt           *time.Time                                                                           `json:"archived_at,omitempty"`
}

// IntegrationsSetupSessionsCreateOutputIntegrationInstance represents the integrations setup sessions create output integration instance type.
type IntegrationsSetupSessionsCreateOutputIntegrationInstance struct {
	Object          string                                                                  `json:"object"`
	Id              string                                                                  `json:"id"`
	Status          string                                                                  `json:"status"`
	Name            string                                                                  `json:"name"`
	Description     *string                                                                 `json:"description,omitempty"`
	Metadata        *map[string]any                                                         `json:"metadata,omitempty"`
	IntegrationId   string                                                                  `json:"integration_id"`
	IdentityActorId *string                                                                 `json:"identity_actor_id,omitempty"`
	IdentityId      *string                                                                 `json:"identity_id,omitempty"`
	Implementation  *IntegrationsSetupSessionsCreateOutputIntegrationInstanceImplementation `json:"implementation,omitempty"`
	Providers       []IntegrationsSetupSessionsCreateOutputIntegrationInstanceProviders     `json:"providers"`
	CreatedAt       time.Time                                                               `json:"created_at"`
	UpdatedAt       time.Time                                                               `json:"updated_at"`
	ArchivedAt      *time.Time                                                              `json:"archived_at,omitempty"`
}

// IntegrationsSetupSessionsCreateOutput represents the integrations setup sessions create output type.
type IntegrationsSetupSessionsCreateOutput struct {
	Object              string                                                   `json:"object"`
	Id                  string                                                   `json:"id"`
	Status              string                                                   `json:"status"`
	Url                 string                                                   `json:"url"`
	Name                *string                                                  `json:"name,omitempty"`
	Description         *string                                                  `json:"description,omitempty"`
	Metadata            *map[string]any                                          `json:"metadata,omitempty"`
	Configuration       *map[string]any                                          `json:"configuration,omitempty"`
	RedirectUrl         *string                                                  `json:"redirect_url,omitempty"`
	IntegrationId       string                                                   `json:"integration_id"`
	IntegrationInstance IntegrationsSetupSessionsCreateOutputIntegrationInstance `json:"integration_instance"`
	CreatedAt           time.Time                                                `json:"created_at"`
	UpdatedAt           time.Time                                                `json:"updated_at"`
	ExpiresAt           time.Time                                                `json:"expires_at"`
}

// MapIntegrationsSetupSessionsCreateOutputFromJSON deserializes JSON data into a IntegrationsSetupSessionsCreateOutput.
func MapIntegrationsSetupSessionsCreateOutputFromJSON(data []byte) (*IntegrationsSetupSessionsCreateOutput, error) {
	var v IntegrationsSetupSessionsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsSetupSessionsCreateOutputToJSON serializes a IntegrationsSetupSessionsCreateOutput to JSON.
func MapIntegrationsSetupSessionsCreateOutputToJSON(v *IntegrationsSetupSessionsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsSetupSessionsCreateBodyConfigurationProviderSearchGroups represents the integrations setup sessions create body configuration provider search groups type.
type IntegrationsSetupSessionsCreateBodyConfigurationProviderSearchGroups struct {
	GroupId string `json:"group_id"`
}

// IntegrationsSetupSessionsCreateBodyConfigurationProviderSearchCollections represents the integrations setup sessions create body configuration provider search collections type.
type IntegrationsSetupSessionsCreateBodyConfigurationProviderSearchCollections struct {
	CollectionId string `json:"collection_id"`
}

// IntegrationsSetupSessionsCreateBodyConfigurationProviderSearchCategories represents the integrations setup sessions create body configuration provider search categories type.
type IntegrationsSetupSessionsCreateBodyConfigurationProviderSearchCategories struct {
	CategoryId string `json:"category_id"`
}

// IntegrationsSetupSessionsCreateBodyConfigurationProviderSearch represents the integrations setup sessions create body configuration provider search type.
type IntegrationsSetupSessionsCreateBodyConfigurationProviderSearch struct {
	Groups      *[]IntegrationsSetupSessionsCreateBodyConfigurationProviderSearchGroups      `json:"groups,omitempty"`
	Collections *[]IntegrationsSetupSessionsCreateBodyConfigurationProviderSearchCollections `json:"collections,omitempty"`
	Categories  *[]IntegrationsSetupSessionsCreateBodyConfigurationProviderSearchCategories  `json:"categories,omitempty"`
}

// IntegrationsSetupSessionsCreateBodyConfigurationToolFilters represents the integrations setup sessions create body configuration tool filters type.
type IntegrationsSetupSessionsCreateBodyConfigurationToolFilters struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// IntegrationsSetupSessionsCreateBodyConfigurationUi represents the integrations setup sessions create body configuration ui type.
type IntegrationsSetupSessionsCreateBodyConfigurationUi struct {
	Layout *string `json:"layout,omitempty"`
}

// IntegrationsSetupSessionsCreateBodyConfiguration represents the integrations setup sessions create body configuration type.
type IntegrationsSetupSessionsCreateBodyConfiguration struct {
	ProviderSearch *IntegrationsSetupSessionsCreateBodyConfigurationProviderSearch `json:"provider_search,omitempty"`
	ToolFilters    *IntegrationsSetupSessionsCreateBodyConfigurationToolFilters    `json:"tool_filters,omitempty"`
	Ui             *IntegrationsSetupSessionsCreateBodyConfigurationUi             `json:"ui,omitempty"`
}

// IntegrationsSetupSessionsCreateBody represents the integrations setup sessions create body type.
type IntegrationsSetupSessionsCreateBody struct {
	IntegrationId   string                                            `json:"integration_id"`
	Name            string                                            `json:"name"`
	Description     *string                                           `json:"description,omitempty"`
	Metadata        *map[string]any                                   `json:"metadata,omitempty"`
	IdentityActorId *string                                           `json:"identity_actor_id,omitempty"`
	IdentityId      *string                                           `json:"identity_id,omitempty"`
	ExpiresAt       *time.Time                                        `json:"expires_at,omitempty"`
	RedirectUrl     *string                                           `json:"redirect_url,omitempty"`
	Configuration   *IntegrationsSetupSessionsCreateBodyConfiguration `json:"configuration,omitempty"`
}

// MapIntegrationsSetupSessionsCreateBodyFromJSON deserializes JSON data into a IntegrationsSetupSessionsCreateBody.
func MapIntegrationsSetupSessionsCreateBodyFromJSON(data []byte) (*IntegrationsSetupSessionsCreateBody, error) {
	var v IntegrationsSetupSessionsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsSetupSessionsCreateBodyToJSON serializes a IntegrationsSetupSessionsCreateBody to JSON.
func MapIntegrationsSetupSessionsCreateBodyToJSON(v *IntegrationsSetupSessionsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
