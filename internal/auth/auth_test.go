package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Test case: No Authorization Header
	headers := http.Header{}
	_, err := GetAPIKey(headers)
	if err == nil {
		t.Errorf("expected error but got none")
	}

	// Test case: Valid Authorization Header
	headers.Set("Authorization", "ApiKey my-api-key")
	key, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("did not expect an error, but got: %v", err)
	}
	if key != "wrong-api-key" {
		t.Errorf("expected API key 'my-api-key', but got: %v", key)
	}
}