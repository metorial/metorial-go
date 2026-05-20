package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/stores/participants"
)

// StoresParticipantsEndpoint provides access to inspect participants assigned to an instance store.
type StoresParticipantsEndpoint struct {
	client *endpoint.Client
}

// NewStoresParticipantsEndpoint creates a new StoresParticipantsEndpoint.
func NewStoresParticipantsEndpoint(client *endpoint.Client) *StoresParticipantsEndpoint {
	return &StoresParticipantsEndpoint{client: client}
}

// StoresParticipantsEndpointListParams contains optional query parameters for List.
type StoresParticipantsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
}

// List returns a paginated list of participants for a specific store.
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

// Get retrieves a specific participant within a store.
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
