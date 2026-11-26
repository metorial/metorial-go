package magicmcpservers

import (
	"encoding/json"
	"time"
)

// MagicMcpServersCreateOutputEndpoints represents the magic mcp servers create output endpoints type.
type MagicMcpServersCreateOutputEndpoints struct {
	// Id - Unique ID.
	Id string `json:"id"`
	// Alias - Connection slug for the server.
	Alias string `json:"alias"`
	// Url - Connection URL for the server.
	Url string `json:"url"`
}

// MagicMcpServersCreateOutputProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type MagicMcpServersCreateOutputProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// MagicMcpServersCreateOutputProvidersProvider represents the magic mcp servers create output providers provider type.
type MagicMcpServersCreateOutputProvidersProvider struct {
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

// MagicMcpServersCreateOutputProvidersDeployment represents the magic mcp servers create output providers deployment type.
type MagicMcpServersCreateOutputProvidersDeployment struct {
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

// MagicMcpServersCreateOutputProvidersAuthMethodInputSchema represents the magic mcp servers create output providers auth method input schema type.
type MagicMcpServersCreateOutputProvidersAuthMethodInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// MagicMcpServersCreateOutputProvidersAuthMethodOutputSchema represents the magic mcp servers create output providers auth method output schema type.
type MagicMcpServersCreateOutputProvidersAuthMethodOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// MagicMcpServersCreateOutputProvidersAuthMethodScopes represents the magic mcp servers create output providers auth method scopes type.
type MagicMcpServersCreateOutputProvidersAuthMethodScopes struct {
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

// MagicMcpServersCreateOutputProvidersAuthMethod represents the magic mcp servers create output providers auth method type.
type MagicMcpServersCreateOutputProvidersAuthMethod struct {
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
	InputSchema  *MagicMcpServersCreateOutputProvidersAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *MagicMcpServersCreateOutputProvidersAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]MagicMcpServersCreateOutputProvidersAuthMethodScopes `json:"scopes,omitempty"`
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

// MagicMcpServersCreateOutputProvidersAuthCredentials represents the magic mcp servers create output providers auth credentials type.
type MagicMcpServersCreateOutputProvidersAuthCredentials struct {
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

// MagicMcpServersCreateOutputProvidersConfig represents the magic mcp servers create output providers config type.
type MagicMcpServersCreateOutputProvidersConfig struct {
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

// MagicMcpServersCreateOutputProvidersAuthConfig represents the magic mcp servers create output providers auth config type.
type MagicMcpServersCreateOutputProvidersAuthConfig struct {
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

// MagicMcpServersCreateOutputProviders represents the magic mcp servers create output providers type.
type MagicMcpServersCreateOutputProviders struct {
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
	ToolFilter      *MagicMcpServersCreateOutputProvidersToolFilter      `json:"tool_filter,omitempty"`
	Provider        MagicMcpServersCreateOutputProvidersProvider         `json:"provider"`
	Deployment      MagicMcpServersCreateOutputProvidersDeployment       `json:"deployment"`
	AuthMethod      *MagicMcpServersCreateOutputProvidersAuthMethod      `json:"auth_method,omitempty"`
	AuthCredentials *MagicMcpServersCreateOutputProvidersAuthCredentials `json:"auth_credentials,omitempty"`
	Config          *MagicMcpServersCreateOutputProvidersConfig          `json:"config,omitempty"`
	AuthConfig      *MagicMcpServersCreateOutputProvidersAuthConfig      `json:"auth_config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// MagicMcpServersCreateOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type MagicMcpServersCreateOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// MagicMcpServersCreateOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type MagicMcpServersCreateOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// MagicMcpServersCreateOutputAccessOwner - The actor that owns this resource, if any.
type MagicMcpServersCreateOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *MagicMcpServersCreateOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *MagicMcpServersCreateOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// MagicMcpServersCreateOutputAccess - Resource ownership and access settings.
type MagicMcpServersCreateOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *MagicMcpServersCreateOutputAccessOwner `json:"owner,omitempty"`
}

// MagicMcpServersCreateOutput represents the magic mcp servers create output type.
type MagicMcpServersCreateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once you delete the server. Archived servers stay listable.
	Status string `json:"status"`
	// Source - How this Magic MCP server was created or configured
	Source string `json:"source"`
	// ProviderManagementMode - Whether providers are managed on the server or inherited.
	ProviderManagementMode string `json:"provider_management_mode"`
	// Endpoints - Connection targets for the server.
	Endpoints []MagicMcpServersCreateOutputEndpoints `json:"endpoints"`
	// ProviderTemplateId - Provider template the server is built from, if any.
	ProviderTemplateId *string `json:"provider_template_id,omitempty"`
	// Providers - Providers the server connects through.
	Providers []MagicMcpServersCreateOutputProviders `json:"providers"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata map[string]any `json:"metadata"`
	// Access - Resource ownership and access settings.
	Access MagicMcpServersCreateOutputAccess `json:"access"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapMagicMcpServersCreateOutputFromJSON deserializes JSON data into a MagicMcpServersCreateOutput.
func MapMagicMcpServersCreateOutputFromJSON(data []byte) (*MagicMcpServersCreateOutput, error) {
	var v MagicMcpServersCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersCreateOutputToJSON serializes a MagicMcpServersCreateOutput to JSON.
func MapMagicMcpServersCreateOutputToJSON(v *MagicMcpServersCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpServersCreateBody represents the magic mcp servers create body type.
type MagicMcpServersCreateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the server. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderTemplateId - Provider template to build the server from.
	ProviderTemplateId *string `json:"provider_template_id,omitempty"`
	// IntegrationInstanceId - Integration instance to build the server from.
	IntegrationInstanceId *string `json:"integration_instance_id,omitempty"`
	// ConsumerProfileId - Consumer profile that will own the server.
	ConsumerProfileId *string `json:"consumer_profile_id,omitempty"`
}

// MapMagicMcpServersCreateBodyFromJSON deserializes JSON data into a MagicMcpServersCreateBody.
func MapMagicMcpServersCreateBodyFromJSON(data []byte) (*MagicMcpServersCreateBody, error) {
	var v MagicMcpServersCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersCreateBodyToJSON serializes a MagicMcpServersCreateBody to JSON.
func MapMagicMcpServersCreateBodyToJSON(v *MagicMcpServersCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
