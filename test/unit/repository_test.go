package unit_test

import (
	"testing"
	"tig/app/domain/entities/repository"
)

func TestRepositoryEntityValidation(t *testing.T) {
	t.Run("empty alias", func(t *testing.T) {
		// Arrange
		draft := repository.Draft{
			Alias:  "",
			Origin: "git@github.com:sneuder/tig.git",
		}

		// Act
		_, err := repository.New(draft)

		// Assert
		if err != repository.ErrEmptyAlias {
			t.Fatalf("expected ErrEmptyAlias, got: %v", err)
		}
	})

	t.Run("empty origin", func(t *testing.T) {
		// Arrange
		draft := repository.Draft{
			Alias:  "tig",
			Origin: "",
		}

		// Act
		_, err := repository.New(draft)

		// Assert
		if err != repository.ErrEmptyOrigin {
			t.Fatalf("expected ErrEmptyOrigin, got: %v", err)
		}
	})

	t.Run("valid repository", func(t *testing.T) {
		// Arrange
		draft := repository.Draft{
			Alias:  "tig",
			Origin: "git@github.com:sneuder/tig.git",
		}

		// Act
		repo, err := repository.New(draft)

		// Assert
		if err != nil || repo == nil {
			t.Fatalf("expected valid repo creation, got error: %v", err)
		}
	})
}
