package sessiontemplates

import (
	"encoding/json"
	"time"
)

// SessionTemplatesListOutputItemsAccessOwnerOrganizationMember - Organization member associated with this owner.
type SessionTemplatesListOutputItemsAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// SessionTemplatesListOutputItemsAccessOwnerConsumer - Consumer and profile associated with this owner.
type SessionTemplatesListOutputItemsAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// SessionTemplatesListOutputItemsAccessOwner - The actor that owns this resource, if any.
type SessionTemplatesListOutputItemsAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *SessionTemplatesListOutputItemsAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *SessionTemplatesListOutputItemsAccessOwnerConsumer `json:"consumer,omitempty"`
}

// SessionTemplatesListOutputItemsAccess - Resource ownership and access settings.
type SessionTemplatesListOutputItemsAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *SessionTemplatesListOutputItemsAccessOwner `json:"owner,omitempty"`
}

// SessionTemplatesListOutputItemsProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SessionTemplatesListOutputItemsProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// SessionTemplatesListOutputItemsProvidersDeployment represents the session templates list output items providers deployment type.
type SessionTemplatesListOutputItemsProvidersDeployment struct {
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

// SessionTemplatesListOutputItemsProvidersConfig represents the session templates list output items providers config type.
type SessionTemplatesListOutputItemsProvidersConfig struct {
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

// SessionTemplatesListOutputItemsProvidersAuthConfig - Auth config the provider runs with. `null` when there is none.
type SessionTemplatesListOutputItemsProvidersAuthConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
}

// SessionTemplatesListOutputItemsProviders represents the session templates list output items providers type.
type SessionTemplatesListOutputItemsProviders struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the template provider is removed. Removed template providers stay listable.
	Status string `json:"status"`
	// ToolFilter - Tool filter configuration
	ToolFilter SessionTemplatesListOutputItemsProvidersToolFilter `json:"tool_filter"`
	// ProviderId - Provider this template provider connects to.
	ProviderId string `json:"provider_id"`
	// SessionTemplateId - Template this provider belongs to.
	SessionTemplateId string                                             `json:"session_template_id"`
	Deployment        SessionTemplatesListOutputItemsProvidersDeployment `json:"deployment"`
	Config            SessionTemplatesListOutputItemsProvidersConfig     `json:"config"`
	// AuthConfig - Auth config the provider runs with. `null` when there is none.
	AuthConfig *SessionTemplatesListOutputItemsProvidersAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// SessionTemplatesListOutputItems represents the session templates list output items type.
type SessionTemplatesListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the template is deleted. Archived templates stay listable.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// IntegrationInstanceId - Integration instance linked to this template. `null` when there is none.
	IntegrationInstanceId *string `json:"integration_instance_id,omitempty"`
	// IntegrationInstanceGroupId - Integration instance group linked to this template. `null` when there is none.
	IntegrationInstanceGroupId *string `json:"integration_instance_group_id,omitempty"`
	// IdentityActorId - Identity actor linked to this template. `null` when there is none.
	IdentityActorId *string `json:"identity_actor_id,omitempty"`
	// Access - Resource ownership and access settings.
	Access SessionTemplatesListOutputItemsAccess `json:"access"`
	// IdentityId - Identity linked to this template. `null` when there is none.
	IdentityId *string `json:"identity_id,omitempty"`
	// Providers - Providers that sessions created from this template connect to.
	Providers []SessionTemplatesListOutputItemsProviders `json:"providers"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// SessionTemplatesListOutputPagination represents the session templates list output pagination type.
type SessionTemplatesListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SessionTemplatesListOutput represents the session templates list output type.
type SessionTemplatesListOutput struct {
	// Object - String representing the object’s type.
	Object     string                               `json:"object"`
	Items      []SessionTemplatesListOutputItems    `json:"items"`
	Pagination SessionTemplatesListOutputPagination `json:"pagination"`
}

// MapSessionTemplatesListOutputFromJSON deserializes JSON data into a SessionTemplatesListOutput.
func MapSessionTemplatesListOutputFromJSON(data []byte) (*SessionTemplatesListOutput, error) {
	var v SessionTemplatesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionTemplatesListOutputToJSON serializes a SessionTemplatesListOutput to JSON.
func MapSessionTemplatesListOutputToJSON(v *SessionTemplatesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SessionTemplatesListQueryCreatedAt - Only return results whose session template creation time is in this range. Set `gt`, `lt`, or both.
type SessionTemplatesListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SessionTemplatesListQueryUpdatedAt - Only return results whose session template last update time is in this range. Set `gt`, `lt`, or both.
type SessionTemplatesListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SessionTemplatesListQuery represents the session templates list query type.
type SessionTemplatesListQuery struct {
	// Limit - Page size, from 1 to 100. Defaults to 20.
	Limit *float64 `json:"limit,omitempty"`
	// After - Return the page after this point. Pass `pagination.after` if the previous response has it, otherwise the `id` of its last item.
	After *string `json:"after,omitempty"`
	// Before - Return the page before this point. Pass `pagination.before` if the previous response has it, otherwise the `id` of its first item.
	Before *string `json:"before,omitempty"`
	// Cursor - A `cur_` cursor from `pagination.after` or `pagination.before`. Alternative to `after` and `before`.
	Cursor *string `json:"cursor,omitempty"`
	// Order - Sort direction.
	Order *string `json:"order,omitempty"`
	// Ownership - Which resources to return under the resource ownership rules: shared (resources without an owner), usable (shared resources plus resources the caller can use), or all. Defaults to shared when ownership is enforced for the project and to all otherwise.
	Ownership *string `json:"ownership,omitempty"`
	// Status - Filter by status (active, archived)
	Status *any `json:"status,omitempty"`
	// Id - Filter by session template ID(s)
	Id *any `json:"id,omitempty"`
	// SessionId - Filter templates that include sessions with these IDs
	SessionId *any `json:"session_id,omitempty"`
	// SessionProviderId - Filter templates that include session providers with these IDs
	SessionProviderId *any `json:"session_provider_id,omitempty"`
	// ProviderId - Filter templates that include providers with these IDs
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderDeploymentId - Filter templates that include provider deployments with these IDs
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Filter templates that include provider configs with these IDs
	ProviderConfigId *any `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Filter templates that include provider auth configs with these IDs
	ProviderAuthConfigId *any `json:"provider_auth_config_id,omitempty"`
	// CreatedAt - Only return results whose session template creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *SessionTemplatesListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose session template last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *SessionTemplatesListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapSessionTemplatesListQueryFromJSON deserializes JSON data into a SessionTemplatesListQuery.
func MapSessionTemplatesListQueryFromJSON(data []byte) (*SessionTemplatesListQuery, error) {
	var v SessionTemplatesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionTemplatesListQueryToJSON serializes a SessionTemplatesListQuery to JSON.
func MapSessionTemplatesListQueryToJSON(v *SessionTemplatesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
