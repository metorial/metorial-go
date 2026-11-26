package specifications

import (
	"encoding/json"
	"time"
)

// ProvidersSpecificationsGetOutputToolsInputSchema represents the providers specifications get output tools input schema type.
type ProvidersSpecificationsGetOutputToolsInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the tool input parameters
	Schema map[string]any `json:"schema"`
}

// ProvidersSpecificationsGetOutputToolsOutputSchema represents the providers specifications get output tools output schema type.
type ProvidersSpecificationsGetOutputToolsOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the tool output format
	Schema map[string]any `json:"schema"`
}

// ProvidersSpecificationsGetOutputToolsTags represents the providers specifications get output tools tags type.
type ProvidersSpecificationsGetOutputToolsTags struct {
	// Destructive - Whether the tool changes or deletes data.
	Destructive *bool `json:"destructive,omitempty"`
	// ReadOnly - Whether the tool only reads data.
	ReadOnly *bool `json:"read_only,omitempty"`
}

// ProvidersSpecificationsGetOutputTools represents the providers specifications get output tools type.
type ProvidersSpecificationsGetOutputTools struct {
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
	Instructions []string                                           `json:"instructions"`
	InputSchema  *ProvidersSpecificationsGetOutputToolsInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *ProvidersSpecificationsGetOutputToolsOutputSchema `json:"output_schema,omitempty"`
	Tags         *ProvidersSpecificationsGetOutputToolsTags         `json:"tags,omitempty"`
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

// ProvidersSpecificationsGetOutputToolCategories represents the providers specifications get output tool categories type.
type ProvidersSpecificationsGetOutputToolCategories struct {
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

// ProvidersSpecificationsGetOutputAuthMethodsInputSchema represents the providers specifications get output auth methods input schema type.
type ProvidersSpecificationsGetOutputAuthMethodsInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// ProvidersSpecificationsGetOutputAuthMethodsOutputSchema represents the providers specifications get output auth methods output schema type.
type ProvidersSpecificationsGetOutputAuthMethodsOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// ProvidersSpecificationsGetOutputAuthMethodsScopes represents the providers specifications get output auth methods scopes type.
type ProvidersSpecificationsGetOutputAuthMethodsScopes struct {
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

// ProvidersSpecificationsGetOutputAuthMethods represents the providers specifications get output auth methods type.
type ProvidersSpecificationsGetOutputAuthMethods struct {
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
	Capabilities map[string]any                                           `json:"capabilities"`
	InputSchema  *ProvidersSpecificationsGetOutputAuthMethodsInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *ProvidersSpecificationsGetOutputAuthMethodsOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]ProvidersSpecificationsGetOutputAuthMethodsScopes `json:"scopes,omitempty"`
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

// ProvidersSpecificationsGetOutput represents the providers specifications get output type.
type ProvidersSpecificationsGetOutput struct {
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
	Tools []ProvidersSpecificationsGetOutputTools `json:"tools"`
	// ToolCategories - Tool categories that group the tools of this specification
	ToolCategories []ProvidersSpecificationsGetOutputToolCategories `json:"tool_categories"`
	// AuthMethods - Auth methods this specification offers.
	AuthMethods []ProvidersSpecificationsGetOutputAuthMethods `json:"auth_methods"`
	// ProviderId - Provider this specification belongs to.
	ProviderId string `json:"provider_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapProvidersSpecificationsGetOutputFromJSON deserializes JSON data into a ProvidersSpecificationsGetOutput.
func MapProvidersSpecificationsGetOutputFromJSON(data []byte) (*ProvidersSpecificationsGetOutput, error) {
	var v ProvidersSpecificationsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProvidersSpecificationsGetOutputToJSON serializes a ProvidersSpecificationsGetOutput to JSON.
func MapProvidersSpecificationsGetOutputToJSON(v *ProvidersSpecificationsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
