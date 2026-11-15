package providerissues

import (
	"encoding/json"
	"time"
)

// ProviderIssuesListOutputItemsDailyOccurrences represents the provider issues list output items daily occurrences type.
type ProviderIssuesListOutputItemsDailyOccurrences struct {
	// StartsAt - Start of the day
	StartsAt time.Time `json:"starts_at"`
	// EndsAt - End of the day
	EndsAt time.Time `json:"ends_at"`
	// Count - Occurrences during the day
	Count float64 `json:"count"`
}

// ProviderIssuesListOutputItems represents the provider issues list output items type.
type ProviderIssuesListOutputItems struct {
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
	DailyOccurrences []ProviderIssuesListOutputItemsDailyOccurrences `json:"daily_occurrences"`
	// FirstSeenAt - When the issue first occurred
	FirstSeenAt time.Time `json:"first_seen_at"`
	// LastSeenAt - When the issue most recently occurred
	LastSeenAt time.Time `json:"last_seen_at"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
}

// ProviderIssuesListOutputPagination represents the provider issues list output pagination type.
type ProviderIssuesListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// ProviderIssuesListOutput represents the provider issues list output type.
type ProviderIssuesListOutput struct {
	Items      []ProviderIssuesListOutputItems    `json:"items"`
	Pagination ProviderIssuesListOutputPagination `json:"pagination"`
}

// MapProviderIssuesListOutputFromJSON deserializes JSON data into a ProviderIssuesListOutput.
func MapProviderIssuesListOutputFromJSON(data []byte) (*ProviderIssuesListOutput, error) {
	var v ProviderIssuesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderIssuesListOutputToJSON serializes a ProviderIssuesListOutput to JSON.
func MapProviderIssuesListOutputToJSON(v *ProviderIssuesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderIssuesListQueryLastSeenAt - Filter provider issue last occurrence time by date range
type ProviderIssuesListQueryLastSeenAt struct {
	// Gt - Only include records after this timestamp for provider issue last occurrence time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for provider issue last occurrence time
	Lt *time.Time `json:"lt,omitempty"`
}

// ProviderIssuesListQueryCreatedAt - Filter provider issue creation time by date range
type ProviderIssuesListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for provider issue creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for provider issue creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// ProviderIssuesListQuery represents the provider issues list query type.
type ProviderIssuesListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Type - Filter by issue type(s)
	Type *any `json:"type,omitempty"`
	// Severity - Filter by severity
	Severity *any `json:"severity,omitempty"`
	// Id - Filter by provider issue ID(s)
	Id *any `json:"id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// LastSeenAt - Filter provider issue last occurrence time by date range
	LastSeenAt *ProviderIssuesListQueryLastSeenAt `json:"last_seen_at,omitempty"`
	// CreatedAt - Filter provider issue creation time by date range
	CreatedAt *ProviderIssuesListQueryCreatedAt `json:"created_at,omitempty"`
	// Search - Search by issue ID, code, subtype or message, or by the name of the related provider
	Search *string `json:"search,omitempty"`
}

// MapProviderIssuesListQueryFromJSON deserializes JSON data into a ProviderIssuesListQuery.
func MapProviderIssuesListQueryFromJSON(data []byte) (*ProviderIssuesListQuery, error) {
	var v ProviderIssuesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderIssuesListQueryToJSON serializes a ProviderIssuesListQuery to JSON.
func MapProviderIssuesListQueryToJSON(v *ProviderIssuesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
