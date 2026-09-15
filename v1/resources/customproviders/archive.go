package customproviders

import (
	"encoding/json"
	"time"
)

// CustomProvidersArchiveOutputDraftContainerImage represents the custom providers archive output draft container image type.
type CustomProvidersArchiveOutputDraftContainerImage struct {
	// Object - String representing the container image draft type
	Object string `json:"object"`
	// ContainerRegistry - URL of the container registry
	ContainerRegistry string `json:"container_registry"`
	// ContainerImageTag - Tag of the container image
	ContainerImageTag string `json:"container_image_tag"`
	// ContainerImage - Name of the container image
	ContainerImage string `json:"container_image"`
}

// CustomProvidersArchiveOutputDraftRemoteMcpServer represents the custom providers archive output draft remote mcp server type.
type CustomProvidersArchiveOutputDraftRemoteMcpServer struct {
	// Object - String representing the remote MCP server draft type
	Object string `json:"object"`
	// Url - URL of the remote MCP server
	Url string `json:"url"`
	// Transport - Transport protocol for connecting to the remote MCP server
	Transport string `json:"transport"`
}

// CustomProvidersArchiveOutputDraftConfigSchema represents the custom providers archive output draft config schema type.
type CustomProvidersArchiveOutputDraftConfigSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the configuration fields for the custom provider
	Schema map[string]any `json:"schema"`
}

// CustomProvidersArchiveOutputDraftConfig represents the custom providers archive output draft config type.
type CustomProvidersArchiveOutputDraftConfig struct {
	// Object - String representing the custom provider config draft type
	Object string                                        `json:"object"`
	Schema CustomProvidersArchiveOutputDraftConfigSchema `json:"schema"`
	// Transformer - Optional jsonata transformer function for the configuration.
	Transformer string `json:"transformer"`
}

// CustomProvidersArchiveOutputDraft represents the custom providers archive output draft type.
type CustomProvidersArchiveOutputDraft struct {
	// Object - String representing the draft's type
	Object          string                                            `json:"object"`
	ContainerImage  *CustomProvidersArchiveOutputDraftContainerImage  `json:"container_image,omitempty"`
	RemoteMcpServer *CustomProvidersArchiveOutputDraftRemoteMcpServer `json:"remote_mcp_server,omitempty"`
	Config          CustomProvidersArchiveOutputDraftConfig           `json:"config"`
}

// CustomProvidersArchiveOutputScmRepoProvider represents the custom providers archive output scm repo provider type.
type CustomProvidersArchiveOutputScmRepoProvider struct {
	Object string `json:"object"`
	// Type - SCM provider type
	Type string `json:"type"`
	// Id - External provider identifier
	Id string `json:"id"`
	// Name - Repository name on the provider
	Name string `json:"name"`
	// Owner - Repository owner on the provider
	Owner string `json:"owner"`
}

// CustomProvidersArchiveOutputScmRepo represents the custom providers archive output scm repo type.
type CustomProvidersArchiveOutputScmRepo struct {
	Object string `json:"object"`
	// Id - Unique repository identifier
	Id       string                                      `json:"id"`
	Provider CustomProvidersArchiveOutputScmRepoProvider `json:"provider"`
	// Url - Repository URL
	Url string `json:"url"`
	// IsPrivate - Whether the repository is private
	IsPrivate bool `json:"is_private"`
	// DefaultBranch - Default branch name
	DefaultBranch string `json:"default_branch"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersArchiveOutputProviderPublisher represents the custom providers archive output provider publisher type.
type CustomProvidersArchiveOutputProviderPublisher struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique publisher identifier
	Id string `json:"id"`
	// Name - Display name of the publisher
	Name string `json:"name"`
	// Description - Brief description of the publisher
	Description *string `json:"description,omitempty"`
	// ImageUrl - URL of the publisher logo
	ImageUrl string `json:"image_url"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomProvidersArchiveOutputProviderCurrentVersion represents the custom providers archive output provider current version type.
type CustomProvidersArchiveOutputProviderCurrentVersion struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique version identifier
	Id string `json:"id"`
	// Version - Version identifier string
	Version string `json:"version"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// IsCurrent - Whether this is the current version
	IsCurrent bool `json:"is_current"`
	// Name - Version name
	Name string `json:"name"`
	// Description - Version description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// SpecificationId - Specification ID
	SpecificationId *string `json:"specification_id,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomProvidersArchiveOutputProviderOauthAutoRegistration represents the custom providers archive output provider oauth auto registration type.
type CustomProvidersArchiveOutputProviderOauthAutoRegistration struct {
	// Status - Auto-registration status
	Status string `json:"status"`
}

// CustomProvidersArchiveOutputProviderOauth represents the custom providers archive output provider oauth type.
type CustomProvidersArchiveOutputProviderOauth struct {
	// Status - OAuth status
	Status string `json:"status"`
	// CallbackUrl - OAuth callback URL
	CallbackUrl      *string                                                   `json:"callback_url,omitempty"`
	AutoRegistration CustomProvidersArchiveOutputProviderOauthAutoRegistration `json:"auto_registration"`
}

// CustomProvidersArchiveOutputProvider represents the custom providers archive output provider type.
type CustomProvidersArchiveOutputProvider struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique provider identifier
	Id string `json:"id"`
	// Access - Access level of the provider
	Access string `json:"access"`
	// Status - Current status of the provider
	Status         string                                              `json:"status"`
	Publisher      CustomProvidersArchiveOutputProviderPublisher       `json:"publisher"`
	CurrentVersion *CustomProvidersArchiveOutputProviderCurrentVersion `json:"current_version,omitempty"`
	Oauth          *CustomProvidersArchiveOutputProviderOauth          `json:"oauth,omitempty"`
	// Identifier - Provider identifier
	Identifier string `json:"identifier"`
	// Name - Display name of the provider
	Name string `json:"name"`
	// Description - Brief description of the provider
	Description *string `json:"description,omitempty"`
	// Slug - URL-friendly identifier
	Slug string `json:"slug"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomProvidersArchiveOutput represents the custom providers archive output type.
type CustomProvidersArchiveOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique custom provider identifier
	Id string `json:"id"`
	// Status - Current status of the custom provider
	Status string `json:"status"`
	// Type - Type of the custom provider
	Type string `json:"type"`
	// Name - Display name of the custom provider
	Name string `json:"name"`
	// Description - Brief description of the custom provider
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any                       `json:"metadata,omitempty"`
	Draft    CustomProvidersArchiveOutputDraft     `json:"draft"`
	ScmRepo  *CustomProvidersArchiveOutputScmRepo  `json:"scm_repo,omitempty"`
	Provider *CustomProvidersArchiveOutputProvider `json:"provider,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapCustomProvidersArchiveOutputFromJSON deserializes JSON data into a CustomProvidersArchiveOutput.
func MapCustomProvidersArchiveOutputFromJSON(data []byte) (*CustomProvidersArchiveOutput, error) {
	var v CustomProvidersArchiveOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCustomProvidersArchiveOutputToJSON serializes a CustomProvidersArchiveOutput to JSON.
func MapCustomProvidersArchiveOutputToJSON(v *CustomProvidersArchiveOutput) ([]byte, error) {
	return json.Marshal(v)
}
