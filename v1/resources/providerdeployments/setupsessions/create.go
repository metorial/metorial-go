package setupsessions

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsSetupSessionsCreateOutputAuthMethodInputSchema represents the provider deployments setup sessions create output auth method input schema type.
type ProviderDeploymentsSetupSessionsCreateOutputAuthMethodInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// ProviderDeploymentsSetupSessionsCreateOutputAuthMethodOutputSchema represents the provider deployments setup sessions create output auth method output schema type.
type ProviderDeploymentsSetupSessionsCreateOutputAuthMethodOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// ProviderDeploymentsSetupSessionsCreateOutputAuthMethodScopes represents the provider deployments setup sessions create output auth method scopes type.
type ProviderDeploymentsSetupSessionsCreateOutputAuthMethodScopes struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique scope identifier
	Id string `json:"id"`
	// Scope - OAuth scope string
	Scope string `json:"scope"`
	// Name - Display name of the scope
	Name string `json:"name"`
	// Description - Scope description
	Description *string `json:"description,omitempty"`
}

// ProviderDeploymentsSetupSessionsCreateOutputAuthMethod represents the provider deployments setup sessions create output auth method type.
type ProviderDeploymentsSetupSessionsCreateOutputAuthMethod struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique auth method identifier
	Id string `json:"id"`
	// Type - Authentication type
	Type string `json:"type"`
	// Key - Auth method key
	Key string `json:"key"`
	// Name - Display name
	Name string `json:"name"`
	// Description - Auth method description
	Description *string `json:"description,omitempty"`
	// Capabilities - Auth method capabilities
	Capabilities map[string]any                                                      `json:"capabilities"`
	InputSchema  *ProviderDeploymentsSetupSessionsCreateOutputAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *ProviderDeploymentsSetupSessionsCreateOutputAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]ProviderDeploymentsSetupSessionsCreateOutputAuthMethodScopes `json:"scopes,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// ProviderSpecificationId - Specification ID
	ProviderSpecificationId string `json:"provider_specification_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsSetupSessionsCreateOutputDeployment represents the provider deployments setup sessions create output deployment type.
type ProviderDeploymentsSetupSessionsCreateOutputDeployment struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Deployment ID
	Id string `json:"id"`
	// IsDefault - Whether this is the default deployment
	IsDefault bool `json:"is_default"`
	// Name - Deployment name
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

// ProviderDeploymentsSetupSessionsCreateOutputCredentials represents the provider deployments setup sessions create output credentials type.
type ProviderDeploymentsSetupSessionsCreateOutputCredentials struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique credentials identifier
	Id   string `json:"id"`
	Type string `json:"type"`
	// IsDefault - Whether this is the default credentials for the provider
	IsDefault bool `json:"is_default"`
	// IsManaged - Whether these credentials are managed by Metorial
	IsManaged bool `json:"is_managed"`
	// Name - Display name
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

// ProviderDeploymentsSetupSessionsCreateOutputAuthConfigToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ProviderDeploymentsSetupSessionsCreateOutputAuthConfigToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// ProviderDeploymentsSetupSessionsCreateOutputAuthConfigDeployment represents the provider deployments setup sessions create output auth config deployment type.
type ProviderDeploymentsSetupSessionsCreateOutputAuthConfigDeployment struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Deployment ID
	Id string `json:"id"`
	// IsDefault - Whether this is the default deployment
	IsDefault bool `json:"is_default"`
	// Name - Deployment name
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

// ProviderDeploymentsSetupSessionsCreateOutputAuthConfigCredentials represents the provider deployments setup sessions create output auth config credentials type.
type ProviderDeploymentsSetupSessionsCreateOutputAuthConfigCredentials struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique credentials identifier
	Id   string `json:"id"`
	Type string `json:"type"`
	// IsDefault - Whether this is the default credentials for the provider
	IsDefault bool `json:"is_default"`
	// IsManaged - Whether these credentials are managed by Metorial
	IsManaged bool `json:"is_managed"`
	// Name - Display name
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

// ProviderDeploymentsSetupSessionsCreateOutputAuthConfigAuthMethodInputSchema represents the provider deployments setup sessions create output auth config auth method input schema type.
type ProviderDeploymentsSetupSessionsCreateOutputAuthConfigAuthMethodInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// ProviderDeploymentsSetupSessionsCreateOutputAuthConfigAuthMethodOutputSchema represents the provider deployments setup sessions create output auth config auth method output schema type.
type ProviderDeploymentsSetupSessionsCreateOutputAuthConfigAuthMethodOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// ProviderDeploymentsSetupSessionsCreateOutputAuthConfigAuthMethodScopes represents the provider deployments setup sessions create output auth config auth method scopes type.
type ProviderDeploymentsSetupSessionsCreateOutputAuthConfigAuthMethodScopes struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique scope identifier
	Id string `json:"id"`
	// Scope - OAuth scope string
	Scope string `json:"scope"`
	// Name - Display name of the scope
	Name string `json:"name"`
	// Description - Scope description
	Description *string `json:"description,omitempty"`
}

// ProviderDeploymentsSetupSessionsCreateOutputAuthConfigAuthMethod represents the provider deployments setup sessions create output auth config auth method type.
type ProviderDeploymentsSetupSessionsCreateOutputAuthConfigAuthMethod struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique auth method identifier
	Id string `json:"id"`
	// Type - Authentication type
	Type string `json:"type"`
	// Key - Auth method key
	Key string `json:"key"`
	// Name - Display name
	Name string `json:"name"`
	// Description - Auth method description
	Description *string `json:"description,omitempty"`
	// Capabilities - Auth method capabilities
	Capabilities map[string]any                                                                `json:"capabilities"`
	InputSchema  *ProviderDeploymentsSetupSessionsCreateOutputAuthConfigAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *ProviderDeploymentsSetupSessionsCreateOutputAuthConfigAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]ProviderDeploymentsSetupSessionsCreateOutputAuthConfigAuthMethodScopes `json:"scopes,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// ProviderSpecificationId - Specification ID
	ProviderSpecificationId string `json:"provider_specification_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsSetupSessionsCreateOutputAuthConfig represents the provider deployments setup sessions create output auth config type.
type ProviderDeploymentsSetupSessionsCreateOutputAuthConfig struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique auth config identifier
	Id string `json:"id"`
	// Type - Authentication type
	Type string `json:"type"`
	// Source - Auth config source
	Source string `json:"source"`
	// Status - Auth config status
	Status string `json:"status"`
	// IsDefault - Whether this is the default auth config
	IsDefault bool `json:"is_default"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// Name - Display name
	Name *string `json:"name,omitempty"`
	// Description - Description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter  ProviderDeploymentsSetupSessionsCreateOutputAuthConfigToolFilter   `json:"tool_filter"`
	Deployment  *ProviderDeploymentsSetupSessionsCreateOutputAuthConfigDeployment  `json:"deployment,omitempty"`
	Credentials *ProviderDeploymentsSetupSessionsCreateOutputAuthConfigCredentials `json:"credentials,omitempty"`
	AuthMethod  ProviderDeploymentsSetupSessionsCreateOutputAuthConfigAuthMethod   `json:"auth_method"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsSetupSessionsCreateOutputConfigToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ProviderDeploymentsSetupSessionsCreateOutputConfigToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// ProviderDeploymentsSetupSessionsCreateOutputConfigDeployment represents the provider deployments setup sessions create output config deployment type.
type ProviderDeploymentsSetupSessionsCreateOutputConfigDeployment struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Deployment ID
	Id string `json:"id"`
	// IsDefault - Whether this is the default deployment
	IsDefault bool `json:"is_default"`
	// Name - Deployment name
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

// ProviderDeploymentsSetupSessionsCreateOutputConfigFromVaultDeployment represents the provider deployments setup sessions create output config from vault deployment type.
type ProviderDeploymentsSetupSessionsCreateOutputConfigFromVaultDeployment struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Deployment ID
	Id string `json:"id"`
	// IsDefault - Whether this is the default deployment
	IsDefault bool `json:"is_default"`
	// Name - Deployment name
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

// ProviderDeploymentsSetupSessionsCreateOutputConfigFromVault represents the provider deployments setup sessions create output config from vault type.
type ProviderDeploymentsSetupSessionsCreateOutputConfigFromVault struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique config vault identifier
	Id string `json:"id"`
	// Name - Display name
	Name string `json:"name"`
	// Description - Description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider ID
	ProviderId string                                                                 `json:"provider_id"`
	Deployment *ProviderDeploymentsSetupSessionsCreateOutputConfigFromVaultDeployment `json:"deployment,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsSetupSessionsCreateOutputConfig represents the provider deployments setup sessions create output config type.
type ProviderDeploymentsSetupSessionsCreateOutputConfig struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique config identifier
	Id string `json:"id"`
	// IsDefault - Whether this is the default config
	IsDefault bool `json:"is_default"`
	// Name - Display name
	Name *string `json:"name,omitempty"`
	// Description - Description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter ProviderDeploymentsSetupSessionsCreateOutputConfigToolFilter `json:"tool_filter"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// SpecificationId - Specification ID
	SpecificationId string                                                        `json:"specification_id"`
	Deployment      *ProviderDeploymentsSetupSessionsCreateOutputConfigDeployment `json:"deployment,omitempty"`
	FromVault       *ProviderDeploymentsSetupSessionsCreateOutputConfigFromVault  `json:"from_vault,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsSetupSessionsCreateOutput represents the provider deployments setup sessions create output type.
type ProviderDeploymentsSetupSessionsCreateOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique setup session identifier
	Id string `json:"id"`
	// Type - Setup session type
	Type string `json:"type"`
	// Status - Session status
	Status string `json:"status"`
	// Url - URL where user completes authentication
	Url string `json:"url"`
	// Name - Display name
	Name *string `json:"name,omitempty"`
	// Description - Description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Configuration - Setup session configuration
	Configuration *map[string]any `json:"configuration,omitempty"`
	// ProviderId - Provider ID
	ProviderId *string `json:"provider_id,omitempty"`
	// IdentityId - Linked identity ID
	IdentityId *string `json:"identity_id,omitempty"`
	// IdentityCredentialId - Identity credential created from this setup session
	IdentityCredentialId *string                                                  `json:"identity_credential_id,omitempty"`
	AuthMethod           *ProviderDeploymentsSetupSessionsCreateOutputAuthMethod  `json:"auth_method,omitempty"`
	Deployment           *ProviderDeploymentsSetupSessionsCreateOutputDeployment  `json:"deployment,omitempty"`
	Credentials          *ProviderDeploymentsSetupSessionsCreateOutputCredentials `json:"credentials,omitempty"`
	AuthConfig           *ProviderDeploymentsSetupSessionsCreateOutputAuthConfig  `json:"auth_config,omitempty"`
	Config               *ProviderDeploymentsSetupSessionsCreateOutputConfig      `json:"config,omitempty"`
	// UiMode - UI mode for setup
	UiMode string `json:"ui_mode"`
	// RedirectUrl - URL to redirect after setup
	RedirectUrl *string `json:"redirect_url,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
	// ExpiresAt - Timestamp when the session expires
	ExpiresAt time.Time `json:"expires_at"`
}

// MapProviderDeploymentsSetupSessionsCreateOutputFromJSON deserializes JSON data into a ProviderDeploymentsSetupSessionsCreateOutput.
func MapProviderDeploymentsSetupSessionsCreateOutputFromJSON(data []byte) (*ProviderDeploymentsSetupSessionsCreateOutput, error) {
	var v ProviderDeploymentsSetupSessionsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsSetupSessionsCreateOutputToJSON serializes a ProviderDeploymentsSetupSessionsCreateOutput to JSON.
func MapProviderDeploymentsSetupSessionsCreateOutputToJSON(v *ProviderDeploymentsSetupSessionsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderDeploymentsSetupSessionsCreateBodyConfigurationProviderSearchGroups represents the provider deployments setup sessions create body configuration provider search groups type.
type ProviderDeploymentsSetupSessionsCreateBodyConfigurationProviderSearchGroups struct {
	GroupId string `json:"group_id"`
}

// ProviderDeploymentsSetupSessionsCreateBodyConfigurationProviderSearchCollections represents the provider deployments setup sessions create body configuration provider search collections type.
type ProviderDeploymentsSetupSessionsCreateBodyConfigurationProviderSearchCollections struct {
	CollectionId string `json:"collection_id"`
}

// ProviderDeploymentsSetupSessionsCreateBodyConfigurationProviderSearchCategories represents the provider deployments setup sessions create body configuration provider search categories type.
type ProviderDeploymentsSetupSessionsCreateBodyConfigurationProviderSearchCategories struct {
	CategoryId string `json:"category_id"`
}

// ProviderDeploymentsSetupSessionsCreateBodyConfigurationProviderSearch represents the provider deployments setup sessions create body configuration provider search type.
type ProviderDeploymentsSetupSessionsCreateBodyConfigurationProviderSearch struct {
	Groups      *[]ProviderDeploymentsSetupSessionsCreateBodyConfigurationProviderSearchGroups      `json:"groups,omitempty"`
	Collections *[]ProviderDeploymentsSetupSessionsCreateBodyConfigurationProviderSearchCollections `json:"collections,omitempty"`
	Categories  *[]ProviderDeploymentsSetupSessionsCreateBodyConfigurationProviderSearchCategories  `json:"categories,omitempty"`
}

// ProviderDeploymentsSetupSessionsCreateBodyConfigurationToolFilters represents the provider deployments setup sessions create body configuration tool filters type.
type ProviderDeploymentsSetupSessionsCreateBodyConfigurationToolFilters struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// ProviderDeploymentsSetupSessionsCreateBodyConfigurationUi represents the provider deployments setup sessions create body configuration ui type.
type ProviderDeploymentsSetupSessionsCreateBodyConfigurationUi struct {
	Layout *string `json:"layout,omitempty"`
}

// ProviderDeploymentsSetupSessionsCreateBodyConfiguration represents the provider deployments setup sessions create body configuration type.
type ProviderDeploymentsSetupSessionsCreateBodyConfiguration struct {
	ProviderSearch *ProviderDeploymentsSetupSessionsCreateBodyConfigurationProviderSearch `json:"provider_search,omitempty"`
	ToolFilters    *ProviderDeploymentsSetupSessionsCreateBodyConfigurationToolFilters    `json:"tool_filters,omitempty"`
	Ui             *ProviderDeploymentsSetupSessionsCreateBodyConfigurationUi             `json:"ui,omitempty"`
}

// ProviderDeploymentsSetupSessionsCreateBody represents the provider deployments setup sessions create body type.
type ProviderDeploymentsSetupSessionsCreateBody struct {
	// ProviderId - The provider ID
	ProviderId *string `json:"provider_id,omitempty"`
	// ProviderDeploymentId - Optional provider deployment ID
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	Name                 *string `json:"name,omitempty"`
	Description          *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderAuthMethodId - The authentication method to use (e.g., OAuth flow)
	ProviderAuthMethodId *string `json:"provider_auth_method_id,omitempty"`
	// ProviderAuthCredentialsId - Optional OAuth app credentials to use instead of defaults
	ProviderAuthCredentialsId *string `json:"provider_auth_credentials_id,omitempty"`
	// IdentityId - Optional identity to link this setup session to
	IdentityId *string `json:"identity_id,omitempty"`
	// ConsumerId - Optional consumer to link this setup session to
	ConsumerId  *string `json:"consumer_id,omitempty"`
	RedirectUrl *string `json:"redirect_url,omitempty"`
	// Type - The type of setup session, determining the flow and outcome of the session
	Type          *string                                                  `json:"type,omitempty"`
	Configuration *ProviderDeploymentsSetupSessionsCreateBodyConfiguration `json:"configuration,omitempty"`
}

// MapProviderDeploymentsSetupSessionsCreateBodyFromJSON deserializes JSON data into a ProviderDeploymentsSetupSessionsCreateBody.
func MapProviderDeploymentsSetupSessionsCreateBodyFromJSON(data []byte) (*ProviderDeploymentsSetupSessionsCreateBody, error) {
	var v ProviderDeploymentsSetupSessionsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsSetupSessionsCreateBodyToJSON serializes a ProviderDeploymentsSetupSessionsCreateBody to JSON.
func MapProviderDeploymentsSetupSessionsCreateBodyToJSON(v *ProviderDeploymentsSetupSessionsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
