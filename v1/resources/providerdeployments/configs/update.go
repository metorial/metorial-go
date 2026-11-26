package configs

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsConfigsUpdateOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ProviderDeploymentsConfigsUpdateOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// ProviderDeploymentsConfigsUpdateOutputDeployment represents the provider deployments configs update output deployment type.
type ProviderDeploymentsConfigsUpdateOutputDeployment struct {
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

// ProviderDeploymentsConfigsUpdateOutputFromVaultDeployment represents the provider deployments configs update output from vault deployment type.
type ProviderDeploymentsConfigsUpdateOutputFromVaultDeployment struct {
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

// ProviderDeploymentsConfigsUpdateOutputFromVault represents the provider deployments configs update output from vault type.
type ProviderDeploymentsConfigsUpdateOutputFromVault struct {
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
	Deployment *ProviderDeploymentsConfigsUpdateOutputFromVaultDeployment `json:"deployment,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsConfigsUpdateOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type ProviderDeploymentsConfigsUpdateOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// ProviderDeploymentsConfigsUpdateOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type ProviderDeploymentsConfigsUpdateOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// ProviderDeploymentsConfigsUpdateOutputAccessOwner - The actor that owns this resource, if any.
type ProviderDeploymentsConfigsUpdateOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *ProviderDeploymentsConfigsUpdateOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *ProviderDeploymentsConfigsUpdateOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// ProviderDeploymentsConfigsUpdateOutputAccess - Resource ownership and access settings.
type ProviderDeploymentsConfigsUpdateOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *ProviderDeploymentsConfigsUpdateOutputAccessOwner `json:"owner,omitempty"`
}

// ProviderDeploymentsConfigsUpdateOutput represents the provider deployments configs update output type.
type ProviderDeploymentsConfigsUpdateOutput struct {
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
	ToolFilter ProviderDeploymentsConfigsUpdateOutputToolFilter `json:"tool_filter"`
	// ProviderId - Provider this config belongs to.
	ProviderId string `json:"provider_id"`
	// SpecificationId - Specification this config follows.
	SpecificationId string                                            `json:"specification_id"`
	Deployment      *ProviderDeploymentsConfigsUpdateOutputDeployment `json:"deployment,omitempty"`
	FromVault       *ProviderDeploymentsConfigsUpdateOutputFromVault  `json:"from_vault,omitempty"`
	// Access - Resource ownership and access settings.
	Access ProviderDeploymentsConfigsUpdateOutputAccess `json:"access"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapProviderDeploymentsConfigsUpdateOutputFromJSON deserializes JSON data into a ProviderDeploymentsConfigsUpdateOutput.
func MapProviderDeploymentsConfigsUpdateOutputFromJSON(data []byte) (*ProviderDeploymentsConfigsUpdateOutput, error) {
	var v ProviderDeploymentsConfigsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsConfigsUpdateOutputToJSON serializes a ProviderDeploymentsConfigsUpdateOutput to JSON.
func MapProviderDeploymentsConfigsUpdateOutputToJSON(v *ProviderDeploymentsConfigsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderDeploymentsConfigsUpdateBody represents the provider deployments configs update body type.
type ProviderDeploymentsConfigsUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the provider config. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
}

// MapProviderDeploymentsConfigsUpdateBodyFromJSON deserializes JSON data into a ProviderDeploymentsConfigsUpdateBody.
func MapProviderDeploymentsConfigsUpdateBodyFromJSON(data []byte) (*ProviderDeploymentsConfigsUpdateBody, error) {
	var v ProviderDeploymentsConfigsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsConfigsUpdateBodyToJSON serializes a ProviderDeploymentsConfigsUpdateBody to JSON.
func MapProviderDeploymentsConfigsUpdateBodyToJSON(v *ProviderDeploymentsConfigsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
