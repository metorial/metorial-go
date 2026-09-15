package forksyncs

import (
	"encoding/json"
	"time"
)

// SkillsForkSyncsGetOutput represents the skills fork syncs get output type.
type SkillsForkSyncsGetOutput struct {
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

// MapSkillsForkSyncsGetOutputFromJSON deserializes JSON data into a SkillsForkSyncsGetOutput.
func MapSkillsForkSyncsGetOutputFromJSON(data []byte) (*SkillsForkSyncsGetOutput, error) {
	var v SkillsForkSyncsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsForkSyncsGetOutputToJSON serializes a SkillsForkSyncsGetOutput to JSON.
func MapSkillsForkSyncsGetOutputToJSON(v *SkillsForkSyncsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
