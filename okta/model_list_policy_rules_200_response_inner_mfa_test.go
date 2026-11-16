package okta

import (
	"encoding/json"
	"testing"
)

// Regression test: MFA_ENROLL policy rules must not be silently dropped.
// The v5 fork carried its own MfaEnrollPolicyRule because the ListPolicyRules
// oneOf had no MFA_ENROLL case. Upstream now maps MFA_ENROLL to
// AuthenticatorEnrollmentPolicyRule, so only this test is carried forward.
func TestListPolicyRules200ResponseInner_MfaEnroll(t *testing.T) {
	data := []byte(`{
		"id": "rul123mfaenroll",
		"type": "MFA_ENROLL",
		"name": "Default Rule",
		"status": "ACTIVE",
		"priority": 1,
		"system": true,
		"conditions": {"people": {"users": {"exclude": []}}},
		"actions": {"enroll": {"self": "CHALLENGE"}}
	}`)

	var inner ListPolicyRules200ResponseInner
	if err := json.Unmarshal(data, &inner); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	rule := inner.AuthenticatorEnrollmentPolicyRule
	if rule == nil {
		t.Fatal("AuthenticatorEnrollmentPolicyRule is nil; MFA_ENROLL rule was dropped")
	}
	if got := rule.GetId(); got != "rul123mfaenroll" {
		t.Errorf("id = %q, want %q", got, "rul123mfaenroll")
	}
	if got := rule.GetType(); got != "MFA_ENROLL" {
		t.Errorf("type = %q, want %q", got, "MFA_ENROLL")
	}
	if inst := inner.GetActualInstance(); inst == nil {
		t.Error("GetActualInstance returned nil for MFA_ENROLL rule")
	}
	// actions is a typed field on the generated AuthenticatorEnrollmentPolicyRule
	if !rule.HasActions() {
		t.Fatal("expected actions to be populated on AuthenticatorEnrollmentPolicyRule")
	}
	actions := rule.GetActions()
	enroll := actions.GetEnroll()
	if got := enroll.GetSelf(); got != "CHALLENGE" {
		t.Errorf("actions.enroll.self = %q, want %q", got, "CHALLENGE")
	}
}
