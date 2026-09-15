package edittoken

import (
	"encoding/json"
	"time"
)

// DocumentsEditTokenGetOutput represents the documents edit token get output type.
type DocumentsEditTokenGetOutput struct {
	// Object - String representing the object's type
	Object     string    `json:"object"`
	Token      string    `json:"token"`
	ExpiresAt  time.Time `json:"expires_at"`
	DocumentId string    `json:"document_id"`
}

// MapDocumentsEditTokenGetOutputFromJSON deserializes JSON data into a DocumentsEditTokenGetOutput.
func MapDocumentsEditTokenGetOutputFromJSON(data []byte) (*DocumentsEditTokenGetOutput, error) {
	var v DocumentsEditTokenGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapDocumentsEditTokenGetOutputToJSON serializes a DocumentsEditTokenGetOutput to JSON.
func MapDocumentsEditTokenGetOutputToJSON(v *DocumentsEditTokenGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
