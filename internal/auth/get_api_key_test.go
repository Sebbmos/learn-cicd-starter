package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// We define a "table" of different test cases
	tests := []struct {
		name       string
		headers    http.Header
		wantKey    string
		wantErr    bool
	}{
		{
			name:    "Valid API Key",
			headers: http.Header{"Authorization": []string{"ApiKey my-secret-key"}},
			wantKey: "my-secret-key",
			wantErr: false,
		},
		{
			name:    "Missing Header",
			headers: http.Header{}, // Empty headers
			wantKey: "",
			wantErr: true,
		},
		{
			name:    "Malformed Header",
			headers: http.Header{"Authorization": []string{"Bearer my-secret-key"}},
			wantKey: "",
			wantErr: true,
		},
	}

	// Loop through each test case and run it
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotKey, err := GetAPIKey(tc.headers)

			// Check if we got an error when we weren't expecting one (or vice versa)
			if (err != nil) != tc.wantErr {
				t.Fatalf("GetAPIKey() error = %v, wantErr %v", err, tc.wantErr)
			}
			
			// Check if the key matches what we expected
			if gotKey != tc.wantKey {
				t.Fatalf("GetAPIKey() got = %v, want %v", gotKey, tc.wantKey)
			}
		})
	}
}