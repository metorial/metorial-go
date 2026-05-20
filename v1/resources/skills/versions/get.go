package versions

import (
	"encoding/json"
	"time"
)

// SkillsVersionsGetOutput represents the skills versions get output type.
type SkillsVersionsGetOutput struct {
	Object         string    `json:"object"`
	Id             string    `json:"id"`
	SkillId        string    `json:"skill_id"`
	StoreId        string    `json:"store_id"`
	StoreVersionId string    `json:"store_version_id"`
	VersionNumber  float64   `json:"version_number"`
	CreatedAt      time.Time `json:"created_at"`
}

// MapSkillsVersionsGetOutputFromJSON deserializes JSON data into a SkillsVersionsGetOutput.
func MapSkillsVersionsGetOutputFromJSON(data []byte) (*SkillsVersionsGetOutput, error) {
	var v SkillsVersionsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsVersionsGetOutputToJSON serializes a SkillsVersionsGetOutput to JSON.
func MapSkillsVersionsGetOutputToJSON(v *SkillsVersionsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
