package unit_test

import (
	"testing"
	"tig/app/domain/entities/organization"
)

func TestOrganizationEntityCreation(t *testing.T) {
	// Arrange & Act
	org := organization.New()

	// Assert
	if org == nil || len(org.GetUsers()) != 0 {
		t.Fatalf("expected empty initialized organization")
	}
}
