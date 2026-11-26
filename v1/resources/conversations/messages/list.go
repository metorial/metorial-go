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

// ConversationsMessagesListOutputItemsRequestActorOrganizationActorMember represents the conversations messages list output items request actor organization actor member type.
type ConversationsMessagesListOutputItemsRequestActorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
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

// ConversationsMessagesListOutputItemsRequestActorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
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
	Member   *ConversationsMessagesListOutputItemsRequestActorOrganizationActorMember `json:"member,omitempty"`
	Teams    []ConversationsMessagesListOutputItemsRequestActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// ConversationsMessagesListOutputItemsRequestActorConsumer - Consumer details. `null` unless `type` is `consumer`.
type ConversationsMessagesListOutputItemsRequestActorConsumer struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Email - Email address.
	Email string `json:"email"`
	// ImageUrl - Avatar image URL.
	ImageUrl string `json:"image_url"`
	// UserId - User this consumer is linked to. `null` when unlinked.
	UserId *string `json:"user_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ConversationsMessagesListOutputItemsRequestActor - Actor behind a participant.
type ConversationsMessagesListOutputItemsRequestActor struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *ConversationsMessagesListOutputItemsRequestActorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *ConversationsMessagesListOutputItemsRequestActorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// ConversationsMessagesListOutputItemsRequestContext represents the conversations messages list output items request context type.
type ConversationsMessagesListOutputItemsRequestContext struct {
	// Mode - The mode the assistant ran in after permission checks
	Mode string `json:"mode"`
	// RequestedMode - The mode requested by the client
	RequestedMode string `json:"requested_mode"`
	// CanWrite - Whether the assistant could change the underlying resource
	CanWrite bool `json:"can_write"`
}

// ConversationsMessagesListOutputItemsRequest represents the conversations messages list output items request type.
type ConversationsMessagesListOutputItemsRequest struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Status string `json:"status"`
	// Actor - Actor behind a participant.
	Actor     *ConversationsMessagesListOutputItemsRequestActor   `json:"actor,omitempty"`
	Context   *ConversationsMessagesListOutputItemsRequestContext `json:"context,omitempty"`
	CreatedAt time.Time                                           `json:"created_at"`
	UpdatedAt time.Time                                           `json:"updated_at"`
}

// ConversationsMessagesListOutputItems represents the conversations messages list output items type.
type ConversationsMessagesListOutputItems struct {
	Object             string                                      `json:"object"`
	Id                 string                                      `json:"id"`
	ConversationItemId string                                      `json:"conversation_item_id"`
	Type               string                                      `json:"type"`
	Status             string                                      `json:"status"`
	AssistantId        *string                                     `json:"assistant_id,omitempty"`
	ParentMessageId    *string                                     `json:"parent_message_id,omitempty"`
	Model              *ConversationsMessagesListOutputItemsModel  `json:"model,omitempty"`
	Request            ConversationsMessagesListOutputItemsRequest `json:"request"`
	Items              []map[string]any                            `json:"items"`
	CreatedAt          time.Time                                   `json:"created_at"`
}

// ConversationsMessagesListOutputPagination represents the conversations messages list output pagination type.
type ConversationsMessagesListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ConversationsMessagesListOutput represents the conversations messages list output type.
type ConversationsMessagesListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                    `json:"object"`
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
