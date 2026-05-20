package conversations

import (
	"encoding/json"
	"time"
)

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

// ConversationsListOutputItemsCreatedByActorOrganizationActor represents the conversations list output items created by actor organization actor type.
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
	Teams    []ConversationsListOutputItemsCreatedByActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// ConversationsListOutputItemsCreatedByActorConsumer represents the conversations list output items created by actor consumer type.
type ConversationsListOutputItemsCreatedByActorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ConversationsListOutputItemsCreatedByActor represents the conversations list output items created by actor type.
type ConversationsListOutputItemsCreatedByActor struct {
	Type              string                                                       `json:"type"`
	Name              string                                                       `json:"name"`
	ImageUrl          *string                                                      `json:"image_url,omitempty"`
	Email             *string                                                      `json:"email,omitempty"`
	OrganizationActor *ConversationsListOutputItemsCreatedByActorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *ConversationsListOutputItemsCreatedByActorConsumer          `json:"consumer,omitempty"`
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
	Object         string                                     `json:"object"`
	Id             string                                     `json:"id"`
	Title          *string                                    `json:"title,omitempty"`
	AssistantId    string                                     `json:"assistant_id"`
	InstanceId     string                                     `json:"instance_id"`
	OrganizationId string                                     `json:"organization_id"`
	CreatedByActor ConversationsListOutputItemsCreatedByActor `json:"created_by_actor"`
	RootMessageId  string                                     `json:"root_message_id"`
	Assistant      ConversationsListOutputItemsAssistant      `json:"assistant"`
	CreatedAt      time.Time                                  `json:"created_at"`
	UpdatedAt      time.Time                                  `json:"updated_at"`
}

// ConversationsListOutputPagination represents the conversations list output pagination type.
type ConversationsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// ConversationsListOutput represents the conversations list output type.
type ConversationsListOutput struct {
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
	Limit       *float64 `json:"limit,omitempty"`
	After       *string  `json:"after,omitempty"`
	Before      *string  `json:"before,omitempty"`
	Cursor      *string  `json:"cursor,omitempty"`
	Order       *string  `json:"order,omitempty"`
	AssistantId *any     `json:"assistant_id,omitempty"`
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
