package providerspecificationchangenotifications

import (
	"encoding/json"
	"time"
)

// ProviderSpecificationChangeNotificationsGetOutputFromSpecification - Specification on one side of the change.
type ProviderSpecificationChangeNotificationsGetOutputFromSpecification struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderSpecificationChangeNotificationsGetOutputToSpecification - Specification on one side of the change.
type ProviderSpecificationChangeNotificationsGetOutputToSpecification struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderSpecificationChangeNotificationsGetOutputFromProviderVersion - Provider version on one side of the change.
type ProviderSpecificationChangeNotificationsGetOutputFromProviderVersion struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Version - Release identifier of the provider version.
	Version string `json:"version"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderSpecificationChangeNotificationsGetOutputToProviderVersion - Provider version on one side of the change.
type ProviderSpecificationChangeNotificationsGetOutputToProviderVersion struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Version - Release identifier of the provider version.
	Version string `json:"version"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderSpecificationChangeNotificationsGetOutput represents the provider specification change notifications get output type.
type ProviderSpecificationChangeNotificationsGetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// ProviderId - Provider whose specification changed.
	ProviderId string `json:"provider_id"`
	// ProviderVersionId - Provider version whose specification changed.
	ProviderVersionId string `json:"provider_version_id"`
	// FromSpecification - Specification on one side of the change.
	FromSpecification *ProviderSpecificationChangeNotificationsGetOutputFromSpecification `json:"from_specification,omitempty"`
	// ToSpecification - Specification on one side of the change.
	ToSpecification *ProviderSpecificationChangeNotificationsGetOutputToSpecification `json:"to_specification,omitempty"`
	// FromProviderVersion - Provider version on one side of the change.
	FromProviderVersion *ProviderSpecificationChangeNotificationsGetOutputFromProviderVersion `json:"from_provider_version,omitempty"`
	// ToProviderVersion - Provider version on one side of the change.
	ToProviderVersion *ProviderSpecificationChangeNotificationsGetOutputToProviderVersion `json:"to_provider_version,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// MapProviderSpecificationChangeNotificationsGetOutputFromJSON deserializes JSON data into a ProviderSpecificationChangeNotificationsGetOutput.
func MapProviderSpecificationChangeNotificationsGetOutputFromJSON(data []byte) (*ProviderSpecificationChangeNotificationsGetOutput, error) {
	var v ProviderSpecificationChangeNotificationsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderSpecificationChangeNotificationsGetOutputToJSON serializes a ProviderSpecificationChangeNotificationsGetOutput to JSON.
func MapProviderSpecificationChangeNotificationsGetOutputToJSON(v *ProviderSpecificationChangeNotificationsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
