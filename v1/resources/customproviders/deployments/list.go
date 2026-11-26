package deployments

import (
	"encoding/json"
	"time"
)

// CustomProvidersDeploymentsListOutputItemsCommit - Commit that triggered this deployment. `null` for deployments without a commit.
type CustomProvidersDeploymentsListOutputItemsCommit struct {
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

// CustomProvidersDeploymentsListOutputItemsImmutableBucketScmRepoLinkRepositoryProvider - Provider the repository lives on.
type CustomProvidersDeploymentsListOutputItemsImmutableBucketScmRepoLinkRepositoryProvider struct {
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

// CustomProvidersDeploymentsListOutputItemsImmutableBucketScmRepoLinkRepository represents the custom providers deployments list output items immutable bucket scm repo link repository type.
type CustomProvidersDeploymentsListOutputItemsImmutableBucketScmRepoLinkRepository struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Provider - Provider the repository lives on.
	Provider CustomProvidersDeploymentsListOutputItemsImmutableBucketScmRepoLinkRepositoryProvider `json:"provider"`
	// Url - Web URL of the repository.
	Url string `json:"url"`
	// IsPrivate - Whether the repository is private.
	IsPrivate bool `json:"is_private"`
	// DefaultBranch - Branch used when none is specified.
	DefaultBranch string `json:"default_branch"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersDeploymentsListOutputItemsImmutableBucketScmRepoLink - Repository this bucket syncs with. `null` when not linked.
type CustomProvidersDeploymentsListOutputItemsImmutableBucketScmRepoLink struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// IsLinked - `true` when the bucket is linked to a repository.
	IsLinked string `json:"is_linked"`
	// Path - Path within the repository the bucket syncs.
	Path       *string                                                                       `json:"path,omitempty"`
	Repository CustomProvidersDeploymentsListOutputItemsImmutableBucketScmRepoLinkRepository `json:"repository"`
}

// CustomProvidersDeploymentsListOutputItemsImmutableBucket - Immutable copy of the deployed code.
type CustomProvidersDeploymentsListOutputItemsImmutableBucket struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsImmutable - `true` when the bucket contents can no longer change.
	IsImmutable bool `json:"is_immutable"`
	// IsReadOnly - `true` when the bucket cannot be written to.
	IsReadOnly bool `json:"is_read_only"`
	// ScmRepoLink - Repository this bucket syncs with. `null` when not linked.
	ScmRepoLink *CustomProvidersDeploymentsListOutputItemsImmutableBucketScmRepoLink `json:"scm_repo_link,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersDeploymentsListOutputItemsActor represents the custom providers deployments list output items actor type.
type CustomProvidersDeploymentsListOutputItemsActor struct {
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

// CustomProvidersDeploymentsListOutputItemsScmPushActor - Actor that pushed.
type CustomProvidersDeploymentsListOutputItemsScmPushActor struct {
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

// CustomProvidersDeploymentsListOutputItemsScmPushCommit - Commit that was pushed.
type CustomProvidersDeploymentsListOutputItemsScmPushCommit struct {
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

// CustomProvidersDeploymentsListOutputItemsScmPushRepositoryProvider - Provider the repository lives on.
type CustomProvidersDeploymentsListOutputItemsScmPushRepositoryProvider struct {
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

// CustomProvidersDeploymentsListOutputItemsScmPushRepository represents the custom providers deployments list output items scm push repository type.
type CustomProvidersDeploymentsListOutputItemsScmPushRepository struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Provider - Provider the repository lives on.
	Provider CustomProvidersDeploymentsListOutputItemsScmPushRepositoryProvider `json:"provider"`
	// Url - Web URL of the repository.
	Url string `json:"url"`
	// IsPrivate - Whether the repository is private.
	IsPrivate bool `json:"is_private"`
	// DefaultBranch - Branch used when none is specified.
	DefaultBranch string `json:"default_branch"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersDeploymentsListOutputItemsScmPush - SCM push that triggered the deployment. `null` unless `trigger` is `scm`.
type CustomProvidersDeploymentsListOutputItemsScmPush struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Actor - Actor that pushed.
	Actor CustomProvidersDeploymentsListOutputItemsScmPushActor `json:"actor"`
	// Commit - Commit that was pushed.
	Commit     CustomProvidersDeploymentsListOutputItemsScmPushCommit     `json:"commit"`
	Repository CustomProvidersDeploymentsListOutputItemsScmPushRepository `json:"repository"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CustomProvidersDeploymentsListOutputItems represents the custom providers deployments list output items type.
type CustomProvidersDeploymentsListOutputItems struct {
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
	Commit *CustomProvidersDeploymentsListOutputItemsCommit `json:"commit,omitempty"`
	// ImmutableBucket - Immutable copy of the deployed code.
	ImmutableBucket *CustomProvidersDeploymentsListOutputItemsImmutableBucket `json:"immutable_bucket,omitempty"`
	Actor           CustomProvidersDeploymentsListOutputItemsActor            `json:"actor"`
	// ScmPush - SCM push that triggered the deployment. `null` unless `trigger` is `scm`.
	ScmPush *CustomProvidersDeploymentsListOutputItemsScmPush `json:"scm_push,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomProvidersDeploymentsListOutputPagination represents the custom providers deployments list output pagination type.
type CustomProvidersDeploymentsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// CustomProvidersDeploymentsListOutput represents the custom providers deployments list output type.
type CustomProvidersDeploymentsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                         `json:"object"`
	Items      []CustomProvidersDeploymentsListOutputItems    `json:"items"`
	Pagination CustomProvidersDeploymentsListOutputPagination `json:"pagination"`
}

// MapCustomProvidersDeploymentsListOutputFromJSON deserializes JSON data into a CustomProvidersDeploymentsListOutput.
func MapCustomProvidersDeploymentsListOutputFromJSON(data []byte) (*CustomProvidersDeploymentsListOutput, error) {
	var v CustomProvidersDeploymentsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCustomProvidersDeploymentsListOutputToJSON serializes a CustomProvidersDeploymentsListOutput to JSON.
func MapCustomProvidersDeploymentsListOutputToJSON(v *CustomProvidersDeploymentsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// CustomProvidersDeploymentsListQueryCreatedAt - Only return results whose custom provider deployment creation time is in this range. Set `gt`, `lt`, or both.
type CustomProvidersDeploymentsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// CustomProvidersDeploymentsListQueryUpdatedAt - Only return results whose custom provider deployment last update time is in this range. Set `gt`, `lt`, or both.
type CustomProvidersDeploymentsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// CustomProvidersDeploymentsListQuery represents the custom providers deployments list query type.
type CustomProvidersDeploymentsListQuery struct {
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
	// Status - Only return deployments with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Id - Only return deployments with one of these IDs.
	Id *any `json:"id,omitempty"`
	// CustomProviderVersionId - Only return deployments of one of these custom provider versions.
	CustomProviderVersionId *any `json:"custom_provider_version_id,omitempty"`
	// CustomProviderId - Only return deployments of one of these custom providers.
	CustomProviderId *any `json:"custom_provider_id,omitempty"`
	// CreatedAt - Only return results whose custom provider deployment creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *CustomProvidersDeploymentsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose custom provider deployment last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *CustomProvidersDeploymentsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapCustomProvidersDeploymentsListQueryFromJSON deserializes JSON data into a CustomProvidersDeploymentsListQuery.
func MapCustomProvidersDeploymentsListQueryFromJSON(data []byte) (*CustomProvidersDeploymentsListQuery, error) {
	var v CustomProvidersDeploymentsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCustomProvidersDeploymentsListQueryToJSON serializes a CustomProvidersDeploymentsListQuery to JSON.
func MapCustomProvidersDeploymentsListQueryToJSON(v *CustomProvidersDeploymentsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
