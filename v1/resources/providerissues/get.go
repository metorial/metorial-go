package providerissues

import (
	"encoding/json"
	"time"
)

// ProviderIssuesGetOutputDailyOccurrences represents the provider issues get output daily occurrences type.
type ProviderIssuesGetOutputDailyOccurrences struct {
	// StartsAt - Start of the day
	StartsAt time.Time `json:"starts_at"`
	// EndsAt - End of the day
	EndsAt time.Time `json:"ends_at"`
	// Count - Occurrences during the day
	Count float64 `json:"count"`
}

// ProviderIssuesGetOutput represents the provider issues get output type.
type ProviderIssuesGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique provider issue identifier
	Id string `json:"id"`
	// Type - Where the issue comes from: a failed provider or tool call, a provider authentication failure, or a ProtoGuard detection on session traffic
	Type string `json:"type"`
	// Subtype - The more specific kind within the type: the error type for provider and auth errors, the filter key for ProtoGuard warnings
	Subtype string `json:"subtype"`
	// Severity - How severe the issue is
	Severity string `json:"severity"`
	// Code - Error code, or the ProtoGuard issue type for ProtoGuard warnings
	Code string `json:"code"`
	// Message - Human-readable summary of the issue
	Message string `json:"message"`
	// ProviderId - The provider the issue occurred on, if known
	ProviderId *string `json:"provider_id,omitempty"`
	// ProtoguardFilterId - The ProtoGuard filter that raised the warning, for ProtoGuard warnings
	ProtoguardFilterId *string `json:"protoguard_filter_id,omitempty"`
	// OccurrenceCount - Total number of times the issue occurred
	OccurrenceCount float64 `json:"occurrence_count"`
	// DailyOccurrences - Occurrences per day over the last 30 days, for days with any occurrence
	DailyOccurrences []ProviderIssuesGetOutputDailyOccurrences `json:"daily_occurrences"`
	// FirstSeenAt - When the issue first occurred
	FirstSeenAt time.Time `json:"first_seen_at"`
	// LastSeenAt - When the issue most recently occurred
	LastSeenAt time.Time `json:"last_seen_at"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
}

// MapProviderIssuesGetOutputFromJSON deserializes JSON data into a ProviderIssuesGetOutput.
func MapProviderIssuesGetOutputFromJSON(data []byte) (*ProviderIssuesGetOutput, error) {
	var v ProviderIssuesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderIssuesGetOutputToJSON serializes a ProviderIssuesGetOutput to JSON.
func MapProviderIssuesGetOutputToJSON(v *ProviderIssuesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
