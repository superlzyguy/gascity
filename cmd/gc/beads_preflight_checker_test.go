package main

import (
	"testing"

	"github.com/gastownhall/gascity/internal/beads/contract"
)

func TestPreflightDoltAuthUsesAmbientGCUserAndPassword(t *testing.T) {
	t.Setenv("GC_DOLT_USER", "superlzy")
	t.Setenv("GC_DOLT_PASSWORD", "secret")
	t.Setenv("BEADS_DOLT_PASSWORD", "stale")
	t.Setenv("BEADS_CREDENTIALS_FILE", "")

	target := contract.DoltConnectionTarget{
		Host:           "superlzy-dolt",
		Port:           "3306",
		Database:       "ac",
		EndpointOrigin: contract.EndpointOriginInheritedCity,
		External:       true,
	}

	user, password := preflightDoltAuth(t.TempDir(), t.TempDir(), target)
	if user != "superlzy" {
		t.Fatalf("user = %q, want ambient GC_DOLT_USER", user)
	}
	if password != "secret" {
		t.Fatalf("password = %q, want ambient GC_DOLT_PASSWORD", password)
	}
}
