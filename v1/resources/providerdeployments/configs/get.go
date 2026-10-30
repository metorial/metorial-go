package configs

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsConfigsGetOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ProviderDeploymentsConfigsGetOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// ProviderDeploymentsConfigsGetOutputDeployment represents the provider deployments configs get output deployment type.
type ProviderDeploymentsConfigsGetOutputDeployment struct {
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

// ProviderDeploymentsConfigsGetOutputFromVaultDeployment represents the provider deployments configs get output from vault deployment type.
type ProviderDeploymentsConfigsGetOutputFromVaultDeployment struct {
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

// ProviderDeploymentsConfigsGetOutputFromVault represents the provider deployments configs get output from vault type.
type ProviderDeploymentsConfigsGetOutputFromVault struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique config vault identifier
	Id string `json:"id"`
	// Status - Status of the config vault
	Status string `json:"status"`
	// Name - Display name
	Name string `json:"name"`
	// Description - Description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider ID
	ProviderId string                                                  `json:"provider_id"`
	Deployment *ProviderDeploymentsConfigsGetOutputFromVaultDeployment `json:"deployment,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsConfigsGetOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type ProviderDeploymentsConfigsGetOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// ProviderDeploymentsConfigsGetOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type ProviderDeploymentsConfigsGetOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// ProviderDeploymentsConfigsGetOutputAccessOwner - The actor that owns this resource, if any.
type ProviderDeploymentsConfigsGetOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *ProviderDeploymentsConfigsGetOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *ProviderDeploymentsConfigsGetOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// ProviderDeploymentsConfigsGetOutputAccess - Resource ownership and access settings.
type ProviderDeploymentsConfigsGetOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *ProviderDeploymentsConfigsGetOutputAccessOwner `json:"owner,omitempty"`
}

// ProviderDeploymentsConfigsGetOutput represents the provider deployments configs get output type.
type ProviderDeploymentsConfigsGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique config identifier
	Id string `json:"id"`
	// Status - Config status
	Status string `json:"status"`
	// IsDefault - Whether this is the default config
	IsDefault bool `json:"is_default"`
	// Name - Display name
	Name *string `json:"name,omitempty"`
	// Description - Description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter ProviderDeploymentsConfigsGetOutputToolFilter `json:"tool_filter"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// SpecificationId - Specification ID
	SpecificationId string                                         `json:"specification_id"`
	Deployment      *ProviderDeploymentsConfigsGetOutputDeployment `json:"deployment,omitempty"`
	FromVault       *ProviderDeploymentsConfigsGetOutputFromVault  `json:"from_vault,omitempty"`
	// Access - Resource ownership and access settings.
	Access ProviderDeploymentsConfigsGetOutputAccess `json:"access"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapProviderDeploymentsConfigsGetOutputFromJSON deserializes JSON data into a ProviderDeploymentsConfigsGetOutput.
func MapProviderDeploymentsConfigsGetOutputFromJSON(data []byte) (*ProviderDeploymentsConfigsGetOutput, error) {
	var v ProviderDeploymentsConfigsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsConfigsGetOutputToJSON serializes a ProviderDeploymentsConfigsGetOutput to JSON.
func MapProviderDeploymentsConfigsGetOutputToJSON(v *ProviderDeploymentsConfigsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
