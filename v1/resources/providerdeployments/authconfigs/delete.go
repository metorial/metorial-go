package authconfigs

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsAuthConfigsDeleteOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ProviderDeploymentsAuthConfigsDeleteOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// ProviderDeploymentsAuthConfigsDeleteOutputDeployment represents the provider deployments auth configs delete output deployment type.
type ProviderDeploymentsAuthConfigsDeleteOutputDeployment struct {
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

// ProviderDeploymentsAuthConfigsDeleteOutputCredentials represents the provider deployments auth configs delete output credentials type.
type ProviderDeploymentsAuthConfigsDeleteOutputCredentials struct {
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

// ProviderDeploymentsAuthConfigsDeleteOutputAuthMethodInputSchema represents the provider deployments auth configs delete output auth method input schema type.
type ProviderDeploymentsAuthConfigsDeleteOutputAuthMethodInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// ProviderDeploymentsAuthConfigsDeleteOutputAuthMethodOutputSchema represents the provider deployments auth configs delete output auth method output schema type.
type ProviderDeploymentsAuthConfigsDeleteOutputAuthMethodOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// ProviderDeploymentsAuthConfigsDeleteOutputAuthMethodScopes represents the provider deployments auth configs delete output auth method scopes type.
type ProviderDeploymentsAuthConfigsDeleteOutputAuthMethodScopes struct {
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

// ProviderDeploymentsAuthConfigsDeleteOutputAuthMethod represents the provider deployments auth configs delete output auth method type.
type ProviderDeploymentsAuthConfigsDeleteOutputAuthMethod struct {
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
	InputSchema  *ProviderDeploymentsAuthConfigsDeleteOutputAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *ProviderDeploymentsAuthConfigsDeleteOutputAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]ProviderDeploymentsAuthConfigsDeleteOutputAuthMethodScopes `json:"scopes,omitempty"`
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

// ProviderDeploymentsAuthConfigsDeleteOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type ProviderDeploymentsAuthConfigsDeleteOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// ProviderDeploymentsAuthConfigsDeleteOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type ProviderDeploymentsAuthConfigsDeleteOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// ProviderDeploymentsAuthConfigsDeleteOutputAccessOwner - The actor that owns this resource, if any.
type ProviderDeploymentsAuthConfigsDeleteOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *ProviderDeploymentsAuthConfigsDeleteOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *ProviderDeploymentsAuthConfigsDeleteOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// ProviderDeploymentsAuthConfigsDeleteOutputAccess - Resource ownership and access settings.
type ProviderDeploymentsAuthConfigsDeleteOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *ProviderDeploymentsAuthConfigsDeleteOutputAccessOwner `json:"owner,omitempty"`
}

// ProviderDeploymentsAuthConfigsDeleteOutput represents the provider deployments auth configs delete output type.
type ProviderDeploymentsAuthConfigsDeleteOutput struct {
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
	ToolFilter  ProviderDeploymentsAuthConfigsDeleteOutputToolFilter   `json:"tool_filter"`
	Deployment  *ProviderDeploymentsAuthConfigsDeleteOutputDeployment  `json:"deployment,omitempty"`
	Credentials *ProviderDeploymentsAuthConfigsDeleteOutputCredentials `json:"credentials,omitempty"`
	AuthMethod  ProviderDeploymentsAuthConfigsDeleteOutputAuthMethod   `json:"auth_method"`
	// Access - Resource ownership and access settings.
	Access ProviderDeploymentsAuthConfigsDeleteOutputAccess `json:"access"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapProviderDeploymentsAuthConfigsDeleteOutputFromJSON deserializes JSON data into a ProviderDeploymentsAuthConfigsDeleteOutput.
func MapProviderDeploymentsAuthConfigsDeleteOutputFromJSON(data []byte) (*ProviderDeploymentsAuthConfigsDeleteOutput, error) {
	var v ProviderDeploymentsAuthConfigsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsAuthConfigsDeleteOutputToJSON serializes a ProviderDeploymentsAuthConfigsDeleteOutput to JSON.
func MapProviderDeploymentsAuthConfigsDeleteOutputToJSON(v *ProviderDeploymentsAuthConfigsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
