package utils

import (
	"testing"

	"github.com/google/uuid"
)

func TestValidateJWTConfigRequiresSecret(t *testing.T) {
	t.Setenv("JWT_SECRET", "")

	err := ValidateJWTConfig()
	if err == nil {
		t.Fatal("expected ValidateJWTConfig to fail when JWT_SECRET is empty")
	}
}

func TestGenerateAndValidateToken(t *testing.T) {
	t.Setenv("JWT_SECRET", "test-secret")

	userID := uuid.New()
	email := "user@example.com"

	token, err := GenerateToken(userID, email)
	if err != nil {
		t.Fatalf("GenerateToken returned error: %v", err)
	}

	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("ValidateToken returned error: %v", err)
	}

	if claims.UserID != userID {
		t.Fatalf("expected user id %s, got %s", userID, claims.UserID)
	}

	if claims.Email != email {
		t.Fatalf("expected email %q, got %q", email, claims.Email)
	}

	if claims.ExpiresAt == nil || claims.IssuedAt == nil {
		t.Fatal("expected registered claims to include issued and expiration timestamps")
	}
}

func TestValidateTokenRejectsTokenSignedWithDifferentSecret(t *testing.T) {
	t.Setenv("JWT_SECRET", "first-secret")

	token, err := GenerateToken(uuid.New(), "user@example.com")
	if err != nil {
		t.Fatalf("GenerateToken returned error: %v", err)
	}

	t.Setenv("JWT_SECRET", "second-secret")

	_, err = ValidateToken(token)
	if err == nil {
		t.Fatal("expected ValidateToken to reject a token signed with a different secret")
	}
}
