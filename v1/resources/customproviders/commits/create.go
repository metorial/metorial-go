package commits

import (
	"encoding/json"
	"time"
)

// CustomProvidersCommitsCreateOutputError - Failure details. `null` when the commit did not fail.
type CustomProvidersCommitsCreateOutputError struct {
	// Code - Machine-readable error code.
	Code string `json:"code"`
	// Message - Human-readable error message.
	Message string `json:"message"`
}

// CustomProvidersCommitsCreateOutputToEnvironment represents the custom providers commits create output to environment type.
type CustomProvidersCommitsCreateOutputToEnvironment struct {
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

// CustomProvidersCommitsCreateOutputFromEnvironment - Environment the version is promoted from. `null` for rollbacks.
type CustomProvidersCommitsCreateOutputFromEnvironment struct {
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

// CustomProvidersCommitsCreateOutputTargetCustomProviderVersionConfigSchema - Configuration fields for the custom provider.
type CustomProvidersCommitsCreateOutputTargetCustomProviderVersionConfigSchema struct {
	// Type - Format of the configuration schema.
	Type string `json:"type"`
	// Schema - JSON Schema for the provider's configuration.
	Schema map[string]any `json:"schema"`
}

// CustomProvidersCommitsCreateOutputTargetCustomProviderVersionConfig - Configuration for this version.
type CustomProvidersCommitsCreateOutputTargetCustomProviderVersionConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Schema - Configuration fields for the custom provider.
	Schema CustomProvidersCommitsCreateOutputTargetCustomProviderVersionConfigSchema `json:"schema"`
	// Transformer - Code that transforms the provider's configuration.
	Transformer string `json:"transformer"`
}

// CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentCommit - Commit that triggered this deployment. `null` for deployments without a commit.
type CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentCommit struct {
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

// CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepositoryProvider - Provider the repository lives on.
type CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepositoryProvider struct {
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

// CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepository represents the custom providers commits create output target custom provider version deployment immutable bucket scm repo link repository type.
type CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepository struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Provider - Provider the repository lives on.
	Provider CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepositoryProvider `json:"provider"`
	// Url - Web URL of the repository.
	Url string `json:"url"`
	// IsPrivate - Whether the repository is private.
	IsPrivate bool `json:"is_private"`
	// DefaultBranch - Branch used when none is specified.
	DefaultBranch string `json:"default_branch"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentImmutableBucketScmRepoLink - Repository this bucket syncs with. `null` when not linked.
type CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentImmutableBucketScmRepoLink struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// IsLinked - `true` when the bucket is linked to a repository.
	IsLinked string `json:"is_linked"`
	// Path - Path within the repository the bucket syncs.
	Path       *string                                                                                                     `json:"path,omitempty"`
	Repository CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepository `json:"repository"`
}

// CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentImmutableBucket - Immutable copy of the deployed code.
type CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentImmutableBucket struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsImmutable - `true` when the bucket contents can no longer change.
	IsImmutable bool `json:"is_immutable"`
	// IsReadOnly - `true` when the bucket cannot be written to.
	IsReadOnly bool `json:"is_read_only"`
	// ScmRepoLink - Repository this bucket syncs with. `null` when not linked.
	ScmRepoLink *CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentImmutableBucketScmRepoLink `json:"scm_repo_link,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentActor represents the custom providers commits create output target custom provider version deployment actor type.
type CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentActor struct {
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

// CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentScmPushActor - Actor that pushed.
type CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentScmPushActor struct {
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

// CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentScmPushCommit - Commit that was pushed.
type CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentScmPushCommit struct {
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

// CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentScmPushRepositoryProvider - Provider the repository lives on.
type CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentScmPushRepositoryProvider struct {
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

// CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentScmPushRepository represents the custom providers commits create output target custom provider version deployment scm push repository type.
type CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentScmPushRepository struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Provider - Provider the repository lives on.
	Provider CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentScmPushRepositoryProvider `json:"provider"`
	// Url - Web URL of the repository.
	Url string `json:"url"`
	// IsPrivate - Whether the repository is private.
	IsPrivate bool `json:"is_private"`
	// DefaultBranch - Branch used when none is specified.
	DefaultBranch string `json:"default_branch"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentScmPush - SCM push that triggered the deployment. `null` unless `trigger` is `scm`.
type CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentScmPush struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Actor - Actor that pushed.
	Actor CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentScmPushActor `json:"actor"`
	// Commit - Commit that was pushed.
	Commit     CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentScmPushCommit     `json:"commit"`
	Repository CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentScmPushRepository `json:"repository"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeployment represents the custom providers commits create output target custom provider version deployment type.
type CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeployment struct {
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
	Commit *CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentCommit `json:"commit,omitempty"`
	// ImmutableBucket - Immutable copy of the deployed code.
	ImmutableBucket *CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentImmutableBucket `json:"immutable_bucket,omitempty"`
	Actor           CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentActor            `json:"actor"`
	// ScmPush - SCM push that triggered the deployment. `null` unless `trigger` is `scm`.
	ScmPush *CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeploymentScmPush `json:"scm_push,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomProvidersCommitsCreateOutputTargetCustomProviderVersionEnvironmentsEnvironment represents the custom providers commits create output target custom provider version environments environment type.
type CustomProvidersCommitsCreateOutputTargetCustomProviderVersionEnvironmentsEnvironment struct {
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

// CustomProvidersCommitsCreateOutputTargetCustomProviderVersionEnvironments - Environment this version is deployed to.
type CustomProvidersCommitsCreateOutputTargetCustomProviderVersionEnvironments struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsCurrentVersionForEnvironment - Whether this version is the current one for the environment.
	IsCurrentVersionForEnvironment bool                                                                                 `json:"is_current_version_for_environment"`
	Environment                    CustomProvidersCommitsCreateOutputTargetCustomProviderVersionEnvironmentsEnvironment `json:"environment"`
}

// CustomProvidersCommitsCreateOutputTargetCustomProviderVersionActor represents the custom providers commits create output target custom provider version actor type.
type CustomProvidersCommitsCreateOutputTargetCustomProviderVersionActor struct {
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

// CustomProvidersCommitsCreateOutputTargetCustomProviderVersionContainerImage - Container image this version deploys. Only set for `container` providers.
type CustomProvidersCommitsCreateOutputTargetCustomProviderVersionContainerImage struct {
	// ContainerRegistry - Registry the image is pulled from.
	ContainerRegistry string `json:"container_registry"`
	// ContainerImageTag - Tag selecting the image version.
	ContainerImageTag string `json:"container_image_tag"`
	// ContainerImage - Name of the image in the registry.
	ContainerImage string `json:"container_image"`
}

// CustomProvidersCommitsCreateOutputTargetCustomProviderVersionRemoteMcpServer - Remote MCP server this version connects to. Only set for `remote` providers.
type CustomProvidersCommitsCreateOutputTargetCustomProviderVersionRemoteMcpServer struct {
	// Url - URL where the remote MCP server is reachable.
	Url string `json:"url"`
	// Transport - MCP transport the remote server speaks.
	Transport string `json:"transport"`
}

// CustomProvidersCommitsCreateOutputTargetCustomProviderVersion represents the custom providers commits create output target custom provider version type.
type CustomProvidersCommitsCreateOutputTargetCustomProviderVersion struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Build status of this version.
	Status string `json:"status"`
	// Config - Configuration for this version.
	Config *CustomProvidersCommitsCreateOutputTargetCustomProviderVersionConfig `json:"config,omitempty"`
	// Index - Sequential number of this version.
	Index float64 `json:"index"`
	// Identifier - Human-readable label for this version.
	Identifier string                                                                  `json:"identifier"`
	Deployment CustomProvidersCommitsCreateOutputTargetCustomProviderVersionDeployment `json:"deployment"`
	// Environments - Environments this version is deployed to.
	Environments []CustomProvidersCommitsCreateOutputTargetCustomProviderVersionEnvironments `json:"environments"`
	// CustomProviderId - Custom provider this version belongs to.
	CustomProviderId string `json:"custom_provider_id"`
	// ProviderId - Associated provider, by ID. `null` when there is none.
	ProviderId *string                                                            `json:"provider_id,omitempty"`
	Actor      CustomProvidersCommitsCreateOutputTargetCustomProviderVersionActor `json:"actor"`
	// ContainerImage - Container image this version deploys. Only set for `container` providers.
	ContainerImage *CustomProvidersCommitsCreateOutputTargetCustomProviderVersionContainerImage `json:"container_image,omitempty"`
	// RemoteMcpServer - Remote MCP server this version connects to. Only set for `remote` providers.
	RemoteMcpServer *CustomProvidersCommitsCreateOutputTargetCustomProviderVersionRemoteMcpServer `json:"remote_mcp_server,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionConfigSchema - Configuration fields for the custom provider.
type CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionConfigSchema struct {
	// Type - Format of the configuration schema.
	Type string `json:"type"`
	// Schema - JSON Schema for the provider's configuration.
	Schema map[string]any `json:"schema"`
}

// CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionConfig - Configuration for this version.
type CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Schema - Configuration fields for the custom provider.
	Schema CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionConfigSchema `json:"schema"`
	// Transformer - Code that transforms the provider's configuration.
	Transformer string `json:"transformer"`
}

// CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentCommit - Commit that triggered this deployment. `null` for deployments without a commit.
type CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentCommit struct {
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

// CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepositoryProvider - Provider the repository lives on.
type CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepositoryProvider struct {
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

// CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepository represents the custom providers commits create output previous custom provider version deployment immutable bucket scm repo link repository type.
type CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepository struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Provider - Provider the repository lives on.
	Provider CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepositoryProvider `json:"provider"`
	// Url - Web URL of the repository.
	Url string `json:"url"`
	// IsPrivate - Whether the repository is private.
	IsPrivate bool `json:"is_private"`
	// DefaultBranch - Branch used when none is specified.
	DefaultBranch string `json:"default_branch"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentImmutableBucketScmRepoLink - Repository this bucket syncs with. `null` when not linked.
type CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentImmutableBucketScmRepoLink struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// IsLinked - `true` when the bucket is linked to a repository.
	IsLinked string `json:"is_linked"`
	// Path - Path within the repository the bucket syncs.
	Path       *string                                                                                                       `json:"path,omitempty"`
	Repository CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepository `json:"repository"`
}

// CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentImmutableBucket - Immutable copy of the deployed code.
type CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentImmutableBucket struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsImmutable - `true` when the bucket contents can no longer change.
	IsImmutable bool `json:"is_immutable"`
	// IsReadOnly - `true` when the bucket cannot be written to.
	IsReadOnly bool `json:"is_read_only"`
	// ScmRepoLink - Repository this bucket syncs with. `null` when not linked.
	ScmRepoLink *CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentImmutableBucketScmRepoLink `json:"scm_repo_link,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentActor represents the custom providers commits create output previous custom provider version deployment actor type.
type CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentActor struct {
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

// CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentScmPushActor - Actor that pushed.
type CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentScmPushActor struct {
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

// CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentScmPushCommit - Commit that was pushed.
type CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentScmPushCommit struct {
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

// CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentScmPushRepositoryProvider - Provider the repository lives on.
type CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentScmPushRepositoryProvider struct {
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

// CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentScmPushRepository represents the custom providers commits create output previous custom provider version deployment scm push repository type.
type CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentScmPushRepository struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Provider - Provider the repository lives on.
	Provider CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentScmPushRepositoryProvider `json:"provider"`
	// Url - Web URL of the repository.
	Url string `json:"url"`
	// IsPrivate - Whether the repository is private.
	IsPrivate bool `json:"is_private"`
	// DefaultBranch - Branch used when none is specified.
	DefaultBranch string `json:"default_branch"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentScmPush - SCM push that triggered the deployment. `null` unless `trigger` is `scm`.
type CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentScmPush struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Actor - Actor that pushed.
	Actor CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentScmPushActor `json:"actor"`
	// Commit - Commit that was pushed.
	Commit     CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentScmPushCommit     `json:"commit"`
	Repository CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentScmPushRepository `json:"repository"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeployment represents the custom providers commits create output previous custom provider version deployment type.
type CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeployment struct {
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
	Commit *CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentCommit `json:"commit,omitempty"`
	// ImmutableBucket - Immutable copy of the deployed code.
	ImmutableBucket *CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentImmutableBucket `json:"immutable_bucket,omitempty"`
	Actor           CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentActor            `json:"actor"`
	// ScmPush - SCM push that triggered the deployment. `null` unless `trigger` is `scm`.
	ScmPush *CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeploymentScmPush `json:"scm_push,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionEnvironmentsEnvironment represents the custom providers commits create output previous custom provider version environments environment type.
type CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionEnvironmentsEnvironment struct {
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

// CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionEnvironments - Environment this version is deployed to.
type CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionEnvironments struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsCurrentVersionForEnvironment - Whether this version is the current one for the environment.
	IsCurrentVersionForEnvironment bool                                                                                   `json:"is_current_version_for_environment"`
	Environment                    CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionEnvironmentsEnvironment `json:"environment"`
}

// CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionActor represents the custom providers commits create output previous custom provider version actor type.
type CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionActor struct {
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

// CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionContainerImage - Container image this version deploys. Only set for `container` providers.
type CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionContainerImage struct {
	// ContainerRegistry - Registry the image is pulled from.
	ContainerRegistry string `json:"container_registry"`
	// ContainerImageTag - Tag selecting the image version.
	ContainerImageTag string `json:"container_image_tag"`
	// ContainerImage - Name of the image in the registry.
	ContainerImage string `json:"container_image"`
}

// CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionRemoteMcpServer - Remote MCP server this version connects to. Only set for `remote` providers.
type CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionRemoteMcpServer struct {
	// Url - URL where the remote MCP server is reachable.
	Url string `json:"url"`
	// Transport - MCP transport the remote server speaks.
	Transport string `json:"transport"`
}

// CustomProvidersCommitsCreateOutputPreviousCustomProviderVersion - Version previously in the environment. `null` when there was none.
type CustomProvidersCommitsCreateOutputPreviousCustomProviderVersion struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Build status of this version.
	Status string `json:"status"`
	// Config - Configuration for this version.
	Config *CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionConfig `json:"config,omitempty"`
	// Index - Sequential number of this version.
	Index float64 `json:"index"`
	// Identifier - Human-readable label for this version.
	Identifier string                                                                    `json:"identifier"`
	Deployment CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionDeployment `json:"deployment"`
	// Environments - Environments this version is deployed to.
	Environments []CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionEnvironments `json:"environments"`
	// CustomProviderId - Custom provider this version belongs to.
	CustomProviderId string `json:"custom_provider_id"`
	// ProviderId - Associated provider, by ID. `null` when there is none.
	ProviderId *string                                                              `json:"provider_id,omitempty"`
	Actor      CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionActor `json:"actor"`
	// ContainerImage - Container image this version deploys. Only set for `container` providers.
	ContainerImage *CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionContainerImage `json:"container_image,omitempty"`
	// RemoteMcpServer - Remote MCP server this version connects to. Only set for `remote` providers.
	RemoteMcpServer *CustomProvidersCommitsCreateOutputPreviousCustomProviderVersionRemoteMcpServer `json:"remote_mcp_server,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomProvidersCommitsCreateOutputActor represents the custom providers commits create output actor type.
type CustomProvidersCommitsCreateOutputActor struct {
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

// CustomProvidersCommitsCreateOutputScmPushActor - Actor that pushed.
type CustomProvidersCommitsCreateOutputScmPushActor struct {
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

// CustomProvidersCommitsCreateOutputScmPushCommit - Commit that was pushed.
type CustomProvidersCommitsCreateOutputScmPushCommit struct {
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

// CustomProvidersCommitsCreateOutputScmPushRepositoryProvider - Provider the repository lives on.
type CustomProvidersCommitsCreateOutputScmPushRepositoryProvider struct {
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

// CustomProvidersCommitsCreateOutputScmPushRepository represents the custom providers commits create output scm push repository type.
type CustomProvidersCommitsCreateOutputScmPushRepository struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Provider - Provider the repository lives on.
	Provider CustomProvidersCommitsCreateOutputScmPushRepositoryProvider `json:"provider"`
	// Url - Web URL of the repository.
	Url string `json:"url"`
	// IsPrivate - Whether the repository is private.
	IsPrivate bool `json:"is_private"`
	// DefaultBranch - Branch used when none is specified.
	DefaultBranch string `json:"default_branch"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsCreateOutputScmPush - SCM push that triggered the commit. `null` unless `trigger` is `scm`.
type CustomProvidersCommitsCreateOutputScmPush struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Actor - Actor that pushed.
	Actor CustomProvidersCommitsCreateOutputScmPushActor `json:"actor"`
	// Commit - Commit that was pushed.
	Commit     CustomProvidersCommitsCreateOutputScmPushCommit     `json:"commit"`
	Repository CustomProvidersCommitsCreateOutputScmPushRepository `json:"repository"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsCreateOutput represents the custom providers commits create output type.
type CustomProvidersCommitsCreateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `applied` once the version promotion takes effect.
	Status string `json:"status"`
	// Trigger - How the commit was created: manually, by the system, or by an SCM push.
	Trigger string `json:"trigger"`
	// Error - Failure details. `null` when the commit did not fail.
	Error *CustomProvidersCommitsCreateOutputError `json:"error,omitempty"`
	// CustomProviderId - Custom provider this commit belongs to.
	CustomProviderId string `json:"custom_provider_id"`
	// ProviderId - Associated provider, by ID. `null` when there is none.
	ProviderId *string `json:"provider_id,omitempty"`
	// CustomProviderDeploymentId - Deployment this commit built. `null` when there is none.
	CustomProviderDeploymentId *string                                         `json:"custom_provider_deployment_id,omitempty"`
	ToEnvironment              CustomProvidersCommitsCreateOutputToEnvironment `json:"to_environment"`
	// FromEnvironment - Environment the version is promoted from. `null` for rollbacks.
	FromEnvironment             *CustomProvidersCommitsCreateOutputFromEnvironment            `json:"from_environment,omitempty"`
	TargetCustomProviderVersion CustomProvidersCommitsCreateOutputTargetCustomProviderVersion `json:"target_custom_provider_version"`
	// PreviousCustomProviderVersion - Version previously in the environment. `null` when there was none.
	PreviousCustomProviderVersion *CustomProvidersCommitsCreateOutputPreviousCustomProviderVersion `json:"previous_custom_provider_version,omitempty"`
	Actor                         CustomProvidersCommitsCreateOutputActor                          `json:"actor"`
	// ScmPush - SCM push that triggered the commit. `null` unless `trigger` is `scm`.
	ScmPush *CustomProvidersCommitsCreateOutputScmPush `json:"scm_push,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// AppliedAt - When the commit took effect. `null` while `status` is `pending`.
	AppliedAt *time.Time `json:"applied_at,omitempty"`
}

// MapCustomProvidersCommitsCreateOutputFromJSON deserializes JSON data into a CustomProvidersCommitsCreateOutput.
func MapCustomProvidersCommitsCreateOutputFromJSON(data []byte) (*CustomProvidersCommitsCreateOutput, error) {
	var v CustomProvidersCommitsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCustomProvidersCommitsCreateOutputToJSON serializes a CustomProvidersCommitsCreateOutput to JSON.
func MapCustomProvidersCommitsCreateOutputToJSON(v *CustomProvidersCommitsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// CustomProvidersCommitsCreateBodyAction represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type CustomProvidersCommitsCreateBodyAction struct {
	// Type - Rolls an environment back to a previous version.
	Type *string `json:"type,omitempty"`
	// FromEnvironmentId - Environment to promote the version from.
	FromEnvironmentId *string `json:"from_environment_id,omitempty"`
	// ToEnvironmentId - Environment to promote the version into.
	ToEnvironmentId *string `json:"to_environment_id,omitempty"`
	// EnvironmentId - Environment to roll back.
	EnvironmentId *string `json:"environment_id,omitempty"`
	// VersionId - Version to restore in the environment.
	VersionId *string `json:"version_id,omitempty"`
}

// CustomProvidersCommitsCreateBody represents the custom providers commits create body type.
type CustomProvidersCommitsCreateBody struct {
	// Message - Commit message
	Message string `json:"message"`
	// Action - The commit action to perform
	Action CustomProvidersCommitsCreateBodyAction `json:"action"`
}

// MapCustomProvidersCommitsCreateBodyFromJSON deserializes JSON data into a CustomProvidersCommitsCreateBody.
func MapCustomProvidersCommitsCreateBodyFromJSON(data []byte) (*CustomProvidersCommitsCreateBody, error) {
	var v CustomProvidersCommitsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCustomProvidersCommitsCreateBodyToJSON serializes a CustomProvidersCommitsCreateBody to JSON.
func MapCustomProvidersCommitsCreateBodyToJSON(v *CustomProvidersCommitsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
