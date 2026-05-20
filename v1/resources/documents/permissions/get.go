package permissions

import (
	"encoding/json"
)

// DocumentsPermissionsGetOutput represents the documents permissions get output type.
type DocumentsPermissionsGetOutput struct {
	// Object - String representing the object's type
	Object           string   `json:"object"`
	DocumentId       string   `json:"document_id"`
	IsOwner          bool     `json:"is_owner"`
	HasFullAccess    bool     `json:"has_full_access"`
	Permissions      []string `json:"permissions"`
	RelevantStoreIds []string `json:"relevant_store_ids"`
	ReadableStoreIds []string `json:"readable_store_ids"`
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
