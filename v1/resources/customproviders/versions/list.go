package versions

import (
	"encoding/json"
	"time"
)

// CustomProvidersVersionsListOutputItemsConfigSchema - Configuration fields for the custom provider.
type CustomProvidersVersionsListOutputItemsConfigSchema struct {
	// Type - Format of the configuration schema.
	Type string `json:"type"`
	// Schema - JSON Schema for the provider's configuration.
	Schema map[string]any `json:"schema"`
}

// CustomProvidersVersionsListOutputItemsConfig - Configuration for this version.
type CustomProvidersVersionsListOutputItemsConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Schema - Configuration fields for the custom provider.
	Schema CustomProvidersVersionsListOutputItemsConfigSchema `json:"schema"`
	// Transformer - Code that transforms the provider's configuration.
	Transformer string `json:"transformer"`
}

// CustomProvidersVersionsListOutputItemsDeploymentCommit - Commit that triggered this deployment. `null` for deployments without a commit.
type CustomProvidersVersionsListOutputItemsDeploymentCommit struct {
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

// CustomProvidersVersionsListOutputItemsDeploymentImmutableBucketScmRepoLinkRepositoryProvider - Provider the repository lives on.
type CustomProvidersVersionsListOutputItemsDeploymentImmutableBucketScmRepoLinkRepositoryProvider struct {
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

// CustomProvidersVersionsListOutputItemsDeploymentImmutableBucketScmRepoLinkRepository represents the custom providers versions list output items deployment immutable bucket scm repo link repository type.
type CustomProvidersVersionsListOutputItemsDeploymentImmutableBucketScmRepoLinkRepository struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Provider - Provider the repository lives on.
	Provider CustomProvidersVersionsListOutputItemsDeploymentImmutableBucketScmRepoLinkRepositoryProvider `json:"provider"`
	// Url - Web URL of the repository.
	Url string `json:"url"`
	// IsPrivate - Whether the repository is private.
	IsPrivate bool `json:"is_private"`
	// DefaultBranch - Branch used when none is specified.
	DefaultBranch string `json:"default_branch"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersVersionsListOutputItemsDeploymentImmutableBucketScmRepoLink - Repository this bucket syncs with. `null` when not linked.
type CustomProvidersVersionsListOutputItemsDeploymentImmutableBucketScmRepoLink struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// IsLinked - `true` when the bucket is linked to a repository.
	IsLinked string `json:"is_linked"`
	// Path - Path within the repository the bucket syncs.
	Path       *string                                                                              `json:"path,omitempty"`
	Repository CustomProvidersVersionsListOutputItemsDeploymentImmutableBucketScmRepoLinkRepository `json:"repository"`
}

// CustomProvidersVersionsListOutputItemsDeploymentImmutableBucket - Immutable copy of the deployed code.
type CustomProvidersVersionsListOutputItemsDeploymentImmutableBucket struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsImmutable - `true` when the bucket contents can no longer change.
	IsImmutable bool `json:"is_immutable"`
	// IsReadOnly - `true` when the bucket cannot be written to.
	IsReadOnly bool `json:"is_read_only"`
	// ScmRepoLink - Repository this bucket syncs with. `null` when not linked.
	ScmRepoLink *CustomProvidersVersionsListOutputItemsDeploymentImmutableBucketScmRepoLink `json:"scm_repo_link,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersVersionsListOutputItemsDeploymentActor represents the custom providers versions list output items deployment actor type.
type CustomProvidersVersionsListOutputItemsDeploymentActor struct {
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

// CustomProvidersVersionsListOutputItemsDeploymentScmPushActor - Actor that pushed.
type CustomProvidersVersionsListOutputItemsDeploymentScmPushActor struct {
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

// CustomProvidersVersionsListOutputItemsDeploymentScmPushCommit - Commit that was pushed.
type CustomProvidersVersionsListOutputItemsDeploymentScmPushCommit struct {
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

// CustomProvidersVersionsListOutputItemsDeploymentScmPushRepositoryProvider - Provider the repository lives on.
type CustomProvidersVersionsListOutputItemsDeploymentScmPushRepositoryProvider struct {
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

// CustomProvidersVersionsListOutputItemsDeploymentScmPushRepository represents the custom providers versions list output items deployment scm push repository type.
type CustomProvidersVersionsListOutputItemsDeploymentScmPushRepository struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Provider - Provider the repository lives on.
	Provider CustomProvidersVersionsListOutputItemsDeploymentScmPushRepositoryProvider `json:"provider"`
	// Url - Web URL of the repository.
	Url string `json:"url"`
	// IsPrivate - Whether the repository is private.
	IsPrivate bool `json:"is_private"`
	// DefaultBranch - Branch used when none is specified.
	DefaultBranch string `json:"default_branch"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersVersionsListOutputItemsDeploymentScmPush - SCM push that triggered the deployment. `null` unless `trigger` is `scm`.
type CustomProvidersVersionsListOutputItemsDeploymentScmPush struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Actor - Actor that pushed.
	Actor CustomProvidersVersionsListOutputItemsDeploymentScmPushActor `json:"actor"`
	// Commit - Commit that was pushed.
	Commit     CustomProvidersVersionsListOutputItemsDeploymentScmPushCommit     `json:"commit"`
	Repository CustomProvidersVersionsListOutputItemsDeploymentScmPushRepository `json:"repository"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersVersionsListOutputItemsDeployment represents the custom providers versions list output items deployment type.
type CustomProvidersVersionsListOutputItemsDeployment struct {
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
	Commit *CustomProvidersVersionsListOutputItemsDeploymentCommit `json:"commit,omitempty"`
	// ImmutableBucket - Immutable copy of the deployed code.
	ImmutableBucket *CustomProvidersVersionsListOutputItemsDeploymentImmutableBucket `json:"immutable_bucket,omitempty"`
	Actor           CustomProvidersVersionsListOutputItemsDeploymentActor            `json:"actor"`
	// ScmPush - SCM push that triggered the deployment. `null` unless `trigger` is `scm`.
	ScmPush *CustomProvidersVersionsListOutputItemsDeploymentScmPush `json:"scm_push,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomProvidersVersionsListOutputItemsEnvironmentsEnvironment represents the custom providers versions list output items environments environment type.
type CustomProvidersVersionsListOutputItemsEnvironmentsEnvironment struct {
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

// CustomProvidersVersionsListOutputItemsEnvironments - Environment this version is deployed to.
type CustomProvidersVersionsListOutputItemsEnvironments struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsCurrentVersionForEnvironment - Whether this version is the current one for the environment.
	IsCurrentVersionForEnvironment bool                                                          `json:"is_current_version_for_environment"`
	Environment                    CustomProvidersVersionsListOutputItemsEnvironmentsEnvironment `json:"environment"`
}

// CustomProvidersVersionsListOutputItemsActor represents the custom providers versions list output items actor type.
type CustomProvidersVersionsListOutputItemsActor struct {
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

// CustomProvidersVersionsListOutputItemsContainerImage - Container image this version deploys. Only set for `container` providers.
type CustomProvidersVersionsListOutputItemsContainerImage struct {
	// ContainerRegistry - Registry the image is pulled from.
	ContainerRegistry string `json:"container_registry"`
	// ContainerImageTag - Tag selecting the image version.
	ContainerImageTag string `json:"container_image_tag"`
	// ContainerImage - Name of the image in the registry.
	ContainerImage string `json:"container_image"`
}

// CustomProvidersVersionsListOutputItemsRemoteMcpServer - Remote MCP server this version connects to. Only set for `remote` providers.
type CustomProvidersVersionsListOutputItemsRemoteMcpServer struct {
	// Url - URL where the remote MCP server is reachable.
	Url string `json:"url"`
	// Transport - MCP transport the remote server speaks.
	Transport string `json:"transport"`
}

// CustomProvidersVersionsListOutputItems represents the custom providers versions list output items type.
type CustomProvidersVersionsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Build status of this version.
	Status string `json:"status"`
	// Config - Configuration for this version.
	Config *CustomProvidersVersionsListOutputItemsConfig `json:"config,omitempty"`
	// Index - Sequential number of this version.
	Index float64 `json:"index"`
	// Identifier - Human-readable label for this version.
	Identifier string                                           `json:"identifier"`
	Deployment CustomProvidersVersionsListOutputItemsDeployment `json:"deployment"`
	// Environments - Environments this version is deployed to.
	Environments []CustomProvidersVersionsListOutputItemsEnvironments `json:"environments"`
	// CustomProviderId - Custom provider this version belongs to.
	CustomProviderId string `json:"custom_provider_id"`
	// ProviderId - Associated provider, by ID. `null` when there is none.
	ProviderId *string                                     `json:"provider_id,omitempty"`
	Actor      CustomProvidersVersionsListOutputItemsActor `json:"actor"`
	// ContainerImage - Container image this version deploys. Only set for `container` providers.
	ContainerImage *CustomProvidersVersionsListOutputItemsContainerImage `json:"container_image,omitempty"`
	// RemoteMcpServer - Remote MCP server this version connects to. Only set for `remote` providers.
	RemoteMcpServer *CustomProvidersVersionsListOutputItemsRemoteMcpServer `json:"remote_mcp_server,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomProvidersVersionsListOutputPagination represents the custom providers versions list output pagination type.
type CustomProvidersVersionsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// CustomProvidersVersionsListOutput represents the custom providers versions list output type.
type CustomProvidersVersionsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                      `json:"object"`
	Items      []CustomProvidersVersionsListOutputItems    `json:"items"`
	Pagination CustomProvidersVersionsListOutputPagination `json:"pagination"`
}

// MapCustomProvidersVersionsListOutputFromJSON deserializes JSON data into a CustomProvidersVersionsListOutput.
func MapCustomProvidersVersionsListOutputFromJSON(data []byte) (*CustomProvidersVersionsListOutput, error) {
	var v CustomProvidersVersionsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCustomProvidersVersionsListOutputToJSON serializes a CustomProvidersVersionsListOutput to JSON.
func MapCustomProvidersVersionsListOutputToJSON(v *CustomProvidersVersionsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// CustomProvidersVersionsListQueryCreatedAt - Only return results whose custom provider version creation time is in this range. Set `gt`, `lt`, or both.
type CustomProvidersVersionsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// CustomProvidersVersionsListQueryUpdatedAt - Only return results whose custom provider version last update time is in this range. Set `gt`, `lt`, or both.
type CustomProvidersVersionsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// CustomProvidersVersionsListQuery represents the custom providers versions list query type.
type CustomProvidersVersionsListQuery struct {
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
	// Status - Only return versions with one of these deployment statuses.
	Status *any `json:"status,omitempty"`
	// Id - Only return versions with one of these IDs.
	Id *any `json:"id,omitempty"`
	// ProviderId - Filter by provider IDs (matches providers connected to sessions)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderVersionId - Filter by provider version IDs (matches providers connected to sessions)
	ProviderVersionId *any `json:"provider_version_id,omitempty"`
	// CustomProviderId - Filter by custom provider IDs (matches providers connected to sessions)
	CustomProviderId *any `json:"custom_provider_id,omitempty"`
	// CustomProviderDeploymentId - Filter by custom provider deployment IDs (matches providers connected to sessions)
	CustomProviderDeploymentId *any `json:"custom_provider_deployment_id,omitempty"`
	// CustomProviderEnvironmentId - Filter by custom provider environment IDs (matches providers connected to sessions)
	CustomProviderEnvironmentId *any `json:"custom_provider_environment_id,omitempty"`
	// CreatedAt - Only return results whose custom provider version creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *CustomProvidersVersionsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose custom provider version last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *CustomProvidersVersionsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapCustomProvidersVersionsListQueryFromJSON deserializes JSON data into a CustomProvidersVersionsListQuery.
func MapCustomProvidersVersionsListQueryFromJSON(data []byte) (*CustomProvidersVersionsListQuery, error) {
	var v CustomProvidersVersionsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCustomProvidersVersionsListQueryToJSON serializes a CustomProvidersVersionsListQuery to JSON.
func MapCustomProvidersVersionsListQueryToJSON(v *CustomProvidersVersionsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
