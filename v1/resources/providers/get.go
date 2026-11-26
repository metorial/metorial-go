package providers

import (
	"encoding/json"
	"time"
)

// ProvidersGetOutputPublisher represents the providers get output publisher type.
type ProvidersGetOutputPublisher struct {
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

// ProvidersGetOutputCurrentVersion represents the providers get output current version type.
type ProvidersGetOutputCurrentVersion struct {
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

// ProvidersGetOutputOauthAutoRegistration represents the providers get output oauth auto registration type.
type ProvidersGetOutputOauthAutoRegistration struct {
	// Status - Auto-registration status
	Status string `json:"status"`
}

// ProvidersGetOutputOauth represents the providers get output oauth type.
type ProvidersGetOutputOauth struct {
	// Status - OAuth status
	Status string `json:"status"`
	// CallbackUrl - Callback URL to register on your OAuth app.
	CallbackUrl      *string                                 `json:"callback_url,omitempty"`
	AutoRegistration ProvidersGetOutputOauthAutoRegistration `json:"auto_registration"`
}

// ProvidersGetOutput represents the providers get output type.
type ProvidersGetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Access - `public` providers are available to every instance. `tenant` providers are your own.
	Access string `json:"access"`
	// Status - Archived once the provider is archived.
	Status         string                            `json:"status"`
	Publisher      ProvidersGetOutputPublisher       `json:"publisher"`
	CurrentVersion *ProvidersGetOutputCurrentVersion `json:"current_version,omitempty"`
	Oauth          *ProvidersGetOutputOauth          `json:"oauth,omitempty"`
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

// MapProvidersGetOutputFromJSON deserializes JSON data into a ProvidersGetOutput.
func MapProvidersGetOutputFromJSON(data []byte) (*ProvidersGetOutput, error) {
	var v ProvidersGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProvidersGetOutputToJSON serializes a ProvidersGetOutput to JSON.
func MapProvidersGetOutputToJSON(v *ProvidersGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
