package simulate

import (
	"context"
	"os"
	"testing"
)

var (
	// Common identity policy used across SCP tests
	testIdentityPolicyEC2RunInstances = []byte(`{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Effect": "Allow",
				"Action": "ec2:RunInstances",
				"Resource": "*"
			}
		]
	}`)
)

// loadSCPExamples loads the allow-all and deny-by-region SCP example files
func loadSCPExamples(t *testing.T) (scpAllowAll, scpDenyRegion []byte) {
	t.Helper()

	scpAllowAll, err := os.ReadFile("../../examples/scp-allow-all.json")
	if err != nil {
		t.Fatalf("Failed to read allow-all SCP file: %v", err)
	}

	scpDenyRegion, err = os.ReadFile("../../examples/scp-deny-by-region.json")
	if err != nil {
		t.Fatalf("Failed to read deny-by-region SCP file: %v", err)
	}

	return scpAllowAll, scpDenyRegion
}

func TestRunSimulation_WithSCP_AllowedInEURegion(t *testing.T) {
	scpAllowAll, scpDenyRegion := loadSCPExamples(t)

	sim := Simulation{
		Action:           "ec2:RunInstances",
		Resource:         "*",
		Principal:        "arn:aws:iam::123456789012:user/testuser",
		IdentityPolicies: [][]byte{testIdentityPolicyEC2RunInstances},
		SCPs:             [][]byte{scpAllowAll, scpDenyRegion},
		Context: map[string]any{
			"aws:RequestedRegion": "eu-west-1",
			"aws:PrincipalARN":    "arn:aws:iam::123456789012:user/testuser",
		},
	}

	result, err := RunSimulation(context.Background(), sim)
	if err != nil {
		t.Fatalf("RunSimulation failed: %v", err)
	}

	if result.Result != EvaluationResultAllowed {
		t.Errorf("Expected request to be allowed in EU region, but got %v", result.Result)
	}
}

func TestRunSimulation_WithSCP_DeniedInUSRegion(t *testing.T) {
	scpAllowAll, scpDenyRegion := loadSCPExamples(t)

	sim := Simulation{
		Action:           "ec2:RunInstances",
		Resource:         "*",
		Principal:        "arn:aws:iam::123456789012:user/testuser",
		IdentityPolicies: [][]byte{testIdentityPolicyEC2RunInstances},
		SCPs:             [][]byte{scpAllowAll, scpDenyRegion},
		Context: map[string]any{
			"aws:RequestedRegion": "us-east-1",
			"aws:PrincipalARN":    "arn:aws:iam::123456789012:user/testuser",
		},
	}

	result, err := RunSimulation(context.Background(), sim)
	if err != nil {
		t.Fatalf("RunSimulation failed: %v", err)
	}

	if result.Result != EvaluationResultExplicitlyDenied {
		t.Errorf("Expected request to be explicitly denied in US region, but got %v", result.Result)
	}
}

func TestRunSimulation_WithSCP_AllowedWithBypassRole(t *testing.T) {
	scpAllowAll, scpDenyRegion := loadSCPExamples(t)

	sim := Simulation{
		Action:           "ec2:RunInstances",
		Resource:         "*",
		Principal:        "arn:aws:iam::123456789012:role/Role1AllowedToBypassThisSCP",
		IdentityPolicies: [][]byte{testIdentityPolicyEC2RunInstances},
		SCPs:             [][]byte{scpAllowAll, scpDenyRegion},
		Context: map[string]any{
			"aws:RequestedRegion": "us-east-1",
			"aws:PrincipalARN":    "arn:aws:iam::123456789012:role/Role1AllowedToBypassThisSCP",
		},
	}

	result, err := RunSimulation(context.Background(), sim)
	if err != nil {
		t.Fatalf("RunSimulation failed: %v", err)
	}

	if result.Result != EvaluationResultAllowed {
		t.Errorf("Expected request to be allowed for bypass role, but got %v", result.Result)
	}
}
