package conversations

import (
	"encoding/json"
	"time"
)

// ConversationsListOutputItemsCreatedByActorOrganizationActorMember represents the conversations list output items created by actor organization actor member type.
type ConversationsListOutputItemsCreatedByActorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// ConversationsListOutputItemsCreatedByActorOrganizationActorTeams - The teams the actor belongs to
type ConversationsListOutputItemsCreatedByActorOrganizationActorTeams struct {
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

// ConversationsListOutputItemsCreatedByActorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type ConversationsListOutputItemsCreatedByActorOrganizationActor struct {
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
	ImageUrl string                                                             `json:"image_url"`
	Member   *ConversationsListOutputItemsCreatedByActorOrganizationActorMember `json:"member,omitempty"`
	Teams    []ConversationsListOutputItemsCreatedByActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// ConversationsListOutputItemsCreatedByActorConsumer - Consumer details. `null` unless `type` is `consumer`.
type ConversationsListOutputItemsCreatedByActorConsumer struct {
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

// ConversationsListOutputItemsCreatedByActor - Actor behind a participant.
type ConversationsListOutputItemsCreatedByActor struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *ConversationsListOutputItemsCreatedByActorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *ConversationsListOutputItemsCreatedByActorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// ConversationsListOutputItemsAssistantDefaultModelProvider represents the conversations list output items assistant default model provider type.
type ConversationsListOutputItemsAssistantDefaultModelProvider struct {
	Object   string `json:"object"`
	Id       string `json:"id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

// ConversationsListOutputItemsAssistantDefaultModel represents the conversations list output items assistant default model type.
type ConversationsListOutputItemsAssistantDefaultModel struct {
	Object        string                                                    `json:"object"`
	Id            string                                                    `json:"id"`
	Slug          string                                                    `json:"slug"`
	Name          string                                                    `json:"name"`
	ContextWindow float64                                                   `json:"context_window"`
	Provider      ConversationsListOutputItemsAssistantDefaultModelProvider `json:"provider"`
}

// ConversationsListOutputItemsAssistantAvailableModelsProvider represents the conversations list output items assistant available models provider type.
type ConversationsListOutputItemsAssistantAvailableModelsProvider struct {
	Object   string `json:"object"`
	Id       string `json:"id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

// ConversationsListOutputItemsAssistantAvailableModels represents the conversations list output items assistant available models type.
type ConversationsListOutputItemsAssistantAvailableModels struct {
	Object        string                                                       `json:"object"`
	Id            string                                                       `json:"id"`
	Slug          string                                                       `json:"slug"`
	Name          string                                                       `json:"name"`
	ContextWindow float64                                                      `json:"context_window"`
	Provider      ConversationsListOutputItemsAssistantAvailableModelsProvider `json:"provider"`
}

// ConversationsListOutputItemsAssistant represents the conversations list output items assistant type.
type ConversationsListOutputItemsAssistant struct {
	Object          string                                                 `json:"object"`
	Id              string                                                 `json:"id"`
	Slug            string                                                 `json:"slug"`
	Name            string                                                 `json:"name"`
	OwnerType       string                                                 `json:"owner_type"`
	OrganizationId  *string                                                `json:"organization_id,omitempty"`
	DefaultModel    *ConversationsListOutputItemsAssistantDefaultModel     `json:"default_model,omitempty"`
	AvailableModels []ConversationsListOutputItemsAssistantAvailableModels `json:"available_models"`
	CreatedAt       time.Time                                              `json:"created_at"`
	UpdatedAt       time.Time                                              `json:"updated_at"`
}

// ConversationsListOutputItems represents the conversations list output items type.
type ConversationsListOutputItems struct {
	Object         string  `json:"object"`
	Id             string  `json:"id"`
	Title          *string `json:"title,omitempty"`
	AssistantId    string  `json:"assistant_id"`
	InstanceId     string  `json:"instance_id"`
	OrganizationId string  `json:"organization_id"`
	// CreatedByActor - Actor behind a participant.
	CreatedByActor ConversationsListOutputItemsCreatedByActor `json:"created_by_actor"`
	RootMessageId  string                                     `json:"root_message_id"`
	Assistant      ConversationsListOutputItemsAssistant      `json:"assistant"`
	CreatedAt      time.Time                                  `json:"created_at"`
	UpdatedAt      time.Time                                  `json:"updated_at"`
}

// ConversationsListOutputPagination represents the conversations list output pagination type.
type ConversationsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ConversationsListOutput represents the conversations list output type.
type ConversationsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                            `json:"object"`
	Items      []ConversationsListOutputItems    `json:"items"`
	Pagination ConversationsListOutputPagination `json:"pagination"`
}

// MapConversationsListOutputFromJSON deserializes JSON data into a ConversationsListOutput.
func MapConversationsListOutputFromJSON(data []byte) (*ConversationsListOutput, error) {
	var v ConversationsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConversationsListOutputToJSON serializes a ConversationsListOutput to JSON.
func MapConversationsListOutputToJSON(v *ConversationsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ConversationsListQuery represents the conversations list query type.
type ConversationsListQuery struct {
	// Limit - Page size, from 1 to 100. Defaults to 20.
	Limit *float64 `json:"limit,omitempty"`
	// After - Return the page after this point. Pass `pagination.after` if the previous response has it, otherwise the `id` of its last item.
	After *string `json:"after,omitempty"`
	// Before - Return the page before this point. Pass `pagination.before` if the previous response has it, otherwise the `id` of its first item.
	Before *string `json:"before,omitempty"`
	// Cursor - A `cur_` cursor from `pagination.after` or `pagination.before`. Alternative to `after` and `before`.
	Cursor *string `json:"cursor,omitempty"`
	// Order - Sort direction.
	Order       *string `json:"order,omitempty"`
	AssistantId *any    `json:"assistant_id,omitempty"`
}

// MapConversationsListQueryFromJSON deserializes JSON data into a ConversationsListQuery.
func MapConversationsListQueryFromJSON(data []byte) (*ConversationsListQuery, error) {
	var v ConversationsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConversationsListQueryToJSON serializes a ConversationsListQuery to JSON.
func MapConversationsListQueryToJSON(v *ConversationsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
