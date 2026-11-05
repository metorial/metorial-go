package access

import (
	"encoding/json"
)

// SkillsAccessGetOutputGrantsSubject represents the skills access get output grants subject type.
type SkillsAccessGetOutputGrantsSubject struct {
	Type              string  `json:"type"`
	ConsumerProfileId *string `json:"consumer_profile_id,omitempty"`
	ConsumerGroupId   *string `json:"consumer_group_id,omitempty"`
	Name              *string `json:"name,omitempty"`
}

// SkillsAccessGetOutputGrants represents the skills access get output grants type.
type SkillsAccessGetOutputGrants struct {
	Object  string                             `json:"object"`
	Level   string                             `json:"level"`
	Subject SkillsAccessGetOutputGrantsSubject `json:"subject"`
}

// SkillsAccessGetOutput represents the skills access get output type.
type SkillsAccessGetOutput struct {
	// Object - String representing the object's type
	Object  string                        `json:"object"`
	SkillId string                        `json:"skill_id"`
	Grants  []SkillsAccessGetOutputGrants `json:"grants"`
}

// MapSkillsAccessGetOutputFromJSON deserializes JSON data into a SkillsAccessGetOutput.
func MapSkillsAccessGetOutputFromJSON(data []byte) (*SkillsAccessGetOutput, error) {
	var v SkillsAccessGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsAccessGetOutputToJSON serializes a SkillsAccessGetOutput to JSON.
func MapSkillsAccessGetOutputToJSON(v *SkillsAccessGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
