package agents

import (
	"encoding/json"
	"time"
)

// SkillsAgentsUpdateOutput represents the skills agents update output type.
type SkillsAgentsUpdateOutput struct {
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

// MapSkillsAgentsUpdateOutputFromJSON deserializes JSON data into a SkillsAgentsUpdateOutput.
func MapSkillsAgentsUpdateOutputFromJSON(data []byte) (*SkillsAgentsUpdateOutput, error) {
	var v SkillsAgentsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsAgentsUpdateOutputToJSON serializes a SkillsAgentsUpdateOutput to JSON.
func MapSkillsAgentsUpdateOutputToJSON(v *SkillsAgentsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsAgentsUpdateBody represents the skills agents update body type.
type SkillsAgentsUpdateBody struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// MapSkillsAgentsUpdateBodyFromJSON deserializes JSON data into a SkillsAgentsUpdateBody.
func MapSkillsAgentsUpdateBodyFromJSON(data []byte) (*SkillsAgentsUpdateBody, error) {
	var v SkillsAgentsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsAgentsUpdateBodyToJSON serializes a SkillsAgentsUpdateBody to JSON.
func MapSkillsAgentsUpdateBodyToJSON(v *SkillsAgentsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
