package skills

import (
	"encoding/json"
	"time"
)

// SkillsDuplicateOutputHierarchyCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsDuplicateOutputHierarchyCreatorOrganizationActorTeams struct {
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

// SkillsDuplicateOutputHierarchyCreatorOrganizationActor represents the skills duplicate output hierarchy creator organization actor type.
type SkillsDuplicateOutputHierarchyCreatorOrganizationActor struct {
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
	ImageUrl string                                                        `json:"image_url"`
	Teams    []SkillsDuplicateOutputHierarchyCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsDuplicateOutputHierarchyCreatorConsumer represents the skills duplicate output hierarchy creator consumer type.
type SkillsDuplicateOutputHierarchyCreatorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsDuplicateOutputHierarchyCreator represents the skills duplicate output hierarchy creator type.
type SkillsDuplicateOutputHierarchyCreator struct {
	Type              string                                                  `json:"type"`
	Name              string                                                  `json:"name"`
	ImageUrl          *string                                                 `json:"image_url,omitempty"`
	Email             *string                                                 `json:"email,omitempty"`
	OrganizationActor *SkillsDuplicateOutputHierarchyCreatorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsDuplicateOutputHierarchyCreatorConsumer          `json:"consumer,omitempty"`
}

// SkillsDuplicateOutputHierarchyForkCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsDuplicateOutputHierarchyForkCreatorOrganizationActorTeams struct {
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

// SkillsDuplicateOutputHierarchyForkCreatorOrganizationActor represents the skills duplicate output hierarchy fork creator organization actor type.
type SkillsDuplicateOutputHierarchyForkCreatorOrganizationActor struct {
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
	ImageUrl string                                                            `json:"image_url"`
	Teams    []SkillsDuplicateOutputHierarchyForkCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsDuplicateOutputHierarchyForkCreatorConsumer represents the skills duplicate output hierarchy fork creator consumer type.
type SkillsDuplicateOutputHierarchyForkCreatorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsDuplicateOutputHierarchyForkCreator represents the skills duplicate output hierarchy fork creator type.
type SkillsDuplicateOutputHierarchyForkCreator struct {
	Type              string                                                      `json:"type"`
	Name              string                                                      `json:"name"`
	ImageUrl          *string                                                     `json:"image_url,omitempty"`
	Email             *string                                                     `json:"email,omitempty"`
	OrganizationActor *SkillsDuplicateOutputHierarchyForkCreatorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsDuplicateOutputHierarchyForkCreatorConsumer          `json:"consumer,omitempty"`
}

// SkillsDuplicateOutputHierarchyForkOriginalCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsDuplicateOutputHierarchyForkOriginalCreatorOrganizationActorTeams struct {
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

// SkillsDuplicateOutputHierarchyForkOriginalCreatorOrganizationActor represents the skills duplicate output hierarchy fork original creator organization actor type.
type SkillsDuplicateOutputHierarchyForkOriginalCreatorOrganizationActor struct {
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
	ImageUrl string                                                                    `json:"image_url"`
	Teams    []SkillsDuplicateOutputHierarchyForkOriginalCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsDuplicateOutputHierarchyForkOriginalCreatorConsumer represents the skills duplicate output hierarchy fork original creator consumer type.
type SkillsDuplicateOutputHierarchyForkOriginalCreatorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsDuplicateOutputHierarchyForkOriginalCreator represents the skills duplicate output hierarchy fork original creator type.
type SkillsDuplicateOutputHierarchyForkOriginalCreator struct {
	Type              string                                                              `json:"type"`
	Name              string                                                              `json:"name"`
	ImageUrl          *string                                                             `json:"image_url,omitempty"`
	Email             *string                                                             `json:"email,omitempty"`
	OrganizationActor *SkillsDuplicateOutputHierarchyForkOriginalCreatorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsDuplicateOutputHierarchyForkOriginalCreatorConsumer          `json:"consumer,omitempty"`
}

// SkillsDuplicateOutputHierarchyFork represents the skills duplicate output hierarchy fork type.
type SkillsDuplicateOutputHierarchyFork struct {
	Id              string                                             `json:"id"`
	ParentSkillId   string                                             `json:"parent_skill_id"`
	Creator         *SkillsDuplicateOutputHierarchyForkCreator         `json:"creator,omitempty"`
	OriginalCreator *SkillsDuplicateOutputHierarchyForkOriginalCreator `json:"original_creator,omitempty"`
	CreatedAt       time.Time                                          `json:"created_at"`
}

// SkillsDuplicateOutputHierarchyEntity represents the skills duplicate output hierarchy entity type.
type SkillsDuplicateOutputHierarchyEntity struct {
	Object        string    `json:"object"`
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	Description   *string   `json:"description,omitempty"`
	ParentSkillId string    `json:"parent_skill_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// SkillsDuplicateOutputHierarchy represents the skills duplicate output hierarchy type.
type SkillsDuplicateOutputHierarchy struct {
	Object        string                                 `json:"object"`
	Type          string                                 `json:"type"`
	ParentSkillId *string                                `json:"parent_skill_id,omitempty"`
	Creator       *SkillsDuplicateOutputHierarchyCreator `json:"creator,omitempty"`
	Fork          *SkillsDuplicateOutputHierarchyFork    `json:"fork,omitempty"`
	Entity        SkillsDuplicateOutputHierarchyEntity   `json:"entity"`
}

// SkillsDuplicateOutputIntegrationsConfiguration represents the skills duplicate output integrations configuration type.
type SkillsDuplicateOutputIntegrationsConfiguration struct {
	CanAttachCustomToolFilters    bool `json:"can_attach_custom_tool_filters"`
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	CanOverrideToolFilters        bool `json:"can_override_tool_filters"`
}

// SkillsDuplicateOutputIntegrations represents the skills duplicate output integrations type.
type SkillsDuplicateOutputIntegrations struct {
	Object        string                                         `json:"object"`
	Id            string                                         `json:"id"`
	Slug          string                                         `json:"slug"`
	Name          string                                         `json:"name"`
	Description   *string                                        `json:"description,omitempty"`
	Metadata      *map[string]any                                `json:"metadata,omitempty"`
	Configuration SkillsDuplicateOutputIntegrationsConfiguration `json:"configuration"`
	CreatedAt     time.Time                                      `json:"created_at"`
	UpdatedAt     time.Time                                      `json:"updated_at"`
	ArchivedAt    *time.Time                                     `json:"archived_at,omitempty"`
}

// SkillsDuplicateOutputProviders represents the skills duplicate output providers type.
type SkillsDuplicateOutputProviders struct {
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

// SkillsDuplicateOutput represents the skills duplicate output type.
type SkillsDuplicateOutput struct {
	Object            string                              `json:"object"`
	Id                string                              `json:"id"`
	Status            string                              `json:"status"`
	Slug              string                              `json:"slug"`
	Name              string                              `json:"name"`
	Description       *string                             `json:"description,omitempty"`
	ImageUrl          string                              `json:"image_url"`
	ClientName        string                              `json:"client_name"`
	ClientDescription *string                             `json:"client_description,omitempty"`
	ClientMetadata    *map[string]any                     `json:"client_metadata,omitempty"`
	License           *string                             `json:"license,omitempty"`
	Compatibility     *string                             `json:"compatibility,omitempty"`
	Metadata          map[string]any                      `json:"metadata"`
	StoreId           string                              `json:"store_id"`
	Hierarchy         SkillsDuplicateOutputHierarchy      `json:"hierarchy"`
	Integrations      []SkillsDuplicateOutputIntegrations `json:"integrations"`
	Providers         []SkillsDuplicateOutputProviders    `json:"providers"`
	CreatedAt         time.Time                           `json:"created_at"`
	UpdatedAt         time.Time                           `json:"updated_at"`
}

// MapSkillsDuplicateOutputFromJSON deserializes JSON data into a SkillsDuplicateOutput.
func MapSkillsDuplicateOutputFromJSON(data []byte) (*SkillsDuplicateOutput, error) {
	var v SkillsDuplicateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsDuplicateOutputToJSON serializes a SkillsDuplicateOutput to JSON.
func MapSkillsDuplicateOutputToJSON(v *SkillsDuplicateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsDuplicateBody represents the skills duplicate body type.
type SkillsDuplicateBody struct {
	Name              string          `json:"name"`
	Description       *string         `json:"description,omitempty"`
	ClientName        *string         `json:"client_name,omitempty"`
	ClientDescription *string         `json:"client_description,omitempty"`
	License           *string         `json:"license,omitempty"`
	Compatibility     *string         `json:"compatibility,omitempty"`
	ClientMetadata    *map[string]any `json:"client_metadata,omitempty"`
	Metadata          *map[string]any `json:"metadata,omitempty"`
}

// MapSkillsDuplicateBodyFromJSON deserializes JSON data into a SkillsDuplicateBody.
func MapSkillsDuplicateBodyFromJSON(data []byte) (*SkillsDuplicateBody, error) {
	var v SkillsDuplicateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsDuplicateBodyToJSON serializes a SkillsDuplicateBody to JSON.
func MapSkillsDuplicateBodyToJSON(v *SkillsDuplicateBody) ([]byte, error) {
	return json.Marshal(v)
}
