package config

import (
	"encoding/json"
	"time"
)

// AiGatewayConfigGetOutputSettingsContentLogging - Storing prompts and responses. Inference logs only return content when this was enabled at the time of the request.
type AiGatewayConfigGetOutputSettingsContentLogging struct {
	// Enabled - Whether request and response bodies are stored with inference logs
	Enabled bool `json:"enabled"`
	// RetentionDays - How many days stored request and response bodies are kept
	RetentionDays float64 `json:"retention_days"`
}

// AiGatewayConfigGetOutputSettingsMetorialProvider - The Metorial provider
type AiGatewayConfigGetOutputSettingsMetorialProvider struct {
	// Enabled - Whether requests can be served with Metorial's own provider accounts, billed by Metorial. Your own provider connections are always tried first.
	Enabled bool `json:"enabled"`
}

// AiGatewayConfigGetOutputSettingsFallbacks - Provider fallbacks
type AiGatewayConfigGetOutputSettingsFallbacks struct {
	// Enabled - Whether a failed request is retried with the next provider that serves the model. When disabled, only the first candidate is tried.
	Enabled bool `json:"enabled"`
}

// AiGatewayConfigGetOutputSettingsZeroDataRetention - Zero data retention
type AiGatewayConfigGetOutputSettingsZeroDataRetention struct {
	// Required - Whether requests are only routed to providers that do not retain prompts or responses. Models without such a provider become unavailable.
	Required bool `json:"required"`
}

// AiGatewayConfigGetOutputSettings - Settings for every request made through the AI gateway
type AiGatewayConfigGetOutputSettings struct {
	// ContentLogging - Storing prompts and responses. Inference logs only return content when this was enabled at the time of the request.
	ContentLogging AiGatewayConfigGetOutputSettingsContentLogging `json:"content_logging"`
	// MetorialProvider - The Metorial provider
	MetorialProvider AiGatewayConfigGetOutputSettingsMetorialProvider `json:"metorial_provider"`
	// Fallbacks - Provider fallbacks
	Fallbacks AiGatewayConfigGetOutputSettingsFallbacks `json:"fallbacks"`
	// ZeroDataRetention - Zero data retention
	ZeroDataRetention AiGatewayConfigGetOutputSettingsZeroDataRetention `json:"zero_data_retention"`
}

// AiGatewayConfigGetOutputLimits represents the ai gateway config get output limits type.
type AiGatewayConfigGetOutputLimits struct {
	// Metric - What the limit counts. spend_usd counts what Metorial bills for requests.
	Metric string `json:"metric"`
	// Amount - The most that can be used per interval
	Amount float64 `json:"amount"`
	// Interval - The window the amount applies to
	Interval string `json:"interval"`
}

// AiGatewayConfigGetOutputPlanLimits represents the ai gateway config get output plan limits type.
type AiGatewayConfigGetOutputPlanLimits struct {
	// Metric - What the limit counts. metorial_spend_usd only counts what Metorial bills, not traffic over your own provider connections.
	Metric string `json:"metric"`
	// Amount - The most that can be used per interval
	Amount float64 `json:"amount"`
	// IntervalSeconds - The window the amount applies to
	IntervalSeconds float64 `json:"interval_seconds"`
}

// AiGatewayConfigGetOutputDisabled - Set when the AI gateway is disabled for the project, e.g. because it is not on your plan. Every request is rejected.
type AiGatewayConfigGetOutputDisabled struct {
	// Reason - Why the restriction applies
	Reason *string `json:"reason,omitempty"`
	// Since - When the restriction started
	Since time.Time `json:"since"`
}

// AiGatewayConfigGetOutputBillingRestriction - Set when the Metorial provider is unavailable because of billing, e.g. when you are out of credits. Requests over your own provider connections still work.
type AiGatewayConfigGetOutputBillingRestriction struct {
	// Reason - Why the restriction applies
	Reason *string `json:"reason,omitempty"`
	// Since - When the restriction started
	Since time.Time `json:"since"`
}

// AiGatewayConfigGetOutput represents the ai gateway config get output type.
type AiGatewayConfigGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// LogRetentionDays - How many days inference logs (metadata, usage and cost, without content) are kept. Usage timelines are built from these logs, so they reach back no further.
	LogRetentionDays float64 `json:"log_retention_days"`
	// Settings - Settings for every request made through the AI gateway
	Settings AiGatewayConfigGetOutputSettings `json:"settings"`
	// Limits - Limits that reject requests once reached. Every limit is enforced independently.
	Limits []AiGatewayConfigGetOutputLimits `json:"limits"`
	// PlanLimits - Limits from your Metorial plan. They apply to all of the project's instances combined.
	PlanLimits []AiGatewayConfigGetOutputPlanLimits `json:"plan_limits"`
	// Disabled - Set when the AI gateway is disabled for the project, e.g. because it is not on your plan. Every request is rejected.
	Disabled *AiGatewayConfigGetOutputDisabled `json:"disabled,omitempty"`
	// BillingRestriction - Set when the Metorial provider is unavailable because of billing, e.g. when you are out of credits. Requests over your own provider connections still work.
	BillingRestriction *AiGatewayConfigGetOutputBillingRestriction `json:"billing_restriction,omitempty"`
	// UpdatedAt - When the configuration was last changed
	UpdatedAt time.Time `json:"updated_at"`
}

// MapAiGatewayConfigGetOutputFromJSON deserializes JSON data into a AiGatewayConfigGetOutput.
func MapAiGatewayConfigGetOutputFromJSON(data []byte) (*AiGatewayConfigGetOutput, error) {
	var v AiGatewayConfigGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayConfigGetOutputToJSON serializes a AiGatewayConfigGetOutput to JSON.
func MapAiGatewayConfigGetOutputToJSON(v *AiGatewayConfigGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
