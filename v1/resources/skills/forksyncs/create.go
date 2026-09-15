package forksyncs

import (
	"encoding/json"
	"time"
)

// SkillsForkSyncsCreateOutput represents the skills fork syncs create output type.
type SkillsForkSyncsCreateOutput struct {
	Object              string     `json:"object"`
	Id                  string     `json:"id"`
	Status              string     `json:"status"`
	ForkSkillId         string     `json:"fork_skill_id"`
	UpstreamSkillId     string     `json:"upstream_skill_id"`
	MergeRequestId      *string    `json:"merge_request_id,omitempty"`
	Error               *string    `json:"error,omitempty"`
	ProcessingStartedAt *time.Time `json:"processing_started_at,omitempty"`
	ActionRequiredAt    *time.Time `json:"action_required_at,omitempty"`
	CompletedAt         *time.Time `json:"completed_at,omitempty"`
	FailedAt            *time.Time `json:"failed_at,omitempty"`
	CancelledAt         *time.Time `json:"cancelled_at,omitempty"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
}

// MapSkillsForkSyncsCreateOutputFromJSON deserializes JSON data into a SkillsForkSyncsCreateOutput.
func MapSkillsForkSyncsCreateOutputFromJSON(data []byte) (*SkillsForkSyncsCreateOutput, error) {
	var v SkillsForkSyncsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsForkSyncsCreateOutputToJSON serializes a SkillsForkSyncsCreateOutput to JSON.
func MapSkillsForkSyncsCreateOutputToJSON(v *SkillsForkSyncsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsForkSyncsCreateBody represents the skills fork syncs create body type.
type SkillsForkSyncsCreateBody struct {
	SkillId string `json:"skill_id"`
}

// MapSkillsForkSyncsCreateBodyFromJSON deserializes JSON data into a SkillsForkSyncsCreateBody.
func MapSkillsForkSyncsCreateBodyFromJSON(data []byte) (*SkillsForkSyncsCreateBody, error) {
	var v SkillsForkSyncsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsForkSyncsCreateBodyToJSON serializes a SkillsForkSyncsCreateBody to JSON.
func MapSkillsForkSyncsCreateBodyToJSON(v *SkillsForkSyncsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
