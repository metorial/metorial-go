package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/documents/participants"
)

// DocumentsParticipantsEndpoint provides access to inspect document participants and their linked Metorial resources.
type DocumentsParticipantsEndpoint struct {
	client *endpoint.Client
}

// NewDocumentsParticipantsEndpoint creates a new DocumentsParticipantsEndpoint.
func NewDocumentsParticipantsEndpoint(client *endpoint.Client) *DocumentsParticipantsEndpoint {
	return &DocumentsParticipantsEndpoint{client: client}
}

// DocumentsParticipantsEndpointListParams contains optional query parameters for List.
type DocumentsParticipantsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Id - Filter by document participant ID
	Id *any `json:"id,omitempty"`
	// CreatedAt - Filter Filter by creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter Filter by update time by date range
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// List returns a paginated list of participants for a specific document.
func (e *DocumentsParticipantsEndpoint) List(instanceId string, documentId string, params *DocumentsParticipantsEndpointListParams) (*participants.DocumentsParticipantsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "documents", documentId, "participants"},
		Query: query,
	}
	var result participants.DocumentsParticipantsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific document participant by its ID.
func (e *DocumentsParticipantsEndpoint) Get(instanceId string, documentId string, documentParticipantId string) (*participants.DocumentsParticipantsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "documents", documentId, "participants", documentParticipantId},
	}
	var result participants.DocumentsParticipantsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
