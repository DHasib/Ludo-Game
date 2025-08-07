package auth

import "testing"

func TestPasswordHash(t *testing.T) {
	password := "secret"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("hash error: %v", err)
	}
	if !CheckPasswordHash(password, hash) {
		t.Fatal("password should match")
	}
}
