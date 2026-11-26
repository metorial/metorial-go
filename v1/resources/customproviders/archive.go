package customproviders

import (
	"encoding/json"
	"time"
)

// CustomProvidersArchiveOutputDraftContainerImage - Container image this provider deploys. Only set for `container` providers.
type CustomProvidersArchiveOutputDraftContainerImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// ContainerRegistry - Registry the image is pulled from.
	ContainerRegistry string `json:"container_registry"`
	// ContainerImageTag - Tag selecting the image version.
	ContainerImageTag string `json:"container_image_tag"`
	// ContainerImage - Name of the image in the registry.
	ContainerImage string `json:"container_image"`
}

// CustomProvidersArchiveOutputDraftRemoteMcpServer - Remote MCP server this provider connects to. Only set for `remote` providers.
type CustomProvidersArchiveOutputDraftRemoteMcpServer struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Url - URL where the remote MCP server is reachable.
	Url string `json:"url"`
	// Transport - MCP transport the remote server speaks.
	Transport string `json:"transport"`
}

// CustomProvidersArchiveOutputDraftConfigSchema - Configuration fields for the custom provider.
type CustomProvidersArchiveOutputDraftConfigSchema struct {
	// Type - Format of the configuration schema.
	Type string `json:"type"`
	// Schema - JSON Schema for the provider's configuration.
	Schema map[string]any `json:"schema"`
}

// CustomProvidersArchiveOutputDraftConfig - Configuration for the custom provider.
type CustomProvidersArchiveOutputDraftConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Schema - Configuration fields for the custom provider.
	Schema CustomProvidersArchiveOutputDraftConfigSchema `json:"schema"`
	// Transformer - Code that transforms the provider's configuration.
	Transformer string `json:"transformer"`
}

// CustomProvidersArchiveOutputDraft - Draft source and configuration of the custom provider.
type CustomProvidersArchiveOutputDraft struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// ContainerImage - Container image this provider deploys. Only set for `container` providers.
	ContainerImage *CustomProvidersArchiveOutputDraftContainerImage `json:"container_image,omitempty"`
	// RemoteMcpServer - Remote MCP server this provider connects to. Only set for `remote` providers.
	RemoteMcpServer *CustomProvidersArchiveOutputDraftRemoteMcpServer `json:"remote_mcp_server,omitempty"`
	// Config - Configuration for the custom provider.
	Config CustomProvidersArchiveOutputDraftConfig `json:"config"`
}

// CustomProvidersArchiveOutputScmRepoProvider - Provider the repository lives on.
type CustomProvidersArchiveOutputScmRepoProvider struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Provider the repository lives on.
	Type string `json:"type"`
	// Id - ID of the repository on the provider.
	Id string `json:"id"`
	// Name - Name of the repository on the provider.
	Name string `json:"name"`
	// Owner - Owner of the repository on the provider.
	Owner string `json:"owner"`
}

// CustomProvidersArchiveOutputScmRepo - Repository the provider code is deployed from. `null` when not linked.
type CustomProvidersArchiveOutputScmRepo struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Provider - Provider the repository lives on.
	Provider CustomProvidersArchiveOutputScmRepoProvider `json:"provider"`
	// Url - Web URL of the repository.
	Url string `json:"url"`
	// IsPrivate - Whether the repository is private.
	IsPrivate bool `json:"is_private"`
	// DefaultBranch - Branch used when none is specified.
	DefaultBranch string `json:"default_branch"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersArchiveOutputProviderPublisher represents the custom providers archive output provider publisher type.
type CustomProvidersArchiveOutputProviderPublisher struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// ImageUrl - URL of the publisher logo
	ImageUrl string `json:"image_url"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomProvidersArchiveOutputProviderCurrentVersion represents the custom providers archive output provider current version type.
type CustomProvidersArchiveOutputProviderCurrentVersion struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Version - Release label for the version.
	Version string `json:"version"`
	// ProviderId - Provider this version belongs to.
	ProviderId string `json:"provider_id"`
	// IsCurrent - The version deployments follow unless locked to another.
	IsCurrent bool `json:"is_current"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// SpecificationId - Specification this version belongs to.
	SpecificationId *string `json:"specification_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
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
	// CallbackUrl - Callback URL to register on your OAuth app.
	CallbackUrl      *string                                                   `json:"callback_url,omitempty"`
	AutoRegistration CustomProvidersArchiveOutputProviderOauthAutoRegistration `json:"auto_registration"`
}

// CustomProvidersArchiveOutputProvider - Session provider backed by this custom provider. `null` when none is linked.
type CustomProvidersArchiveOutputProvider struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Access - `public` providers are available to every instance. `tenant` providers are your own.
	Access string `json:"access"`
	// Status - Archived once the provider is archived.
	Status         string                                              `json:"status"`
	Publisher      CustomProvidersArchiveOutputProviderPublisher       `json:"publisher"`
	CurrentVersion *CustomProvidersArchiveOutputProviderCurrentVersion `json:"current_version,omitempty"`
	Oauth          *CustomProvidersArchiveOutputProviderOauth          `json:"oauth,omitempty"`
	// Identifier - Provider identifier
	Identifier string `json:"identifier"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Slug - URL-safe name for the provider.
	Slug string `json:"slug"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomProvidersArchiveOutput represents the custom providers archive output type.
type CustomProvidersArchiveOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once archived. Archived providers reject new connections.
	Status string `json:"status"`
	// Type - How the provider runs: `function`, `container`, or `remote`.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Draft - Draft source and configuration of the custom provider.
	Draft CustomProvidersArchiveOutputDraft `json:"draft"`
	// ScmRepo - Repository the provider code is deployed from. `null` when not linked.
	ScmRepo *CustomProvidersArchiveOutputScmRepo `json:"scm_repo,omitempty"`
	// Provider - Session provider backed by this custom provider. `null` when none is linked.
	Provider *CustomProvidersArchiveOutputProvider `json:"provider,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
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
