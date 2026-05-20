package snapshot

import (
	"encoding/json"
	"time"
)

// SkillsVersionsSnapshotGetOutputItems represents the skills versions snapshot get output items type.
type SkillsVersionsSnapshotGetOutputItems struct {
	Object            string    `json:"object"`
	Id                string    `json:"id"`
	Kind              string    `json:"kind"`
	Path              string    `json:"path"`
	FileId            *string   `json:"file_id,omitempty"`
	DocumentId        *string   `json:"document_id,omitempty"`
	DocumentVersionId *string   `json:"document_version_id,omitempty"`
	Content           *string   `json:"content,omitempty"`
	CreatedAt         time.Time `json:"created_at"`
}

// SkillsVersionsSnapshotGetOutput represents the skills versions snapshot get output type.
type SkillsVersionsSnapshotGetOutput struct {
	Object         string                                 `json:"object"`
	Id             string                                 `json:"id"`
	SkillId        string                                 `json:"skill_id"`
	StoreId        string                                 `json:"store_id"`
	StoreVersionId string                                 `json:"store_version_id"`
	VersionNumber  float64                                `json:"version_number"`
	Items          []SkillsVersionsSnapshotGetOutputItems `json:"items"`
	CreatedAt      time.Time                              `json:"created_at"`
}

// MapSkillsVersionsSnapshotGetOutputFromJSON deserializes JSON data into a SkillsVersionsSnapshotGetOutput.
func MapSkillsVersionsSnapshotGetOutputFromJSON(data []byte) (*SkillsVersionsSnapshotGetOutput, error) {
	var v SkillsVersionsSnapshotGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsVersionsSnapshotGetOutputToJSON serializes a SkillsVersionsSnapshotGetOutput to JSON.
func MapSkillsVersionsSnapshotGetOutputToJSON(v *SkillsVersionsSnapshotGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
