package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/stores/participants"
)

// StoresParticipantsEndpoint provides access to a store participant is an actor with access to a store and its contents.
type StoresParticipantsEndpoint struct {
	client *endpoint.Client
}

// NewStoresParticipantsEndpoint creates a new StoresParticipantsEndpoint.
func NewStoresParticipantsEndpoint(client *endpoint.Client) *StoresParticipantsEndpoint {
	return &StoresParticipantsEndpoint{client: client}
}

// StoresParticipantsEndpointListParams contains optional query parameters for List.
type StoresParticipantsEndpointListParams struct {
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
}

// List returns a paginated list of participants in a store.
func (e *StoresParticipantsEndpoint) List(storeId string, params *StoresParticipantsEndpointListParams) (*participants.StoresParticipantsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"stores", storeId, "participants"},
		Query: query,
	}
	var result participants.StoresParticipantsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a participant in a store.
func (e *StoresParticipantsEndpoint) Get(storeId string, storeParticipantId string) (*participants.StoresParticipantsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"stores", storeId, "participants", storeParticipantId},
	}
	var result participants.StoresParticipantsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
