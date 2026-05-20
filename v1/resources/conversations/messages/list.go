package messages

import (
	"encoding/json"
	"time"
)

// ConversationsMessagesListOutputItemsModelProvider represents the conversations messages list output items model provider type.
type ConversationsMessagesListOutputItemsModelProvider struct {
	Object   string `json:"object"`
	Id       string `json:"id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

// ConversationsMessagesListOutputItemsModel represents the conversations messages list output items model type.
type ConversationsMessagesListOutputItemsModel struct {
	Object        string                                            `json:"object"`
	Id            string                                            `json:"id"`
	Slug          string                                            `json:"slug"`
	Name          string                                            `json:"name"`
	ContextWindow float64                                           `json:"context_window"`
	Provider      ConversationsMessagesListOutputItemsModelProvider `json:"provider"`
}

// ConversationsMessagesListOutputItemsRequestActorOrganizationActorTeams - The teams the actor belongs to
type ConversationsMessagesListOutputItemsRequestActorOrganizationActorTeams struct {
	// Id - The team ID
	Id string `json:"id"`
	// Name - The team name
	Name string `json:"name"`
	// Slug - The team slug
	Slug string `json:"slug"`
	// AssignmentId - The team assignment ID
	AssignmentId string `json:"assignment_id"`
	// CreatedAt - The team assignment creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The team assignment last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// ConversationsMessagesListOutputItemsRequestActorOrganizationActor represents the conversations messages list output items request actor organization actor type.
type ConversationsMessagesListOutputItemsRequestActorOrganizationActor struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Type - The organization member's type
	Type string `json:"type"`
	// OrganizationId - The organization member's organization ID
	OrganizationId string `json:"organization_id"`
	// Name - The organization member's name
	Name string `json:"name"`
	// Email - The organization member's email
	Email *string `json:"email,omitempty"`
	// ImageUrl - The organization member's image URL
	ImageUrl string                                                                   `json:"image_url"`
	Teams    []ConversationsMessagesListOutputItemsRequestActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// ConversationsMessagesListOutputItemsRequestActorConsumer represents the conversations messages list output items request actor consumer type.
type ConversationsMessagesListOutputItemsRequestActorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ConversationsMessagesListOutputItemsRequestActor represents the conversations messages list output items request actor type.
type ConversationsMessagesListOutputItemsRequestActor struct {
	Type              string                                                             `json:"type"`
	Name              string                                                             `json:"name"`
	ImageUrl          *string                                                            `json:"image_url,omitempty"`
	Email             *string                                                            `json:"email,omitempty"`
	OrganizationActor *ConversationsMessagesListOutputItemsRequestActorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *ConversationsMessagesListOutputItemsRequestActorConsumer          `json:"consumer,omitempty"`
}

// ConversationsMessagesListOutputItemsRequest represents the conversations messages list output items request type.
type ConversationsMessagesListOutputItemsRequest struct {
	Object    string                                            `json:"object"`
	Id        string                                            `json:"id"`
	Status    string                                            `json:"status"`
	Actor     *ConversationsMessagesListOutputItemsRequestActor `json:"actor,omitempty"`
	CreatedAt time.Time                                         `json:"created_at"`
	UpdatedAt time.Time                                         `json:"updated_at"`
}

// ConversationsMessagesListOutputItems represents the conversations messages list output items type.
type ConversationsMessagesListOutputItems struct {
	Object             string                                      `json:"object"`
	Id                 string                                      `json:"id"`
	ConversationItemId string                                      `json:"conversation_item_id"`
	Type               string                                      `json:"type"`
	AssistantId        *string                                     `json:"assistant_id,omitempty"`
	ParentMessageId    *string                                     `json:"parent_message_id,omitempty"`
	Model              *ConversationsMessagesListOutputItemsModel  `json:"model,omitempty"`
	Request            ConversationsMessagesListOutputItemsRequest `json:"request"`
	Items              []map[string]any                            `json:"items"`
	CreatedAt          time.Time                                   `json:"created_at"`
}

// ConversationsMessagesListOutputPagination represents the conversations messages list output pagination type.
type ConversationsMessagesListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// ConversationsMessagesListOutput represents the conversations messages list output type.
type ConversationsMessagesListOutput struct {
	Items      []ConversationsMessagesListOutputItems    `json:"items"`
	Pagination ConversationsMessagesListOutputPagination `json:"pagination"`
}

// MapConversationsMessagesListOutputFromJSON deserializes JSON data into a ConversationsMessagesListOutput.
func MapConversationsMessagesListOutputFromJSON(data []byte) (*ConversationsMessagesListOutput, error) {
	var v ConversationsMessagesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConversationsMessagesListOutputToJSON serializes a ConversationsMessagesListOutput to JSON.
func MapConversationsMessagesListOutputToJSON(v *ConversationsMessagesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ConversationsMessagesListQuery represents the conversations messages list query type.
type ConversationsMessagesListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
}

// MapConversationsMessagesListQueryFromJSON deserializes JSON data into a ConversationsMessagesListQuery.
func MapConversationsMessagesListQueryFromJSON(data []byte) (*ConversationsMessagesListQuery, error) {
	var v ConversationsMessagesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConversationsMessagesListQueryToJSON serializes a ConversationsMessagesListQuery to JSON.
func MapConversationsMessagesListQueryToJSON(v *ConversationsMessagesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
