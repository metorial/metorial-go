package authconfigs

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsAuthConfigsCreateOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ProviderDeploymentsAuthConfigsCreateOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// ProviderDeploymentsAuthConfigsCreateOutputDeployment represents the provider deployments auth configs create output deployment type.
type ProviderDeploymentsAuthConfigsCreateOutputDeployment struct {
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

// ProviderDeploymentsAuthConfigsCreateOutputCredentials represents the provider deployments auth configs create output credentials type.
type ProviderDeploymentsAuthConfigsCreateOutputCredentials struct {
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

// ProviderDeploymentsAuthConfigsCreateOutputAuthMethodInputSchema represents the provider deployments auth configs create output auth method input schema type.
type ProviderDeploymentsAuthConfigsCreateOutputAuthMethodInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// ProviderDeploymentsAuthConfigsCreateOutputAuthMethodOutputSchema represents the provider deployments auth configs create output auth method output schema type.
type ProviderDeploymentsAuthConfigsCreateOutputAuthMethodOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// ProviderDeploymentsAuthConfigsCreateOutputAuthMethodScopes represents the provider deployments auth configs create output auth method scopes type.
type ProviderDeploymentsAuthConfigsCreateOutputAuthMethodScopes struct {
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

// ProviderDeploymentsAuthConfigsCreateOutputAuthMethod represents the provider deployments auth configs create output auth method type.
type ProviderDeploymentsAuthConfigsCreateOutputAuthMethod struct {
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
	Capabilities map[string]any                                                    `json:"capabilities"`
	InputSchema  *ProviderDeploymentsAuthConfigsCreateOutputAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *ProviderDeploymentsAuthConfigsCreateOutputAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]ProviderDeploymentsAuthConfigsCreateOutputAuthMethodScopes `json:"scopes,omitempty"`
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

// ProviderDeploymentsAuthConfigsCreateOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type ProviderDeploymentsAuthConfigsCreateOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// ProviderDeploymentsAuthConfigsCreateOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type ProviderDeploymentsAuthConfigsCreateOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// ProviderDeploymentsAuthConfigsCreateOutputAccessOwner - The actor that owns this resource, if any.
type ProviderDeploymentsAuthConfigsCreateOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *ProviderDeploymentsAuthConfigsCreateOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *ProviderDeploymentsAuthConfigsCreateOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// ProviderDeploymentsAuthConfigsCreateOutputAccess - Resource ownership and access settings.
type ProviderDeploymentsAuthConfigsCreateOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *ProviderDeploymentsAuthConfigsCreateOutputAccessOwner `json:"owner,omitempty"`
}

// ProviderDeploymentsAuthConfigsCreateOutput represents the provider deployments auth configs create output type.
type ProviderDeploymentsAuthConfigsCreateOutput struct {
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
	ToolFilter  ProviderDeploymentsAuthConfigsCreateOutputToolFilter   `json:"tool_filter"`
	Deployment  *ProviderDeploymentsAuthConfigsCreateOutputDeployment  `json:"deployment,omitempty"`
	Credentials *ProviderDeploymentsAuthConfigsCreateOutputCredentials `json:"credentials,omitempty"`
	AuthMethod  ProviderDeploymentsAuthConfigsCreateOutputAuthMethod   `json:"auth_method"`
	// Access - Resource ownership and access settings.
	Access ProviderDeploymentsAuthConfigsCreateOutputAccess `json:"access"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapProviderDeploymentsAuthConfigsCreateOutputFromJSON deserializes JSON data into a ProviderDeploymentsAuthConfigsCreateOutput.
func MapProviderDeploymentsAuthConfigsCreateOutputFromJSON(data []byte) (*ProviderDeploymentsAuthConfigsCreateOutput, error) {
	var v ProviderDeploymentsAuthConfigsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsAuthConfigsCreateOutputToJSON serializes a ProviderDeploymentsAuthConfigsCreateOutput to JSON.
func MapProviderDeploymentsAuthConfigsCreateOutputToJSON(v *ProviderDeploymentsAuthConfigsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderDeploymentsAuthConfigsCreateBody represents the provider deployments auth configs create body type.
type ProviderDeploymentsAuthConfigsCreateBody struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata    *map[string]any `json:"metadata,omitempty"`
	ToolFilters *any            `json:"tool_filters,omitempty"`
	// ProviderAuthMethodId - The authentication method this config uses (e.g., OAuth, API key)
	ProviderAuthMethodId string `json:"provider_auth_method_id"`
	// ProviderDeploymentId - The provider deployment this auth config is associated with (if applicable)
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	// Value - Authentication config payload
	Value map[string]any `json:"value"`
}

// MapProviderDeploymentsAuthConfigsCreateBodyFromJSON deserializes JSON data into a ProviderDeploymentsAuthConfigsCreateBody.
func MapProviderDeploymentsAuthConfigsCreateBodyFromJSON(data []byte) (*ProviderDeploymentsAuthConfigsCreateBody, error) {
	var v ProviderDeploymentsAuthConfigsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsAuthConfigsCreateBodyToJSON serializes a ProviderDeploymentsAuthConfigsCreateBody to JSON.
func MapProviderDeploymentsAuthConfigsCreateBodyToJSON(v *ProviderDeploymentsAuthConfigsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
