package linux

import "testing"

func TestValidFromStdOut(t *testing.T) {
	result, err := fromStdOut("Private key: uHMCmvXiUFxdniESp-k_lZHwnpkxNxeqjPAdPgKhNWo\nPublic key: 2M_r3tjVelo8SGo2zaWOjGeMkHkDFO9TxtDeMIoa2ns")
	if err != nil {
		t.Error("Invalid response, expected no errors")
	}
	if result.Pub != "2M_r3tjVelo8SGo2zaWOjGeMkHkDFO9TxtDeMIoa2ns" {
		t.Error("Unexpected public key")
	}
	if result.Private != "uHMCmvXiUFxdniESp-k_lZHwnpkxNxeqjPAdPgKhNWo" {
		t.Error("Unexpected private key")
	}
}

func TestInvalidStdOut(t *testing.T) {
	result, err := fromStdOut("joajpajospajasposadjo\n")
	if err == nil {
		t.Error("Expected error got response")
	}
	if result != nil {
		t.Error("Unexpected result")
	}
}
