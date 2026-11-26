package tools

import (
	"encoding/json"
	"time"
)

// ProvidersToolsListOutputItemsInputSchema represents the providers tools list output items input schema type.
type ProvidersToolsListOutputItemsInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the tool input parameters
	Schema map[string]any `json:"schema"`
}

// ProvidersToolsListOutputItemsOutputSchema represents the providers tools list output items output schema type.
type ProvidersToolsListOutputItemsOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the tool output format
	Schema map[string]any `json:"schema"`
}

// ProvidersToolsListOutputItemsTags represents the providers tools list output items tags type.
type ProvidersToolsListOutputItemsTags struct {
	// Destructive - Whether the tool changes or deletes data.
	Destructive *bool `json:"destructive,omitempty"`
	// ReadOnly - Whether the tool only reads data.
	ReadOnly *bool `json:"read_only,omitempty"`
}

// ProvidersToolsListOutputItems represents the providers tools list output items type.
type ProvidersToolsListOutputItems struct {
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
	Instructions []string                                   `json:"instructions"`
	InputSchema  *ProvidersToolsListOutputItemsInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *ProvidersToolsListOutputItemsOutputSchema `json:"output_schema,omitempty"`
	Tags         *ProvidersToolsListOutputItemsTags         `json:"tags,omitempty"`
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

// ProvidersToolsListOutputPagination represents the providers tools list output pagination type.
type ProvidersToolsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ProvidersToolsListOutput represents the providers tools list output type.
type ProvidersToolsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                             `json:"object"`
	Items      []ProvidersToolsListOutputItems    `json:"items"`
	Pagination ProvidersToolsListOutputPagination `json:"pagination"`
}

// MapProvidersToolsListOutputFromJSON deserializes JSON data into a ProvidersToolsListOutput.
func MapProvidersToolsListOutputFromJSON(data []byte) (*ProvidersToolsListOutput, error) {
	var v ProvidersToolsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProvidersToolsListOutputToJSON serializes a ProvidersToolsListOutput to JSON.
func MapProvidersToolsListOutputToJSON(v *ProvidersToolsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProvidersToolsListQuery represents the providers tools list query type.
type ProvidersToolsListQuery struct {
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
	// ProviderVersionId - Provider version to list tools for.
	ProviderVersionId string `json:"provider_version_id"`
	// ProviderAuthMethodId - Filter to tools compatible with one auth method ID, or the common compatible subset for multiple auth method IDs.
	ProviderAuthMethodId *any `json:"provider_auth_method_id,omitempty"`
}

// MapProvidersToolsListQueryFromJSON deserializes JSON data into a ProvidersToolsListQuery.
func MapProvidersToolsListQueryFromJSON(data []byte) (*ProvidersToolsListQuery, error) {
	var v ProvidersToolsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProvidersToolsListQueryToJSON serializes a ProvidersToolsListQuery to JSON.
func MapProvidersToolsListQueryToJSON(v *ProvidersToolsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
