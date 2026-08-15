package utils

import "testing"

func TestGenerateAndValidateToken(t *testing.T) {
	token, err := GenerateToken("alice", "secret-key", 1)
	if err != nil {
		t.Fatalf("GenerateToken() error = %v", err)
	}
	if token == "" {
		t.Fatal("GenerateToken() returned empty token")
	}

	claims, err := ValidateToken(token, "secret-key")
	if err != nil {
		t.Fatalf("ValidateToken() error = %v", err)
	}
	if claims.Username != "alice" {
		t.Fatalf("Username mismatch: got %q, want %q", claims.Username, "alice")
	}
}

func TestValidateToken_InvalidToken(t *testing.T) {
	if _, err := ValidateToken("not-a-valid-token", "secret-key"); err == nil {
		t.Fatal("ValidateToken() expected an error for invalid token")
	}
}

func TestHashPasswordAndCheckPasswordHash(t *testing.T) {
	hash, err := HashPassword("P@ssw0rd")
	if err != nil {
		t.Fatalf("HashPassword() error = %v", err)
	}
	if hash == "" {
		t.Fatal("HashPassword() returned empty hash")
	}

	if !CheckPasswordHash("P@ssw0rd", hash) {
		t.Fatal("CheckPasswordHash() should return true for the correct password")
	}

	if CheckPasswordHash("wrong-password", hash) {
		t.Fatal("CheckPasswordHash() should return false for the wrong password")
	}
}
