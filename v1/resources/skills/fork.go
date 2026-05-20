package skills

import (
	"encoding/json"
	"time"
)

// SkillsForkOutputHierarchyCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsForkOutputHierarchyCreatorOrganizationActorTeams struct {
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

// SkillsForkOutputHierarchyCreatorOrganizationActor represents the skills fork output hierarchy creator organization actor type.
type SkillsForkOutputHierarchyCreatorOrganizationActor struct {
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
	ImageUrl string                                                   `json:"image_url"`
	Teams    []SkillsForkOutputHierarchyCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsForkOutputHierarchyCreatorConsumer represents the skills fork output hierarchy creator consumer type.
type SkillsForkOutputHierarchyCreatorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsForkOutputHierarchyCreator represents the skills fork output hierarchy creator type.
type SkillsForkOutputHierarchyCreator struct {
	Type              string                                             `json:"type"`
	Name              string                                             `json:"name"`
	ImageUrl          *string                                            `json:"image_url,omitempty"`
	Email             *string                                            `json:"email,omitempty"`
	OrganizationActor *SkillsForkOutputHierarchyCreatorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsForkOutputHierarchyCreatorConsumer          `json:"consumer,omitempty"`
}

// SkillsForkOutputHierarchyForkCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsForkOutputHierarchyForkCreatorOrganizationActorTeams struct {
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

// SkillsForkOutputHierarchyForkCreatorOrganizationActor represents the skills fork output hierarchy fork creator organization actor type.
type SkillsForkOutputHierarchyForkCreatorOrganizationActor struct {
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
	Teams    []SkillsForkOutputHierarchyForkCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsForkOutputHierarchyForkCreatorConsumer represents the skills fork output hierarchy fork creator consumer type.
type SkillsForkOutputHierarchyForkCreatorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsForkOutputHierarchyForkCreator represents the skills fork output hierarchy fork creator type.
type SkillsForkOutputHierarchyForkCreator struct {
	Type              string                                                 `json:"type"`
	Name              string                                                 `json:"name"`
	ImageUrl          *string                                                `json:"image_url,omitempty"`
	Email             *string                                                `json:"email,omitempty"`
	OrganizationActor *SkillsForkOutputHierarchyForkCreatorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsForkOutputHierarchyForkCreatorConsumer          `json:"consumer,omitempty"`
}

// SkillsForkOutputHierarchyForkOriginalCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsForkOutputHierarchyForkOriginalCreatorOrganizationActorTeams struct {
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

// SkillsForkOutputHierarchyForkOriginalCreatorOrganizationActor represents the skills fork output hierarchy fork original creator organization actor type.
type SkillsForkOutputHierarchyForkOriginalCreatorOrganizationActor struct {
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
	ImageUrl string                                                               `json:"image_url"`
	Teams    []SkillsForkOutputHierarchyForkOriginalCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsForkOutputHierarchyForkOriginalCreatorConsumer represents the skills fork output hierarchy fork original creator consumer type.
type SkillsForkOutputHierarchyForkOriginalCreatorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsForkOutputHierarchyForkOriginalCreator represents the skills fork output hierarchy fork original creator type.
type SkillsForkOutputHierarchyForkOriginalCreator struct {
	Type              string                                                         `json:"type"`
	Name              string                                                         `json:"name"`
	ImageUrl          *string                                                        `json:"image_url,omitempty"`
	Email             *string                                                        `json:"email,omitempty"`
	OrganizationActor *SkillsForkOutputHierarchyForkOriginalCreatorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsForkOutputHierarchyForkOriginalCreatorConsumer          `json:"consumer,omitempty"`
}

// SkillsForkOutputHierarchyFork represents the skills fork output hierarchy fork type.
type SkillsForkOutputHierarchyFork struct {
	Id              string                                        `json:"id"`
	ParentSkillId   string                                        `json:"parent_skill_id"`
	Creator         *SkillsForkOutputHierarchyForkCreator         `json:"creator,omitempty"`
	OriginalCreator *SkillsForkOutputHierarchyForkOriginalCreator `json:"original_creator,omitempty"`
	CreatedAt       time.Time                                     `json:"created_at"`
}

// SkillsForkOutputHierarchyEntity represents the skills fork output hierarchy entity type.
type SkillsForkOutputHierarchyEntity struct {
	Object        string    `json:"object"`
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	Description   *string   `json:"description,omitempty"`
	ParentSkillId string    `json:"parent_skill_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// SkillsForkOutputHierarchy represents the skills fork output hierarchy type.
type SkillsForkOutputHierarchy struct {
	Object        string                            `json:"object"`
	Type          string                            `json:"type"`
	ParentSkillId *string                           `json:"parent_skill_id,omitempty"`
	Creator       *SkillsForkOutputHierarchyCreator `json:"creator,omitempty"`
	Fork          *SkillsForkOutputHierarchyFork    `json:"fork,omitempty"`
	Entity        SkillsForkOutputHierarchyEntity   `json:"entity"`
}

// SkillsForkOutputIntegrationsConfiguration represents the skills fork output integrations configuration type.
type SkillsForkOutputIntegrationsConfiguration struct {
	CanAttachCustomToolFilters    bool `json:"can_attach_custom_tool_filters"`
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	CanOverrideToolFilters        bool `json:"can_override_tool_filters"`
}

// SkillsForkOutputIntegrations represents the skills fork output integrations type.
type SkillsForkOutputIntegrations struct {
	Object        string                                    `json:"object"`
	Id            string                                    `json:"id"`
	Slug          string                                    `json:"slug"`
	Name          string                                    `json:"name"`
	Description   *string                                   `json:"description,omitempty"`
	Metadata      *map[string]any                           `json:"metadata,omitempty"`
	Configuration SkillsForkOutputIntegrationsConfiguration `json:"configuration"`
	CreatedAt     time.Time                                 `json:"created_at"`
	UpdatedAt     time.Time                                 `json:"updated_at"`
	ArchivedAt    *time.Time                                `json:"archived_at,omitempty"`
}

// SkillsForkOutputProviders represents the skills fork output providers type.
type SkillsForkOutputProviders struct {
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

// SkillsForkOutput represents the skills fork output type.
type SkillsForkOutput struct {
	Object            string                         `json:"object"`
	Id                string                         `json:"id"`
	Status            string                         `json:"status"`
	Slug              string                         `json:"slug"`
	Name              string                         `json:"name"`
	Description       *string                        `json:"description,omitempty"`
	ImageUrl          string                         `json:"image_url"`
	ClientName        string                         `json:"client_name"`
	ClientDescription *string                        `json:"client_description,omitempty"`
	ClientMetadata    *map[string]any                `json:"client_metadata,omitempty"`
	License           *string                        `json:"license,omitempty"`
	Compatibility     *string                        `json:"compatibility,omitempty"`
	Metadata          map[string]any                 `json:"metadata"`
	StoreId           string                         `json:"store_id"`
	Hierarchy         SkillsForkOutputHierarchy      `json:"hierarchy"`
	Integrations      []SkillsForkOutputIntegrations `json:"integrations"`
	Providers         []SkillsForkOutputProviders    `json:"providers"`
	CreatedAt         time.Time                      `json:"created_at"`
	UpdatedAt         time.Time                      `json:"updated_at"`
}

// MapSkillsForkOutputFromJSON deserializes JSON data into a SkillsForkOutput.
func MapSkillsForkOutputFromJSON(data []byte) (*SkillsForkOutput, error) {
	var v SkillsForkOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsForkOutputToJSON serializes a SkillsForkOutput to JSON.
func MapSkillsForkOutputToJSON(v *SkillsForkOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsForkBody represents the skills fork body type.
type SkillsForkBody struct {
	Name              string          `json:"name"`
	Description       *string         `json:"description,omitempty"`
	ClientName        *string         `json:"client_name,omitempty"`
	ClientDescription *string         `json:"client_description,omitempty"`
	License           *string         `json:"license,omitempty"`
	Compatibility     *string         `json:"compatibility,omitempty"`
	ClientMetadata    *map[string]any `json:"client_metadata,omitempty"`
	Metadata          *map[string]any `json:"metadata,omitempty"`
	ImageFileId       *string         `json:"image_file_id,omitempty"`
}

// MapSkillsForkBodyFromJSON deserializes JSON data into a SkillsForkBody.
func MapSkillsForkBodyFromJSON(data []byte) (*SkillsForkBody, error) {
	var v SkillsForkBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsForkBodyToJSON serializes a SkillsForkBody to JSON.
func MapSkillsForkBodyToJSON(v *SkillsForkBody) ([]byte, error) {
	return json.Marshal(v)
}
