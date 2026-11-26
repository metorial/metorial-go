package providers

import (
	"encoding/json"
	"github.com/metorial/metorial-go/v1/nullable"
	"time"
)

// MagicMcpServersProvidersCreateOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type MagicMcpServersProvidersCreateOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// MagicMcpServersProvidersCreateOutputProvider represents the magic mcp servers providers create output provider type.
type MagicMcpServersProvidersCreateOutputProvider struct {
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

// MagicMcpServersProvidersCreateOutputDeployment represents the magic mcp servers providers create output deployment type.
type MagicMcpServersProvidersCreateOutputDeployment struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsDefault - Used when a request doesn't name a deployment.
	IsDefault bool `json:"is_default"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider this deployment belongs to.
	ProviderId string `json:"provider_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpServersProvidersCreateOutputAuthMethodInputSchema represents the magic mcp servers providers create output auth method input schema type.
type MagicMcpServersProvidersCreateOutputAuthMethodInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// MagicMcpServersProvidersCreateOutputAuthMethodOutputSchema represents the magic mcp servers providers create output auth method output schema type.
type MagicMcpServersProvidersCreateOutputAuthMethodOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// MagicMcpServersProvidersCreateOutputAuthMethodScopes represents the magic mcp servers providers create output auth method scopes type.
type MagicMcpServersProvidersCreateOutputAuthMethodScopes struct {
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

// MagicMcpServersProvidersCreateOutputAuthMethod represents the magic mcp servers providers create output auth method type.
type MagicMcpServersProvidersCreateOutputAuthMethod struct {
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
	// Capabilities - Capabilities of this auth method in the provider version the deployment currently uses: its locked version, otherwise the provider's current version. The other fields describe the pinned auth method and may come from an older specification.
	Capabilities map[string]any                                              `json:"capabilities"`
	InputSchema  *MagicMcpServersProvidersCreateOutputAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *MagicMcpServersProvidersCreateOutputAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]MagicMcpServersProvidersCreateOutputAuthMethodScopes `json:"scopes,omitempty"`
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

// MagicMcpServersProvidersCreateOutputAuthCredentials represents the magic mcp servers providers create output auth credentials type.
type MagicMcpServersProvidersCreateOutputAuthCredentials struct {
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

// MagicMcpServersProvidersCreateOutputConfig represents the magic mcp servers providers create output config type.
type MagicMcpServersProvidersCreateOutputConfig struct {
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

// MagicMcpServersProvidersCreateOutputAuthConfig represents the magic mcp servers providers create output auth config type.
type MagicMcpServersProvidersCreateOutputAuthConfig struct {
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

// MagicMcpServersProvidersCreateOutput represents the magic mcp servers providers create output type.
type MagicMcpServersProvidersCreateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once you remove the provider from the server.
	Status string `json:"status"`
	// MagicMcpServerId - Server the provider belongs to.
	MagicMcpServerId string `json:"magic_mcp_server_id"`
	// ProviderManagementMode - Whether the provider is managed here or inherited.
	ProviderManagementMode string `json:"provider_management_mode"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter      *MagicMcpServersProvidersCreateOutputToolFilter      `json:"tool_filter,omitempty"`
	Provider        MagicMcpServersProvidersCreateOutputProvider         `json:"provider"`
	Deployment      MagicMcpServersProvidersCreateOutputDeployment       `json:"deployment"`
	AuthMethod      *MagicMcpServersProvidersCreateOutputAuthMethod      `json:"auth_method,omitempty"`
	AuthCredentials *MagicMcpServersProvidersCreateOutputAuthCredentials `json:"auth_credentials,omitempty"`
	Config          *MagicMcpServersProvidersCreateOutputConfig          `json:"config,omitempty"`
	AuthConfig      *MagicMcpServersProvidersCreateOutputAuthConfig      `json:"auth_config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// MapMagicMcpServersProvidersCreateOutputFromJSON deserializes JSON data into a MagicMcpServersProvidersCreateOutput.
func MapMagicMcpServersProvidersCreateOutputFromJSON(data []byte) (*MagicMcpServersProvidersCreateOutput, error) {
	var v MagicMcpServersProvidersCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersProvidersCreateOutputToJSON serializes a MagicMcpServersProvidersCreateOutput to JSON.
func MapMagicMcpServersProvidersCreateOutputToJSON(v *MagicMcpServersProvidersCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpServersProvidersCreateBody represents the magic mcp servers providers create body type.
type MagicMcpServersProvidersCreateBody struct {
	// ProviderId - Provider to add to the server.
	ProviderId string `json:"provider_id"`
	// ProviderDeploymentId - Deployment of the provider to use.
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Configuration of the provider to use. `null` clears it.
	ProviderConfigId *nullable.Nullable[string] `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Auth configuration of the provider to use. `null` clears it.
	ProviderAuthConfigId *nullable.Nullable[string] `json:"provider_auth_config_id,omitempty"`
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
}

// MapMagicMcpServersProvidersCreateBodyFromJSON deserializes JSON data into a MagicMcpServersProvidersCreateBody.
func MapMagicMcpServersProvidersCreateBodyFromJSON(data []byte) (*MagicMcpServersProvidersCreateBody, error) {
	var v MagicMcpServersProvidersCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersProvidersCreateBodyToJSON serializes a MagicMcpServersProvidersCreateBody to JSON.
func MapMagicMcpServersProvidersCreateBodyToJSON(v *MagicMcpServersProvidersCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
