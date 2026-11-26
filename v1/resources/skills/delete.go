package skills

import (
	"encoding/json"
	"time"
)

// SkillsDeleteOutputImage - Image shown for the skill.
type SkillsDeleteOutputImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Where the skill image comes from.
	Type string `json:"type"`
	// Icon - Library icon ID. `null` unless `type` is `icon`.
	Icon *string `json:"icon,omitempty"`
	// Color - Tile color. `null` unless `type` is `icon`.
	Color *string `json:"color,omitempty"`
}

// SkillsDeleteOutputHierarchyCreatorOrganizationActorMember represents the skills delete output hierarchy creator organization actor member type.
type SkillsDeleteOutputHierarchyCreatorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsDeleteOutputHierarchyCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsDeleteOutputHierarchyCreatorOrganizationActorTeams struct {
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

// SkillsDeleteOutputHierarchyCreatorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type SkillsDeleteOutputHierarchyCreatorOrganizationActor struct {
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
	Member   *SkillsDeleteOutputHierarchyCreatorOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsDeleteOutputHierarchyCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsDeleteOutputHierarchyCreatorConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsDeleteOutputHierarchyCreatorConsumer struct {
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

// SkillsDeleteOutputHierarchyCreator - Actor behind a participant.
type SkillsDeleteOutputHierarchyCreator struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsDeleteOutputHierarchyCreatorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsDeleteOutputHierarchyCreatorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsDeleteOutputHierarchyForkCreatorOrganizationActorMember represents the skills delete output hierarchy fork creator organization actor member type.
type SkillsDeleteOutputHierarchyForkCreatorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsDeleteOutputHierarchyForkCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsDeleteOutputHierarchyForkCreatorOrganizationActorTeams struct {
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

// SkillsDeleteOutputHierarchyForkCreatorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type SkillsDeleteOutputHierarchyForkCreatorOrganizationActor struct {
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
	Member   *SkillsDeleteOutputHierarchyForkCreatorOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsDeleteOutputHierarchyForkCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsDeleteOutputHierarchyForkCreatorConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsDeleteOutputHierarchyForkCreatorConsumer struct {
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

// SkillsDeleteOutputHierarchyForkCreator - Actor behind a participant.
type SkillsDeleteOutputHierarchyForkCreator struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsDeleteOutputHierarchyForkCreatorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsDeleteOutputHierarchyForkCreatorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsDeleteOutputHierarchyForkOriginalCreatorOrganizationActorMember represents the skills delete output hierarchy fork original creator organization actor member type.
type SkillsDeleteOutputHierarchyForkOriginalCreatorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

// SkillsDeleteOutputHierarchyForkOriginalCreatorOrganizationActorTeams - The teams the actor belongs to
type SkillsDeleteOutputHierarchyForkOriginalCreatorOrganizationActorTeams struct {
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

// SkillsDeleteOutputHierarchyForkOriginalCreatorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
type SkillsDeleteOutputHierarchyForkOriginalCreatorOrganizationActor struct {
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
	Member   *SkillsDeleteOutputHierarchyForkOriginalCreatorOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsDeleteOutputHierarchyForkOriginalCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsDeleteOutputHierarchyForkOriginalCreatorConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsDeleteOutputHierarchyForkOriginalCreatorConsumer struct {
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

// SkillsDeleteOutputHierarchyForkOriginalCreator - Actor behind a participant.
type SkillsDeleteOutputHierarchyForkOriginalCreator struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsDeleteOutputHierarchyForkOriginalCreatorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsDeleteOutputHierarchyForkOriginalCreatorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsDeleteOutputHierarchyFork - Fork details. `null` unless the skill is a fork.
type SkillsDeleteOutputHierarchyFork struct {
	// Id - Unique ID.
	Id string `json:"id"`
	// ParentSkillId - Upstream skill the fork was created from, by ID.
	ParentSkillId string `json:"parent_skill_id"`
	// Creator - Actor behind a participant.
	Creator *SkillsDeleteOutputHierarchyForkCreator `json:"creator,omitempty"`
	// OriginalCreator - Actor behind a participant.
	OriginalCreator *SkillsDeleteOutputHierarchyForkOriginalCreator `json:"original_creator,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// SkillsDeleteOutputHierarchyEntity - Skill entity the versioned skill content belongs to.
type SkillsDeleteOutputHierarchyEntity struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Slug - URL-safe name, unique within the instance.
	Slug string `json:"slug"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// ParentSkillId - Upstream skill, by ID.
	ParentSkillId string `json:"parent_skill_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsDeleteOutputHierarchy - Where the skill came from.
type SkillsDeleteOutputHierarchy struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - How the skill was created. `root` for original skills, `fork` for consumer forks, `duplicated` for copies.
	Type string `json:"type"`
	// ParentSkillId - Parent skill, by ID. `null` for root skills.
	ParentSkillId *string `json:"parent_skill_id,omitempty"`
	// Creator - Actor behind a participant.
	Creator *SkillsDeleteOutputHierarchyCreator `json:"creator,omitempty"`
	// Fork - Fork details. `null` unless the skill is a fork.
	Fork *SkillsDeleteOutputHierarchyFork `json:"fork,omitempty"`
	// Entity - Skill entity the versioned skill content belongs to.
	Entity SkillsDeleteOutputHierarchyEntity `json:"entity"`
}

// SkillsDeleteOutputIntegrationsConfiguration - Which settings integration instances may customize.
type SkillsDeleteOutputIntegrationsConfiguration struct {
	// CanAttachCustomToolFilters - Lets integration instances attach their own tool filters.
	CanAttachCustomToolFilters bool `json:"can_attach_custom_tool_filters"`
	// CanAttachCustomProviderConfig - Lets integration instances use their own provider configs.
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	// CanOverrideToolFilters - Lets integration instances ignore the integration tool filters.
	CanOverrideToolFilters bool `json:"can_override_tool_filters"`
	// EnableCallbackTools - Makes callback event tools available for this integration.
	EnableCallbackTools bool `json:"enable_callback_tools"`
	// UseIntegrationNameInToolNames - Names session providers after the integration instead of the provider. `null` uses the default.
	UseIntegrationNameInToolNames *bool `json:"use_integration_name_in_tool_names,omitempty"`
}

// SkillsDeleteOutputIntegrations represents the skills delete output integrations type.
type SkillsDeleteOutputIntegrations struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Slug - URL-safe name, unique across Metorial.
	Slug string `json:"slug"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Configuration - Which settings integration instances may customize.
	Configuration SkillsDeleteOutputIntegrationsConfiguration `json:"configuration"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// SkillsDeleteOutputProviders represents the skills delete output providers type.
type SkillsDeleteOutputProviders struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Slug - URL-safe name for the provider.
	Slug string `json:"slug"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsDeleteOutput represents the skills delete output type.
type SkillsDeleteOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Whether the skill is available, archived, or deleted.
	Status string `json:"status"`
	// Slug - URL-safe name, unique within the instance.
	Slug string `json:"slug"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// ImageUrl - Public URL of the skill image.
	ImageUrl string `json:"image_url"`
	// Image - Image shown for the skill.
	Image SkillsDeleteOutputImage `json:"image"`
	// ClientName - Client-facing skill name.
	ClientName string `json:"client_name"`
	// ClientDescription - Client-facing description.
	ClientDescription *string `json:"client_description,omitempty"`
	// ClientMetadata - Custom data for the client-facing skill.
	ClientMetadata *map[string]any `json:"client_metadata,omitempty"`
	// License - License terms for the skill.
	License *string `json:"license,omitempty"`
	// Compatibility - Compatibility notes for the skill.
	Compatibility *string `json:"compatibility,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata map[string]any `json:"metadata"`
	// StoreId - Store holding the skill, by ID.
	StoreId string `json:"store_id"`
	// Hierarchy - Where the skill came from.
	Hierarchy SkillsDeleteOutputHierarchy `json:"hierarchy"`
	// Integrations - Integrations attached to the skill.
	Integrations []SkillsDeleteOutputIntegrations `json:"integrations"`
	// Providers - Providers attached to the skill.
	Providers []SkillsDeleteOutputProviders `json:"providers"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapSkillsDeleteOutputFromJSON deserializes JSON data into a SkillsDeleteOutput.
func MapSkillsDeleteOutputFromJSON(data []byte) (*SkillsDeleteOutput, error) {
	var v SkillsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsDeleteOutputToJSON serializes a SkillsDeleteOutput to JSON.
func MapSkillsDeleteOutputToJSON(v *SkillsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
