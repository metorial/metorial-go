package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/documents/participants"
)

// DocumentsParticipantsEndpoint provides access to a document participant is an actor with access to a document.
type DocumentsParticipantsEndpoint struct {
	client *endpoint.Client
}

// NewDocumentsParticipantsEndpoint creates a new DocumentsParticipantsEndpoint.
func NewDocumentsParticipantsEndpoint(client *endpoint.Client) *DocumentsParticipantsEndpoint {
	return &DocumentsParticipantsEndpoint{client: client}
}

// DocumentsParticipantsEndpointListParams contains optional query parameters for List.
type DocumentsParticipantsEndpointListParams struct {
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
	// Id - Only return participants with one of these IDs.
	Id *any `json:"id,omitempty"`
	// CreatedAt - Only return results whose participant creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose participant last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// List returns a paginated list of participants in a document.
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

// Get retrieves a participant in a document.
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
