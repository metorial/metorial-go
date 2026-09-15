package providerspecificationchangenotifications

import (
	"encoding/json"
	"time"
)

// ProviderSpecificationChangeNotificationsGetOutputFromSpecification represents the provider specification change notifications get output from specification type.
type ProviderSpecificationChangeNotificationsGetOutputFromSpecification struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ProviderSpecificationChangeNotificationsGetOutputToSpecification represents the provider specification change notifications get output to specification type.
type ProviderSpecificationChangeNotificationsGetOutputToSpecification struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ProviderSpecificationChangeNotificationsGetOutputFromProviderVersion represents the provider specification change notifications get output from provider version type.
type ProviderSpecificationChangeNotificationsGetOutputFromProviderVersion struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Version     string    `json:"version"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ProviderSpecificationChangeNotificationsGetOutputToProviderVersion represents the provider specification change notifications get output to provider version type.
type ProviderSpecificationChangeNotificationsGetOutputToProviderVersion struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Version     string    `json:"version"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ProviderSpecificationChangeNotificationsGetOutput represents the provider specification change notifications get output type.
type ProviderSpecificationChangeNotificationsGetOutput struct {
	Object              string                                                                `json:"object"`
	Id                  string                                                                `json:"id"`
	ProviderId          string                                                                `json:"provider_id"`
	ProviderVersionId   string                                                                `json:"provider_version_id"`
	FromSpecification   *ProviderSpecificationChangeNotificationsGetOutputFromSpecification   `json:"from_specification,omitempty"`
	ToSpecification     *ProviderSpecificationChangeNotificationsGetOutputToSpecification     `json:"to_specification,omitempty"`
	FromProviderVersion *ProviderSpecificationChangeNotificationsGetOutputFromProviderVersion `json:"from_provider_version,omitempty"`
	ToProviderVersion   *ProviderSpecificationChangeNotificationsGetOutputToProviderVersion   `json:"to_provider_version,omitempty"`
	CreatedAt           time.Time                                                             `json:"created_at"`
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
