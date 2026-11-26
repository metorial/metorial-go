package commits

import (
	"encoding/json"
	"time"
)

// CustomProvidersCommitsGetOutputError - Failure details. `null` when the commit did not fail.
type CustomProvidersCommitsGetOutputError struct {
	// Code - Machine-readable error code.
	Code string `json:"code"`
	// Message - Human-readable error message.
	Message string `json:"message"`
}

// CustomProvidersCommitsGetOutputToEnvironment represents the custom providers commits get output to environment type.
type CustomProvidersCommitsGetOutputToEnvironment struct {
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

// CustomProvidersCommitsGetOutputFromEnvironment - Environment the version is promoted from. `null` for rollbacks.
type CustomProvidersCommitsGetOutputFromEnvironment struct {
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

// CustomProvidersCommitsGetOutputTargetCustomProviderVersionConfigSchema - Configuration fields for the custom provider.
type CustomProvidersCommitsGetOutputTargetCustomProviderVersionConfigSchema struct {
	// Type - Format of the configuration schema.
	Type string `json:"type"`
	// Schema - JSON Schema for the provider's configuration.
	Schema map[string]any `json:"schema"`
}

// CustomProvidersCommitsGetOutputTargetCustomProviderVersionConfig - Configuration for this version.
type CustomProvidersCommitsGetOutputTargetCustomProviderVersionConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Schema - Configuration fields for the custom provider.
	Schema CustomProvidersCommitsGetOutputTargetCustomProviderVersionConfigSchema `json:"schema"`
	// Transformer - Code that transforms the provider's configuration.
	Transformer string `json:"transformer"`
}

// CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentCommit - Commit that triggered this deployment. `null` for deployments without a commit.
type CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentCommit struct {
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

// CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepositoryProvider - Provider the repository lives on.
type CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepositoryProvider struct {
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

// CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepository represents the custom providers commits get output target custom provider version deployment immutable bucket scm repo link repository type.
type CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepository struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Provider - Provider the repository lives on.
	Provider CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepositoryProvider `json:"provider"`
	// Url - Web URL of the repository.
	Url string `json:"url"`
	// IsPrivate - Whether the repository is private.
	IsPrivate bool `json:"is_private"`
	// DefaultBranch - Branch used when none is specified.
	DefaultBranch string `json:"default_branch"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentImmutableBucketScmRepoLink - Repository this bucket syncs with. `null` when not linked.
type CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentImmutableBucketScmRepoLink struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// IsLinked - `true` when the bucket is linked to a repository.
	IsLinked string `json:"is_linked"`
	// Path - Path within the repository the bucket syncs.
	Path       *string                                                                                                  `json:"path,omitempty"`
	Repository CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepository `json:"repository"`
}

// CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentImmutableBucket - Immutable copy of the deployed code.
type CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentImmutableBucket struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsImmutable - `true` when the bucket contents can no longer change.
	IsImmutable bool `json:"is_immutable"`
	// IsReadOnly - `true` when the bucket cannot be written to.
	IsReadOnly bool `json:"is_read_only"`
	// ScmRepoLink - Repository this bucket syncs with. `null` when not linked.
	ScmRepoLink *CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentImmutableBucketScmRepoLink `json:"scm_repo_link,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentActor represents the custom providers commits get output target custom provider version deployment actor type.
type CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentActor struct {
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

// CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentScmPushActor - Actor that pushed.
type CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentScmPushActor struct {
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

// CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentScmPushCommit - Commit that was pushed.
type CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentScmPushCommit struct {
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

// CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentScmPushRepositoryProvider - Provider the repository lives on.
type CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentScmPushRepositoryProvider struct {
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

// CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentScmPushRepository represents the custom providers commits get output target custom provider version deployment scm push repository type.
type CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentScmPushRepository struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Provider - Provider the repository lives on.
	Provider CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentScmPushRepositoryProvider `json:"provider"`
	// Url - Web URL of the repository.
	Url string `json:"url"`
	// IsPrivate - Whether the repository is private.
	IsPrivate bool `json:"is_private"`
	// DefaultBranch - Branch used when none is specified.
	DefaultBranch string `json:"default_branch"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentScmPush - SCM push that triggered the deployment. `null` unless `trigger` is `scm`.
type CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentScmPush struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Actor - Actor that pushed.
	Actor CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentScmPushActor `json:"actor"`
	// Commit - Commit that was pushed.
	Commit     CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentScmPushCommit     `json:"commit"`
	Repository CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentScmPushRepository `json:"repository"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeployment represents the custom providers commits get output target custom provider version deployment type.
type CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeployment struct {
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
	Commit *CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentCommit `json:"commit,omitempty"`
	// ImmutableBucket - Immutable copy of the deployed code.
	ImmutableBucket *CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentImmutableBucket `json:"immutable_bucket,omitempty"`
	Actor           CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentActor            `json:"actor"`
	// ScmPush - SCM push that triggered the deployment. `null` unless `trigger` is `scm`.
	ScmPush *CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeploymentScmPush `json:"scm_push,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomProvidersCommitsGetOutputTargetCustomProviderVersionEnvironmentsEnvironment represents the custom providers commits get output target custom provider version environments environment type.
type CustomProvidersCommitsGetOutputTargetCustomProviderVersionEnvironmentsEnvironment struct {
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

// CustomProvidersCommitsGetOutputTargetCustomProviderVersionEnvironments - Environment this version is deployed to.
type CustomProvidersCommitsGetOutputTargetCustomProviderVersionEnvironments struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsCurrentVersionForEnvironment - Whether this version is the current one for the environment.
	IsCurrentVersionForEnvironment bool                                                                              `json:"is_current_version_for_environment"`
	Environment                    CustomProvidersCommitsGetOutputTargetCustomProviderVersionEnvironmentsEnvironment `json:"environment"`
}

// CustomProvidersCommitsGetOutputTargetCustomProviderVersionActor represents the custom providers commits get output target custom provider version actor type.
type CustomProvidersCommitsGetOutputTargetCustomProviderVersionActor struct {
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

// CustomProvidersCommitsGetOutputTargetCustomProviderVersionContainerImage - Container image this version deploys. Only set for `container` providers.
type CustomProvidersCommitsGetOutputTargetCustomProviderVersionContainerImage struct {
	// ContainerRegistry - Registry the image is pulled from.
	ContainerRegistry string `json:"container_registry"`
	// ContainerImageTag - Tag selecting the image version.
	ContainerImageTag string `json:"container_image_tag"`
	// ContainerImage - Name of the image in the registry.
	ContainerImage string `json:"container_image"`
}

// CustomProvidersCommitsGetOutputTargetCustomProviderVersionRemoteMcpServer - Remote MCP server this version connects to. Only set for `remote` providers.
type CustomProvidersCommitsGetOutputTargetCustomProviderVersionRemoteMcpServer struct {
	// Url - URL where the remote MCP server is reachable.
	Url string `json:"url"`
	// Transport - MCP transport the remote server speaks.
	Transport string `json:"transport"`
}

// CustomProvidersCommitsGetOutputTargetCustomProviderVersion represents the custom providers commits get output target custom provider version type.
type CustomProvidersCommitsGetOutputTargetCustomProviderVersion struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Build status of this version.
	Status string `json:"status"`
	// Config - Configuration for this version.
	Config *CustomProvidersCommitsGetOutputTargetCustomProviderVersionConfig `json:"config,omitempty"`
	// Index - Sequential number of this version.
	Index float64 `json:"index"`
	// Identifier - Human-readable label for this version.
	Identifier string                                                               `json:"identifier"`
	Deployment CustomProvidersCommitsGetOutputTargetCustomProviderVersionDeployment `json:"deployment"`
	// Environments - Environments this version is deployed to.
	Environments []CustomProvidersCommitsGetOutputTargetCustomProviderVersionEnvironments `json:"environments"`
	// CustomProviderId - Custom provider this version belongs to.
	CustomProviderId string `json:"custom_provider_id"`
	// ProviderId - Associated provider, by ID. `null` when there is none.
	ProviderId *string                                                         `json:"provider_id,omitempty"`
	Actor      CustomProvidersCommitsGetOutputTargetCustomProviderVersionActor `json:"actor"`
	// ContainerImage - Container image this version deploys. Only set for `container` providers.
	ContainerImage *CustomProvidersCommitsGetOutputTargetCustomProviderVersionContainerImage `json:"container_image,omitempty"`
	// RemoteMcpServer - Remote MCP server this version connects to. Only set for `remote` providers.
	RemoteMcpServer *CustomProvidersCommitsGetOutputTargetCustomProviderVersionRemoteMcpServer `json:"remote_mcp_server,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomProvidersCommitsGetOutputPreviousCustomProviderVersionConfigSchema - Configuration fields for the custom provider.
type CustomProvidersCommitsGetOutputPreviousCustomProviderVersionConfigSchema struct {
	// Type - Format of the configuration schema.
	Type string `json:"type"`
	// Schema - JSON Schema for the provider's configuration.
	Schema map[string]any `json:"schema"`
}

// CustomProvidersCommitsGetOutputPreviousCustomProviderVersionConfig - Configuration for this version.
type CustomProvidersCommitsGetOutputPreviousCustomProviderVersionConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Schema - Configuration fields for the custom provider.
	Schema CustomProvidersCommitsGetOutputPreviousCustomProviderVersionConfigSchema `json:"schema"`
	// Transformer - Code that transforms the provider's configuration.
	Transformer string `json:"transformer"`
}

// CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentCommit - Commit that triggered this deployment. `null` for deployments without a commit.
type CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentCommit struct {
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

// CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepositoryProvider - Provider the repository lives on.
type CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepositoryProvider struct {
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

// CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepository represents the custom providers commits get output previous custom provider version deployment immutable bucket scm repo link repository type.
type CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepository struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Provider - Provider the repository lives on.
	Provider CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepositoryProvider `json:"provider"`
	// Url - Web URL of the repository.
	Url string `json:"url"`
	// IsPrivate - Whether the repository is private.
	IsPrivate bool `json:"is_private"`
	// DefaultBranch - Branch used when none is specified.
	DefaultBranch string `json:"default_branch"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentImmutableBucketScmRepoLink - Repository this bucket syncs with. `null` when not linked.
type CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentImmutableBucketScmRepoLink struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// IsLinked - `true` when the bucket is linked to a repository.
	IsLinked string `json:"is_linked"`
	// Path - Path within the repository the bucket syncs.
	Path       *string                                                                                                    `json:"path,omitempty"`
	Repository CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepository `json:"repository"`
}

// CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentImmutableBucket - Immutable copy of the deployed code.
type CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentImmutableBucket struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsImmutable - `true` when the bucket contents can no longer change.
	IsImmutable bool `json:"is_immutable"`
	// IsReadOnly - `true` when the bucket cannot be written to.
	IsReadOnly bool `json:"is_read_only"`
	// ScmRepoLink - Repository this bucket syncs with. `null` when not linked.
	ScmRepoLink *CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentImmutableBucketScmRepoLink `json:"scm_repo_link,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentActor represents the custom providers commits get output previous custom provider version deployment actor type.
type CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentActor struct {
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

// CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentScmPushActor - Actor that pushed.
type CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentScmPushActor struct {
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

// CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentScmPushCommit - Commit that was pushed.
type CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentScmPushCommit struct {
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

// CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentScmPushRepositoryProvider - Provider the repository lives on.
type CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentScmPushRepositoryProvider struct {
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

// CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentScmPushRepository represents the custom providers commits get output previous custom provider version deployment scm push repository type.
type CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentScmPushRepository struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Provider - Provider the repository lives on.
	Provider CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentScmPushRepositoryProvider `json:"provider"`
	// Url - Web URL of the repository.
	Url string `json:"url"`
	// IsPrivate - Whether the repository is private.
	IsPrivate bool `json:"is_private"`
	// DefaultBranch - Branch used when none is specified.
	DefaultBranch string `json:"default_branch"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentScmPush - SCM push that triggered the deployment. `null` unless `trigger` is `scm`.
type CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentScmPush struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Actor - Actor that pushed.
	Actor CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentScmPushActor `json:"actor"`
	// Commit - Commit that was pushed.
	Commit     CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentScmPushCommit     `json:"commit"`
	Repository CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentScmPushRepository `json:"repository"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeployment represents the custom providers commits get output previous custom provider version deployment type.
type CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeployment struct {
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
	Commit *CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentCommit `json:"commit,omitempty"`
	// ImmutableBucket - Immutable copy of the deployed code.
	ImmutableBucket *CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentImmutableBucket `json:"immutable_bucket,omitempty"`
	Actor           CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentActor            `json:"actor"`
	// ScmPush - SCM push that triggered the deployment. `null` unless `trigger` is `scm`.
	ScmPush *CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeploymentScmPush `json:"scm_push,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomProvidersCommitsGetOutputPreviousCustomProviderVersionEnvironmentsEnvironment represents the custom providers commits get output previous custom provider version environments environment type.
type CustomProvidersCommitsGetOutputPreviousCustomProviderVersionEnvironmentsEnvironment struct {
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

// CustomProvidersCommitsGetOutputPreviousCustomProviderVersionEnvironments - Environment this version is deployed to.
type CustomProvidersCommitsGetOutputPreviousCustomProviderVersionEnvironments struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsCurrentVersionForEnvironment - Whether this version is the current one for the environment.
	IsCurrentVersionForEnvironment bool                                                                                `json:"is_current_version_for_environment"`
	Environment                    CustomProvidersCommitsGetOutputPreviousCustomProviderVersionEnvironmentsEnvironment `json:"environment"`
}

// CustomProvidersCommitsGetOutputPreviousCustomProviderVersionActor represents the custom providers commits get output previous custom provider version actor type.
type CustomProvidersCommitsGetOutputPreviousCustomProviderVersionActor struct {
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

// CustomProvidersCommitsGetOutputPreviousCustomProviderVersionContainerImage - Container image this version deploys. Only set for `container` providers.
type CustomProvidersCommitsGetOutputPreviousCustomProviderVersionContainerImage struct {
	// ContainerRegistry - Registry the image is pulled from.
	ContainerRegistry string `json:"container_registry"`
	// ContainerImageTag - Tag selecting the image version.
	ContainerImageTag string `json:"container_image_tag"`
	// ContainerImage - Name of the image in the registry.
	ContainerImage string `json:"container_image"`
}

// CustomProvidersCommitsGetOutputPreviousCustomProviderVersionRemoteMcpServer - Remote MCP server this version connects to. Only set for `remote` providers.
type CustomProvidersCommitsGetOutputPreviousCustomProviderVersionRemoteMcpServer struct {
	// Url - URL where the remote MCP server is reachable.
	Url string `json:"url"`
	// Transport - MCP transport the remote server speaks.
	Transport string `json:"transport"`
}

// CustomProvidersCommitsGetOutputPreviousCustomProviderVersion - Version previously in the environment. `null` when there was none.
type CustomProvidersCommitsGetOutputPreviousCustomProviderVersion struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Build status of this version.
	Status string `json:"status"`
	// Config - Configuration for this version.
	Config *CustomProvidersCommitsGetOutputPreviousCustomProviderVersionConfig `json:"config,omitempty"`
	// Index - Sequential number of this version.
	Index float64 `json:"index"`
	// Identifier - Human-readable label for this version.
	Identifier string                                                                 `json:"identifier"`
	Deployment CustomProvidersCommitsGetOutputPreviousCustomProviderVersionDeployment `json:"deployment"`
	// Environments - Environments this version is deployed to.
	Environments []CustomProvidersCommitsGetOutputPreviousCustomProviderVersionEnvironments `json:"environments"`
	// CustomProviderId - Custom provider this version belongs to.
	CustomProviderId string `json:"custom_provider_id"`
	// ProviderId - Associated provider, by ID. `null` when there is none.
	ProviderId *string                                                           `json:"provider_id,omitempty"`
	Actor      CustomProvidersCommitsGetOutputPreviousCustomProviderVersionActor `json:"actor"`
	// ContainerImage - Container image this version deploys. Only set for `container` providers.
	ContainerImage *CustomProvidersCommitsGetOutputPreviousCustomProviderVersionContainerImage `json:"container_image,omitempty"`
	// RemoteMcpServer - Remote MCP server this version connects to. Only set for `remote` providers.
	RemoteMcpServer *CustomProvidersCommitsGetOutputPreviousCustomProviderVersionRemoteMcpServer `json:"remote_mcp_server,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomProvidersCommitsGetOutputActor represents the custom providers commits get output actor type.
type CustomProvidersCommitsGetOutputActor struct {
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

// CustomProvidersCommitsGetOutputScmPushActor - Actor that pushed.
type CustomProvidersCommitsGetOutputScmPushActor struct {
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

// CustomProvidersCommitsGetOutputScmPushCommit - Commit that was pushed.
type CustomProvidersCommitsGetOutputScmPushCommit struct {
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

// CustomProvidersCommitsGetOutputScmPushRepositoryProvider - Provider the repository lives on.
type CustomProvidersCommitsGetOutputScmPushRepositoryProvider struct {
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

// CustomProvidersCommitsGetOutputScmPushRepository represents the custom providers commits get output scm push repository type.
type CustomProvidersCommitsGetOutputScmPushRepository struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Provider - Provider the repository lives on.
	Provider CustomProvidersCommitsGetOutputScmPushRepositoryProvider `json:"provider"`
	// Url - Web URL of the repository.
	Url string `json:"url"`
	// IsPrivate - Whether the repository is private.
	IsPrivate bool `json:"is_private"`
	// DefaultBranch - Branch used when none is specified.
	DefaultBranch string `json:"default_branch"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsGetOutputScmPush - SCM push that triggered the commit. `null` unless `trigger` is `scm`.
type CustomProvidersCommitsGetOutputScmPush struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Actor - Actor that pushed.
	Actor CustomProvidersCommitsGetOutputScmPushActor `json:"actor"`
	// Commit - Commit that was pushed.
	Commit     CustomProvidersCommitsGetOutputScmPushCommit     `json:"commit"`
	Repository CustomProvidersCommitsGetOutputScmPushRepository `json:"repository"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsGetOutput represents the custom providers commits get output type.
type CustomProvidersCommitsGetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `applied` once the version promotion takes effect.
	Status string `json:"status"`
	// Trigger - How the commit was created: manually, by the system, or by an SCM push.
	Trigger string `json:"trigger"`
	// Error - Failure details. `null` when the commit did not fail.
	Error *CustomProvidersCommitsGetOutputError `json:"error,omitempty"`
	// CustomProviderId - Custom provider this commit belongs to.
	CustomProviderId string `json:"custom_provider_id"`
	// ProviderId - Associated provider, by ID. `null` when there is none.
	ProviderId *string `json:"provider_id,omitempty"`
	// CustomProviderDeploymentId - Deployment this commit built. `null` when there is none.
	CustomProviderDeploymentId *string                                      `json:"custom_provider_deployment_id,omitempty"`
	ToEnvironment              CustomProvidersCommitsGetOutputToEnvironment `json:"to_environment"`
	// FromEnvironment - Environment the version is promoted from. `null` for rollbacks.
	FromEnvironment             *CustomProvidersCommitsGetOutputFromEnvironment            `json:"from_environment,omitempty"`
	TargetCustomProviderVersion CustomProvidersCommitsGetOutputTargetCustomProviderVersion `json:"target_custom_provider_version"`
	// PreviousCustomProviderVersion - Version previously in the environment. `null` when there was none.
	PreviousCustomProviderVersion *CustomProvidersCommitsGetOutputPreviousCustomProviderVersion `json:"previous_custom_provider_version,omitempty"`
	Actor                         CustomProvidersCommitsGetOutputActor                          `json:"actor"`
	// ScmPush - SCM push that triggered the commit. `null` unless `trigger` is `scm`.
	ScmPush *CustomProvidersCommitsGetOutputScmPush `json:"scm_push,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// AppliedAt - When the commit took effect. `null` while `status` is `pending`.
	AppliedAt *time.Time `json:"applied_at,omitempty"`
}

// MapCustomProvidersCommitsGetOutputFromJSON deserializes JSON data into a CustomProvidersCommitsGetOutput.
func MapCustomProvidersCommitsGetOutputFromJSON(data []byte) (*CustomProvidersCommitsGetOutput, error) {
	var v CustomProvidersCommitsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCustomProvidersCommitsGetOutputToJSON serializes a CustomProvidersCommitsGetOutput to JSON.
func MapCustomProvidersCommitsGetOutputToJSON(v *CustomProvidersCommitsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
