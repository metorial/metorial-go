package agents

import (
	"encoding/json"
	"time"
)

// SkillsAgentsGetOutput represents the skills agents get output type.
type SkillsAgentsGetOutput struct {
	// Object - String representing the object's type
	Object      string     `json:"object"`
	Id          string     `json:"id"`
	SkillId     string     `json:"skill_id"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	Slug        string     `json:"slug"`
	Status      string     `json:"status"`
	StoreId     string     `json:"store_id"`
	StoreItemId *string    `json:"store_item_id,omitempty"`
	Path        *string    `json:"path,omitempty"`
	DocumentId  string     `json:"document_id"`
	ArchivedAt  *time.Time `json:"archived_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// MapSkillsAgentsGetOutputFromJSON deserializes JSON data into a SkillsAgentsGetOutput.
func MapSkillsAgentsGetOutputFromJSON(data []byte) (*SkillsAgentsGetOutput, error) {
	var v SkillsAgentsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsAgentsGetOutputToJSON serializes a SkillsAgentsGetOutput to JSON.
func MapSkillsAgentsGetOutputToJSON(v *SkillsAgentsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
