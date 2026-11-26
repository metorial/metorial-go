package customproviders

import (
	"encoding/json"
	"time"
)

// CustomProvidersUpdateOutputDraftContainerImage - Container image this provider deploys. Only set for `container` providers.
type CustomProvidersUpdateOutputDraftContainerImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// ContainerRegistry - Registry the image is pulled from.
	ContainerRegistry string `json:"container_registry"`
	// ContainerImageTag - Tag selecting the image version.
	ContainerImageTag string `json:"container_image_tag"`
	// ContainerImage - Name of the image in the registry.
	ContainerImage string `json:"container_image"`
}

// CustomProvidersUpdateOutputDraftRemoteMcpServer - Remote MCP server this provider connects to. Only set for `remote` providers.
type CustomProvidersUpdateOutputDraftRemoteMcpServer struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Url - URL where the remote MCP server is reachable.
	Url string `json:"url"`
	// Transport - MCP transport the remote server speaks.
	Transport string `json:"transport"`
}

// CustomProvidersUpdateOutputDraftConfigSchema - Configuration fields for the custom provider.
type CustomProvidersUpdateOutputDraftConfigSchema struct {
	// Type - Format of the configuration schema.
	Type string `json:"type"`
	// Schema - JSON Schema for the provider's configuration.
	Schema map[string]any `json:"schema"`
}

// CustomProvidersUpdateOutputDraftConfig - Configuration for the custom provider.
type CustomProvidersUpdateOutputDraftConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Schema - Configuration fields for the custom provider.
	Schema CustomProvidersUpdateOutputDraftConfigSchema `json:"schema"`
	// Transformer - Code that transforms the provider's configuration.
	Transformer string `json:"transformer"`
}

// CustomProvidersUpdateOutputDraft - Draft source and configuration of the custom provider.
type CustomProvidersUpdateOutputDraft struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// ContainerImage - Container image this provider deploys. Only set for `container` providers.
	ContainerImage *CustomProvidersUpdateOutputDraftContainerImage `json:"container_image,omitempty"`
	// RemoteMcpServer - Remote MCP server this provider connects to. Only set for `remote` providers.
	RemoteMcpServer *CustomProvidersUpdateOutputDraftRemoteMcpServer `json:"remote_mcp_server,omitempty"`
	// Config - Configuration for the custom provider.
	Config CustomProvidersUpdateOutputDraftConfig `json:"config"`
}

// CustomProvidersUpdateOutputScmRepoProvider - Provider the repository lives on.
type CustomProvidersUpdateOutputScmRepoProvider struct {
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

// CustomProvidersUpdateOutputScmRepo - Repository the provider code is deployed from. `null` when not linked.
type CustomProvidersUpdateOutputScmRepo struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Provider - Provider the repository lives on.
	Provider CustomProvidersUpdateOutputScmRepoProvider `json:"provider"`
	// Url - Web URL of the repository.
	Url string `json:"url"`
	// IsPrivate - Whether the repository is private.
	IsPrivate bool `json:"is_private"`
	// DefaultBranch - Branch used when none is specified.
	DefaultBranch string `json:"default_branch"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersUpdateOutputProviderPublisher represents the custom providers update output provider publisher type.
type CustomProvidersUpdateOutputProviderPublisher struct {
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

// CustomProvidersUpdateOutputProviderCurrentVersion represents the custom providers update output provider current version type.
type CustomProvidersUpdateOutputProviderCurrentVersion struct {
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

// CustomProvidersUpdateOutputProviderOauthAutoRegistration represents the custom providers update output provider oauth auto registration type.
type CustomProvidersUpdateOutputProviderOauthAutoRegistration struct {
	// Status - Auto-registration status
	Status string `json:"status"`
}

// CustomProvidersUpdateOutputProviderOauth represents the custom providers update output provider oauth type.
type CustomProvidersUpdateOutputProviderOauth struct {
	// Status - OAuth status
	Status string `json:"status"`
	// CallbackUrl - Callback URL to register on your OAuth app.
	CallbackUrl      *string                                                  `json:"callback_url,omitempty"`
	AutoRegistration CustomProvidersUpdateOutputProviderOauthAutoRegistration `json:"auto_registration"`
}

// CustomProvidersUpdateOutputProvider - Session provider backed by this custom provider. `null` when none is linked.
type CustomProvidersUpdateOutputProvider struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Access - `public` providers are available to every instance. `tenant` providers are your own.
	Access string `json:"access"`
	// Status - Archived once the provider is archived.
	Status         string                                             `json:"status"`
	Publisher      CustomProvidersUpdateOutputProviderPublisher       `json:"publisher"`
	CurrentVersion *CustomProvidersUpdateOutputProviderCurrentVersion `json:"current_version,omitempty"`
	Oauth          *CustomProvidersUpdateOutputProviderOauth          `json:"oauth,omitempty"`
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

// CustomProvidersUpdateOutput represents the custom providers update output type.
type CustomProvidersUpdateOutput struct {
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
	Draft CustomProvidersUpdateOutputDraft `json:"draft"`
	// ScmRepo - Repository the provider code is deployed from. `null` when not linked.
	ScmRepo *CustomProvidersUpdateOutputScmRepo `json:"scm_repo,omitempty"`
	// Provider - Session provider backed by this custom provider. `null` when none is linked.
	Provider *CustomProvidersUpdateOutputProvider `json:"provider,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapCustomProvidersUpdateOutputFromJSON deserializes JSON data into a CustomProvidersUpdateOutput.
func MapCustomProvidersUpdateOutputFromJSON(data []byte) (*CustomProvidersUpdateOutput, error) {
	var v CustomProvidersUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCustomProvidersUpdateOutputToJSON serializes a CustomProvidersUpdateOutput to JSON.
func MapCustomProvidersUpdateOutputToJSON(v *CustomProvidersUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// CustomProvidersUpdateBody represents the custom providers update body type.
type CustomProvidersUpdateBody struct {
	// Name - Display name for the custom provider.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the custom provider.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Readme - README content in markdown format
	Readme *string `json:"readme,omitempty"`
}

// MapCustomProvidersUpdateBodyFromJSON deserializes JSON data into a CustomProvidersUpdateBody.
func MapCustomProvidersUpdateBodyFromJSON(data []byte) (*CustomProvidersUpdateBody, error) {
	var v CustomProvidersUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCustomProvidersUpdateBodyToJSON serializes a CustomProvidersUpdateBody to JSON.
func MapCustomProvidersUpdateBodyToJSON(v *CustomProvidersUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
