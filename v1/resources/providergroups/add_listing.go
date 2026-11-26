package providergroups

import (
	"encoding/json"
	"time"
)

// ProviderGroupsAddListingOutput represents the provider groups add listing output type.
type ProviderGroupsAddListingOutput struct {
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

// MapProviderGroupsAddListingOutputFromJSON deserializes JSON data into a ProviderGroupsAddListingOutput.
func MapProviderGroupsAddListingOutputFromJSON(data []byte) (*ProviderGroupsAddListingOutput, error) {
	var v ProviderGroupsAddListingOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderGroupsAddListingOutputToJSON serializes a ProviderGroupsAddListingOutput to JSON.
func MapProviderGroupsAddListingOutputToJSON(v *ProviderGroupsAddListingOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderGroupsAddListingBody represents the provider groups add listing body type.
type ProviderGroupsAddListingBody struct {
	// ProviderListingId - Listing to add to the group.
	ProviderListingId string `json:"provider_listing_id"`
}

// MapProviderGroupsAddListingBodyFromJSON deserializes JSON data into a ProviderGroupsAddListingBody.
func MapProviderGroupsAddListingBodyFromJSON(data []byte) (*ProviderGroupsAddListingBody, error) {
	var v ProviderGroupsAddListingBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderGroupsAddListingBodyToJSON serializes a ProviderGroupsAddListingBody to JSON.
func MapProviderGroupsAddListingBodyToJSON(v *ProviderGroupsAddListingBody) ([]byte, error) {
	return json.Marshal(v)
}
