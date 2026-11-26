package config

import (
	"encoding/json"
	"time"
)

// AiGatewayConfigUpdateOutputSettingsContentLogging - Storing prompts and responses. Inference logs only return content when this was enabled at the time of the request.
type AiGatewayConfigUpdateOutputSettingsContentLogging struct {
	// Enabled - Whether request and response bodies are stored with inference logs
	Enabled bool `json:"enabled"`
	// RetentionDays - How many days stored request and response bodies are kept
	RetentionDays float64 `json:"retention_days"`
}

// AiGatewayConfigUpdateOutputSettingsMetorialProvider - The Metorial provider
type AiGatewayConfigUpdateOutputSettingsMetorialProvider struct {
	// Enabled - Whether requests can be served with Metorial's own provider accounts, billed by Metorial. Your own provider connections are always tried first.
	Enabled bool `json:"enabled"`
}

// AiGatewayConfigUpdateOutputSettingsFallbacks - Provider fallbacks
type AiGatewayConfigUpdateOutputSettingsFallbacks struct {
	// Enabled - Whether a failed request is retried with the next provider that serves the model. When disabled, only the first candidate is tried.
	Enabled bool `json:"enabled"`
}

// AiGatewayConfigUpdateOutputSettingsZeroDataRetention - Zero data retention
type AiGatewayConfigUpdateOutputSettingsZeroDataRetention struct {
	// Required - Whether requests are only routed to providers that do not retain prompts or responses. Models without such a provider become unavailable.
	Required bool `json:"required"`
}

// AiGatewayConfigUpdateOutputSettings - Settings for every request made through the AI gateway
type AiGatewayConfigUpdateOutputSettings struct {
	// ContentLogging - Storing prompts and responses. Inference logs only return content when this was enabled at the time of the request.
	ContentLogging AiGatewayConfigUpdateOutputSettingsContentLogging `json:"content_logging"`
	// MetorialProvider - The Metorial provider
	MetorialProvider AiGatewayConfigUpdateOutputSettingsMetorialProvider `json:"metorial_provider"`
	// Fallbacks - Provider fallbacks
	Fallbacks AiGatewayConfigUpdateOutputSettingsFallbacks `json:"fallbacks"`
	// ZeroDataRetention - Zero data retention
	ZeroDataRetention AiGatewayConfigUpdateOutputSettingsZeroDataRetention `json:"zero_data_retention"`
}

// AiGatewayConfigUpdateOutputLimits represents the ai gateway config update output limits type.
type AiGatewayConfigUpdateOutputLimits struct {
	// Metric - What the limit counts. spend_usd counts what Metorial bills for requests.
	Metric string `json:"metric"`
	// Amount - The most that can be used per interval
	Amount float64 `json:"amount"`
	// Interval - The window the amount applies to
	Interval string `json:"interval"`
}

// AiGatewayConfigUpdateOutputPlanLimits represents the ai gateway config update output plan limits type.
type AiGatewayConfigUpdateOutputPlanLimits struct {
	// Metric - What the limit counts. metorial_spend_usd only counts what Metorial bills, not traffic over your own provider connections.
	Metric string `json:"metric"`
	// Amount - The most that can be used per interval
	Amount float64 `json:"amount"`
	// IntervalSeconds - The window the amount applies to
	IntervalSeconds float64 `json:"interval_seconds"`
}

// AiGatewayConfigUpdateOutputDisabled - Set when the AI gateway is disabled for the project, e.g. because it is not on your plan. Every request is rejected.
type AiGatewayConfigUpdateOutputDisabled struct {
	// Reason - Why the restriction applies
	Reason *string `json:"reason,omitempty"`
	// Since - When the restriction started
	Since time.Time `json:"since"`
}

// AiGatewayConfigUpdateOutputBillingRestriction - Set when the Metorial provider is unavailable because of billing, e.g. when you are out of credits. Requests over your own provider connections still work.
type AiGatewayConfigUpdateOutputBillingRestriction struct {
	// Reason - Why the restriction applies
	Reason *string `json:"reason,omitempty"`
	// Since - When the restriction started
	Since time.Time `json:"since"`
}

// AiGatewayConfigUpdateOutput represents the ai gateway config update output type.
type AiGatewayConfigUpdateOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// LogRetentionDays - How many days inference logs (metadata, usage and cost, without content) are kept. Usage timelines are built from these logs, so they reach back no further.
	LogRetentionDays float64 `json:"log_retention_days"`
	// Settings - Settings for every request made through the AI gateway
	Settings AiGatewayConfigUpdateOutputSettings `json:"settings"`
	// Limits - Limits that reject requests once reached. Every limit is enforced independently.
	Limits []AiGatewayConfigUpdateOutputLimits `json:"limits"`
	// PlanLimits - Limits from your Metorial plan. They apply to all of the project's instances combined.
	PlanLimits []AiGatewayConfigUpdateOutputPlanLimits `json:"plan_limits"`
	// Disabled - Set when the AI gateway is disabled for the project, e.g. because it is not on your plan. Every request is rejected.
	Disabled *AiGatewayConfigUpdateOutputDisabled `json:"disabled,omitempty"`
	// BillingRestriction - Set when the Metorial provider is unavailable because of billing, e.g. when you are out of credits. Requests over your own provider connections still work.
	BillingRestriction *AiGatewayConfigUpdateOutputBillingRestriction `json:"billing_restriction,omitempty"`
	// UpdatedAt - When the configuration was last changed
	UpdatedAt time.Time `json:"updated_at"`
}

// MapAiGatewayConfigUpdateOutputFromJSON deserializes JSON data into a AiGatewayConfigUpdateOutput.
func MapAiGatewayConfigUpdateOutputFromJSON(data []byte) (*AiGatewayConfigUpdateOutput, error) {
	var v AiGatewayConfigUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayConfigUpdateOutputToJSON serializes a AiGatewayConfigUpdateOutput to JSON.
func MapAiGatewayConfigUpdateOutputToJSON(v *AiGatewayConfigUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// AiGatewayConfigUpdateBodySettingsContentLogging - Storing prompts and responses. Applies to requests made after the change.
type AiGatewayConfigUpdateBodySettingsContentLogging struct {
	// Enabled - Store request and response bodies with inference logs
	Enabled *bool `json:"enabled,omitempty"`
	// RetentionDays - How many days stored bodies are kept, from 1 to 90. Defaults to 7.
	RetentionDays *float64 `json:"retention_days,omitempty"`
}

// AiGatewayConfigUpdateBodySettingsMetorialProvider - The Metorial provider
type AiGatewayConfigUpdateBodySettingsMetorialProvider struct {
	// Enabled - Allow requests to be served with Metorial's own provider accounts, billed by Metorial. Your own provider connections are always tried first.
	Enabled *bool `json:"enabled,omitempty"`
}

// AiGatewayConfigUpdateBodySettingsFallbacks - Provider fallbacks
type AiGatewayConfigUpdateBodySettingsFallbacks struct {
	// Enabled - Retry a failed request with the next provider that serves the model. When disabled, only the first candidate is tried.
	Enabled *bool `json:"enabled,omitempty"`
}

// AiGatewayConfigUpdateBodySettingsZeroDataRetention - Zero data retention
type AiGatewayConfigUpdateBodySettingsZeroDataRetention struct {
	// Required - Only route requests to providers that do not retain prompts or responses. Models without such a provider become unavailable.
	Required *bool `json:"required,omitempty"`
}

// AiGatewayConfigUpdateBodySettings - The settings to change. Fields you omit keep their current value.
type AiGatewayConfigUpdateBodySettings struct {
	// ContentLogging - Storing prompts and responses. Applies to requests made after the change.
	ContentLogging *AiGatewayConfigUpdateBodySettingsContentLogging `json:"content_logging,omitempty"`
	// MetorialProvider - The Metorial provider
	MetorialProvider *AiGatewayConfigUpdateBodySettingsMetorialProvider `json:"metorial_provider,omitempty"`
	// Fallbacks - Provider fallbacks
	Fallbacks *AiGatewayConfigUpdateBodySettingsFallbacks `json:"fallbacks,omitempty"`
	// ZeroDataRetention - Zero data retention
	ZeroDataRetention *AiGatewayConfigUpdateBodySettingsZeroDataRetention `json:"zero_data_retention,omitempty"`
}

// AiGatewayConfigUpdateBodyLimits represents the ai gateway config update body limits type.
type AiGatewayConfigUpdateBodyLimits struct {
	// Metric - What the limit counts. spend_usd counts what Metorial bills for requests.
	Metric string `json:"metric"`
	// Amount - The most that can be used per interval
	Amount float64 `json:"amount"`
	// Interval - The window the amount applies to
	Interval string `json:"interval"`
}

// AiGatewayConfigUpdateBody represents the ai gateway config update body type.
type AiGatewayConfigUpdateBody struct {
	// LogRetentionDays - How many days inference logs are kept, from 1 to 400. Usage timelines reach back no further. Applies to requests made after the change.
	LogRetentionDays *float64 `json:"log_retention_days,omitempty"`
	// Settings - The settings to change. Fields you omit keep their current value.
	Settings *AiGatewayConfigUpdateBodySettings `json:"settings,omitempty"`
	// Limits - Limits that reject requests once reached. Replaces the existing limits when updated.
	Limits *[]AiGatewayConfigUpdateBodyLimits `json:"limits,omitempty"`
}

// MapAiGatewayConfigUpdateBodyFromJSON deserializes JSON data into a AiGatewayConfigUpdateBody.
func MapAiGatewayConfigUpdateBodyFromJSON(data []byte) (*AiGatewayConfigUpdateBody, error) {
	var v AiGatewayConfigUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayConfigUpdateBodyToJSON serializes a AiGatewayConfigUpdateBody to JSON.
func MapAiGatewayConfigUpdateBodyToJSON(v *AiGatewayConfigUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
