package callbacks

import (
	"encoding/json"
	"time"
)

// CallbacksListOutputItemsProvider represents the callbacks list output items provider type.
type CallbacksListOutputItemsProvider struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Slug - URL-safe name for the provider.
	Slug string `json:"slug"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// CallbacksListOutputItemsTriggerGroups - A trigger group that accepts a selection
type CallbacksListOutputItemsTriggerGroups struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Key - Trigger group key to use in `trigger_group_selections`
	Key string `json:"key"`
	// Name - Display name of the trigger group
	Name string `json:"name"`
	// Description - What the trigger group delivers
	Description *string `json:"description,omitempty"`
	// SelectionSchema - JSON Schema the selection for this trigger group must match. Fields marked `x-metorial-options` have a pick list in the dashboard.
	SelectionSchema map[string]any `json:"selection_schema"`
	// Status - `selection_required` when the trigger group needs a selection before it delivers events.
	Status string `json:"status"`
}

// CallbacksListOutputItems represents the callbacks list output items type.
type CallbacksListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Callback lifecycle status. Archived once callbacks are disabled on the integration provider.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description. `null` when none was set.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// IntegrationId - Integration this callback belongs to
	IntegrationId string `json:"integration_id"`
	// IntegrationProviderId - Integration provider this callback receives events for.
	IntegrationProviderId string                           `json:"integration_provider_id"`
	Provider              CallbacksListOutputItemsProvider `json:"provider"`
	// TriggerGroupSelections - Selections for trigger groups that need to know what to watch, keyed by trigger group key
	TriggerGroupSelections *map[string]map[string]any `json:"trigger_group_selections,omitempty"`
	// TriggerGroups - Trigger groups that accept a selection, with their schema and whether one is still required
	TriggerGroups []CallbacksListOutputItemsTriggerGroups `json:"trigger_groups"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// CallbacksListOutputPagination represents the callbacks list output pagination type.
type CallbacksListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// CallbacksListOutput represents the callbacks list output type.
type CallbacksListOutput struct {
	// Object - String representing the object’s type.
	Object     string                        `json:"object"`
	Items      []CallbacksListOutputItems    `json:"items"`
	Pagination CallbacksListOutputPagination `json:"pagination"`
}

// MapCallbacksListOutputFromJSON deserializes JSON data into a CallbacksListOutput.
func MapCallbacksListOutputFromJSON(data []byte) (*CallbacksListOutput, error) {
	var v CallbacksListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCallbacksListOutputToJSON serializes a CallbacksListOutput to JSON.
func MapCallbacksListOutputToJSON(v *CallbacksListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// CallbacksListQueryCreatedAt - Only return results whose callback creation time is in this range. Set `gt`, `lt`, or both.
type CallbacksListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// CallbacksListQueryUpdatedAt - Only return results whose callback last update time is in this range. Set `gt`, `lt`, or both.
type CallbacksListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// CallbacksListQuery represents the callbacks list query type.
type CallbacksListQuery struct {
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
	// Id - Only return callbacks with one of these IDs.
	Id *any `json:"id,omitempty"`
	// IntegrationId - Only return callbacks for one of these integrations.
	IntegrationId *any `json:"integration_id,omitempty"`
	// IntegrationProviderId - Only return callbacks for one of these integration providers.
	IntegrationProviderId *any `json:"integration_provider_id,omitempty"`
	// ProviderId - Only return callbacks for one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// Search - Search callbacks by name, description, integration, or provider.
	Search *string `json:"search,omitempty"`
	// Status - Only return callbacks with one of these statuses. `deleted` callbacks are never listed.
	Status *any `json:"status,omitempty"`
	// CreatedAt - Only return results whose callback creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *CallbacksListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose callback last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *CallbacksListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapCallbacksListQueryFromJSON deserializes JSON data into a CallbacksListQuery.
func MapCallbacksListQueryFromJSON(data []byte) (*CallbacksListQuery, error) {
	var v CallbacksListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCallbacksListQueryToJSON serializes a CallbacksListQuery to JSON.
func MapCallbacksListQueryToJSON(v *CallbacksListQuery) ([]byte, error) {
	return json.Marshal(v)
}
