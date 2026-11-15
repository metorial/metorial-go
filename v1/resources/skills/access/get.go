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

// SkillsAccessGetOutputGrantsSourceSkillGroup - Set when the access is inherited from a skill group.
type SkillsAccessGetOutputGrantsSourceSkillGroup struct {
	// Id - The skill group the access is inherited from.
	Id string `json:"id"`
	// Name - The name of the skill group.
	Name string `json:"name"`
}

// SkillsAccessGetOutputGrantsSource - Where this access grant comes from.
type SkillsAccessGetOutputGrantsSource struct {
	// Type - Whether access was granted on the skill itself or is inherited from a skill group that contains it.
	Type string `json:"type"`
	// SkillGroup - Set when the access is inherited from a skill group.
	SkillGroup *SkillsAccessGetOutputGrantsSourceSkillGroup `json:"skill_group,omitempty"`
}

// SkillsAccessGetOutputGrants represents the skills access get output grants type.
type SkillsAccessGetOutputGrants struct {
	Object  string                             `json:"object"`
	Level   string                             `json:"level"`
	Subject SkillsAccessGetOutputGrantsSubject `json:"subject"`
	// Source - Where this access grant comes from.
	Source SkillsAccessGetOutputGrantsSource `json:"source"`
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
