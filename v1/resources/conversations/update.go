package conversations

import (
	"encoding/json"
	"time"
)

// ConversationsUpdateOutputCreatedByActorOrganizationActorMember represents the conversations update output created by actor organization actor member type.
type ConversationsUpdateOutputCreatedByActorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// ConversationsUpdateOutputCreatedByActorOrganizationActorTeams - The teams the actor belongs to
type ConversationsUpdateOutputCreatedByActorOrganizationActorTeams struct {
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

// ConversationsUpdateOutputCreatedByActorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type ConversationsUpdateOutputCreatedByActorOrganizationActor struct {
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
	ImageUrl string                                                          `json:"image_url"`
	Member   *ConversationsUpdateOutputCreatedByActorOrganizationActorMember `json:"member,omitempty"`
	Teams    []ConversationsUpdateOutputCreatedByActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// ConversationsUpdateOutputCreatedByActorConsumer - Consumer details. `null` unless `type` is `consumer`.
type ConversationsUpdateOutputCreatedByActorConsumer struct {
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

// ConversationsUpdateOutputCreatedByActor - Actor behind a participant.
type ConversationsUpdateOutputCreatedByActor struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *ConversationsUpdateOutputCreatedByActorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *ConversationsUpdateOutputCreatedByActorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// ConversationsUpdateOutputAssistantDefaultModelProvider represents the conversations update output assistant default model provider type.
type ConversationsUpdateOutputAssistantDefaultModelProvider struct {
	Object   string `json:"object"`
	Id       string `json:"id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

// ConversationsUpdateOutputAssistantDefaultModel represents the conversations update output assistant default model type.
type ConversationsUpdateOutputAssistantDefaultModel struct {
	Object        string                                                 `json:"object"`
	Id            string                                                 `json:"id"`
	Slug          string                                                 `json:"slug"`
	Name          string                                                 `json:"name"`
	ContextWindow float64                                                `json:"context_window"`
	Provider      ConversationsUpdateOutputAssistantDefaultModelProvider `json:"provider"`
}

// ConversationsUpdateOutputAssistantAvailableModelsProvider represents the conversations update output assistant available models provider type.
type ConversationsUpdateOutputAssistantAvailableModelsProvider struct {
	Object   string `json:"object"`
	Id       string `json:"id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

// ConversationsUpdateOutputAssistantAvailableModels represents the conversations update output assistant available models type.
type ConversationsUpdateOutputAssistantAvailableModels struct {
	Object        string                                                    `json:"object"`
	Id            string                                                    `json:"id"`
	Slug          string                                                    `json:"slug"`
	Name          string                                                    `json:"name"`
	ContextWindow float64                                                   `json:"context_window"`
	Provider      ConversationsUpdateOutputAssistantAvailableModelsProvider `json:"provider"`
}

// ConversationsUpdateOutputAssistant represents the conversations update output assistant type.
type ConversationsUpdateOutputAssistant struct {
	Object          string                                              `json:"object"`
	Id              string                                              `json:"id"`
	Slug            string                                              `json:"slug"`
	Name            string                                              `json:"name"`
	OwnerType       string                                              `json:"owner_type"`
	OrganizationId  *string                                             `json:"organization_id,omitempty"`
	DefaultModel    *ConversationsUpdateOutputAssistantDefaultModel     `json:"default_model,omitempty"`
	AvailableModels []ConversationsUpdateOutputAssistantAvailableModels `json:"available_models"`
	CreatedAt       time.Time                                           `json:"created_at"`
	UpdatedAt       time.Time                                           `json:"updated_at"`
}

// ConversationsUpdateOutput represents the conversations update output type.
type ConversationsUpdateOutput struct {
	Object         string  `json:"object"`
	Id             string  `json:"id"`
	Title          *string `json:"title,omitempty"`
	AssistantId    string  `json:"assistant_id"`
	InstanceId     string  `json:"instance_id"`
	OrganizationId string  `json:"organization_id"`
	// CreatedByActor - Actor behind a participant.
	CreatedByActor ConversationsUpdateOutputCreatedByActor `json:"created_by_actor"`
	RootMessageId  string                                  `json:"root_message_id"`
	Assistant      ConversationsUpdateOutputAssistant      `json:"assistant"`
	CreatedAt      time.Time                               `json:"created_at"`
	UpdatedAt      time.Time                               `json:"updated_at"`
}

// MapConversationsUpdateOutputFromJSON deserializes JSON data into a ConversationsUpdateOutput.
func MapConversationsUpdateOutputFromJSON(data []byte) (*ConversationsUpdateOutput, error) {
	var v ConversationsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConversationsUpdateOutputToJSON serializes a ConversationsUpdateOutput to JSON.
func MapConversationsUpdateOutputToJSON(v *ConversationsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ConversationsUpdateBody represents the conversations update body type.
type ConversationsUpdateBody struct {
	Title *string `json:"title,omitempty"`
}

// MapConversationsUpdateBodyFromJSON deserializes JSON data into a ConversationsUpdateBody.
func MapConversationsUpdateBodyFromJSON(data []byte) (*ConversationsUpdateBody, error) {
	var v ConversationsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConversationsUpdateBodyToJSON serializes a ConversationsUpdateBody to JSON.
func MapConversationsUpdateBodyToJSON(v *ConversationsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
