package magicmcpservers

import (
	"encoding/json"
	"time"
)

// MagicMcpServersUpdateOutputEndpoints represents the magic mcp servers update output endpoints type.
type MagicMcpServersUpdateOutputEndpoints struct {
	Id    string `json:"id"`
	Alias string `json:"alias"`
	Url   string `json:"url"`
}

// MagicMcpServersUpdateOutputProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type MagicMcpServersUpdateOutputProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// MagicMcpServersUpdateOutputProvidersProvider represents the magic mcp servers update output providers provider type.
type MagicMcpServersUpdateOutputProvidersProvider struct {
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

// MagicMcpServersUpdateOutputProvidersDeployment represents the magic mcp servers update output providers deployment type.
type MagicMcpServersUpdateOutputProvidersDeployment struct {
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

// MagicMcpServersUpdateOutputProvidersAuthMethodInputSchema represents the magic mcp servers update output providers auth method input schema type.
type MagicMcpServersUpdateOutputProvidersAuthMethodInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// MagicMcpServersUpdateOutputProvidersAuthMethodOutputSchema represents the magic mcp servers update output providers auth method output schema type.
type MagicMcpServersUpdateOutputProvidersAuthMethodOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// MagicMcpServersUpdateOutputProvidersAuthMethodScopes represents the magic mcp servers update output providers auth method scopes type.
type MagicMcpServersUpdateOutputProvidersAuthMethodScopes struct {
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

// MagicMcpServersUpdateOutputProvidersAuthMethod represents the magic mcp servers update output providers auth method type.
type MagicMcpServersUpdateOutputProvidersAuthMethod struct {
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
	// Capabilities - Capabilities of this auth method in the provider version the deployment currently uses: its locked version, otherwise the provider's current version. The other fields describe the pinned auth method and may come from an older specification.
	Capabilities map[string]any                                              `json:"capabilities"`
	InputSchema  *MagicMcpServersUpdateOutputProvidersAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *MagicMcpServersUpdateOutputProvidersAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]MagicMcpServersUpdateOutputProvidersAuthMethodScopes `json:"scopes,omitempty"`
	// Adapters - Adapter IDs this auth method can be used with
	Adapters *[]string `json:"adapters,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// ProviderSpecificationId - Specification ID
	ProviderSpecificationId string `json:"provider_specification_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpServersUpdateOutputProvidersAuthCredentials represents the magic mcp servers update output providers auth credentials type.
type MagicMcpServersUpdateOutputProvidersAuthCredentials struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique credentials identifier
	Id   string `json:"id"`
	Type string `json:"type"`
	// Status - Credentials status
	Status string `json:"status"`
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
	// Scopes - OAuth scopes requested by this credential
	Scopes *[]string `json:"scopes,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpServersUpdateOutputProvidersConfig represents the magic mcp servers update output providers config type.
type MagicMcpServersUpdateOutputProvidersConfig struct {
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

// MagicMcpServersUpdateOutputProvidersAuthConfig represents the magic mcp servers update output providers auth config type.
type MagicMcpServersUpdateOutputProvidersAuthConfig struct {
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

// MagicMcpServersUpdateOutputProviders represents the magic mcp servers update output providers type.
type MagicMcpServersUpdateOutputProviders struct {
	Object                 string          `json:"object"`
	Id                     string          `json:"id"`
	Status                 string          `json:"status"`
	MagicMcpServerId       string          `json:"magic_mcp_server_id"`
	ProviderManagementMode string          `json:"provider_management_mode"`
	Name                   string          `json:"name"`
	Description            *string         `json:"description,omitempty"`
	Metadata               *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter      *MagicMcpServersUpdateOutputProvidersToolFilter      `json:"tool_filter,omitempty"`
	Provider        MagicMcpServersUpdateOutputProvidersProvider         `json:"provider"`
	Deployment      MagicMcpServersUpdateOutputProvidersDeployment       `json:"deployment"`
	AuthMethod      *MagicMcpServersUpdateOutputProvidersAuthMethod      `json:"auth_method,omitempty"`
	AuthCredentials *MagicMcpServersUpdateOutputProvidersAuthCredentials `json:"auth_credentials,omitempty"`
	Config          *MagicMcpServersUpdateOutputProvidersConfig          `json:"config,omitempty"`
	AuthConfig      *MagicMcpServersUpdateOutputProvidersAuthConfig      `json:"auth_config,omitempty"`
	CreatedAt       time.Time                                            `json:"created_at"`
	UpdatedAt       time.Time                                            `json:"updated_at"`
	ArchivedAt      *time.Time                                           `json:"archived_at,omitempty"`
}

// MagicMcpServersUpdateOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type MagicMcpServersUpdateOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// MagicMcpServersUpdateOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type MagicMcpServersUpdateOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// MagicMcpServersUpdateOutputAccessOwner - The actor that owns this resource, if any.
type MagicMcpServersUpdateOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *MagicMcpServersUpdateOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *MagicMcpServersUpdateOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// MagicMcpServersUpdateOutputAccess - Resource ownership and access settings.
type MagicMcpServersUpdateOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *MagicMcpServersUpdateOutputAccessOwner `json:"owner,omitempty"`
}

// MagicMcpServersUpdateOutput represents the magic mcp servers update output type.
type MagicMcpServersUpdateOutput struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
	// Source - How this Magic MCP server was created or configured
	Source                 string                                 `json:"source"`
	ProviderManagementMode string                                 `json:"provider_management_mode"`
	Endpoints              []MagicMcpServersUpdateOutputEndpoints `json:"endpoints"`
	ProviderTemplateId     *string                                `json:"provider_template_id,omitempty"`
	Providers              []MagicMcpServersUpdateOutputProviders `json:"providers"`
	Name                   *string                                `json:"name,omitempty"`
	Description            *string                                `json:"description,omitempty"`
	Metadata               map[string]any                         `json:"metadata"`
	// Access - Resource ownership and access settings.
	Access    MagicMcpServersUpdateOutputAccess `json:"access"`
	CreatedAt time.Time                         `json:"created_at"`
	UpdatedAt time.Time                         `json:"updated_at"`
}

// MapMagicMcpServersUpdateOutputFromJSON deserializes JSON data into a MagicMcpServersUpdateOutput.
func MapMagicMcpServersUpdateOutputFromJSON(data []byte) (*MagicMcpServersUpdateOutput, error) {
	var v MagicMcpServersUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersUpdateOutputToJSON serializes a MagicMcpServersUpdateOutput to JSON.
func MapMagicMcpServersUpdateOutputToJSON(v *MagicMcpServersUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpServersUpdateBody represents the magic mcp servers update body type.
type MagicMcpServersUpdateBody struct {
	Name        *string         `json:"name,omitempty"`
	Description *string         `json:"description,omitempty"`
	Metadata    *map[string]any `json:"metadata,omitempty"`
	Aliases     *[]string       `json:"aliases,omitempty"`
}

// MapMagicMcpServersUpdateBodyFromJSON deserializes JSON data into a MagicMcpServersUpdateBody.
func MapMagicMcpServersUpdateBodyFromJSON(data []byte) (*MagicMcpServersUpdateBody, error) {
	var v MagicMcpServersUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersUpdateBodyToJSON serializes a MagicMcpServersUpdateBody to JSON.
func MapMagicMcpServersUpdateBodyToJSON(v *MagicMcpServersUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
