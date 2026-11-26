package permissions

import (
	"encoding/json"
)

// DocumentsPermissionsGetOutput represents the documents permissions get output type.
type DocumentsPermissionsGetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// DocumentId - Document these permissions apply to.
	DocumentId string `json:"document_id"`
	// IsOwner - Whether the caller owns the document.
	IsOwner bool `json:"is_owner"`
	// HasFullAccess - Whether the caller has full access.
	HasFullAccess bool `json:"has_full_access"`
	// Permissions - Permissions the caller has.
	Permissions []string `json:"permissions"`
	// RelevantStoreIds - Stores this resource belongs to.
	RelevantStoreIds []string `json:"relevant_store_ids"`
	// ReadableStoreIds - Relevant stores the caller can read.
	ReadableStoreIds []string `json:"readable_store_ids"`
	// WritableStoreIds - Relevant stores the caller can write to.
	WritableStoreIds []string `json:"writable_store_ids"`
}

// MapDocumentsPermissionsGetOutputFromJSON deserializes JSON data into a DocumentsPermissionsGetOutput.
func MapDocumentsPermissionsGetOutputFromJSON(data []byte) (*DocumentsPermissionsGetOutput, error) {
	var v DocumentsPermissionsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsPermissionsGetOutputToJSON serializes a DocumentsPermissionsGetOutput to JSON.
func MapDocumentsPermissionsGetOutputToJSON(v *DocumentsPermissionsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
