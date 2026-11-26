package conversations

import (
	"encoding/json"
	"time"
)

// ConversationsGetOutputCreatedByActorOrganizationActorMember represents the conversations get output created by actor organization actor member type.
type ConversationsGetOutputCreatedByActorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// ConversationsGetOutputCreatedByActorOrganizationActorTeams - The teams the actor belongs to
type ConversationsGetOutputCreatedByActorOrganizationActorTeams struct {
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

// ConversationsGetOutputCreatedByActorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type ConversationsGetOutputCreatedByActorOrganizationActor struct {
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
	ImageUrl string                                                       `json:"image_url"`
	Member   *ConversationsGetOutputCreatedByActorOrganizationActorMember `json:"member,omitempty"`
	Teams    []ConversationsGetOutputCreatedByActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// ConversationsGetOutputCreatedByActorConsumer - Consumer details. `null` unless `type` is `consumer`.
type ConversationsGetOutputCreatedByActorConsumer struct {
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

// ConversationsGetOutputCreatedByActor - Actor behind a participant.
type ConversationsGetOutputCreatedByActor struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *ConversationsGetOutputCreatedByActorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *ConversationsGetOutputCreatedByActorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// ConversationsGetOutputAssistantDefaultModelProvider represents the conversations get output assistant default model provider type.
type ConversationsGetOutputAssistantDefaultModelProvider struct {
	Object   string `json:"object"`
	Id       string `json:"id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

// ConversationsGetOutputAssistantDefaultModel represents the conversations get output assistant default model type.
type ConversationsGetOutputAssistantDefaultModel struct {
	Object        string                                              `json:"object"`
	Id            string                                              `json:"id"`
	Slug          string                                              `json:"slug"`
	Name          string                                              `json:"name"`
	ContextWindow float64                                             `json:"context_window"`
	Provider      ConversationsGetOutputAssistantDefaultModelProvider `json:"provider"`
}

// ConversationsGetOutputAssistantAvailableModelsProvider represents the conversations get output assistant available models provider type.
type ConversationsGetOutputAssistantAvailableModelsProvider struct {
	Object   string `json:"object"`
	Id       string `json:"id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

// ConversationsGetOutputAssistantAvailableModels represents the conversations get output assistant available models type.
type ConversationsGetOutputAssistantAvailableModels struct {
	Object        string                                                 `json:"object"`
	Id            string                                                 `json:"id"`
	Slug          string                                                 `json:"slug"`
	Name          string                                                 `json:"name"`
	ContextWindow float64                                                `json:"context_window"`
	Provider      ConversationsGetOutputAssistantAvailableModelsProvider `json:"provider"`
}

// ConversationsGetOutputAssistant represents the conversations get output assistant type.
type ConversationsGetOutputAssistant struct {
	Object          string                                           `json:"object"`
	Id              string                                           `json:"id"`
	Slug            string                                           `json:"slug"`
	Name            string                                           `json:"name"`
	OwnerType       string                                           `json:"owner_type"`
	OrganizationId  *string                                          `json:"organization_id,omitempty"`
	DefaultModel    *ConversationsGetOutputAssistantDefaultModel     `json:"default_model,omitempty"`
	AvailableModels []ConversationsGetOutputAssistantAvailableModels `json:"available_models"`
	CreatedAt       time.Time                                        `json:"created_at"`
	UpdatedAt       time.Time                                        `json:"updated_at"`
}

// ConversationsGetOutput represents the conversations get output type.
type ConversationsGetOutput struct {
	Object         string  `json:"object"`
	Id             string  `json:"id"`
	Title          *string `json:"title,omitempty"`
	AssistantId    string  `json:"assistant_id"`
	InstanceId     string  `json:"instance_id"`
	OrganizationId string  `json:"organization_id"`
	// CreatedByActor - Actor behind a participant.
	CreatedByActor ConversationsGetOutputCreatedByActor `json:"created_by_actor"`
	RootMessageId  string                               `json:"root_message_id"`
	Assistant      ConversationsGetOutputAssistant      `json:"assistant"`
	CreatedAt      time.Time                            `json:"created_at"`
	UpdatedAt      time.Time                            `json:"updated_at"`
}

// MapConversationsGetOutputFromJSON deserializes JSON data into a ConversationsGetOutput.
func MapConversationsGetOutputFromJSON(data []byte) (*ConversationsGetOutput, error) {
	var v ConversationsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConversationsGetOutputToJSON serializes a ConversationsGetOutput to JSON.
func MapConversationsGetOutputToJSON(v *ConversationsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
