package providerlistings

import (
	"encoding/json"
	"time"
)

// ProviderListingsListOutputItemsAttributes - Listing attribute flags
type ProviderListingsListOutputItemsAttributes struct {
	// IsPublic - Whether publicly visible
	IsPublic bool `json:"is_public"`
	// IsCustomized - Whether has custom config
	IsCustomized bool `json:"is_customized"`
	// IsMetorial - Whether Metorial-maintained
	IsMetorial bool `json:"is_metorial"`
	// IsVerified - Reviewed and trusted by Metorial.
	IsVerified bool `json:"is_verified"`
	// IsOfficial - Whether official integration
	IsOfficial bool `json:"is_official"`
}

// ProviderListingsListOutputItemsProviderPublisher represents the provider listings list output items provider publisher type.
type ProviderListingsListOutputItemsProviderPublisher struct {
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

// ProviderListingsListOutputItemsProviderCurrentVersion represents the provider listings list output items provider current version type.
type ProviderListingsListOutputItemsProviderCurrentVersion struct {
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

// ProviderListingsListOutputItemsProviderOauthAutoRegistration represents the provider listings list output items provider oauth auto registration type.
type ProviderListingsListOutputItemsProviderOauthAutoRegistration struct {
	// Status - Auto-registration status
	Status string `json:"status"`
}

// ProviderListingsListOutputItemsProviderOauth represents the provider listings list output items provider oauth type.
type ProviderListingsListOutputItemsProviderOauth struct {
	// Status - OAuth status
	Status string `json:"status"`
	// CallbackUrl - Callback URL to register on your OAuth app.
	CallbackUrl      *string                                                      `json:"callback_url,omitempty"`
	AutoRegistration ProviderListingsListOutputItemsProviderOauthAutoRegistration `json:"auto_registration"`
}

// ProviderListingsListOutputItemsProvider represents the provider listings list output items provider type.
type ProviderListingsListOutputItemsProvider struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Access - `public` providers are available to every instance. `tenant` providers are your own.
	Access string `json:"access"`
	// Status - Archived once the provider is archived.
	Status         string                                                 `json:"status"`
	Publisher      ProviderListingsListOutputItemsProviderPublisher       `json:"publisher"`
	CurrentVersion *ProviderListingsListOutputItemsProviderCurrentVersion `json:"current_version,omitempty"`
	Oauth          *ProviderListingsListOutputItemsProviderOauth          `json:"oauth,omitempty"`
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

// ProviderListingsListOutputItemsCategories represents the provider listings list output items categories type.
type ProviderListingsListOutputItemsCategories struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Description of providers in this category
	Description string `json:"description"`
	// Slug - URL-safe name for the category.
	Slug string `json:"slug"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderListingsListOutputItemsCollections represents the provider listings list output items collections type.
type ProviderListingsListOutputItemsCollections struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description string `json:"description"`
	// Slug - URL-safe name for the collection.
	Slug string `json:"slug"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderListingsListOutputItemsGroups represents the provider listings list output items groups type.
type ProviderListingsListOutputItemsGroups struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Slug - URL-safe name for the group.
	Slug string `json:"slug"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderListingsListOutputItems represents the provider listings list output items type.
type ProviderListingsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Attributes - Listing attribute flags
	Attributes ProviderListingsListOutputItemsAttributes `json:"attributes"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Slug - URL-safe name for the provider.
	Slug string `json:"slug"`
	// ImageUrl - URL of the listing logo/icon
	ImageUrl string `json:"image_url"`
	// Readme - README content in markdown
	Readme *string `json:"readme,omitempty"`
	// Skills - Capability tags
	Skills   []string                                `json:"skills"`
	Provider ProviderListingsListOutputItemsProvider `json:"provider"`
	// Categories - Provider categories for organization and filtering
	Categories []ProviderListingsListOutputItemsCategories `json:"categories"`
	// Collections - Provider collections this provider belongs to
	Collections []ProviderListingsListOutputItemsCollections `json:"collections"`
	// Groups - Provider groups this listing is organized into.
	Groups []ProviderListingsListOutputItemsGroups `json:"groups"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderListingsListOutputPagination represents the provider listings list output pagination type.
type ProviderListingsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ProviderListingsListOutput represents the provider listings list output type.
type ProviderListingsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                               `json:"object"`
	Items      []ProviderListingsListOutputItems    `json:"items"`
	Pagination ProviderListingsListOutputPagination `json:"pagination"`
}

// MapProviderListingsListOutputFromJSON deserializes JSON data into a ProviderListingsListOutput.
func MapProviderListingsListOutputFromJSON(data []byte) (*ProviderListingsListOutput, error) {
	var v ProviderListingsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderListingsListOutputToJSON serializes a ProviderListingsListOutput to JSON.
func MapProviderListingsListOutputToJSON(v *ProviderListingsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderListingsListQueryCreatedAt - Only return results whose provider listing creation time is in this range. Set `gt`, `lt`, or both.
type ProviderListingsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ProviderListingsListQueryUpdatedAt - Only return results whose provider listing last update time is in this range. Set `gt`, `lt`, or both.
type ProviderListingsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ProviderListingsListQuery represents the provider listings list query type.
type ProviderListingsListQuery struct {
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
	// Search - Search listings by name, description, or readme.
	Search *string `json:"search,omitempty"`
	// Id - Filter by listing ID(s).
	Id *any `json:"id,omitempty"`
	// ProviderCategoryId - Only return listings in these categories.
	ProviderCategoryId *any `json:"provider_category_id,omitempty"`
	// ProviderCollectionId - Only return listings in these collections.
	ProviderCollectionId *any `json:"provider_collection_id,omitempty"`
	// ProviderGroupId - Only return listings in these groups.
	ProviderGroupId *any `json:"provider_group_id,omitempty"`
	// PublisherId - Only return listings from these publishers.
	PublisherId *any `json:"publisher_id,omitempty"`
	// IsOwner - Only return listings published by your organization.
	IsOwner *bool `json:"is_owner,omitempty"`
	// IsPublic - Only return public listings.
	IsPublic *bool `json:"is_public,omitempty"`
	// IsVerified - Only return listings reviewed and trusted by Metorial.
	IsVerified *bool `json:"is_verified,omitempty"`
	// IsOfficial - Only return official listings.
	IsOfficial *bool `json:"is_official,omitempty"`
	// IsMetorial - Only return Metorial-maintained listings.
	IsMetorial *bool `json:"is_metorial,omitempty"`
	// CreatedAt - Only return results whose provider listing creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *ProviderListingsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose provider listing last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *ProviderListingsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapProviderListingsListQueryFromJSON deserializes JSON data into a ProviderListingsListQuery.
func MapProviderListingsListQueryFromJSON(data []byte) (*ProviderListingsListQuery, error) {
	var v ProviderListingsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderListingsListQueryToJSON serializes a ProviderListingsListQuery to JSON.
func MapProviderListingsListQueryToJSON(v *ProviderListingsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
