package tokens

import (
	"encoding/json"
	"time"
)

// AiGatewayTokensCreateOutputConsumer - The consumer the token was created for. Inference logs and usage can be filtered by it. Null if the token has no consumer or the consumer no longer exists.
type AiGatewayTokensCreateOutputConsumer struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The consumer's unique identifier
	Id string `json:"id"`
	// Name - The consumer's name
	Name string `json:"name"`
	// Email - The consumer's email address
	Email string `json:"email"`
}

// AiGatewayTokensCreateOutputLimits represents the ai gateway tokens create output limits type.
type AiGatewayTokensCreateOutputLimits struct {
	// Metric - What the limit counts. spend_usd counts what Metorial bills for requests.
	Metric string `json:"metric"`
	// Amount - The most that can be used per interval
	Amount float64 `json:"amount"`
	// Interval - The window the amount applies to
	Interval string `json:"interval"`
}

// AiGatewayTokensCreateOutput represents the ai gateway tokens create output type.
type AiGatewayTokensCreateOutput struct {
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
	Consumer *AiGatewayTokensCreateOutputConsumer `json:"consumer,omitempty"`
	// AllowedModels - Models and virtual models the token can call. An empty list allows every model.
	AllowedModels []string `json:"allowed_models"`
	// Limits - Limits that reject requests once reached. Every limit is enforced independently.
	Limits []AiGatewayTokensCreateOutputLimits `json:"limits"`
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

// MapAiGatewayTokensCreateOutputFromJSON deserializes JSON data into a AiGatewayTokensCreateOutput.
func MapAiGatewayTokensCreateOutputFromJSON(data []byte) (*AiGatewayTokensCreateOutput, error) {
	var v AiGatewayTokensCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayTokensCreateOutputToJSON serializes a AiGatewayTokensCreateOutput to JSON.
func MapAiGatewayTokensCreateOutputToJSON(v *AiGatewayTokensCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// AiGatewayTokensCreateBodyLimits represents the ai gateway tokens create body limits type.
type AiGatewayTokensCreateBodyLimits struct {
	// Metric - What the limit counts. spend_usd counts what Metorial bills for requests.
	Metric string `json:"metric"`
	// Amount - The most that can be used per interval
	Amount float64 `json:"amount"`
	// Interval - The window the amount applies to
	Interval string `json:"interval"`
}

// AiGatewayTokensCreateBody represents the ai gateway tokens create body type.
type AiGatewayTokensCreateBody struct {
	// Name - A name for the token
	Name string `json:"name"`
	// ConsumerId - The consumer the token is for. Requests made with it can be filtered by this consumer in inference logs and usage, and virtual model rules can route on it. Cannot be changed afterwards.
	ConsumerId *string `json:"consumer_id,omitempty"`
	// AllowedModels - Models and virtual models the token can call, by slug. An empty list allows every model.
	AllowedModels *[]string `json:"allowed_models,omitempty"`
	// Limits - Limits that reject requests once reached. Replaces the existing limits when updated.
	Limits *[]AiGatewayTokensCreateBodyLimits `json:"limits,omitempty"`
	// Metadata - User-defined metadata that virtual model rules can read. Replaces existing metadata when updated.
	Metadata *map[string]string `json:"metadata,omitempty"`
	// ExpiresAt - When the token stops working. Omit for no expiry.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// MapAiGatewayTokensCreateBodyFromJSON deserializes JSON data into a AiGatewayTokensCreateBody.
func MapAiGatewayTokensCreateBodyFromJSON(data []byte) (*AiGatewayTokensCreateBody, error) {
	var v AiGatewayTokensCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayTokensCreateBodyToJSON serializes a AiGatewayTokensCreateBody to JSON.
func MapAiGatewayTokensCreateBodyToJSON(v *AiGatewayTokensCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
