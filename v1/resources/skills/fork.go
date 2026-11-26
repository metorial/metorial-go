package skills

import (
	"encoding/json"
	"time"
)

// SkillsForkOutputImage - Image shown for the skill.
type SkillsForkOutputImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Where the skill image comes from.
	Type string `json:"type"`
	// Icon - Library icon ID. `null` unless `type` is `icon`.
	Icon *string `json:"icon,omitempty"`
	// Color - Tile color. `null` unless `type` is `icon`.
	Color *string `json:"color,omitempty"`
}

// SkillsForkOutputHierarchyCreatorOrganizationActorMember represents the skills fork output hierarchy creator organization actor member type.
type SkillsForkOutputHierarchyCreatorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
}

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

// SkillsForkOutputHierarchyCreatorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
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
	Member   *SkillsForkOutputHierarchyCreatorOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsForkOutputHierarchyCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsForkOutputHierarchyCreatorConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsForkOutputHierarchyCreatorConsumer struct {
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

// SkillsForkOutputHierarchyCreator - Actor behind a participant.
type SkillsForkOutputHierarchyCreator struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsForkOutputHierarchyCreatorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsForkOutputHierarchyCreatorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsForkOutputHierarchyForkCreatorOrganizationActorMember represents the skills fork output hierarchy fork creator organization actor member type.
type SkillsForkOutputHierarchyForkCreatorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
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

// SkillsForkOutputHierarchyForkCreatorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
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
	Member   *SkillsForkOutputHierarchyForkCreatorOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsForkOutputHierarchyForkCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsForkOutputHierarchyForkCreatorConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsForkOutputHierarchyForkCreatorConsumer struct {
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

// SkillsForkOutputHierarchyForkCreator - Actor behind a participant.
type SkillsForkOutputHierarchyForkCreator struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsForkOutputHierarchyForkCreatorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsForkOutputHierarchyForkCreatorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsForkOutputHierarchyForkOriginalCreatorOrganizationActorMember represents the skills fork output hierarchy fork original creator organization actor member type.
type SkillsForkOutputHierarchyForkOriginalCreatorOrganizationActorMember struct {
	// Object - String representing the organization's member preview type
	Object string `json:"object"`
	// Id - The organization member's unique identifier
	Id string `json:"id"`
	// Status - The organization member's status
	Status string `json:"status"`
	// Role - The organization member's role
	Role string `json:"role"`
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

// SkillsForkOutputHierarchyForkOriginalCreatorOrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
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
	Member   *SkillsForkOutputHierarchyForkOriginalCreatorOrganizationActorMember `json:"member,omitempty"`
	Teams    []SkillsForkOutputHierarchyForkOriginalCreatorOrganizationActorTeams `json:"teams"`
	// CreatedAt - The organization member's creation date
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - The organization member's last update date
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsForkOutputHierarchyForkOriginalCreatorConsumer - Consumer details. `null` unless `type` is `consumer`.
type SkillsForkOutputHierarchyForkOriginalCreatorConsumer struct {
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

// SkillsForkOutputHierarchyForkOriginalCreator - Actor behind a participant.
type SkillsForkOutputHierarchyForkOriginalCreator struct {
	// Type - What kind of actor this is.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// ImageUrl - Profile image URL.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address.
	Email *string `json:"email,omitempty"`
	// OrganizationActor - Organization actor details. `null` unless `type` is `organization_actor`.
	OrganizationActor *SkillsForkOutputHierarchyForkOriginalCreatorOrganizationActor `json:"organization_actor,omitempty"`
	// Consumer - Consumer details. `null` unless `type` is `consumer`.
	Consumer *SkillsForkOutputHierarchyForkOriginalCreatorConsumer `json:"consumer,omitempty"`
	// ConsumerProfile - Consumer profile details. `null` unless `type` is `consumer`.
	ConsumerProfile *map[string]any `json:"consumer_profile,omitempty"`
}

// SkillsForkOutputHierarchyFork - Fork details. `null` unless the skill is a fork.
type SkillsForkOutputHierarchyFork struct {
	// Id - Unique ID.
	Id string `json:"id"`
	// ParentSkillId - Upstream skill the fork was created from, by ID.
	ParentSkillId string `json:"parent_skill_id"`
	// Creator - Actor behind a participant.
	Creator *SkillsForkOutputHierarchyForkCreator `json:"creator,omitempty"`
	// OriginalCreator - Actor behind a participant.
	OriginalCreator *SkillsForkOutputHierarchyForkOriginalCreator `json:"original_creator,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// SkillsForkOutputHierarchyEntity - Skill entity the versioned skill content belongs to.
type SkillsForkOutputHierarchyEntity struct {
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

// SkillsForkOutputHierarchy - Where the skill came from.
type SkillsForkOutputHierarchy struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - How the skill was created. `root` for original skills, `fork` for consumer forks, `duplicated` for copies.
	Type string `json:"type"`
	// ParentSkillId - Parent skill, by ID. `null` for root skills.
	ParentSkillId *string `json:"parent_skill_id,omitempty"`
	// Creator - Actor behind a participant.
	Creator *SkillsForkOutputHierarchyCreator `json:"creator,omitempty"`
	// Fork - Fork details. `null` unless the skill is a fork.
	Fork *SkillsForkOutputHierarchyFork `json:"fork,omitempty"`
	// Entity - Skill entity the versioned skill content belongs to.
	Entity SkillsForkOutputHierarchyEntity `json:"entity"`
}

// SkillsForkOutputIntegrationsConfiguration - Which settings integration instances may customize.
type SkillsForkOutputIntegrationsConfiguration struct {
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

// SkillsForkOutputIntegrations represents the skills fork output integrations type.
type SkillsForkOutputIntegrations struct {
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
	Configuration SkillsForkOutputIntegrationsConfiguration `json:"configuration"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// SkillsForkOutputProviders represents the skills fork output providers type.
type SkillsForkOutputProviders struct {
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

// SkillsForkOutput represents the skills fork output type.
type SkillsForkOutput struct {
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
	Image SkillsForkOutputImage `json:"image"`
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
	Hierarchy SkillsForkOutputHierarchy `json:"hierarchy"`
	// Integrations - Integrations attached to the skill.
	Integrations []SkillsForkOutputIntegrations `json:"integrations"`
	// Providers - Providers attached to the skill.
	Providers []SkillsForkOutputProviders `json:"providers"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
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
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// ClientName - Client-facing skill name. Defaults to `name`. Lowercase letters, numbers, and hyphens.
	ClientName *string `json:"client_name,omitempty"`
	// ClientDescription - Client-facing description.
	ClientDescription *string `json:"client_description,omitempty"`
	// License - License terms for the skill.
	License *string `json:"license,omitempty"`
	// Compatibility - Compatibility notes for the skill.
	Compatibility *string `json:"compatibility,omitempty"`
	// ClientMetadata - Custom data for the client-facing skill.
	ClientMetadata *map[string]any `json:"client_metadata,omitempty"`
	// Metadata - Custom data to identify or organize the skill.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ImageFileId - Image file for the skill, by ID.
	ImageFileId *string `json:"image_file_id,omitempty"`
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
