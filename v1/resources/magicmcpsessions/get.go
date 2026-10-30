package magicmcpsessions

import (
	"encoding/json"
	"time"
)

// MagicMcpSessionsGetOutputMagicMcpServerEndpoints represents the magic mcp sessions get output magic mcp server endpoints type.
type MagicMcpSessionsGetOutputMagicMcpServerEndpoints struct {
	Id    string `json:"id"`
	Alias string `json:"alias"`
	Url   string `json:"url"`
}

// MagicMcpSessionsGetOutputMagicMcpServerProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type MagicMcpSessionsGetOutputMagicMcpServerProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// MagicMcpSessionsGetOutputMagicMcpServerProvidersProvider represents the magic mcp sessions get output magic mcp server providers provider type.
type MagicMcpSessionsGetOutputMagicMcpServerProvidersProvider struct {
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

// MagicMcpSessionsGetOutputMagicMcpServerProvidersDeployment represents the magic mcp sessions get output magic mcp server providers deployment type.
type MagicMcpSessionsGetOutputMagicMcpServerProvidersDeployment struct {
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

// MagicMcpSessionsGetOutputMagicMcpServerProvidersAuthMethodInputSchema represents the magic mcp sessions get output magic mcp server providers auth method input schema type.
type MagicMcpSessionsGetOutputMagicMcpServerProvidersAuthMethodInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// MagicMcpSessionsGetOutputMagicMcpServerProvidersAuthMethodOutputSchema represents the magic mcp sessions get output magic mcp server providers auth method output schema type.
type MagicMcpSessionsGetOutputMagicMcpServerProvidersAuthMethodOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// MagicMcpSessionsGetOutputMagicMcpServerProvidersAuthMethodScopes represents the magic mcp sessions get output magic mcp server providers auth method scopes type.
type MagicMcpSessionsGetOutputMagicMcpServerProvidersAuthMethodScopes struct {
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

// MagicMcpSessionsGetOutputMagicMcpServerProvidersAuthMethod represents the magic mcp sessions get output magic mcp server providers auth method type.
type MagicMcpSessionsGetOutputMagicMcpServerProvidersAuthMethod struct {
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
	Capabilities map[string]any                                                          `json:"capabilities"`
	InputSchema  *MagicMcpSessionsGetOutputMagicMcpServerProvidersAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *MagicMcpSessionsGetOutputMagicMcpServerProvidersAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]MagicMcpSessionsGetOutputMagicMcpServerProvidersAuthMethodScopes `json:"scopes,omitempty"`
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

// MagicMcpSessionsGetOutputMagicMcpServerProvidersAuthCredentials represents the magic mcp sessions get output magic mcp server providers auth credentials type.
type MagicMcpSessionsGetOutputMagicMcpServerProvidersAuthCredentials struct {
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

// MagicMcpSessionsGetOutputMagicMcpServerProvidersConfig represents the magic mcp sessions get output magic mcp server providers config type.
type MagicMcpSessionsGetOutputMagicMcpServerProvidersConfig struct {
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

// MagicMcpSessionsGetOutputMagicMcpServerProvidersAuthConfig represents the magic mcp sessions get output magic mcp server providers auth config type.
type MagicMcpSessionsGetOutputMagicMcpServerProvidersAuthConfig struct {
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

// MagicMcpSessionsGetOutputMagicMcpServerProviders represents the magic mcp sessions get output magic mcp server providers type.
type MagicMcpSessionsGetOutputMagicMcpServerProviders struct {
	Object                 string          `json:"object"`
	Id                     string          `json:"id"`
	Status                 string          `json:"status"`
	MagicMcpServerId       string          `json:"magic_mcp_server_id"`
	ProviderManagementMode string          `json:"provider_management_mode"`
	Name                   string          `json:"name"`
	Description            *string         `json:"description,omitempty"`
	Metadata               *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter      *MagicMcpSessionsGetOutputMagicMcpServerProvidersToolFilter      `json:"tool_filter,omitempty"`
	Provider        MagicMcpSessionsGetOutputMagicMcpServerProvidersProvider         `json:"provider"`
	Deployment      MagicMcpSessionsGetOutputMagicMcpServerProvidersDeployment       `json:"deployment"`
	AuthMethod      *MagicMcpSessionsGetOutputMagicMcpServerProvidersAuthMethod      `json:"auth_method,omitempty"`
	AuthCredentials *MagicMcpSessionsGetOutputMagicMcpServerProvidersAuthCredentials `json:"auth_credentials,omitempty"`
	Config          *MagicMcpSessionsGetOutputMagicMcpServerProvidersConfig          `json:"config,omitempty"`
	AuthConfig      *MagicMcpSessionsGetOutputMagicMcpServerProvidersAuthConfig      `json:"auth_config,omitempty"`
	CreatedAt       time.Time                                                        `json:"created_at"`
	UpdatedAt       time.Time                                                        `json:"updated_at"`
	ArchivedAt      *time.Time                                                       `json:"archived_at,omitempty"`
}

// MagicMcpSessionsGetOutputMagicMcpServerAccessOwnerOrganizationMember - Organization member associated with this owner.
type MagicMcpSessionsGetOutputMagicMcpServerAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// MagicMcpSessionsGetOutputMagicMcpServerAccessOwnerConsumer - Consumer and profile associated with this owner.
type MagicMcpSessionsGetOutputMagicMcpServerAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// MagicMcpSessionsGetOutputMagicMcpServerAccessOwner - The actor that owns this resource, if any.
type MagicMcpSessionsGetOutputMagicMcpServerAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *MagicMcpSessionsGetOutputMagicMcpServerAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *MagicMcpSessionsGetOutputMagicMcpServerAccessOwnerConsumer `json:"consumer,omitempty"`
}

// MagicMcpSessionsGetOutputMagicMcpServerAccess - Resource ownership and access settings.
type MagicMcpSessionsGetOutputMagicMcpServerAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *MagicMcpSessionsGetOutputMagicMcpServerAccessOwner `json:"owner,omitempty"`
}

// MagicMcpSessionsGetOutputMagicMcpServer represents the magic mcp sessions get output magic mcp server type.
type MagicMcpSessionsGetOutputMagicMcpServer struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
	// Source - How this Magic MCP server was created or configured
	Source                 string                                             `json:"source"`
	ProviderManagementMode string                                             `json:"provider_management_mode"`
	Endpoints              []MagicMcpSessionsGetOutputMagicMcpServerEndpoints `json:"endpoints"`
	ProviderTemplateId     *string                                            `json:"provider_template_id,omitempty"`
	Providers              []MagicMcpSessionsGetOutputMagicMcpServerProviders `json:"providers"`
	Name                   *string                                            `json:"name,omitempty"`
	Description            *string                                            `json:"description,omitempty"`
	Metadata               map[string]any                                     `json:"metadata"`
	// Access - Resource ownership and access settings.
	Access    MagicMcpSessionsGetOutputMagicMcpServerAccess `json:"access"`
	CreatedAt time.Time                                     `json:"created_at"`
	UpdatedAt time.Time                                     `json:"updated_at"`
}

// MagicMcpSessionsGetOutputMagicMcpEndpointAccessOwnerOrganizationMember - Organization member associated with this owner.
type MagicMcpSessionsGetOutputMagicMcpEndpointAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// MagicMcpSessionsGetOutputMagicMcpEndpointAccessOwnerConsumer - Consumer and profile associated with this owner.
type MagicMcpSessionsGetOutputMagicMcpEndpointAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// MagicMcpSessionsGetOutputMagicMcpEndpointAccessOwner - The actor that owns this resource, if any.
type MagicMcpSessionsGetOutputMagicMcpEndpointAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *MagicMcpSessionsGetOutputMagicMcpEndpointAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *MagicMcpSessionsGetOutputMagicMcpEndpointAccessOwnerConsumer `json:"consumer,omitempty"`
}

// MagicMcpSessionsGetOutputMagicMcpEndpointAccess - Resource ownership and access settings.
type MagicMcpSessionsGetOutputMagicMcpEndpointAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *MagicMcpSessionsGetOutputMagicMcpEndpointAccessOwner `json:"owner,omitempty"`
}

// MagicMcpSessionsGetOutputMagicMcpEndpoint represents the magic mcp sessions get output magic mcp endpoint type.
type MagicMcpSessionsGetOutputMagicMcpEndpoint struct {
	Object      string           `json:"object"`
	Id          string           `json:"id"`
	Status      string           `json:"status"`
	Slug        string           `json:"slug"`
	Url         string           `json:"url"`
	Servers     []map[string]any `json:"servers"`
	Name        *string          `json:"name,omitempty"`
	Description *string          `json:"description,omitempty"`
	Metadata    map[string]any   `json:"metadata"`
	// EnableSkillEditing - Whether consumers connecting through this endpoint with a consumer magic MCP token get tools to read and edit the skills they have access to. The tools are hidden while skills are disabled for the instance.
	EnableSkillEditing bool `json:"enable_skill_editing"`
	// Access - Resource ownership and access settings.
	Access    MagicMcpSessionsGetOutputMagicMcpEndpointAccess `json:"access"`
	CreatedAt time.Time                                       `json:"created_at"`
	UpdatedAt time.Time                                       `json:"updated_at"`
}

// MagicMcpSessionsGetOutput represents the magic mcp sessions get output type.
type MagicMcpSessionsGetOutput struct {
	Object                 string                                     `json:"object"`
	Id                     string                                     `json:"id"`
	MagicMcpServer         *MagicMcpSessionsGetOutputMagicMcpServer   `json:"magic_mcp_server,omitempty"`
	MagicMcpEndpoint       *MagicMcpSessionsGetOutputMagicMcpEndpoint `json:"magic_mcp_endpoint,omitempty"`
	ConsumerProfileId      *string                                    `json:"consumer_profile_id,omitempty"`
	ConsumerIntegrationIds []string                                   `json:"consumer_integration_ids"`
	SessionId              string                                     `json:"session_id"`
	ExpiresAt              *time.Time                                 `json:"expires_at,omitempty"`
	CreatedAt              time.Time                                  `json:"created_at"`
	UpdatedAt              time.Time                                  `json:"updated_at"`
}

// MapMagicMcpSessionsGetOutputFromJSON deserializes JSON data into a MagicMcpSessionsGetOutput.
func MapMagicMcpSessionsGetOutputFromJSON(data []byte) (*MagicMcpSessionsGetOutput, error) {
	var v MagicMcpSessionsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpSessionsGetOutputToJSON serializes a MagicMcpSessionsGetOutput to JSON.
func MapMagicMcpSessionsGetOutputToJSON(v *MagicMcpSessionsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
