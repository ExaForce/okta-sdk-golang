package okta

import (
	"encoding/json"
	"testing"
)

// Regression test: MFA_ENROLL policy rules must not be silently dropped.
// Before the fix, the oneOf UnmarshalJSON had no MFA_ENROLL case, so these
// rules deserialized into an all-nil struct and callers discarded them.
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

	if inner.MfaEnrollPolicyRule == nil {
		t.Fatal("MfaEnrollPolicyRule is nil; MFA_ENROLL rule was dropped")
	}
	if got := inner.MfaEnrollPolicyRule.GetId(); got != "rul123mfaenroll" {
		t.Errorf("id = %q, want %q", got, "rul123mfaenroll")
	}
	if got := inner.MfaEnrollPolicyRule.GetType(); got != "MFA_ENROLL" {
		t.Errorf("type = %q, want %q", got, "MFA_ENROLL")
	}
	if inst := inner.GetActualInstance(); inst == nil {
		t.Error("GetActualInstance returned nil for MFA_ENROLL rule")
	}
	// actions is a typed field on the generated MfaEnrollPolicyRule
	if inner.MfaEnrollPolicyRule.GetActions() == nil {
		t.Error("expected actions to be populated on MfaEnrollPolicyRule")
	}
}
