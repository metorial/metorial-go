package tokens

import (
	"encoding/json"
	"time"
)

// AiGatewayTokensListOutputItemsConsumer - The consumer the token was created for. Inference logs and usage can be filtered by it. Null if the token has no consumer or the consumer no longer exists.
type AiGatewayTokensListOutputItemsConsumer struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The consumer's unique identifier
	Id string `json:"id"`
	// Name - The consumer's name
	Name string `json:"name"`
	// Email - The consumer's email address
	Email string `json:"email"`
}

// AiGatewayTokensListOutputItemsLimits represents the ai gateway tokens list output items limits type.
type AiGatewayTokensListOutputItemsLimits struct {
	// Metric - What the limit counts. spend_usd counts what Metorial bills for requests.
	Metric string `json:"metric"`
	// Amount - The most that can be used per interval
	Amount float64 `json:"amount"`
	// Interval - The window the amount applies to
	Interval string `json:"interval"`
}

// AiGatewayTokensListOutputItems represents the ai gateway tokens list output items type.
type AiGatewayTokensListOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The token's unique identifier
	Id string `json:"id"`
	// Name - The token's name
	Name string `json:"name"`
	// Prefix - The first characters of the secret, to recognize the token by
	Prefix string `json:"prefix"`
	// Secret - The secret to send as the API key of AI gateway requests. Only returned when the token is created.
	Secret *string `json:"secret,omitempty"`
	// Consumer - The consumer the token was created for. Inference logs and usage can be filtered by it. Null if the token has no consumer or the consumer no longer exists.
	Consumer *AiGatewayTokensListOutputItemsConsumer `json:"consumer,omitempty"`
	// AllowedModels - Models and virtual models the token can call. An empty list allows every model.
	AllowedModels []string `json:"allowed_models"`
	// Limits - Limits that reject requests once reached. Every limit is enforced independently.
	Limits []AiGatewayTokensListOutputItemsLimits `json:"limits"`
	// Metadata - User-defined metadata. Virtual model rules can read it as token.metadata.
	Metadata map[string]string `json:"metadata"`
	// LastUsedAt - When the token last made a request
	LastUsedAt *time.Time `json:"last_used_at,omitempty"`
	// ExpiresAt - When the token stops working. Null if it never expires.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// RevokedAt - When the token was revoked
	RevokedAt *time.Time `json:"revoked_at,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
}

// AiGatewayTokensListOutputPagination represents the ai gateway tokens list output pagination type.
type AiGatewayTokensListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// AiGatewayTokensListOutput represents the ai gateway tokens list output type.
type AiGatewayTokensListOutput struct {
	// Object - String representing the object’s type.
	Object     string                              `json:"object"`
	Items      []AiGatewayTokensListOutputItems    `json:"items"`
	Pagination AiGatewayTokensListOutputPagination `json:"pagination"`
}

// MapAiGatewayTokensListOutputFromJSON deserializes JSON data into a AiGatewayTokensListOutput.
func MapAiGatewayTokensListOutputFromJSON(data []byte) (*AiGatewayTokensListOutput, error) {
	var v AiGatewayTokensListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayTokensListOutputToJSON serializes a AiGatewayTokensListOutput to JSON.
func MapAiGatewayTokensListOutputToJSON(v *AiGatewayTokensListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// AiGatewayTokensListQuery represents the ai gateway tokens list query type.
type AiGatewayTokensListQuery struct {
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
	// ConsumerId - Only tokens created for this consumer
	ConsumerId *string `json:"consumer_id,omitempty"`
}

// MapAiGatewayTokensListQueryFromJSON deserializes JSON data into a AiGatewayTokensListQuery.
func MapAiGatewayTokensListQueryFromJSON(data []byte) (*AiGatewayTokensListQuery, error) {
	var v AiGatewayTokensListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayTokensListQueryToJSON serializes a AiGatewayTokensListQuery to JSON.
func MapAiGatewayTokensListQueryToJSON(v *AiGatewayTokensListQuery) ([]byte, error) {
	return json.Marshal(v)
}
