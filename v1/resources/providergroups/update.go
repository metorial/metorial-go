package providergroups

import (
	"encoding/json"
	"time"
)

// ProviderGroupsUpdateOutput represents the provider groups update output type.
type ProviderGroupsUpdateOutput struct {
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

// MapProviderGroupsUpdateOutputFromJSON deserializes JSON data into a ProviderGroupsUpdateOutput.
func MapProviderGroupsUpdateOutputFromJSON(data []byte) (*ProviderGroupsUpdateOutput, error) {
	var v ProviderGroupsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderGroupsUpdateOutputToJSON serializes a ProviderGroupsUpdateOutput to JSON.
func MapProviderGroupsUpdateOutputToJSON(v *ProviderGroupsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderGroupsUpdateBody represents the provider groups update body type.
type ProviderGroupsUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
}

// MapProviderGroupsUpdateBodyFromJSON deserializes JSON data into a ProviderGroupsUpdateBody.
func MapProviderGroupsUpdateBodyFromJSON(data []byte) (*ProviderGroupsUpdateBody, error) {
	var v ProviderGroupsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderGroupsUpdateBodyToJSON serializes a ProviderGroupsUpdateBody to JSON.
func MapProviderGroupsUpdateBodyToJSON(v *ProviderGroupsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
