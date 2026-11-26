package providerconnections

import (
	"encoding/json"
	"time"
)

// AiGatewayProviderConnectionsGetOutputProvider - The provider the connection sends requests to
type AiGatewayProviderConnectionsGetOutputProvider struct {
	// Id - The provider's unique identifier
	Id string `json:"id"`
	// Slug - The provider's slug, e.g. openai or anthropic
	Slug string `json:"slug"`
	// Name - The provider's display name
	Name string `json:"name"`
}

// AiGatewayProviderConnectionsGetOutputFeatures - Capabilities of the provider account behind the connection
type AiGatewayProviderConnectionsGetOutputFeatures struct {
	// Zdr - Whether the provider account does not retain prompts or completions. Requests that require zero data retention only route to connections with it enabled.
	Zdr bool `json:"zdr"`
}

// AiGatewayProviderConnectionsGetOutputPricingOverride - Prices you pay this provider when they differ from the catalog, e.g. because of a negotiated discount. Used to calculate the cost of requests routed through the connection. Null to use catalog prices.
type AiGatewayProviderConnectionsGetOutputPricingOverride struct {
	// Multiplier - Factor applied to the catalog price of every model without its own price below, e.g. 0.8 for a 20% discount
	Multiplier *float64 `json:"multiplier,omitempty"`
	// Models - Prices that replace the catalog price of a model, keyed by model slug
	Models map[string]map[string]any `json:"models"`
}

// AiGatewayProviderConnectionsGetOutput represents the ai gateway provider connections get output type.
type AiGatewayProviderConnectionsGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The provider connection's unique identifier
	Id string `json:"id"`
	// Name - The connection's name
	Name string `json:"name"`
	// Description - The connection's description
	Description *string `json:"description,omitempty"`
	// Priority - Connections with a higher priority are tried first when several connections can serve a request. Defaults to 0.
	Priority float64 `json:"priority"`
	// Provider - The provider the connection sends requests to
	Provider AiGatewayProviderConnectionsGetOutputProvider `json:"provider"`
	// Config - Provider-specific settings that are not secret, e.g. the resource name of an Azure deployment
	Config map[string]any `json:"config"`
	// Features - Capabilities of the provider account behind the connection
	Features AiGatewayProviderConnectionsGetOutputFeatures `json:"features"`
	// PricingOverride - Prices you pay this provider when they differ from the catalog, e.g. because of a negotiated discount. Used to calculate the cost of requests routed through the connection. Null to use catalog prices.
	PricingOverride *AiGatewayProviderConnectionsGetOutputPricingOverride `json:"pricing_override,omitempty"`
	// CredentialHint - The last characters of the credentials, to recognize them by. The credentials themselves are never returned.
	CredentialHint *string `json:"credential_hint,omitempty"`
	// Status - Whether the gateway can use the connection. valid: the provider accepted the credentials at the last check. pending: the credentials have not been confirmed yet, e.g. because the provider could not be reached; the connection is still used for routing, after valid ones. invalid: the provider rejected the credentials; routing skips the connection, and it is re-checked daily for a week before it waits for new credentials. disabled: you turned the connection off; it is neither routed to nor checked.
	Status string `json:"status"`
	// StatusMessage - Why the last check failed, e.g. the error the provider returned. Null after a successful check.
	StatusMessage *string `json:"status_message,omitempty"`
	// ConsecutiveFailures - How many checks in a row have failed, either because the provider rejected the credentials or could not be reached. Reset by a successful check.
	ConsecutiveFailures float64 `json:"consecutive_failures"`
	// LastCheckedAt - When the credentials were last checked with the provider
	LastCheckedAt *time.Time `json:"last_checked_at,omitempty"`
	// LastValidAt - When the provider last accepted the credentials
	LastValidAt *time.Time `json:"last_valid_at,omitempty"`
	// NextCheckAt - When the credentials are checked next. Null when no check is scheduled, e.g. for disabled or deleted connections.
	NextCheckAt *time.Time `json:"next_check_at,omitempty"`
	// DeletedAt - When the connection was deleted
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// MapAiGatewayProviderConnectionsGetOutputFromJSON deserializes JSON data into a AiGatewayProviderConnectionsGetOutput.
func MapAiGatewayProviderConnectionsGetOutputFromJSON(data []byte) (*AiGatewayProviderConnectionsGetOutput, error) {
	var v AiGatewayProviderConnectionsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayProviderConnectionsGetOutputToJSON serializes a AiGatewayProviderConnectionsGetOutput to JSON.
func MapAiGatewayProviderConnectionsGetOutputToJSON(v *AiGatewayProviderConnectionsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
