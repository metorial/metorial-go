package versions

import (
	"encoding/json"
	"time"
)

// CustomProvidersVersionsCreateOutputConfigSchema - Configuration fields for the custom provider.
type CustomProvidersVersionsCreateOutputConfigSchema struct {
	// Type - Format of the configuration schema.
	Type string `json:"type"`
	// Schema - JSON Schema for the provider's configuration.
	Schema map[string]any `json:"schema"`
}

// CustomProvidersVersionsCreateOutputConfig - Configuration for this version.
type CustomProvidersVersionsCreateOutputConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Schema - Configuration fields for the custom provider.
	Schema CustomProvidersVersionsCreateOutputConfigSchema `json:"schema"`
	// Transformer - Code that transforms the provider's configuration.
	Transformer string `json:"transformer"`
}

// CustomProvidersVersionsCreateOutputDeploymentCommit - Commit that triggered this deployment. `null` for deployments without a commit.
type CustomProvidersVersionsCreateOutputDeploymentCommit struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - What the commit did: created a version, merged one into an environment, or rolled one back.
	Type string `json:"type"`
	// Message - Message of the commit. `null` when the commit has no message.
	Message *string `json:"message,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersVersionsCreateOutputDeploymentImmutableBucketScmRepoLinkRepositoryProvider - Provider the repository lives on.
type CustomProvidersVersionsCreateOutputDeploymentImmutableBucketScmRepoLinkRepositoryProvider struct {
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

// CustomProvidersVersionsCreateOutputDeploymentImmutableBucketScmRepoLinkRepository represents the custom providers versions create output deployment immutable bucket scm repo link repository type.
type CustomProvidersVersionsCreateOutputDeploymentImmutableBucketScmRepoLinkRepository struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Provider - Provider the repository lives on.
	Provider CustomProvidersVersionsCreateOutputDeploymentImmutableBucketScmRepoLinkRepositoryProvider `json:"provider"`
	// Url - Web URL of the repository.
	Url string `json:"url"`
	// IsPrivate - Whether the repository is private.
	IsPrivate bool `json:"is_private"`
	// DefaultBranch - Branch used when none is specified.
	DefaultBranch string `json:"default_branch"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersVersionsCreateOutputDeploymentImmutableBucketScmRepoLink - Repository this bucket syncs with. `null` when not linked.
type CustomProvidersVersionsCreateOutputDeploymentImmutableBucketScmRepoLink struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// IsLinked - `true` when the bucket is linked to a repository.
	IsLinked string `json:"is_linked"`
	// Path - Path within the repository the bucket syncs.
	Path       *string                                                                           `json:"path,omitempty"`
	Repository CustomProvidersVersionsCreateOutputDeploymentImmutableBucketScmRepoLinkRepository `json:"repository"`
}

// CustomProvidersVersionsCreateOutputDeploymentImmutableBucket - Immutable copy of the deployed code.
type CustomProvidersVersionsCreateOutputDeploymentImmutableBucket struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsImmutable - `true` when the bucket contents can no longer change.
	IsImmutable bool `json:"is_immutable"`
	// IsReadOnly - `true` when the bucket cannot be written to.
	IsReadOnly bool `json:"is_read_only"`
	// ScmRepoLink - Repository this bucket syncs with. `null` when not linked.
	ScmRepoLink *CustomProvidersVersionsCreateOutputDeploymentImmutableBucketScmRepoLink `json:"scm_repo_link,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersVersionsCreateOutputDeploymentActor represents the custom providers versions create output deployment actor type.
type CustomProvidersVersionsCreateOutputDeploymentActor struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - `external` for user actors, `system` for automated ones.
	Type string `json:"type"`
	// Identifier - Stable identifier for this actor.
	Identifier string `json:"identifier"`
	// Name - Display name.
	Name string `json:"name"`
	// OrganizationActorId - Linked organization actor, by ID. `null` when not linked.
	OrganizationActorId *string `json:"organization_actor_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersVersionsCreateOutputDeploymentScmPushActor - Actor that pushed.
type CustomProvidersVersionsCreateOutputDeploymentScmPushActor struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Stable ID for the actor.
	Id string `json:"id"`
	// ExternalId - ID of the actor on the provider.
	ExternalId *string `json:"external_id,omitempty"`
	// Name - Name of the actor.
	Name *string `json:"name,omitempty"`
	// Email - Email of the actor.
	Email *string `json:"email,omitempty"`
}

// CustomProvidersVersionsCreateOutputDeploymentScmPushCommit - Commit that was pushed.
type CustomProvidersVersionsCreateOutputDeploymentScmPushCommit struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Sha - SHA hash of the commit.
	Sha string `json:"sha"`
	// Branch - Branch the commit was pushed to.
	Branch string `json:"branch"`
	// Message - Full message of the commit.
	Message *string `json:"message,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersVersionsCreateOutputDeploymentScmPushRepositoryProvider - Provider the repository lives on.
type CustomProvidersVersionsCreateOutputDeploymentScmPushRepositoryProvider struct {
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

// CustomProvidersVersionsCreateOutputDeploymentScmPushRepository represents the custom providers versions create output deployment scm push repository type.
type CustomProvidersVersionsCreateOutputDeploymentScmPushRepository struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Provider - Provider the repository lives on.
	Provider CustomProvidersVersionsCreateOutputDeploymentScmPushRepositoryProvider `json:"provider"`
	// Url - Web URL of the repository.
	Url string `json:"url"`
	// IsPrivate - Whether the repository is private.
	IsPrivate bool `json:"is_private"`
	// DefaultBranch - Branch used when none is specified.
	DefaultBranch string `json:"default_branch"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersVersionsCreateOutputDeploymentScmPush - SCM push that triggered the deployment. `null` unless `trigger` is `scm`.
type CustomProvidersVersionsCreateOutputDeploymentScmPush struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Actor - Actor that pushed.
	Actor CustomProvidersVersionsCreateOutputDeploymentScmPushActor `json:"actor"`
	// Commit - Commit that was pushed.
	Commit     CustomProvidersVersionsCreateOutputDeploymentScmPushCommit     `json:"commit"`
	Repository CustomProvidersVersionsCreateOutputDeploymentScmPushRepository `json:"repository"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersVersionsCreateOutputDeployment represents the custom providers versions create output deployment type.
type CustomProvidersVersionsCreateOutputDeployment struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Build and rollout status.
	Status string `json:"status"`
	// Trigger - How the deployment was triggered: manually, by the system, or by an SCM push.
	Trigger string `json:"trigger"`
	// CustomProviderId - Custom provider being deployed.
	CustomProviderId string `json:"custom_provider_id"`
	// ProviderId - Associated provider, by ID. `null` when there is none.
	ProviderId *string `json:"provider_id,omitempty"`
	// CustomProviderVersionId - Version being deployed, by ID.
	CustomProviderVersionId *string `json:"custom_provider_version_id,omitempty"`
	// Commit - Commit that triggered this deployment. `null` for deployments without a commit.
	Commit *CustomProvidersVersionsCreateOutputDeploymentCommit `json:"commit,omitempty"`
	// ImmutableBucket - Immutable copy of the deployed code.
	ImmutableBucket *CustomProvidersVersionsCreateOutputDeploymentImmutableBucket `json:"immutable_bucket,omitempty"`
	Actor           CustomProvidersVersionsCreateOutputDeploymentActor            `json:"actor"`
	// ScmPush - SCM push that triggered the deployment. `null` unless `trigger` is `scm`.
	ScmPush *CustomProvidersVersionsCreateOutputDeploymentScmPush `json:"scm_push,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomProvidersVersionsCreateOutputEnvironmentsEnvironment represents the custom providers versions create output environments environment type.
type CustomProvidersVersionsCreateOutputEnvironmentsEnvironment struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// CustomProviderId - Custom provider this environment belongs to.
	CustomProviderId string `json:"custom_provider_id"`
	// ProviderId - Associated provider, by ID. `null` when there is none.
	ProviderId *string `json:"provider_id,omitempty"`
	// CurrentProviderVersionId - Version currently live in this environment. `null` when no version is live.
	CurrentProviderVersionId *string `json:"current_provider_version_id,omitempty"`
	// InstanceId - ID of the instance this environment deploys into.
	InstanceId string `json:"instance_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomProvidersVersionsCreateOutputEnvironments - Environment this version is deployed to.
type CustomProvidersVersionsCreateOutputEnvironments struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsCurrentVersionForEnvironment - Whether this version is the current one for the environment.
	IsCurrentVersionForEnvironment bool                                                       `json:"is_current_version_for_environment"`
	Environment                    CustomProvidersVersionsCreateOutputEnvironmentsEnvironment `json:"environment"`
}

// CustomProvidersVersionsCreateOutputActor represents the custom providers versions create output actor type.
type CustomProvidersVersionsCreateOutputActor struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - `external` for user actors, `system` for automated ones.
	Type string `json:"type"`
	// Identifier - Stable identifier for this actor.
	Identifier string `json:"identifier"`
	// Name - Display name.
	Name string `json:"name"`
	// OrganizationActorId - Linked organization actor, by ID. `null` when not linked.
	OrganizationActorId *string `json:"organization_actor_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersVersionsCreateOutputContainerImage - Container image this version deploys. Only set for `container` providers.
type CustomProvidersVersionsCreateOutputContainerImage struct {
	// ContainerRegistry - Registry the image is pulled from.
	ContainerRegistry string `json:"container_registry"`
	// ContainerImageTag - Tag selecting the image version.
	ContainerImageTag string `json:"container_image_tag"`
	// ContainerImage - Name of the image in the registry.
	ContainerImage string `json:"container_image"`
}

// CustomProvidersVersionsCreateOutputRemoteMcpServer - Remote MCP server this version connects to. Only set for `remote` providers.
type CustomProvidersVersionsCreateOutputRemoteMcpServer struct {
	// Url - URL where the remote MCP server is reachable.
	Url string `json:"url"`
	// Transport - MCP transport the remote server speaks.
	Transport string `json:"transport"`
}

// CustomProvidersVersionsCreateOutput represents the custom providers versions create output type.
type CustomProvidersVersionsCreateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Build status of this version.
	Status string `json:"status"`
	// Config - Configuration for this version.
	Config *CustomProvidersVersionsCreateOutputConfig `json:"config,omitempty"`
	// Index - Sequential number of this version.
	Index float64 `json:"index"`
	// Identifier - Human-readable label for this version.
	Identifier string                                        `json:"identifier"`
	Deployment CustomProvidersVersionsCreateOutputDeployment `json:"deployment"`
	// Environments - Environments this version is deployed to.
	Environments []CustomProvidersVersionsCreateOutputEnvironments `json:"environments"`
	// CustomProviderId - Custom provider this version belongs to.
	CustomProviderId string `json:"custom_provider_id"`
	// ProviderId - Associated provider, by ID. `null` when there is none.
	ProviderId *string                                  `json:"provider_id,omitempty"`
	Actor      CustomProvidersVersionsCreateOutputActor `json:"actor"`
	// ContainerImage - Container image this version deploys. Only set for `container` providers.
	ContainerImage *CustomProvidersVersionsCreateOutputContainerImage `json:"container_image,omitempty"`
	// RemoteMcpServer - Remote MCP server this version connects to. Only set for `remote` providers.
	RemoteMcpServer *CustomProvidersVersionsCreateOutputRemoteMcpServer `json:"remote_mcp_server,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapCustomProvidersVersionsCreateOutputFromJSON deserializes JSON data into a CustomProvidersVersionsCreateOutput.
func MapCustomProvidersVersionsCreateOutputFromJSON(data []byte) (*CustomProvidersVersionsCreateOutput, error) {
	var v CustomProvidersVersionsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCustomProvidersVersionsCreateOutputToJSON serializes a CustomProvidersVersionsCreateOutput to JSON.
func MapCustomProvidersVersionsCreateOutputToJSON(v *CustomProvidersVersionsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// CustomProvidersVersionsCreateBodyFromFiles represents the custom providers versions create body from files type.
type CustomProvidersVersionsCreateBodyFromFiles struct {
	// Filename - File path within the deployment.
	Filename string `json:"filename"`
	// Content - Content of the file.
	Content string `json:"content"`
	// Encoding - Content encoding.
	Encoding *string `json:"encoding,omitempty"`
}

// CustomProvidersVersionsCreateBodyFromRuntime represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type CustomProvidersVersionsCreateBodyFromRuntime struct {
	// Identifier - Python runtime.
	Identifier *string `json:"identifier,omitempty"`
	// Version - Python version.
	Version *string `json:"version,omitempty"`
}

// CustomProvidersVersionsCreateBodyFromRepository represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type CustomProvidersVersionsCreateBodyFromRepository struct {
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

// CustomProvidersVersionsCreateBodyFrom represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type CustomProvidersVersionsCreateBodyFrom struct {
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
	Files *[]CustomProvidersVersionsCreateBodyFromFiles `json:"files,omitempty"`
	// Env - Environment variables available to the function.
	Env *map[string]string `json:"env,omitempty"`
	// Runtime - Runtime the function runs on.
	Runtime *CustomProvidersVersionsCreateBodyFromRuntime `json:"runtime,omitempty"`
	// Repository - Repository the function code is deployed from.
	Repository *CustomProvidersVersionsCreateBodyFromRepository `json:"repository,omitempty"`
}

// CustomProvidersVersionsCreateBodyConfig - Configuration schema and transformer for the provider.
type CustomProvidersVersionsCreateBodyConfig struct {
	// Schema - JSON Schema for the provider's configuration.
	Schema map[string]any `json:"schema"`
	// Transformer - Code that transforms the provider's configuration.
	Transformer string `json:"transformer"`
}

// CustomProvidersVersionsCreateBody represents the custom providers versions create body type.
type CustomProvidersVersionsCreateBody struct {
	// CustomProviderId - Custom provider to release a new version for.
	CustomProviderId string `json:"custom_provider_id"`
	// From - Source the provider runs from.
	From CustomProvidersVersionsCreateBodyFrom `json:"from"`
	// Config - Configuration schema and transformer for the provider.
	Config *CustomProvidersVersionsCreateBodyConfig `json:"config,omitempty"`
}

// MapCustomProvidersVersionsCreateBodyFromJSON deserializes JSON data into a CustomProvidersVersionsCreateBody.
func MapCustomProvidersVersionsCreateBodyFromJSON(data []byte) (*CustomProvidersVersionsCreateBody, error) {
	var v CustomProvidersVersionsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCustomProvidersVersionsCreateBodyToJSON serializes a CustomProvidersVersionsCreateBody to JSON.
func MapCustomProvidersVersionsCreateBodyToJSON(v *CustomProvidersVersionsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
