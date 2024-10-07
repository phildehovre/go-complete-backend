package auth

import (
	"testing"
)

func TestPasswordGeneration(t *testing.T) {
	t.Run("generates a hashed password", func(*testing.T) {
		hashed, err := HashPassword([]byte("password"))

		if err != nil {
			t.Errorf("error hashing password")
		}

		if hashed == "" {
			t.Errorf("expecting password not to be empty")
		}

		if hashed == "password" {
			t.Errorf("expecting password to be different: %s, %s", hashed, "password")
		}
	})
}
