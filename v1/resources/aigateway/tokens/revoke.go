package tokens

import (
	"encoding/json"
	"time"
)

// AiGatewayTokensRevokeOutputConsumer - The consumer the token was created for. Inference logs and usage can be filtered by it. Null if the token has no consumer or the consumer no longer exists.
type AiGatewayTokensRevokeOutputConsumer struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The consumer's unique identifier
	Id string `json:"id"`
	// Name - The consumer's name
	Name string `json:"name"`
	// Email - The consumer's email address
	Email string `json:"email"`
}

// AiGatewayTokensRevokeOutputLimits represents the ai gateway tokens revoke output limits type.
type AiGatewayTokensRevokeOutputLimits struct {
	// Metric - What the limit counts. spend_usd counts what Metorial bills for requests.
	Metric string `json:"metric"`
	// Amount - The most that can be used per interval
	Amount float64 `json:"amount"`
	// Interval - The window the amount applies to
	Interval string `json:"interval"`
}

// AiGatewayTokensRevokeOutput represents the ai gateway tokens revoke output type.
type AiGatewayTokensRevokeOutput struct {
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
	Consumer *AiGatewayTokensRevokeOutputConsumer `json:"consumer,omitempty"`
	// AllowedModels - Models and virtual models the token can call. An empty list allows every model.
	AllowedModels []string `json:"allowed_models"`
	// Limits - Limits that reject requests once reached. Every limit is enforced independently.
	Limits []AiGatewayTokensRevokeOutputLimits `json:"limits"`
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

// MapAiGatewayTokensRevokeOutputFromJSON deserializes JSON data into a AiGatewayTokensRevokeOutput.
func MapAiGatewayTokensRevokeOutputFromJSON(data []byte) (*AiGatewayTokensRevokeOutput, error) {
	var v AiGatewayTokensRevokeOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayTokensRevokeOutputToJSON serializes a AiGatewayTokensRevokeOutput to JSON.
func MapAiGatewayTokensRevokeOutputToJSON(v *AiGatewayTokensRevokeOutput) ([]byte, error) {
	return json.Marshal(v)
}
