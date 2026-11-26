package customproviders

import (
	"encoding/json"
	"time"
)

// CustomProvidersCreateOutputDraftContainerImage - Container image this provider deploys. Only set for `container` providers.
type CustomProvidersCreateOutputDraftContainerImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// ContainerRegistry - Registry the image is pulled from.
	ContainerRegistry string `json:"container_registry"`
	// ContainerImageTag - Tag selecting the image version.
	ContainerImageTag string `json:"container_image_tag"`
	// ContainerImage - Name of the image in the registry.
	ContainerImage string `json:"container_image"`
}

// CustomProvidersCreateOutputDraftRemoteMcpServer - Remote MCP server this provider connects to. Only set for `remote` providers.
type CustomProvidersCreateOutputDraftRemoteMcpServer struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Url - URL where the remote MCP server is reachable.
	Url string `json:"url"`
	// Transport - MCP transport the remote server speaks.
	Transport string `json:"transport"`
}

// CustomProvidersCreateOutputDraftConfigSchema - Configuration fields for the custom provider.
type CustomProvidersCreateOutputDraftConfigSchema struct {
	// Type - Format of the configuration schema.
	Type string `json:"type"`
	// Schema - JSON Schema for the provider's configuration.
	Schema map[string]any `json:"schema"`
}

// CustomProvidersCreateOutputDraftConfig - Configuration for the custom provider.
type CustomProvidersCreateOutputDraftConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Schema - Configuration fields for the custom provider.
	Schema CustomProvidersCreateOutputDraftConfigSchema `json:"schema"`
	// Transformer - Code that transforms the provider's configuration.
	Transformer string `json:"transformer"`
}

// CustomProvidersCreateOutputDraft - Draft source and configuration of the custom provider.
type CustomProvidersCreateOutputDraft struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// ContainerImage - Container image this provider deploys. Only set for `container` providers.
	ContainerImage *CustomProvidersCreateOutputDraftContainerImage `json:"container_image,omitempty"`
	// RemoteMcpServer - Remote MCP server this provider connects to. Only set for `remote` providers.
	RemoteMcpServer *CustomProvidersCreateOutputDraftRemoteMcpServer `json:"remote_mcp_server,omitempty"`
	// Config - Configuration for the custom provider.
	Config CustomProvidersCreateOutputDraftConfig `json:"config"`
}

// CustomProvidersCreateOutputScmRepoProvider - Provider the repository lives on.
type CustomProvidersCreateOutputScmRepoProvider struct {
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

// CustomProvidersCreateOutputScmRepo - Repository the provider code is deployed from. `null` when not linked.
type CustomProvidersCreateOutputScmRepo struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Provider - Provider the repository lives on.
	Provider CustomProvidersCreateOutputScmRepoProvider `json:"provider"`
	// Url - Web URL of the repository.
	Url string `json:"url"`
	// IsPrivate - Whether the repository is private.
	IsPrivate bool `json:"is_private"`
	// DefaultBranch - Branch used when none is specified.
	DefaultBranch string `json:"default_branch"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCreateOutputProviderPublisher represents the custom providers create output provider publisher type.
type CustomProvidersCreateOutputProviderPublisher struct {
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

// CustomProvidersCreateOutputProviderCurrentVersion represents the custom providers create output provider current version type.
type CustomProvidersCreateOutputProviderCurrentVersion struct {
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

// CustomProvidersCreateOutputProviderOauthAutoRegistration represents the custom providers create output provider oauth auto registration type.
type CustomProvidersCreateOutputProviderOauthAutoRegistration struct {
	// Status - Auto-registration status
	Status string `json:"status"`
}

// CustomProvidersCreateOutputProviderOauth represents the custom providers create output provider oauth type.
type CustomProvidersCreateOutputProviderOauth struct {
	// Status - OAuth status
	Status string `json:"status"`
	// CallbackUrl - Callback URL to register on your OAuth app.
	CallbackUrl      *string                                                  `json:"callback_url,omitempty"`
	AutoRegistration CustomProvidersCreateOutputProviderOauthAutoRegistration `json:"auto_registration"`
}

// CustomProvidersCreateOutputProvider - Session provider backed by this custom provider. `null` when none is linked.
type CustomProvidersCreateOutputProvider struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Access - `public` providers are available to every instance. `tenant` providers are your own.
	Access string `json:"access"`
	// Status - Archived once the provider is archived.
	Status         string                                             `json:"status"`
	Publisher      CustomProvidersCreateOutputProviderPublisher       `json:"publisher"`
	CurrentVersion *CustomProvidersCreateOutputProviderCurrentVersion `json:"current_version,omitempty"`
	Oauth          *CustomProvidersCreateOutputProviderOauth          `json:"oauth,omitempty"`
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

// CustomProvidersCreateOutput represents the custom providers create output type.
type CustomProvidersCreateOutput struct {
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
	Draft CustomProvidersCreateOutputDraft `json:"draft"`
	// ScmRepo - Repository the provider code is deployed from. `null` when not linked.
	ScmRepo *CustomProvidersCreateOutputScmRepo `json:"scm_repo,omitempty"`
	// Provider - Session provider backed by this custom provider. `null` when none is linked.
	Provider *CustomProvidersCreateOutputProvider `json:"provider,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapCustomProvidersCreateOutputFromJSON deserializes JSON data into a CustomProvidersCreateOutput.
func MapCustomProvidersCreateOutputFromJSON(data []byte) (*CustomProvidersCreateOutput, error) {
	var v CustomProvidersCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCustomProvidersCreateOutputToJSON serializes a CustomProvidersCreateOutput to JSON.
func MapCustomProvidersCreateOutputToJSON(v *CustomProvidersCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// CustomProvidersCreateBodyFromFiles represents the custom providers create body from files type.
type CustomProvidersCreateBodyFromFiles struct {
	// Filename - File path within the deployment.
	Filename string `json:"filename"`
	// Content - Content of the file.
	Content string `json:"content"`
	// Encoding - Content encoding.
	Encoding *string `json:"encoding,omitempty"`
}

// CustomProvidersCreateBodyFromRuntime represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type CustomProvidersCreateBodyFromRuntime struct {
	// Identifier - Python runtime.
	Identifier *string `json:"identifier,omitempty"`
	// Version - Python version.
	Version *string `json:"version,omitempty"`
}

// CustomProvidersCreateBodyFromRepository represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type CustomProvidersCreateBodyFromRepository struct {
	// RepositoryId - Linked SCM repository to deploy, by ID.
	RepositoryId *string `json:"repository_id,omitempty"`
	// Branch - Branch to deploy.
	Branch *string `json:"branch,omitempty"`
	// Path - Subdirectory of the branch to deploy.
	Path *string `json:"path,omitempty"`
	// Type - Deploys from a Git repository URL.
	Type *string `json:"type,omitempty"`
	// RepositoryUrl - URL of the Git repository to deploy.
	RepositoryUrl *string `json:"repository_url,omitempty"`
}

// CustomProvidersCreateBodyFrom represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type CustomProvidersCreateBodyFrom struct {
	// Type - Deploys the provider from serverless function code.
	Type *string `json:"type,omitempty"`
	// ImageRef - Container image to deploy, with tag.
	ImageRef *string `json:"image_ref,omitempty"`
	// Username - Username for the container registry. Omit for public images.
	Username *string `json:"username,omitempty"`
	// Password - Password for the container registry. Omit for public images.
	Password *string `json:"password,omitempty"`
	// RemoteUrl - URL of the remote MCP server.
	RemoteUrl *string `json:"remote_url,omitempty"`
	// OauthConfig - OAuth configuration sent to the remote server.
	OauthConfig *map[string]any `json:"oauth_config,omitempty"`
	// Protocol - MCP transport to use: `sse` (HTTP+SSE, 2024-11-05), `streamable_http` (session-based Streamable HTTP, 2025-03-26 to 2025-11-25), or `stateless` (sessionless Streamable HTTP, 2026-07-28)
	Protocol *string `json:"protocol,omitempty"`
	// Files - Source files to deploy.
	Files *[]CustomProvidersCreateBodyFromFiles `json:"files,omitempty"`
	// Env - Environment variables available to the function.
	Env *map[string]string `json:"env,omitempty"`
	// Runtime - Runtime the function runs on.
	Runtime *CustomProvidersCreateBodyFromRuntime `json:"runtime,omitempty"`
	// Repository - Repository the function code is deployed from.
	Repository *CustomProvidersCreateBodyFromRepository `json:"repository,omitempty"`
}

// CustomProvidersCreateBodyConfig - Configuration schema and transformer for the provider.
type CustomProvidersCreateBodyConfig struct {
	// Schema - JSON Schema for the provider's configuration.
	Schema map[string]any `json:"schema"`
	// Transformer - Code that transforms the provider's configuration.
	Transformer string `json:"transformer"`
}

// CustomProvidersCreateBody represents the custom providers create body type.
type CustomProvidersCreateBody struct {
	// Name - Display name for the custom provider.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the custom provider.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// From - Source the provider runs from.
	From CustomProvidersCreateBodyFrom `json:"from"`
	// Config - Configuration schema and transformer for the provider.
	Config *CustomProvidersCreateBodyConfig `json:"config,omitempty"`
}

// MapCustomProvidersCreateBodyFromJSON deserializes JSON data into a CustomProvidersCreateBody.
func MapCustomProvidersCreateBodyFromJSON(data []byte) (*CustomProvidersCreateBody, error) {
	var v CustomProvidersCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCustomProvidersCreateBodyToJSON serializes a CustomProvidersCreateBody to JSON.
func MapCustomProvidersCreateBodyToJSON(v *CustomProvidersCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
