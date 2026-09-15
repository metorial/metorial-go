package skills

import (
	"encoding/json"
	"time"
)

// SkillsShareOutputHierarchyCreatorOrganizationActorMember represents the skills share output hierarchy creator organization actor member type.
type SkillsShareOutputHierarchyCreatorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsShareOutputHierarchyCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsShareOutputHierarchyCreatorOrganizationActorTeams struct {
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

// SkillsShareOutputHierarchyCreatorOrganizationActor represents the skills share output hierarchy creator organization actor type.
type SkillsShareOutputHierarchyCreatorOrganizationActor struct {
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
	ImageUrl string                                                    `json:"image_url"`
	Member   *SkillsShareOutputHierarchyCreatorOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsShareOutputHierarchyCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsShareOutputHierarchyCreatorConsumer represents the skills share output hierarchy creator consumer type.
type SkillsShareOutputHierarchyCreatorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsShareOutputHierarchyCreator represents the skills share output hierarchy creator type.
type SkillsShareOutputHierarchyCreator struct {
	Type              string                                              `json:"type"`
	Name              string                                              `json:"name"`
	ImageUrl          *string                                             `json:"image_url,omitempty"`
	Email             *string                                             `json:"email,omitempty"`
	OrganizationActor *SkillsShareOutputHierarchyCreatorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsShareOutputHierarchyCreatorConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                     `json:"consumer_profile,omitempty"`
}

// SkillsShareOutputHierarchyForkCreatorOrganizationActorMember represents the skills share output hierarchy fork creator organization actor member type.
type SkillsShareOutputHierarchyForkCreatorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsShareOutputHierarchyForkCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsShareOutputHierarchyForkCreatorOrganizationActorTeams struct {
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

// SkillsShareOutputHierarchyForkCreatorOrganizationActor represents the skills share output hierarchy fork creator organization actor type.
type SkillsShareOutputHierarchyForkCreatorOrganizationActor struct {
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
	Member   *SkillsShareOutputHierarchyForkCreatorOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsShareOutputHierarchyForkCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsShareOutputHierarchyForkCreatorConsumer represents the skills share output hierarchy fork creator consumer type.
type SkillsShareOutputHierarchyForkCreatorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsShareOutputHierarchyForkCreator represents the skills share output hierarchy fork creator type.
type SkillsShareOutputHierarchyForkCreator struct {
	Type              string                                                  `json:"type"`
	Name              string                                                  `json:"name"`
	ImageUrl          *string                                                 `json:"image_url,omitempty"`
	Email             *string                                                 `json:"email,omitempty"`
	OrganizationActor *SkillsShareOutputHierarchyForkCreatorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsShareOutputHierarchyForkCreatorConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                         `json:"consumer_profile,omitempty"`
}

// SkillsShareOutputHierarchyForkOriginalCreatorOrganizationActorMember represents the skills share output hierarchy fork original creator organization actor member type.
type SkillsShareOutputHierarchyForkOriginalCreatorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsShareOutputHierarchyForkOriginalCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsShareOutputHierarchyForkOriginalCreatorOrganizationActorTeams struct {
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

// SkillsShareOutputHierarchyForkOriginalCreatorOrganizationActor represents the skills share output hierarchy fork original creator organization actor type.
type SkillsShareOutputHierarchyForkOriginalCreatorOrganizationActor struct {
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
	ImageUrl string                                                                `json:"image_url"`
	Member   *SkillsShareOutputHierarchyForkOriginalCreatorOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsShareOutputHierarchyForkOriginalCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsShareOutputHierarchyForkOriginalCreatorConsumer represents the skills share output hierarchy fork original creator consumer type.
type SkillsShareOutputHierarchyForkOriginalCreatorConsumer struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageUrl  string    `json:"image_url"`
	UserId    *string   `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsShareOutputHierarchyForkOriginalCreator represents the skills share output hierarchy fork original creator type.
type SkillsShareOutputHierarchyForkOriginalCreator struct {
	Type              string                                                          `json:"type"`
	Name              string                                                          `json:"name"`
	ImageUrl          *string                                                         `json:"image_url,omitempty"`
	Email             *string                                                         `json:"email,omitempty"`
	OrganizationActor *SkillsShareOutputHierarchyForkOriginalCreatorOrganizationActor `json:"organization_actor,omitempty"`
	Consumer          *SkillsShareOutputHierarchyForkOriginalCreatorConsumer          `json:"consumer,omitempty"`
	ConsumerProfile   *map[string]any                                                 `json:"consumer_profile,omitempty"`
}

// SkillsShareOutputHierarchyFork represents the skills share output hierarchy fork type.
type SkillsShareOutputHierarchyFork struct {
	Id              string                                         `json:"id"`
	ParentSkillId   string                                         `json:"parent_skill_id"`
	Creator         *SkillsShareOutputHierarchyForkCreator         `json:"creator,omitempty"`
	OriginalCreator *SkillsShareOutputHierarchyForkOriginalCreator `json:"original_creator,omitempty"`
	CreatedAt       time.Time                                      `json:"created_at"`
}

// SkillsShareOutputHierarchyEntity represents the skills share output hierarchy entity type.
type SkillsShareOutputHierarchyEntity struct {
	Object        string    `json:"object"`
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug"`
	Description   *string   `json:"description,omitempty"`
	ParentSkillId string    `json:"parent_skill_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// SkillsShareOutputHierarchy represents the skills share output hierarchy type.
type SkillsShareOutputHierarchy struct {
	Object        string                             `json:"object"`
	Type          string                             `json:"type"`
	ParentSkillId *string                            `json:"parent_skill_id,omitempty"`
	Creator       *SkillsShareOutputHierarchyCreator `json:"creator,omitempty"`
	Fork          *SkillsShareOutputHierarchyFork    `json:"fork,omitempty"`
	Entity        SkillsShareOutputHierarchyEntity   `json:"entity"`
}

// SkillsShareOutputIntegrationsConfiguration represents the skills share output integrations configuration type.
type SkillsShareOutputIntegrationsConfiguration struct {
	CanAttachCustomToolFilters    bool  `json:"can_attach_custom_tool_filters"`
	CanAttachCustomProviderConfig bool  `json:"can_attach_custom_provider_config"`
	CanOverrideToolFilters        bool  `json:"can_override_tool_filters"`
	UseIntegrationNameInToolNames *bool `json:"use_integration_name_in_tool_names,omitempty"`
}

// SkillsShareOutputIntegrations represents the skills share output integrations type.
type SkillsShareOutputIntegrations struct {
	Object        string                                     `json:"object"`
	Id            string                                     `json:"id"`
	Slug          string                                     `json:"slug"`
	Name          string                                     `json:"name"`
	Description   *string                                    `json:"description,omitempty"`
	Metadata      *map[string]any                            `json:"metadata,omitempty"`
	Configuration SkillsShareOutputIntegrationsConfiguration `json:"configuration"`
	CreatedAt     time.Time                                  `json:"created_at"`
	UpdatedAt     time.Time                                  `json:"updated_at"`
	ArchivedAt    *time.Time                                 `json:"archived_at,omitempty"`
}

// SkillsShareOutputProviders represents the skills share output providers type.
type SkillsShareOutputProviders struct {
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

// SkillsShareOutput represents the skills share output type.
type SkillsShareOutput struct {
	Object            string                          `json:"object"`
	Id                string                          `json:"id"`
	Status            string                          `json:"status"`
	Slug              string                          `json:"slug"`
	Name              string                          `json:"name"`
	Description       *string                         `json:"description,omitempty"`
	ImageUrl          string                          `json:"image_url"`
	ClientName        string                          `json:"client_name"`
	ClientDescription *string                         `json:"client_description,omitempty"`
	ClientMetadata    *map[string]any                 `json:"client_metadata,omitempty"`
	License           *string                         `json:"license,omitempty"`
	Compatibility     *string                         `json:"compatibility,omitempty"`
	Metadata          map[string]any                  `json:"metadata"`
	StoreId           string                          `json:"store_id"`
	Hierarchy         SkillsShareOutputHierarchy      `json:"hierarchy"`
	Integrations      []SkillsShareOutputIntegrations `json:"integrations"`
	Providers         []SkillsShareOutputProviders    `json:"providers"`
	CreatedAt         time.Time                       `json:"created_at"`
	UpdatedAt         time.Time                       `json:"updated_at"`
}

// MapSkillsShareOutputFromJSON deserializes JSON data into a SkillsShareOutput.
func MapSkillsShareOutputFromJSON(data []byte) (*SkillsShareOutput, error) {
	var v SkillsShareOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsShareOutputToJSON serializes a SkillsShareOutput to JSON.
func MapSkillsShareOutputToJSON(v *SkillsShareOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsShareBody represents the skills share body type.
type SkillsShareBody struct {
	ConsumerProfileIds    *[]string `json:"consumer_profile_ids,omitempty"`
	OrganizationMemberIds *[]string `json:"organization_member_ids,omitempty"`
	Permission            string    `json:"permission"`
}

// MapSkillsShareBodyFromJSON deserializes JSON data into a SkillsShareBody.
func MapSkillsShareBodyFromJSON(data []byte) (*SkillsShareBody, error) {
	var v SkillsShareBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsShareBodyToJSON serializes a SkillsShareBody to JSON.
func MapSkillsShareBodyToJSON(v *SkillsShareBody) ([]byte, error) {
	return json.Marshal(v)
}
