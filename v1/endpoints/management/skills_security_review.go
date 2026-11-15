package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/securityreview"
)

// SkillsSecurityReviewEndpoint provides access to a security review assesses what a skill makes an agent do and whether its files can be trusted. Everyone who can read a skill can read its review.
type SkillsSecurityReviewEndpoint struct {
	client *endpoint.Client
}

// NewSkillsSecurityReviewEndpoint creates a new SkillsSecurityReviewEndpoint.
func NewSkillsSecurityReviewEndpoint(client *endpoint.Client) *SkillsSecurityReviewEndpoint {
	return &SkillsSecurityReviewEndpoint{client: client}
}

// Get returns the security review of a skill's latest version. Requests a review if none exists yet.
func (e *SkillsSecurityReviewEndpoint) Get(instanceId string, skillId string) (*securityreview.SkillsSecurityReviewGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skills", skillId, "security-review"},
	}
	var result securityreview.SkillsSecurityReviewGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Reassess starts a new review of the skill. Pending edits are sealed into a new version first.
func (e *SkillsSecurityReviewEndpoint) Reassess(instanceId string, skillId string) (*securityreview.SkillsSecurityReviewReassessOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skills", skillId, "security-review", "reassess"},
	}
	var result securityreview.SkillsSecurityReviewReassessOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Dismiss stops reviewing the skill. Its verdict becomes `dismissed` and no longer holds back the skill.
func (e *SkillsSecurityReviewEndpoint) Dismiss(instanceId string, skillId string) (*securityreview.SkillsSecurityReviewDismissOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skills", skillId, "security-review", "dismiss"},
	}
	var result securityreview.SkillsSecurityReviewDismissOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Restore resumes reviewing a dismissed skill and starts a new review.
func (e *SkillsSecurityReviewEndpoint) Restore(instanceId string, skillId string) (*securityreview.SkillsSecurityReviewRestoreOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skills", skillId, "security-review", "restore"},
	}
	var result securityreview.SkillsSecurityReviewRestoreOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
