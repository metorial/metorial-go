package commits

import (
	"encoding/json"
	"time"
)

// CustomProvidersCommitsListOutputItemsError - Failure details. `null` when the commit did not fail.
type CustomProvidersCommitsListOutputItemsError struct {
	// Code - Machine-readable error code.
	Code string `json:"code"`
	// Message - Human-readable error message.
	Message string `json:"message"`
}

// CustomProvidersCommitsListOutputItemsToEnvironment represents the custom providers commits list output items to environment type.
type CustomProvidersCommitsListOutputItemsToEnvironment struct {
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

// CustomProvidersCommitsListOutputItemsFromEnvironment - Environment the version is promoted from. `null` for rollbacks.
type CustomProvidersCommitsListOutputItemsFromEnvironment struct {
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

// CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionConfigSchema - Configuration fields for the custom provider.
type CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionConfigSchema struct {
	// Type - Format of the configuration schema.
	Type string `json:"type"`
	// Schema - JSON Schema for the provider's configuration.
	Schema map[string]any `json:"schema"`
}

// CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionConfig - Configuration for this version.
type CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Schema - Configuration fields for the custom provider.
	Schema CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionConfigSchema `json:"schema"`
	// Transformer - Code that transforms the provider's configuration.
	Transformer string `json:"transformer"`
}

// CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentCommit - Commit that triggered this deployment. `null` for deployments without a commit.
type CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentCommit struct {
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

// CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepositoryProvider - Provider the repository lives on.
type CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepositoryProvider struct {
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

// CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepository represents the custom providers commits list output items target custom provider version deployment immutable bucket scm repo link repository type.
type CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepository struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Provider - Provider the repository lives on.
	Provider CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepositoryProvider `json:"provider"`
	// Url - Web URL of the repository.
	Url string `json:"url"`
	// IsPrivate - Whether the repository is private.
	IsPrivate bool `json:"is_private"`
	// DefaultBranch - Branch used when none is specified.
	DefaultBranch string `json:"default_branch"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentImmutableBucketScmRepoLink - Repository this bucket syncs with. `null` when not linked.
type CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentImmutableBucketScmRepoLink struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// IsLinked - `true` when the bucket is linked to a repository.
	IsLinked string `json:"is_linked"`
	// Path - Path within the repository the bucket syncs.
	Path       *string                                                                                                        `json:"path,omitempty"`
	Repository CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepository `json:"repository"`
}

// CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentImmutableBucket - Immutable copy of the deployed code.
type CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentImmutableBucket struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsImmutable - `true` when the bucket contents can no longer change.
	IsImmutable bool `json:"is_immutable"`
	// IsReadOnly - `true` when the bucket cannot be written to.
	IsReadOnly bool `json:"is_read_only"`
	// ScmRepoLink - Repository this bucket syncs with. `null` when not linked.
	ScmRepoLink *CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentImmutableBucketScmRepoLink `json:"scm_repo_link,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentActor represents the custom providers commits list output items target custom provider version deployment actor type.
type CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentActor struct {
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

// CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentScmPushActor - Actor that pushed.
type CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentScmPushActor struct {
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

// CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentScmPushCommit - Commit that was pushed.
type CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentScmPushCommit struct {
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

// CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentScmPushRepositoryProvider - Provider the repository lives on.
type CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentScmPushRepositoryProvider struct {
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

// CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentScmPushRepository represents the custom providers commits list output items target custom provider version deployment scm push repository type.
type CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentScmPushRepository struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Provider - Provider the repository lives on.
	Provider CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentScmPushRepositoryProvider `json:"provider"`
	// Url - Web URL of the repository.
	Url string `json:"url"`
	// IsPrivate - Whether the repository is private.
	IsPrivate bool `json:"is_private"`
	// DefaultBranch - Branch used when none is specified.
	DefaultBranch string `json:"default_branch"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentScmPush - SCM push that triggered the deployment. `null` unless `trigger` is `scm`.
type CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentScmPush struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Actor - Actor that pushed.
	Actor CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentScmPushActor `json:"actor"`
	// Commit - Commit that was pushed.
	Commit     CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentScmPushCommit     `json:"commit"`
	Repository CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentScmPushRepository `json:"repository"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeployment represents the custom providers commits list output items target custom provider version deployment type.
type CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeployment struct {
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
	Commit *CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentCommit `json:"commit,omitempty"`
	// ImmutableBucket - Immutable copy of the deployed code.
	ImmutableBucket *CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentImmutableBucket `json:"immutable_bucket,omitempty"`
	Actor           CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentActor            `json:"actor"`
	// ScmPush - SCM push that triggered the deployment. `null` unless `trigger` is `scm`.
	ScmPush *CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeploymentScmPush `json:"scm_push,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionEnvironmentsEnvironment represents the custom providers commits list output items target custom provider version environments environment type.
type CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionEnvironmentsEnvironment struct {
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

// CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionEnvironments - Environment this version is deployed to.
type CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionEnvironments struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsCurrentVersionForEnvironment - Whether this version is the current one for the environment.
	IsCurrentVersionForEnvironment bool                                                                                    `json:"is_current_version_for_environment"`
	Environment                    CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionEnvironmentsEnvironment `json:"environment"`
}

// CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionActor represents the custom providers commits list output items target custom provider version actor type.
type CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionActor struct {
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

// CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionContainerImage - Container image this version deploys. Only set for `container` providers.
type CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionContainerImage struct {
	// ContainerRegistry - Registry the image is pulled from.
	ContainerRegistry string `json:"container_registry"`
	// ContainerImageTag - Tag selecting the image version.
	ContainerImageTag string `json:"container_image_tag"`
	// ContainerImage - Name of the image in the registry.
	ContainerImage string `json:"container_image"`
}

// CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionRemoteMcpServer - Remote MCP server this version connects to. Only set for `remote` providers.
type CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionRemoteMcpServer struct {
	// Url - URL where the remote MCP server is reachable.
	Url string `json:"url"`
	// Transport - MCP transport the remote server speaks.
	Transport string `json:"transport"`
}

// CustomProvidersCommitsListOutputItemsTargetCustomProviderVersion represents the custom providers commits list output items target custom provider version type.
type CustomProvidersCommitsListOutputItemsTargetCustomProviderVersion struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Build status of this version.
	Status string `json:"status"`
	// Config - Configuration for this version.
	Config *CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionConfig `json:"config,omitempty"`
	// Index - Sequential number of this version.
	Index float64 `json:"index"`
	// Identifier - Human-readable label for this version.
	Identifier string                                                                     `json:"identifier"`
	Deployment CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionDeployment `json:"deployment"`
	// Environments - Environments this version is deployed to.
	Environments []CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionEnvironments `json:"environments"`
	// CustomProviderId - Custom provider this version belongs to.
	CustomProviderId string `json:"custom_provider_id"`
	// ProviderId - Associated provider, by ID. `null` when there is none.
	ProviderId *string                                                               `json:"provider_id,omitempty"`
	Actor      CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionActor `json:"actor"`
	// ContainerImage - Container image this version deploys. Only set for `container` providers.
	ContainerImage *CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionContainerImage `json:"container_image,omitempty"`
	// RemoteMcpServer - Remote MCP server this version connects to. Only set for `remote` providers.
	RemoteMcpServer *CustomProvidersCommitsListOutputItemsTargetCustomProviderVersionRemoteMcpServer `json:"remote_mcp_server,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionConfigSchema - Configuration fields for the custom provider.
type CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionConfigSchema struct {
	// Type - Format of the configuration schema.
	Type string `json:"type"`
	// Schema - JSON Schema for the provider's configuration.
	Schema map[string]any `json:"schema"`
}

// CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionConfig - Configuration for this version.
type CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Schema - Configuration fields for the custom provider.
	Schema CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionConfigSchema `json:"schema"`
	// Transformer - Code that transforms the provider's configuration.
	Transformer string `json:"transformer"`
}

// CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentCommit - Commit that triggered this deployment. `null` for deployments without a commit.
type CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentCommit struct {
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

// CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepositoryProvider - Provider the repository lives on.
type CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepositoryProvider struct {
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

// CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepository represents the custom providers commits list output items previous custom provider version deployment immutable bucket scm repo link repository type.
type CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepository struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Provider - Provider the repository lives on.
	Provider CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepositoryProvider `json:"provider"`
	// Url - Web URL of the repository.
	Url string `json:"url"`
	// IsPrivate - Whether the repository is private.
	IsPrivate bool `json:"is_private"`
	// DefaultBranch - Branch used when none is specified.
	DefaultBranch string `json:"default_branch"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentImmutableBucketScmRepoLink - Repository this bucket syncs with. `null` when not linked.
type CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentImmutableBucketScmRepoLink struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// IsLinked - `true` when the bucket is linked to a repository.
	IsLinked string `json:"is_linked"`
	// Path - Path within the repository the bucket syncs.
	Path       *string                                                                                                          `json:"path,omitempty"`
	Repository CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentImmutableBucketScmRepoLinkRepository `json:"repository"`
}

// CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentImmutableBucket - Immutable copy of the deployed code.
type CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentImmutableBucket struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsImmutable - `true` when the bucket contents can no longer change.
	IsImmutable bool `json:"is_immutable"`
	// IsReadOnly - `true` when the bucket cannot be written to.
	IsReadOnly bool `json:"is_read_only"`
	// ScmRepoLink - Repository this bucket syncs with. `null` when not linked.
	ScmRepoLink *CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentImmutableBucketScmRepoLink `json:"scm_repo_link,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentActor represents the custom providers commits list output items previous custom provider version deployment actor type.
type CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentActor struct {
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

// CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentScmPushActor - Actor that pushed.
type CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentScmPushActor struct {
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

// CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentScmPushCommit - Commit that was pushed.
type CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentScmPushCommit struct {
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

// CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentScmPushRepositoryProvider - Provider the repository lives on.
type CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentScmPushRepositoryProvider struct {
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

// CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentScmPushRepository represents the custom providers commits list output items previous custom provider version deployment scm push repository type.
type CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentScmPushRepository struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Provider - Provider the repository lives on.
	Provider CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentScmPushRepositoryProvider `json:"provider"`
	// Url - Web URL of the repository.
	Url string `json:"url"`
	// IsPrivate - Whether the repository is private.
	IsPrivate bool `json:"is_private"`
	// DefaultBranch - Branch used when none is specified.
	DefaultBranch string `json:"default_branch"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentScmPush - SCM push that triggered the deployment. `null` unless `trigger` is `scm`.
type CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentScmPush struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Actor - Actor that pushed.
	Actor CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentScmPushActor `json:"actor"`
	// Commit - Commit that was pushed.
	Commit     CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentScmPushCommit     `json:"commit"`
	Repository CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentScmPushRepository `json:"repository"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeployment represents the custom providers commits list output items previous custom provider version deployment type.
type CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeployment struct {
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
	Commit *CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentCommit `json:"commit,omitempty"`
	// ImmutableBucket - Immutable copy of the deployed code.
	ImmutableBucket *CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentImmutableBucket `json:"immutable_bucket,omitempty"`
	Actor           CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentActor            `json:"actor"`
	// ScmPush - SCM push that triggered the deployment. `null` unless `trigger` is `scm`.
	ScmPush *CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeploymentScmPush `json:"scm_push,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionEnvironmentsEnvironment represents the custom providers commits list output items previous custom provider version environments environment type.
type CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionEnvironmentsEnvironment struct {
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

// CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionEnvironments - Environment this version is deployed to.
type CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionEnvironments struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsCurrentVersionForEnvironment - Whether this version is the current one for the environment.
	IsCurrentVersionForEnvironment bool                                                                                      `json:"is_current_version_for_environment"`
	Environment                    CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionEnvironmentsEnvironment `json:"environment"`
}

// CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionActor represents the custom providers commits list output items previous custom provider version actor type.
type CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionActor struct {
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

// CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionContainerImage - Container image this version deploys. Only set for `container` providers.
type CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionContainerImage struct {
	// ContainerRegistry - Registry the image is pulled from.
	ContainerRegistry string `json:"container_registry"`
	// ContainerImageTag - Tag selecting the image version.
	ContainerImageTag string `json:"container_image_tag"`
	// ContainerImage - Name of the image in the registry.
	ContainerImage string `json:"container_image"`
}

// CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionRemoteMcpServer - Remote MCP server this version connects to. Only set for `remote` providers.
type CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionRemoteMcpServer struct {
	// Url - URL where the remote MCP server is reachable.
	Url string `json:"url"`
	// Transport - MCP transport the remote server speaks.
	Transport string `json:"transport"`
}

// CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersion - Version previously in the environment. `null` when there was none.
type CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersion struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Build status of this version.
	Status string `json:"status"`
	// Config - Configuration for this version.
	Config *CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionConfig `json:"config,omitempty"`
	// Index - Sequential number of this version.
	Index float64 `json:"index"`
	// Identifier - Human-readable label for this version.
	Identifier string                                                                       `json:"identifier"`
	Deployment CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionDeployment `json:"deployment"`
	// Environments - Environments this version is deployed to.
	Environments []CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionEnvironments `json:"environments"`
	// CustomProviderId - Custom provider this version belongs to.
	CustomProviderId string `json:"custom_provider_id"`
	// ProviderId - Associated provider, by ID. `null` when there is none.
	ProviderId *string                                                                 `json:"provider_id,omitempty"`
	Actor      CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionActor `json:"actor"`
	// ContainerImage - Container image this version deploys. Only set for `container` providers.
	ContainerImage *CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionContainerImage `json:"container_image,omitempty"`
	// RemoteMcpServer - Remote MCP server this version connects to. Only set for `remote` providers.
	RemoteMcpServer *CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersionRemoteMcpServer `json:"remote_mcp_server,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomProvidersCommitsListOutputItemsActor represents the custom providers commits list output items actor type.
type CustomProvidersCommitsListOutputItemsActor struct {
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

// CustomProvidersCommitsListOutputItemsScmPushActor - Actor that pushed.
type CustomProvidersCommitsListOutputItemsScmPushActor struct {
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

// CustomProvidersCommitsListOutputItemsScmPushCommit - Commit that was pushed.
type CustomProvidersCommitsListOutputItemsScmPushCommit struct {
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

// CustomProvidersCommitsListOutputItemsScmPushRepositoryProvider - Provider the repository lives on.
type CustomProvidersCommitsListOutputItemsScmPushRepositoryProvider struct {
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

// CustomProvidersCommitsListOutputItemsScmPushRepository represents the custom providers commits list output items scm push repository type.
type CustomProvidersCommitsListOutputItemsScmPushRepository struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Provider - Provider the repository lives on.
	Provider CustomProvidersCommitsListOutputItemsScmPushRepositoryProvider `json:"provider"`
	// Url - Web URL of the repository.
	Url string `json:"url"`
	// IsPrivate - Whether the repository is private.
	IsPrivate bool `json:"is_private"`
	// DefaultBranch - Branch used when none is specified.
	DefaultBranch string `json:"default_branch"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsListOutputItemsScmPush - SCM push that triggered the commit. `null` unless `trigger` is `scm`.
type CustomProvidersCommitsListOutputItemsScmPush struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Actor - Actor that pushed.
	Actor CustomProvidersCommitsListOutputItemsScmPushActor `json:"actor"`
	// Commit - Commit that was pushed.
	Commit     CustomProvidersCommitsListOutputItemsScmPushCommit     `json:"commit"`
	Repository CustomProvidersCommitsListOutputItemsScmPushRepository `json:"repository"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersCommitsListOutputItems represents the custom providers commits list output items type.
type CustomProvidersCommitsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `applied` once the version promotion takes effect.
	Status string `json:"status"`
	// Trigger - How the commit was created: manually, by the system, or by an SCM push.
	Trigger string `json:"trigger"`
	// Error - Failure details. `null` when the commit did not fail.
	Error *CustomProvidersCommitsListOutputItemsError `json:"error,omitempty"`
	// CustomProviderId - Custom provider this commit belongs to.
	CustomProviderId string `json:"custom_provider_id"`
	// ProviderId - Associated provider, by ID. `null` when there is none.
	ProviderId *string `json:"provider_id,omitempty"`
	// CustomProviderDeploymentId - Deployment this commit built. `null` when there is none.
	CustomProviderDeploymentId *string                                            `json:"custom_provider_deployment_id,omitempty"`
	ToEnvironment              CustomProvidersCommitsListOutputItemsToEnvironment `json:"to_environment"`
	// FromEnvironment - Environment the version is promoted from. `null` for rollbacks.
	FromEnvironment             *CustomProvidersCommitsListOutputItemsFromEnvironment            `json:"from_environment,omitempty"`
	TargetCustomProviderVersion CustomProvidersCommitsListOutputItemsTargetCustomProviderVersion `json:"target_custom_provider_version"`
	// PreviousCustomProviderVersion - Version previously in the environment. `null` when there was none.
	PreviousCustomProviderVersion *CustomProvidersCommitsListOutputItemsPreviousCustomProviderVersion `json:"previous_custom_provider_version,omitempty"`
	Actor                         CustomProvidersCommitsListOutputItemsActor                          `json:"actor"`
	// ScmPush - SCM push that triggered the commit. `null` unless `trigger` is `scm`.
	ScmPush *CustomProvidersCommitsListOutputItemsScmPush `json:"scm_push,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// AppliedAt - When the commit took effect. `null` while `status` is `pending`.
	AppliedAt *time.Time `json:"applied_at,omitempty"`
}

// CustomProvidersCommitsListOutputPagination represents the custom providers commits list output pagination type.
type CustomProvidersCommitsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// CustomProvidersCommitsListOutput represents the custom providers commits list output type.
type CustomProvidersCommitsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                     `json:"object"`
	Items      []CustomProvidersCommitsListOutputItems    `json:"items"`
	Pagination CustomProvidersCommitsListOutputPagination `json:"pagination"`
}

// MapCustomProvidersCommitsListOutputFromJSON deserializes JSON data into a CustomProvidersCommitsListOutput.
func MapCustomProvidersCommitsListOutputFromJSON(data []byte) (*CustomProvidersCommitsListOutput, error) {
	var v CustomProvidersCommitsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCustomProvidersCommitsListOutputToJSON serializes a CustomProvidersCommitsListOutput to JSON.
func MapCustomProvidersCommitsListOutputToJSON(v *CustomProvidersCommitsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// CustomProvidersCommitsListQueryCreatedAt - Only return results whose custom provider commit creation time is in this range. Set `gt`, `lt`, or both.
type CustomProvidersCommitsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// CustomProvidersCommitsListQueryUpdatedAt - Only return results whose custom provider commit last update time is in this range. Set `gt`, `lt`, or both.
type CustomProvidersCommitsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// CustomProvidersCommitsListQuery represents the custom providers commits list query type.
type CustomProvidersCommitsListQuery struct {
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
	// Id - Only return commits with one of these IDs.
	Id *any `json:"id,omitempty"`
	// CustomProviderVersionId - Only return commits for one of these custom provider versions.
	CustomProviderVersionId *any `json:"custom_provider_version_id,omitempty"`
	// CustomProviderEnvironmentId - Only return commits for one of these custom provider environments.
	CustomProviderEnvironmentId *any `json:"custom_provider_environment_id,omitempty"`
	// CustomProviderId - Only return commits for one of these custom providers.
	CustomProviderId *any `json:"custom_provider_id,omitempty"`
	// ProviderId - Only return commits for one of these provider IDs.
	ProviderId *any `json:"provider_id,omitempty"`
	// CreatedAt - Only return results whose custom provider commit creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *CustomProvidersCommitsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose custom provider commit last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *CustomProvidersCommitsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapCustomProvidersCommitsListQueryFromJSON deserializes JSON data into a CustomProvidersCommitsListQuery.
func MapCustomProvidersCommitsListQueryFromJSON(data []byte) (*CustomProvidersCommitsListQuery, error) {
	var v CustomProvidersCommitsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCustomProvidersCommitsListQueryToJSON serializes a CustomProvidersCommitsListQuery to JSON.
func MapCustomProvidersCommitsListQueryToJSON(v *CustomProvidersCommitsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
