package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills"
)

// SkillsEndpoint provides access to skills group provider and integration capabilities into reusable, owned compositions.
type SkillsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsEndpoint creates a new SkillsEndpoint.
func NewSkillsEndpoint(client *endpoint.Client) *SkillsEndpoint {
	return &SkillsEndpoint{client: client}
}

// SkillsEndpointListParams contains optional query parameters for List.
type SkillsEndpointListParams struct {
	Limit         *float64 `json:"limit,omitempty"`
	After         *string  `json:"after,omitempty"`
	Before        *string  `json:"before,omitempty"`
	Cursor        *string  `json:"cursor,omitempty"`
	Order         *string  `json:"order,omitempty"`
	Search        *string  `json:"search,omitempty"`
	Status        *any     `json:"status,omitempty"`
	Id            *any     `json:"id,omitempty"`
	SkillGroupId  *any     `json:"skill_group_id,omitempty"`
	IntegrationId *any     `json:"integration_id,omitempty"`
	ProviderId    *any     `json:"provider_id,omitempty"`
	// CreatedAt - Filter skill creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter skill last update time by date range
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// SkillsEndpointCreateBody contains the request body for Create.
type SkillsEndpointCreateBody struct {
	Name              string          `json:"name"`
	Description       *string         `json:"description,omitempty"`
	Metadata          *map[string]any `json:"metadata,omitempty"`
	ClientName        *string         `json:"client_name,omitempty"`
	ClientDescription *string         `json:"client_description,omitempty"`
	License           *string         `json:"license,omitempty"`
	Compatibility     *string         `json:"compatibility,omitempty"`
	ClientMetadata    *map[string]any `json:"client_metadata,omitempty"`
	ImageFileId       *string         `json:"image_file_id,omitempty"`
	TemplateId        *string         `json:"template_id,omitempty"`
}

// SkillsEndpointUpdateBody contains the request body for Update.
type SkillsEndpointUpdateBody struct {
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

// SkillsEndpointForkBody contains the request body for Fork.
type SkillsEndpointForkBody struct {
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

// SkillsEndpointDuplicateBody contains the request body for Duplicate.
type SkillsEndpointDuplicateBody struct {
	Name              string          `json:"name"`
	Description       *string         `json:"description,omitempty"`
	ClientName        *string         `json:"client_name,omitempty"`
	ClientDescription *string         `json:"client_description,omitempty"`
	License           *string         `json:"license,omitempty"`
	Compatibility     *string         `json:"compatibility,omitempty"`
	ClientMetadata    *map[string]any `json:"client_metadata,omitempty"`
	Metadata          *map[string]any `json:"metadata,omitempty"`
}

// List returns a paginated list of skills.
func (e *SkillsEndpoint) List(instanceId string, params *SkillsEndpointListParams) (*skills.SkillsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "skills"},
		Query: query,
	}
	var result skills.SkillsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific skill.
func (e *SkillsEndpoint) Get(instanceId string, skillId string) (*skills.SkillsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skills", skillId},
	}
	var result skills.SkillsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new skill.
func (e *SkillsEndpoint) Create(instanceId string, body *SkillsEndpointCreateBody) (*skills.SkillsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skills"},
		Body: body,
	}
	var result skills.SkillsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a specific skill.
func (e *SkillsEndpoint) Update(instanceId string, skillId string, body *SkillsEndpointUpdateBody) (*skills.SkillsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skills", skillId},
		Body: body,
	}
	var result skills.SkillsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a specific skill.
func (e *SkillsEndpoint) Delete(instanceId string, skillId string) (*skills.SkillsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skills", skillId},
	}
	var result skills.SkillsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fork forks a skill for the current consumer. Non-consumer callers duplicate the skill instead.
func (e *SkillsEndpoint) Fork(instanceId string, skillId string, body *SkillsEndpointForkBody) (*skills.SkillsForkOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skills", skillId, "fork"},
		Body: body,
	}
	var result skills.SkillsForkOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PublishConsumerSkill publishes a consumer-owned skill to the consumer groups they belong to.
func (e *SkillsEndpoint) PublishConsumerSkill(instanceId string, skillId string) (*skills.SkillsPublishConsumerSkillOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skills", skillId, "publish"},
	}
	var result skills.SkillsPublishConsumerSkillOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Duplicate duplicates a skill.
func (e *SkillsEndpoint) Duplicate(instanceId string, skillId string, body *SkillsEndpointDuplicateBody) (*skills.SkillsDuplicateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skills", skillId, "duplicate"},
		Body: body,
	}
	var result skills.SkillsDuplicateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
