package configs

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsConfigsDeleteOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ProviderDeploymentsConfigsDeleteOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// ProviderDeploymentsConfigsDeleteOutputDeployment represents the provider deployments configs delete output deployment type.
type ProviderDeploymentsConfigsDeleteOutputDeployment struct {
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

// ProviderDeploymentsConfigsDeleteOutputFromVaultDeployment represents the provider deployments configs delete output from vault deployment type.
type ProviderDeploymentsConfigsDeleteOutputFromVaultDeployment struct {
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

// ProviderDeploymentsConfigsDeleteOutputFromVault represents the provider deployments configs delete output from vault type.
type ProviderDeploymentsConfigsDeleteOutputFromVault struct {
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
	ProviderId string                                                     `json:"provider_id"`
	Deployment *ProviderDeploymentsConfigsDeleteOutputFromVaultDeployment `json:"deployment,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsConfigsDeleteOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type ProviderDeploymentsConfigsDeleteOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// ProviderDeploymentsConfigsDeleteOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type ProviderDeploymentsConfigsDeleteOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// ProviderDeploymentsConfigsDeleteOutputAccessOwner - The actor that owns this resource, if any.
type ProviderDeploymentsConfigsDeleteOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *ProviderDeploymentsConfigsDeleteOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *ProviderDeploymentsConfigsDeleteOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// ProviderDeploymentsConfigsDeleteOutputAccess - Resource ownership and access settings.
type ProviderDeploymentsConfigsDeleteOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *ProviderDeploymentsConfigsDeleteOutputAccessOwner `json:"owner,omitempty"`
}

// ProviderDeploymentsConfigsDeleteOutput represents the provider deployments configs delete output type.
type ProviderDeploymentsConfigsDeleteOutput struct {
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
	ToolFilter ProviderDeploymentsConfigsDeleteOutputToolFilter `json:"tool_filter"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// SpecificationId - Specification ID
	SpecificationId string                                            `json:"specification_id"`
	Deployment      *ProviderDeploymentsConfigsDeleteOutputDeployment `json:"deployment,omitempty"`
	FromVault       *ProviderDeploymentsConfigsDeleteOutputFromVault  `json:"from_vault,omitempty"`
	// Access - Resource ownership and access settings.
	Access ProviderDeploymentsConfigsDeleteOutputAccess `json:"access"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapProviderDeploymentsConfigsDeleteOutputFromJSON deserializes JSON data into a ProviderDeploymentsConfigsDeleteOutput.
func MapProviderDeploymentsConfigsDeleteOutputFromJSON(data []byte) (*ProviderDeploymentsConfigsDeleteOutput, error) {
	var v ProviderDeploymentsConfigsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsConfigsDeleteOutputToJSON serializes a ProviderDeploymentsConfigsDeleteOutput to JSON.
func MapProviderDeploymentsConfigsDeleteOutputToJSON(v *ProviderDeploymentsConfigsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
