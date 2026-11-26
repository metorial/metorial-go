package providerconnections

import (
	"encoding/json"
	"time"
)

// AiGatewayProviderConnectionsListOutputItemsProvider - The provider the connection sends requests to
type AiGatewayProviderConnectionsListOutputItemsProvider struct {
	// Id - The provider's unique identifier
	Id string `json:"id"`
	// Slug - The provider's slug, e.g. openai or anthropic
	Slug string `json:"slug"`
	// Name - The provider's display name
	Name string `json:"name"`
}

// AiGatewayProviderConnectionsListOutputItemsFeatures - Capabilities of the provider account behind the connection
type AiGatewayProviderConnectionsListOutputItemsFeatures struct {
	// Zdr - Whether the provider account does not retain prompts or completions. Requests that require zero data retention only route to connections with it enabled.
	Zdr bool `json:"zdr"`
}

// AiGatewayProviderConnectionsListOutputItemsPricingOverride - Prices you pay this provider when they differ from the catalog, e.g. because of a negotiated discount. Used to calculate the cost of requests routed through the connection. Null to use catalog prices.
type AiGatewayProviderConnectionsListOutputItemsPricingOverride struct {
	// Multiplier - Factor applied to the catalog price of every model without its own price below, e.g. 0.8 for a 20% discount
	Multiplier *float64 `json:"multiplier,omitempty"`
	// Models - Prices that replace the catalog price of a model, keyed by model slug
	Models map[string]map[string]any `json:"models"`
}

// AiGatewayProviderConnectionsListOutputItems represents the ai gateway provider connections list output items type.
type AiGatewayProviderConnectionsListOutputItems struct {
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
	Provider AiGatewayProviderConnectionsListOutputItemsProvider `json:"provider"`
	// Config - Provider-specific settings that are not secret, e.g. the resource name of an Azure deployment
	Config map[string]any `json:"config"`
	// Features - Capabilities of the provider account behind the connection
	Features AiGatewayProviderConnectionsListOutputItemsFeatures `json:"features"`
	// PricingOverride - Prices you pay this provider when they differ from the catalog, e.g. because of a negotiated discount. Used to calculate the cost of requests routed through the connection. Null to use catalog prices.
	PricingOverride *AiGatewayProviderConnectionsListOutputItemsPricingOverride `json:"pricing_override,omitempty"`
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

// AiGatewayProviderConnectionsListOutputPagination represents the ai gateway provider connections list output pagination type.
type AiGatewayProviderConnectionsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// AiGatewayProviderConnectionsListOutput represents the ai gateway provider connections list output type.
type AiGatewayProviderConnectionsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                           `json:"object"`
	Items      []AiGatewayProviderConnectionsListOutputItems    `json:"items"`
	Pagination AiGatewayProviderConnectionsListOutputPagination `json:"pagination"`
}

// MapAiGatewayProviderConnectionsListOutputFromJSON deserializes JSON data into a AiGatewayProviderConnectionsListOutput.
func MapAiGatewayProviderConnectionsListOutputFromJSON(data []byte) (*AiGatewayProviderConnectionsListOutput, error) {
	var v AiGatewayProviderConnectionsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayProviderConnectionsListOutputToJSON serializes a AiGatewayProviderConnectionsListOutput to JSON.
func MapAiGatewayProviderConnectionsListOutputToJSON(v *AiGatewayProviderConnectionsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// AiGatewayProviderConnectionsListQuery represents the ai gateway provider connections list query type.
type AiGatewayProviderConnectionsListQuery struct {
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
	// Provider - Only connections to the provider with this slug, e.g. openai
	Provider *string `json:"provider,omitempty"`
	// Status - Only connections with one of these statuses
	Status *any `json:"status,omitempty"`
}

// MapAiGatewayProviderConnectionsListQueryFromJSON deserializes JSON data into a AiGatewayProviderConnectionsListQuery.
func MapAiGatewayProviderConnectionsListQueryFromJSON(data []byte) (*AiGatewayProviderConnectionsListQuery, error) {
	var v AiGatewayProviderConnectionsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayProviderConnectionsListQueryToJSON serializes a AiGatewayProviderConnectionsListQuery to JSON.
func MapAiGatewayProviderConnectionsListQueryToJSON(v *AiGatewayProviderConnectionsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
