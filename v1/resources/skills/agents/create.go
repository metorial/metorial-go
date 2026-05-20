package agents

import (
	"encoding/json"
	"time"
)

// SkillsAgentsCreateOutput represents the skills agents create output type.
type SkillsAgentsCreateOutput struct {
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

// MapSkillsAgentsCreateOutputFromJSON deserializes JSON data into a SkillsAgentsCreateOutput.
func MapSkillsAgentsCreateOutputFromJSON(data []byte) (*SkillsAgentsCreateOutput, error) {
	var v SkillsAgentsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsAgentsCreateOutputToJSON serializes a SkillsAgentsCreateOutput to JSON.
func MapSkillsAgentsCreateOutputToJSON(v *SkillsAgentsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsAgentsCreateBody represents the skills agents create body type.
type SkillsAgentsCreateBody struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Content     *string `json:"content,omitempty"`
}

// MapSkillsAgentsCreateBodyFromJSON deserializes JSON data into a SkillsAgentsCreateBody.
func MapSkillsAgentsCreateBodyFromJSON(data []byte) (*SkillsAgentsCreateBody, error) {
	var v SkillsAgentsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsAgentsCreateBodyToJSON serializes a SkillsAgentsCreateBody to JSON.
func MapSkillsAgentsCreateBodyToJSON(v *SkillsAgentsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
