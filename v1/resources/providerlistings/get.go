package providerlistings

import (
	"encoding/json"
	"time"
)

// ProviderListingsGetOutputAttributes - Listing attribute flags
type ProviderListingsGetOutputAttributes struct {
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

// ProviderListingsGetOutputProviderPublisher represents the provider listings get output provider publisher type.
type ProviderListingsGetOutputProviderPublisher struct {
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

// ProviderListingsGetOutputProviderCurrentVersion represents the provider listings get output provider current version type.
type ProviderListingsGetOutputProviderCurrentVersion struct {
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

// ProviderListingsGetOutputProviderOauthAutoRegistration represents the provider listings get output provider oauth auto registration type.
type ProviderListingsGetOutputProviderOauthAutoRegistration struct {
	// Status - Auto-registration status
	Status string `json:"status"`
}

// ProviderListingsGetOutputProviderOauth represents the provider listings get output provider oauth type.
type ProviderListingsGetOutputProviderOauth struct {
	// Status - OAuth status
	Status string `json:"status"`
	// CallbackUrl - Callback URL to register on your OAuth app.
	CallbackUrl      *string                                                `json:"callback_url,omitempty"`
	AutoRegistration ProviderListingsGetOutputProviderOauthAutoRegistration `json:"auto_registration"`
}

// ProviderListingsGetOutputProvider represents the provider listings get output provider type.
type ProviderListingsGetOutputProvider struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Access - `public` providers are available to every instance. `tenant` providers are your own.
	Access string `json:"access"`
	// Status - Archived once the provider is archived.
	Status         string                                           `json:"status"`
	Publisher      ProviderListingsGetOutputProviderPublisher       `json:"publisher"`
	CurrentVersion *ProviderListingsGetOutputProviderCurrentVersion `json:"current_version,omitempty"`
	Oauth          *ProviderListingsGetOutputProviderOauth          `json:"oauth,omitempty"`
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

// ProviderListingsGetOutputCategories represents the provider listings get output categories type.
type ProviderListingsGetOutputCategories struct {
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

// ProviderListingsGetOutputCollections represents the provider listings get output collections type.
type ProviderListingsGetOutputCollections struct {
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

// ProviderListingsGetOutputGroups represents the provider listings get output groups type.
type ProviderListingsGetOutputGroups struct {
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

// ProviderListingsGetOutput represents the provider listings get output type.
type ProviderListingsGetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Attributes - Listing attribute flags
	Attributes ProviderListingsGetOutputAttributes `json:"attributes"`
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
	Skills   []string                          `json:"skills"`
	Provider ProviderListingsGetOutputProvider `json:"provider"`
	// Categories - Provider categories for organization and filtering
	Categories []ProviderListingsGetOutputCategories `json:"categories"`
	// Collections - Provider collections this provider belongs to
	Collections []ProviderListingsGetOutputCollections `json:"collections"`
	// Groups - Provider groups this listing is organized into.
	Groups []ProviderListingsGetOutputGroups `json:"groups"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapProviderListingsGetOutputFromJSON deserializes JSON data into a ProviderListingsGetOutput.
func MapProviderListingsGetOutputFromJSON(data []byte) (*ProviderListingsGetOutput, error) {
	var v ProviderListingsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderListingsGetOutputToJSON serializes a ProviderListingsGetOutput to JSON.
func MapProviderListingsGetOutputToJSON(v *ProviderListingsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
