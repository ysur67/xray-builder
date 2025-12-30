package commands

import "testing"

func TestValidFromStdOut(t *testing.T) {
	result, err := fromStdOut("PrivateKey: uHMCmvXiUFxdniESpk_lZHwnpkxNxeqjPAdPgKhNWo\nPassword: 2M_r3tjVelo8SGo2zaWOjGeMkHkDFO9TxtDeMIoa2ns\nHash32: 7s8O-YBq1hon4rKk-8zbsEo2EBsQjMCjzfGAfGUItvo")
	if err != nil {
		t.Error("Invalid response, expected no errors")
	}
	if result.Pub != "2M_r3tjVelo8SGo2zaWOjGeMkHkDFO9TxtDeMIoa2ns" {
		t.Error("Unexpected public key")
	}
	if result.Private != "uHMCmvXiUFxdniESpk_lZHwnpkxNxeqjPAdPgKhNWo" {
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
