package configs

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsConfigsGetConfigSchemaOutputSchema represents the provider deployments configs get config schema output schema type.
type ProviderDeploymentsConfigsGetConfigSchemaOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema for provider configuration
	Schema map[string]any `json:"schema"`
}

// ProviderDeploymentsConfigsGetConfigSchemaOutput represents the provider deployments configs get config schema output type.
type ProviderDeploymentsConfigsGetConfigSchemaOutput struct {
	// Object - String representing the object’s type.
	Object string                                                 `json:"object"`
	Schema *ProviderDeploymentsConfigsGetConfigSchemaOutputSchema `json:"schema,omitempty"`
	// Visibility - Whether config values are stored encrypted or in plain text.
	Visibility string `json:"visibility"`
	// SpecificationId - Specification this schema belongs to.
	SpecificationId string `json:"specification_id"`
	// ProviderId - Provider this schema belongs to.
	ProviderId string `json:"provider_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapProviderDeploymentsConfigsGetConfigSchemaOutputFromJSON deserializes JSON data into a ProviderDeploymentsConfigsGetConfigSchemaOutput.
func MapProviderDeploymentsConfigsGetConfigSchemaOutputFromJSON(data []byte) (*ProviderDeploymentsConfigsGetConfigSchemaOutput, error) {
	var v ProviderDeploymentsConfigsGetConfigSchemaOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsConfigsGetConfigSchemaOutputToJSON serializes a ProviderDeploymentsConfigsGetConfigSchemaOutput to JSON.
func MapProviderDeploymentsConfigsGetConfigSchemaOutputToJSON(v *ProviderDeploymentsConfigsGetConfigSchemaOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderDeploymentsConfigsGetConfigSchemaQuery represents the provider deployments configs get config schema query type.
type ProviderDeploymentsConfigsGetConfigSchemaQuery struct {
	// ProviderId - Provider to resolve the schema for.
	ProviderId *string `json:"provider_id,omitempty"`
	// ProviderConfigId - Config to resolve the schema for.
	ProviderConfigId *string `json:"provider_config_id,omitempty"`
	// ProviderVersionId - Version to resolve the schema for.
	ProviderVersionId *string `json:"provider_version_id,omitempty"`
	// ProviderDeploymentId - Deployment to resolve the schema for.
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
}

// MapProviderDeploymentsConfigsGetConfigSchemaQueryFromJSON deserializes JSON data into a ProviderDeploymentsConfigsGetConfigSchemaQuery.
func MapProviderDeploymentsConfigsGetConfigSchemaQueryFromJSON(data []byte) (*ProviderDeploymentsConfigsGetConfigSchemaQuery, error) {
	var v ProviderDeploymentsConfigsGetConfigSchemaQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsConfigsGetConfigSchemaQueryToJSON serializes a ProviderDeploymentsConfigsGetConfigSchemaQuery to JSON.
func MapProviderDeploymentsConfigsGetConfigSchemaQueryToJSON(v *ProviderDeploymentsConfigsGetConfigSchemaQuery) ([]byte, error) {
	return json.Marshal(v)
}
