package authmethods

import (
	"encoding/json"
	"time"
)

// ProvidersAuthMethodsListOutputItemsInputSchema represents the providers auth methods list output items input schema type.
type ProvidersAuthMethodsListOutputItemsInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// ProvidersAuthMethodsListOutputItemsOutputSchema represents the providers auth methods list output items output schema type.
type ProvidersAuthMethodsListOutputItemsOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// ProvidersAuthMethodsListOutputItemsScopes represents the providers auth methods list output items scopes type.
type ProvidersAuthMethodsListOutputItemsScopes struct {
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

// ProvidersAuthMethodsListOutputItems represents the providers auth methods list output items type.
type ProvidersAuthMethodsListOutputItems struct {
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
	Capabilities map[string]any                                   `json:"capabilities"`
	InputSchema  *ProvidersAuthMethodsListOutputItemsInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *ProvidersAuthMethodsListOutputItemsOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]ProvidersAuthMethodsListOutputItemsScopes `json:"scopes,omitempty"`
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

// ProvidersAuthMethodsListOutputPagination represents the providers auth methods list output pagination type.
type ProvidersAuthMethodsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ProvidersAuthMethodsListOutput represents the providers auth methods list output type.
type ProvidersAuthMethodsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                   `json:"object"`
	Items      []ProvidersAuthMethodsListOutputItems    `json:"items"`
	Pagination ProvidersAuthMethodsListOutputPagination `json:"pagination"`
}

// MapProvidersAuthMethodsListOutputFromJSON deserializes JSON data into a ProvidersAuthMethodsListOutput.
func MapProvidersAuthMethodsListOutputFromJSON(data []byte) (*ProvidersAuthMethodsListOutput, error) {
	var v ProvidersAuthMethodsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProvidersAuthMethodsListOutputToJSON serializes a ProvidersAuthMethodsListOutput to JSON.
func MapProvidersAuthMethodsListOutputToJSON(v *ProvidersAuthMethodsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProvidersAuthMethodsListQuery represents the providers auth methods list query type.
type ProvidersAuthMethodsListQuery struct {
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
	// ProviderVersionId - Provider version to list auth methods for.
	ProviderVersionId string `json:"provider_version_id"`
	// Adapter - Only return auth methods supporting this adapter.
	Adapter *string `json:"adapter,omitempty"`
}

// MapProvidersAuthMethodsListQueryFromJSON deserializes JSON data into a ProvidersAuthMethodsListQuery.
func MapProvidersAuthMethodsListQueryFromJSON(data []byte) (*ProvidersAuthMethodsListQuery, error) {
	var v ProvidersAuthMethodsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProvidersAuthMethodsListQueryToJSON serializes a ProvidersAuthMethodsListQuery to JSON.
func MapProvidersAuthMethodsListQueryToJSON(v *ProvidersAuthMethodsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
