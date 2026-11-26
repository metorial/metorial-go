package configs

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsConfigsCreateOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ProviderDeploymentsConfigsCreateOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// ProviderDeploymentsConfigsCreateOutputDeployment represents the provider deployments configs create output deployment type.
type ProviderDeploymentsConfigsCreateOutputDeployment struct {
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

// ProviderDeploymentsConfigsCreateOutputFromVaultDeployment represents the provider deployments configs create output from vault deployment type.
type ProviderDeploymentsConfigsCreateOutputFromVaultDeployment struct {
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

// ProviderDeploymentsConfigsCreateOutputFromVault represents the provider deployments configs create output from vault type.
type ProviderDeploymentsConfigsCreateOutputFromVault struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Archived once the vault is archived. Archived vaults stay listable.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider this vault belongs to.
	ProviderId string                                                     `json:"provider_id"`
	Deployment *ProviderDeploymentsConfigsCreateOutputFromVaultDeployment `json:"deployment,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsConfigsCreateOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type ProviderDeploymentsConfigsCreateOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// ProviderDeploymentsConfigsCreateOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type ProviderDeploymentsConfigsCreateOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// ProviderDeploymentsConfigsCreateOutputAccessOwner - The actor that owns this resource, if any.
type ProviderDeploymentsConfigsCreateOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *ProviderDeploymentsConfigsCreateOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *ProviderDeploymentsConfigsCreateOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// ProviderDeploymentsConfigsCreateOutputAccess - Resource ownership and access settings.
type ProviderDeploymentsConfigsCreateOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *ProviderDeploymentsConfigsCreateOutputAccessOwner `json:"owner,omitempty"`
}

// ProviderDeploymentsConfigsCreateOutput represents the provider deployments configs create output type.
type ProviderDeploymentsConfigsCreateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Archived once the config is archived. Archived configs stay listable.
	Status string `json:"status"`
	// IsDefault - The deployment falls back to this config when none is specified.
	IsDefault bool `json:"is_default"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter ProviderDeploymentsConfigsCreateOutputToolFilter `json:"tool_filter"`
	// ProviderId - Provider this config belongs to.
	ProviderId string `json:"provider_id"`
	// SpecificationId - Specification this config follows.
	SpecificationId string                                            `json:"specification_id"`
	Deployment      *ProviderDeploymentsConfigsCreateOutputDeployment `json:"deployment,omitempty"`
	FromVault       *ProviderDeploymentsConfigsCreateOutputFromVault  `json:"from_vault,omitempty"`
	// Access - Resource ownership and access settings.
	Access ProviderDeploymentsConfigsCreateOutputAccess `json:"access"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapProviderDeploymentsConfigsCreateOutputFromJSON deserializes JSON data into a ProviderDeploymentsConfigsCreateOutput.
func MapProviderDeploymentsConfigsCreateOutputFromJSON(data []byte) (*ProviderDeploymentsConfigsCreateOutput, error) {
	var v ProviderDeploymentsConfigsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsConfigsCreateOutputToJSON serializes a ProviderDeploymentsConfigsCreateOutput to JSON.
func MapProviderDeploymentsConfigsCreateOutputToJSON(v *ProviderDeploymentsConfigsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderDeploymentsConfigsCreateBody represents the provider deployments configs create body type.
type ProviderDeploymentsConfigsCreateBody struct {
	// ProviderId - Provider to create the config for.
	ProviderId string `json:"provider_id"`
	// ProviderDeploymentId - Deployment the config is locked to.
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the provider config. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
	// Value - Provider-specific configuration values
	Value *map[string]any `json:"value,omitempty"`
	// ProviderConfigVaultId - Config vault ID to use as template
	ProviderConfigVaultId *string `json:"provider_config_vault_id,omitempty"`
}

// MapProviderDeploymentsConfigsCreateBodyFromJSON deserializes JSON data into a ProviderDeploymentsConfigsCreateBody.
func MapProviderDeploymentsConfigsCreateBodyFromJSON(data []byte) (*ProviderDeploymentsConfigsCreateBody, error) {
	var v ProviderDeploymentsConfigsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsConfigsCreateBodyToJSON serializes a ProviderDeploymentsConfigsCreateBody to JSON.
func MapProviderDeploymentsConfigsCreateBodyToJSON(v *ProviderDeploymentsConfigsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
