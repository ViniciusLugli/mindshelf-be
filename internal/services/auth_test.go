package services

import "testing"

func TestHashPasswordAndCheckPassword(t *testing.T) {
	password := "super-secret-password"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword returned error: %v", err)
	}

	if hash == password {
		t.Fatal("expected hashed password to differ from the original password")
	}

	if !CheckPassword(password, hash) {
		t.Fatal("expected CheckPassword to accept the original password")
	}

	if CheckPassword("wrong-password", hash) {
		t.Fatal("expected CheckPassword to reject an invalid password")
	}
}
