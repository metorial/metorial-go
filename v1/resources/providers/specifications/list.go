package specifications

import (
	"encoding/json"
	"time"
)

// ProvidersSpecificationsListOutputItemsToolsInputSchema represents the providers specifications list output items tools input schema type.
type ProvidersSpecificationsListOutputItemsToolsInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the tool input parameters
	Schema map[string]any `json:"schema"`
}

// ProvidersSpecificationsListOutputItemsToolsOutputSchema represents the providers specifications list output items tools output schema type.
type ProvidersSpecificationsListOutputItemsToolsOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the tool output format
	Schema map[string]any `json:"schema"`
}

// ProvidersSpecificationsListOutputItemsToolsTags represents the providers specifications list output items tools tags type.
type ProvidersSpecificationsListOutputItemsToolsTags struct {
	// Destructive - Whether the tool changes or deletes data.
	Destructive *bool `json:"destructive,omitempty"`
	// ReadOnly - Whether the tool only reads data.
	ReadOnly *bool `json:"read_only,omitempty"`
}

// ProvidersSpecificationsListOutputItemsTools represents the providers specifications list output items tools type.
type ProvidersSpecificationsListOutputItemsTools struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Key - Key used to call the tool over MCP.
	Key string `json:"key"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Capabilities - What the tool can do.
	Capabilities map[string]any `json:"capabilities"`
	// Constraints - Rules for using the tool.
	Constraints []string `json:"constraints"`
	// Instructions - Guidance for calling the tool correctly.
	Instructions []string                                                 `json:"instructions"`
	InputSchema  *ProvidersSpecificationsListOutputItemsToolsInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *ProvidersSpecificationsListOutputItemsToolsOutputSchema `json:"output_schema,omitempty"`
	Tags         *ProvidersSpecificationsListOutputItemsToolsTags         `json:"tags,omitempty"`
	// ToolCategoryKeys - Keys of the tool categories this tool belongs to
	ToolCategoryKeys []string `json:"tool_category_keys"`
	// SpecificationId - Specification this tool belongs to.
	SpecificationId string `json:"specification_id"`
	// ProviderId - Provider this tool belongs to.
	ProviderId string `json:"provider_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProvidersSpecificationsListOutputItemsToolCategories represents the providers specifications list output items tool categories type.
type ProvidersSpecificationsListOutputItemsToolCategories struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Key - Stable key for the tool category.
	Key string `json:"key"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// ImageUrl - URL of an icon or logo for the tool category
	ImageUrl *string `json:"image_url,omitempty"`
	// ToolCount - Number of tools in this specification that belong to the category
	ToolCount float64 `json:"tool_count"`
}

// ProvidersSpecificationsListOutputItemsAuthMethodsInputSchema represents the providers specifications list output items auth methods input schema type.
type ProvidersSpecificationsListOutputItemsAuthMethodsInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// ProvidersSpecificationsListOutputItemsAuthMethodsOutputSchema represents the providers specifications list output items auth methods output schema type.
type ProvidersSpecificationsListOutputItemsAuthMethodsOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// ProvidersSpecificationsListOutputItemsAuthMethodsScopes represents the providers specifications list output items auth methods scopes type.
type ProvidersSpecificationsListOutputItemsAuthMethodsScopes struct {
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

// ProvidersSpecificationsListOutputItemsAuthMethods represents the providers specifications list output items auth methods type.
type ProvidersSpecificationsListOutputItemsAuthMethods struct {
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
	// Capabilities - What this auth method supports.
	Capabilities map[string]any                                                 `json:"capabilities"`
	InputSchema  *ProvidersSpecificationsListOutputItemsAuthMethodsInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *ProvidersSpecificationsListOutputItemsAuthMethodsOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]ProvidersSpecificationsListOutputItemsAuthMethodsScopes `json:"scopes,omitempty"`
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

// ProvidersSpecificationsListOutputItems represents the providers specifications list output items type.
type ProvidersSpecificationsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Key - Stable key identifying the specification.
	Key string `json:"key"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// ConfigSchema - JSON Schema defining the configuration structure
	ConfigSchema map[string]any `json:"config_schema"`
	// ConfigVisibility - Whether config values are stored encrypted or in plain text.
	ConfigVisibility string `json:"config_visibility"`
	// Tools - Tools this specification offers.
	Tools []ProvidersSpecificationsListOutputItemsTools `json:"tools"`
	// ToolCategories - Tool categories that group the tools of this specification
	ToolCategories []ProvidersSpecificationsListOutputItemsToolCategories `json:"tool_categories"`
	// AuthMethods - Auth methods this specification offers.
	AuthMethods []ProvidersSpecificationsListOutputItemsAuthMethods `json:"auth_methods"`
	// ProviderId - Provider this specification belongs to.
	ProviderId string `json:"provider_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProvidersSpecificationsListOutputPagination represents the providers specifications list output pagination type.
type ProvidersSpecificationsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ProvidersSpecificationsListOutput represents the providers specifications list output type.
type ProvidersSpecificationsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                      `json:"object"`
	Items      []ProvidersSpecificationsListOutputItems    `json:"items"`
	Pagination ProvidersSpecificationsListOutputPagination `json:"pagination"`
}

// MapProvidersSpecificationsListOutputFromJSON deserializes JSON data into a ProvidersSpecificationsListOutput.
func MapProvidersSpecificationsListOutputFromJSON(data []byte) (*ProvidersSpecificationsListOutput, error) {
	var v ProvidersSpecificationsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProvidersSpecificationsListOutputToJSON serializes a ProvidersSpecificationsListOutput to JSON.
func MapProvidersSpecificationsListOutputToJSON(v *ProvidersSpecificationsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProvidersSpecificationsListQueryCreatedAt - Only return results whose provider specification creation time is in this range. Set `gt`, `lt`, or both.
type ProvidersSpecificationsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ProvidersSpecificationsListQueryUpdatedAt - Only return results whose provider specification last update time is in this range. Set `gt`, `lt`, or both.
type ProvidersSpecificationsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ProvidersSpecificationsListQuery represents the providers specifications list query type.
type ProvidersSpecificationsListQuery struct {
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
	// Id - Filter by specification ID(s)
	Id *any `json:"id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderVersionId - Filter by provider version ID(s)
	ProviderVersionId *any `json:"provider_version_id,omitempty"`
	// ProviderDeploymentId - Filter by provider deployment ID(s)
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Filter by provider config ID(s)
	ProviderConfigId *any `json:"provider_config_id,omitempty"`
	// Search - Match a provider specification by its exact ID
	Search *string `json:"search,omitempty"`
	// CreatedAt - Only return results whose provider specification creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *ProvidersSpecificationsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose provider specification last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *ProvidersSpecificationsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapProvidersSpecificationsListQueryFromJSON deserializes JSON data into a ProvidersSpecificationsListQuery.
func MapProvidersSpecificationsListQueryFromJSON(data []byte) (*ProvidersSpecificationsListQuery, error) {
	var v ProvidersSpecificationsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProvidersSpecificationsListQueryToJSON serializes a ProvidersSpecificationsListQuery to JSON.
func MapProvidersSpecificationsListQueryToJSON(v *ProvidersSpecificationsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
