package agents

import (
	"encoding/json"
	"time"
)

// SkillsAgentsDeleteOutput represents the skills agents delete output type.
type SkillsAgentsDeleteOutput struct {
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

// MapSkillsAgentsDeleteOutputFromJSON deserializes JSON data into a SkillsAgentsDeleteOutput.
func MapSkillsAgentsDeleteOutputFromJSON(data []byte) (*SkillsAgentsDeleteOutput, error) {
	var v SkillsAgentsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsAgentsDeleteOutputToJSON serializes a SkillsAgentsDeleteOutput to JSON.
func MapSkillsAgentsDeleteOutputToJSON(v *SkillsAgentsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
