package unit_test

import (
	"testing"
	"tig/app/domain/entities/user"
)

func TestUserEntityValidation(t *testing.T) {
	t.Run("empty alias", func(t *testing.T) {
		// Arrange
		draft := user.Draft{
			Alias: "",
			Name:  "Sneuder",
			Email: "sneuder@tig.com",
			Host:  "g",
		}

		// Act
		_, err := user.New(draft)

		// Assert
		if err != user.ErrEmptyAlias {
			t.Fatalf("expected ErrEmptyAlias, got: %v", err)
		}
	})

	t.Run("empty name", func(t *testing.T) {
		// Arrange
		draft := user.Draft{
			Alias: "sneuder",
			Name:  "",
			Email: "sneuder@tig.com",
			Host:  "g",
		}

		// Act
		_, err := user.New(draft)

		// Assert
		if err != user.ErrEmptyName {
			t.Fatalf("expected ErrEmptyName, got: %v", err)
		}
	})

	t.Run("invalid email format", func(t *testing.T) {
		// Arrange
		draft := user.Draft{
			Alias: "sneuder",
			Name:  "Sneuder",
			Email: "not-an-email",
			Host:  "g",
		}

		// Act
		_, err := user.New(draft)

		// Assert
		if err != user.ErrInvalidEmail {
			t.Fatalf("expected ErrInvalidEmail, got: %v", err)
		}
	})

	t.Run("valid creation", func(t *testing.T) {
		// Arrange
		draft := user.Draft{
			Alias: "sneuder",
			Name:  "Sneuder",
			Email: "sneuder@tig.com",
			Host:  "g",
		}

		// Act
		u, err := user.New(draft)

		// Assert
		if err != nil || u == nil {
			t.Fatalf("expected successful user creation, got error: %v", err)
		}
	})
}
