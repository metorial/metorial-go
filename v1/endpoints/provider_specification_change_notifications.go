package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/providerspecificationchangenotifications"
)

// ProviderSpecificationChangeNotificationsEndpoint provides access to provider specification change notifications describe provider schema changes.
type ProviderSpecificationChangeNotificationsEndpoint struct {
	client *endpoint.Client
}

// NewProviderSpecificationChangeNotificationsEndpoint creates a new ProviderSpecificationChangeNotificationsEndpoint.
func NewProviderSpecificationChangeNotificationsEndpoint(client *endpoint.Client) *ProviderSpecificationChangeNotificationsEndpoint {
	return &ProviderSpecificationChangeNotificationsEndpoint{client: client}
}

// ProviderSpecificationChangeNotificationsEndpointListParams contains optional query parameters for List.
type ProviderSpecificationChangeNotificationsEndpointListParams struct {
	Limit                   *float64 `json:"limit,omitempty"`
	After                   *string  `json:"after,omitempty"`
	Before                  *string  `json:"before,omitempty"`
	Cursor                  *string  `json:"cursor,omitempty"`
	Order                   *string  `json:"order,omitempty"`
	Id                      *any     `json:"id,omitempty"`
	Target                  *any     `json:"target,omitempty"`
	ProviderId              *any     `json:"provider_id,omitempty"`
	ProviderVersionId       *any     `json:"provider_version_id,omitempty"`
	ProviderSpecificationId *any     `json:"provider_specification_id,omitempty"`
	// CreatedAt - Filter provider specification change notification time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
}

// List returns a paginated list of provider specification change notifications for this instance.
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

// Get retrieves a provider specification change notification by ID.
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
