package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/providerspecificationchangenotifications"
)

// ProviderSpecificationChangeNotificationsEndpoint provides access to a provider specification change notification is a record of a provider schema change. It captures the before and after specification and version.
type ProviderSpecificationChangeNotificationsEndpoint struct {
	client *endpoint.Client
}

// NewProviderSpecificationChangeNotificationsEndpoint creates a new ProviderSpecificationChangeNotificationsEndpoint.
func NewProviderSpecificationChangeNotificationsEndpoint(client *endpoint.Client) *ProviderSpecificationChangeNotificationsEndpoint {
	return &ProviderSpecificationChangeNotificationsEndpoint{client: client}
}

// ProviderSpecificationChangeNotificationsEndpointListParams contains optional query parameters for List.
type ProviderSpecificationChangeNotificationsEndpointListParams struct {
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
	// Id - Only return provider specification change notifications with one of these IDs.
	Id *any `json:"id,omitempty"`
	// Target - Only return provider specification change notifications with one of these targets.
	Target *any `json:"target,omitempty"`
	// ProviderId - Only return provider specification change notifications for one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderVersionId - Only return provider specification change notifications for one of these provider versions.
	ProviderVersionId *any `json:"provider_version_id,omitempty"`
	// ProviderSpecificationId - Only return provider specification change notifications for one of these provider specifications.
	ProviderSpecificationId *any `json:"provider_specification_id,omitempty"`
	// CreatedAt - Only return results whose provider specification change notification time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
}

// List returns a paginated list of provider specification change notifications in the instance.
func (e *ProviderSpecificationChangeNotificationsEndpoint) List(params *ProviderSpecificationChangeNotificationsEndpointListParams) (*providerspecificationchangenotifications.ProviderSpecificationChangeNotificationsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"provider-specification-change-notifications"},
		Query: query,
	}
	var result providerspecificationchangenotifications.ProviderSpecificationChangeNotificationsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a provider specification change notification.
func (e *ProviderSpecificationChangeNotificationsEndpoint) Get(notificationId string) (*providerspecificationchangenotifications.ProviderSpecificationChangeNotificationsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"provider-specification-change-notifications", notificationId},
	}
	var result providerspecificationchangenotifications.ProviderSpecificationChangeNotificationsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
