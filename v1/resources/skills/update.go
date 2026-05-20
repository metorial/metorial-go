package skills

import (
	"encoding/json"
	"time"
)

// SkillsUpdateOutputHierarchyCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsUpdateOutputHierarchyCreatorOrganizationActorTeams struct {
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

// SkillsUpdateOutputHierarchyCreatorOrganizationActor represents the skills update output hierarchy creator organization actor type.
type SkillsUpdateOutputHierarchyCreatorOrganizationActor struct {
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
	ImageUrl string                                                     `json:"image_url"`
	Teams    []SkillsUpdateOutputHierarchyCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsUpdateOutputHierarchyCreatorConsumer represents the skills update output hierarchy creator consumer type.
type SkillsUpdateOutputHierarchyCreatorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsUpdateOutputHierarchyCreator represents the skills update output hierarchy creator type.
type SkillsUpdateOutputHierarchyCreator struct {
	Type              string                                               `json:"type"`
	Name              string                                               `json:"name"`
	ImageUrl          *string                                              `json:"image_url,omitempty"`
	Email             *string                                              `json:"email,omitempty"`
	OrganizationActor *SkillsUpdateOutputHierarchyCreatorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsUpdateOutputHierarchyCreatorConsumer          `json:"consumer,omitempty"`
}

// SkillsUpdateOutputHierarchyForkCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsUpdateOutputHierarchyForkCreatorOrganizationActorTeams struct {
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

// SkillsUpdateOutputHierarchyForkCreatorOrganizationActor represents the skills update output hierarchy fork creator organization actor type.
type SkillsUpdateOutputHierarchyForkCreatorOrganizationActor struct {
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
	ImageUrl string                                                         `json:"image_url"`
	Teams    []SkillsUpdateOutputHierarchyForkCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsUpdateOutputHierarchyForkCreatorConsumer represents the skills update output hierarchy fork creator consumer type.
type SkillsUpdateOutputHierarchyForkCreatorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsUpdateOutputHierarchyForkCreator represents the skills update output hierarchy fork creator type.
type SkillsUpdateOutputHierarchyForkCreator struct {
	Type              string                                                   `json:"type"`
	Name              string                                                   `json:"name"`
	ImageUrl          *string                                                  `json:"image_url,omitempty"`
	Email             *string                                                  `json:"email,omitempty"`
	OrganizationActor *SkillsUpdateOutputHierarchyForkCreatorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsUpdateOutputHierarchyForkCreatorConsumer          `json:"consumer,omitempty"`
}

// SkillsUpdateOutputHierarchyForkOriginalCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsUpdateOutputHierarchyForkOriginalCreatorOrganizationActorTeams struct {
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

// SkillsUpdateOutputHierarchyForkOriginalCreatorOrganizationActor represents the skills update output hierarchy fork original creator organization actor type.
type SkillsUpdateOutputHierarchyForkOriginalCreatorOrganizationActor struct {
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
	ImageUrl string                                                                 `json:"image_url"`
	Teams    []SkillsUpdateOutputHierarchyForkOriginalCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsUpdateOutputHierarchyForkOriginalCreatorConsumer represents the skills update output hierarchy fork original creator consumer type.
type SkillsUpdateOutputHierarchyForkOriginalCreatorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsUpdateOutputHierarchyForkOriginalCreator represents the skills update output hierarchy fork original creator type.
type SkillsUpdateOutputHierarchyForkOriginalCreator struct {
	Type              string                                                           `json:"type"`
	Name              string                                                           `json:"name"`
	ImageUrl          *string                                                          `json:"image_url,omitempty"`
	Email             *string                                                          `json:"email,omitempty"`
	OrganizationActor *SkillsUpdateOutputHierarchyForkOriginalCreatorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsUpdateOutputHierarchyForkOriginalCreatorConsumer          `json:"consumer,omitempty"`
}

// SkillsUpdateOutputHierarchyFork represents the skills update output hierarchy fork type.
type SkillsUpdateOutputHierarchyFork struct {
	Id              string                                          `json:"id"`
	ParentSkillId   string                                          `json:"parent_skill_id"`
	Creator         *SkillsUpdateOutputHierarchyForkCreator         `json:"creator,omitempty"`
	OriginalCreator *SkillsUpdateOutputHierarchyForkOriginalCreator `json:"original_creator,omitempty"`
	CreatedAt       time.Time                                       `json:"created_at"`
}

// SkillsUpdateOutputHierarchyEntity represents the skills update output hierarchy entity type.
type SkillsUpdateOutputHierarchyEntity struct {
	Object        string    `json:"object"`
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	Description   *string   `json:"description,omitempty"`
	ParentSkillId string    `json:"parent_skill_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// SkillsUpdateOutputHierarchy represents the skills update output hierarchy type.
type SkillsUpdateOutputHierarchy struct {
	Object        string                              `json:"object"`
	Type          string                              `json:"type"`
	ParentSkillId *string                             `json:"parent_skill_id,omitempty"`
	Creator       *SkillsUpdateOutputHierarchyCreator `json:"creator,omitempty"`
	Fork          *SkillsUpdateOutputHierarchyFork    `json:"fork,omitempty"`
	Entity        SkillsUpdateOutputHierarchyEntity   `json:"entity"`
}

// SkillsUpdateOutputIntegrationsConfiguration represents the skills update output integrations configuration type.
type SkillsUpdateOutputIntegrationsConfiguration struct {
	CanAttachCustomToolFilters    bool `json:"can_attach_custom_tool_filters"`
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	CanOverrideToolFilters        bool `json:"can_override_tool_filters"`
}

// SkillsUpdateOutputIntegrations represents the skills update output integrations type.
type SkillsUpdateOutputIntegrations struct {
	Object        string                                      `json:"object"`
	Id            string                                      `json:"id"`
	Slug          string                                      `json:"slug"`
	Name          string                                      `json:"name"`
	Description   *string                                     `json:"description,omitempty"`
	Metadata      *map[string]any                             `json:"metadata,omitempty"`
	Configuration SkillsUpdateOutputIntegrationsConfiguration `json:"configuration"`
	CreatedAt     time.Time                                   `json:"created_at"`
	UpdatedAt     time.Time                                   `json:"updated_at"`
	ArchivedAt    *time.Time                                  `json:"archived_at,omitempty"`
}

// SkillsUpdateOutputProviders represents the skills update output providers type.
type SkillsUpdateOutputProviders struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique provider identifier
	Id string `json:"id"`
	// Name - Display name of the provider
	Name string `json:"name"`
	// Description - Brief description of the provider
	Description *string `json:"description,omitempty"`
	// Slug - URL-friendly identifier
	Slug string `json:"slug"`
	// CreatedAt - Timestamp when the provider was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the provider was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsUpdateOutput represents the skills update output type.
type SkillsUpdateOutput struct {
	Object            string                           `json:"object"`
	Id                string                           `json:"id"`
	Status            string                           `json:"status"`
	Slug              string                           `json:"slug"`
	Name              string                           `json:"name"`
	Description       *string                          `json:"description,omitempty"`
	ImageUrl          string                           `json:"image_url"`
	ClientName        string                           `json:"client_name"`
	ClientDescription *string                          `json:"client_description,omitempty"`
	ClientMetadata    *map[string]any                  `json:"client_metadata,omitempty"`
	License           *string                          `json:"license,omitempty"`
	Compatibility     *string                          `json:"compatibility,omitempty"`
	Metadata          map[string]any                   `json:"metadata"`
	StoreId           string                           `json:"store_id"`
	Hierarchy         SkillsUpdateOutputHierarchy      `json:"hierarchy"`
	Integrations      []SkillsUpdateOutputIntegrations `json:"integrations"`
	Providers         []SkillsUpdateOutputProviders    `json:"providers"`
	CreatedAt         time.Time                        `json:"created_at"`
	UpdatedAt         time.Time                        `json:"updated_at"`
}

// MapSkillsUpdateOutputFromJSON deserializes JSON data into a SkillsUpdateOutput.
func MapSkillsUpdateOutputFromJSON(data []byte) (*SkillsUpdateOutput, error) {
	var v SkillsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsUpdateOutputToJSON serializes a SkillsUpdateOutput to JSON.
func MapSkillsUpdateOutputToJSON(v *SkillsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsUpdateBody represents the skills update body type.
type SkillsUpdateBody struct {
	Name              *string         `json:"name,omitempty"`
	Description       *string         `json:"description,omitempty"`
	ClientName        *string         `json:"client_name,omitempty"`
	ClientDescription *string         `json:"client_description,omitempty"`
	License           *string         `json:"license,omitempty"`
	Compatibility     *string         `json:"compatibility,omitempty"`
	ClientMetadata    *map[string]any `json:"client_metadata,omitempty"`
	Metadata          *map[string]any `json:"metadata,omitempty"`
	ImageFileId       *string         `json:"image_file_id,omitempty"`
}

// MapSkillsUpdateBodyFromJSON deserializes JSON data into a SkillsUpdateBody.
func MapSkillsUpdateBodyFromJSON(data []byte) (*SkillsUpdateBody, error) {
	var v SkillsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsUpdateBodyToJSON serializes a SkillsUpdateBody to JSON.
func MapSkillsUpdateBodyToJSON(v *SkillsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
