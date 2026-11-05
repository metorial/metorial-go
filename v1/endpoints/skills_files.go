package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/files"
)

// SkillsFilesEndpoint provides access to read and edit the files of a skill by path.
type SkillsFilesEndpoint struct {
	client *endpoint.Client
}

// NewSkillsFilesEndpoint creates a new SkillsFilesEndpoint.
func NewSkillsFilesEndpoint(client *endpoint.Client) *SkillsFilesEndpoint {
	return &SkillsFilesEndpoint{client: client}
}

// SkillsFilesEndpointGetParams contains optional query parameters for Get.
type SkillsFilesEndpointGetParams struct {
	// Path - Path of the file inside the skill. A leading "/" is optional. Paths may contain letters, numbers, spaces and the characters . _ - ( ) $.
	Path string `json:"path"`
}

// SkillsFilesEndpointPutBody contains the request body for Put.
type SkillsFilesEndpointPutBody struct {
	// Path - Path of the file inside the skill. A leading "/" is optional. Paths may contain letters, numbers, spaces and the characters . _ - ( ) $.
	Path string `json:"path"`
	// Content - The new text content of the file. Markdown files are stored as documents, other files as plain text files.
	Content *string `json:"content,omitempty"`
	// FileId - The ID of a previously uploaded file to place at the path. Use this for binary files. Files that belong to a document cannot be attached; use content to write markdown documents.
	FileId *string `json:"file_id,omitempty"`
	// IfMatch - Only apply the change if the file still has this etag. Pass null to require that no file exists at the path yet. Omit to skip the check.
	IfMatch *string `json:"if_match,omitempty"`
}

// SkillsFilesEndpointDeleteParams contains optional query parameters for Delete.
type SkillsFilesEndpointDeleteParams struct {
	// Path - Path of the file inside the skill. A leading "/" is optional. Paths may contain letters, numbers, spaces and the characters . _ - ( ) $.
	Path string `json:"path"`
	// IfMatch - Only apply the change if the file still has this etag. Omit to skip the check.
	IfMatch *string `json:"if_match,omitempty"`
}

// SkillsFilesEndpointMoveBody contains the request body for Move.
type SkillsFilesEndpointMoveBody struct {
	// Path - Path of the file inside the skill. A leading "/" is optional. Paths may contain letters, numbers, spaces and the characters . _ - ( ) $.
	Path string `json:"path"`
	// ToPath - The new path of the file.
	ToPath string `json:"to_path"`
	// IfMatch - Only apply the change if the file still has this etag. Omit to skip the check.
	IfMatch *string `json:"if_match,omitempty"`
}

// SkillsFilesEndpointBatchBody contains the request body for Batch.
type SkillsFilesEndpointBatchBody struct {
	// Changes - The changes to apply, in order.
	Changes []any `json:"changes"`
}

// List lists every file of a skill with its path, size and etag. Directories are implied by file paths.
func (e *SkillsFilesEndpoint) List(skillId string) (*files.SkillsFilesListOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skills", skillId, "files"},
	}
	var result files.SkillsFilesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a single skill file by path. Documents include their text content, other files include a download URL.
func (e *SkillsFilesEndpoint) Get(skillId string, params *SkillsFilesEndpointGetParams) (*files.SkillsFilesGetOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"skills", skillId, "files", "content"},
		Query: query,
	}
	var result files.SkillsFilesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Put creates or replaces the file at a path. Parent directories are created automatically.
func (e *SkillsFilesEndpoint) Put(skillId string, body *SkillsFilesEndpointPutBody) (*files.SkillsFilesPutOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skills", skillId, "files"},
		Body: body,
	}
	var result files.SkillsFilesPutOutput
	if err := e.client.Put(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes the file at a path. Directories that were created automatically and become empty are removed. The root SKILL.md cannot be deleted.
func (e *SkillsFilesEndpoint) Delete(skillId string, params *SkillsFilesEndpointDeleteParams) (*files.SkillsFilesDeleteOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"skills", skillId, "files"},
		Query: query,
	}
	var result files.SkillsFilesDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Move moves or renames a file. The root SKILL.md cannot be moved.
func (e *SkillsFilesEndpoint) Move(skillId string, body *SkillsFilesEndpointMoveBody) (*files.SkillsFilesMoveOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skills", skillId, "files", "move"},
		Body: body,
	}
	var result files.SkillsFilesMoveOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Batch applies up to 100 put, delete and move changes one after another, in order. Every change is validated before any is applied, but changes are not applied atomically: if a change fails while being applied, the changes before it stay applied. Every change can carry its own if_match precondition.
func (e *SkillsFilesEndpoint) Batch(skillId string, body *SkillsFilesEndpointBatchBody) (*files.SkillsFilesBatchOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skills", skillId, "files", "batch"},
		Body: body,
	}
	var result files.SkillsFilesBatchOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
