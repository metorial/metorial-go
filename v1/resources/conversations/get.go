package conversations

import (
	"encoding/json"
	"time"
)

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

// ConversationsGetOutputCreatedByActorOrganizationActor represents the conversations get output created by actor organization actor type.
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
	Teams    []ConversationsGetOutputCreatedByActorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// ConversationsGetOutputCreatedByActorConsumer represents the conversations get output created by actor consumer type.
type ConversationsGetOutputCreatedByActorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ConversationsGetOutputCreatedByActor represents the conversations get output created by actor type.
type ConversationsGetOutputCreatedByActor struct {
	Type              string                                                 `json:"type"`
	Name              string                                                 `json:"name"`
	ImageUrl          *string                                                `json:"image_url,omitempty"`
	Email             *string                                                `json:"email,omitempty"`
	OrganizationActor *ConversationsGetOutputCreatedByActorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *ConversationsGetOutputCreatedByActorConsumer          `json:"consumer,omitempty"`
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
	Object         string                               `json:"object"`
	Id             string                               `json:"id"`
	Title          *string                              `json:"title,omitempty"`
	AssistantId    string                               `json:"assistant_id"`
	InstanceId     string                               `json:"instance_id"`
	OrganizationId string                               `json:"organization_id"`
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
