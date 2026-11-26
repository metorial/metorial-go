package tokens

import (
	"encoding/json"
	"github.com/metorial/metorial-go/v1/nullable"
	"time"
)

// AiGatewayTokensUpdateOutputConsumer - The consumer the token was created for. Inference logs and usage can be filtered by it. Null if the token has no consumer or the consumer no longer exists.
type AiGatewayTokensUpdateOutputConsumer struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The consumer's unique identifier
	Id string `json:"id"`
	// Name - The consumer's name
	Name string `json:"name"`
	// Email - The consumer's email address
	Email string `json:"email"`
}

// AiGatewayTokensUpdateOutputLimits represents the ai gateway tokens update output limits type.
type AiGatewayTokensUpdateOutputLimits struct {
	// Metric - What the limit counts. spend_usd counts what Metorial bills for requests.
	Metric string `json:"metric"`
	// Amount - The most that can be used per interval
	Amount float64 `json:"amount"`
	// Interval - The window the amount applies to
	Interval string `json:"interval"`
}

// AiGatewayTokensUpdateOutput represents the ai gateway tokens update output type.
type AiGatewayTokensUpdateOutput struct {
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
	Consumer *AiGatewayTokensUpdateOutputConsumer `json:"consumer,omitempty"`
	// AllowedModels - Models and virtual models the token can call. An empty list allows every model.
	AllowedModels []string `json:"allowed_models"`
	// Limits - Limits that reject requests once reached. Every limit is enforced independently.
	Limits []AiGatewayTokensUpdateOutputLimits `json:"limits"`
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

// MapAiGatewayTokensUpdateOutputFromJSON deserializes JSON data into a AiGatewayTokensUpdateOutput.
func MapAiGatewayTokensUpdateOutputFromJSON(data []byte) (*AiGatewayTokensUpdateOutput, error) {
	var v AiGatewayTokensUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayTokensUpdateOutputToJSON serializes a AiGatewayTokensUpdateOutput to JSON.
func MapAiGatewayTokensUpdateOutputToJSON(v *AiGatewayTokensUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// AiGatewayTokensUpdateBodyLimits represents the ai gateway tokens update body limits type.
type AiGatewayTokensUpdateBodyLimits struct {
	// Metric - What the limit counts. spend_usd counts what Metorial bills for requests.
	Metric string `json:"metric"`
	// Amount - The most that can be used per interval
	Amount float64 `json:"amount"`
	// Interval - The window the amount applies to
	Interval string `json:"interval"`
}

// AiGatewayTokensUpdateBody represents the ai gateway tokens update body type.
type AiGatewayTokensUpdateBody struct {
	// Name - A name for the token
	Name *string `json:"name,omitempty"`
	// AllowedModels - Models and virtual models the token can call, by slug. An empty list allows every model.
	AllowedModels *[]string `json:"allowed_models,omitempty"`
	// Limits - Limits that reject requests once reached. Replaces the existing limits when updated.
	Limits *[]AiGatewayTokensUpdateBodyLimits `json:"limits,omitempty"`
	// Metadata - User-defined metadata that virtual model rules can read. Replaces existing metadata when updated.
	Metadata *map[string]string `json:"metadata,omitempty"`
	// ExpiresAt - When the token stops working. Pass null to never expire.
	ExpiresAt *nullable.Nullable[time.Time] `json:"expires_at,omitempty"`
}

// MapAiGatewayTokensUpdateBodyFromJSON deserializes JSON data into a AiGatewayTokensUpdateBody.
func MapAiGatewayTokensUpdateBodyFromJSON(data []byte) (*AiGatewayTokensUpdateBody, error) {
	var v AiGatewayTokensUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayTokensUpdateBodyToJSON serializes a AiGatewayTokensUpdateBody to JSON.
func MapAiGatewayTokensUpdateBodyToJSON(v *AiGatewayTokensUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
