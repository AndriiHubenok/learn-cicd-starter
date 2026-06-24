package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		expectedKey string
		expectErr   bool
	}{
		{
			name: "Success - Valid API Key",
			headers: http.Header{
				"Authorization": []string{"ApiKey super_secret_key_123"},
			},
			expectedKey: "super_secret_key_123",
			expectErr:   false,
		},
		{
			name:        "Failure - Missing Authorization Header",
			headers:     http.Header{},
			expectedKey: "",
			expectErr:   true,
		},
		{
			name: "Failure - Malformed Header (Wrong Prefix)",
			headers: http.Header{
				"Authorization": []string{"Bearer super_secret_key_123"},
			},
			expectedKey: "",
			expectErr:   true,
		},
		{
			name: "Failure - Malformed Header (No Key)",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expectedKey: "",
			expectErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)

			if (err != nil) != tt.expectErr {
				t.Fatalf("expected error: %v, got error: %v", tt.expectErr, err)
			}

			if key != tt.expectedKey {
				t.Errorf("expected key: %q, got: %q", tt.expectedKey, key)
			}
		})
	}
}
