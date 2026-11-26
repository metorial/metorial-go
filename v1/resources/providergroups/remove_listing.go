package providergroups

import (
	"encoding/json"
	"time"
)

// ProviderGroupsRemoveListingOutput represents the provider groups remove listing output type.
type ProviderGroupsRemoveListingOutput struct {
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

// MapProviderGroupsRemoveListingOutputFromJSON deserializes JSON data into a ProviderGroupsRemoveListingOutput.
func MapProviderGroupsRemoveListingOutputFromJSON(data []byte) (*ProviderGroupsRemoveListingOutput, error) {
	var v ProviderGroupsRemoveListingOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderGroupsRemoveListingOutputToJSON serializes a ProviderGroupsRemoveListingOutput to JSON.
func MapProviderGroupsRemoveListingOutputToJSON(v *ProviderGroupsRemoveListingOutput) ([]byte, error) {
	return json.Marshal(v)
}
