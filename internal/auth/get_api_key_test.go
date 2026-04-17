package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_Valid(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-secret-key")

	key, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if key != "my-secret-key" {
		t.Fatalf("expected 'my-secret-key', got %s", key)
	}
}

func TestGetAPIKey_MissingHeader(t *testing.T) {
	headers := http.Header{}

	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("expected ErrNoAuthHeaderIncluded, got %v", err)
	}
}

func TestGetAPIKey_MalformedHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer my-secret-key") // wrong prefix

	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGetAPIKey_NoKeyProvided(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey") // missing key part

	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
