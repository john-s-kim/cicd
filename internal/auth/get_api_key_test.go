package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		header  string
		wantKey string
		wantErr bool
	}{
		{
			name:    "valid API key",
			header:  "ApiKey abc123",
			wantKey: "abc123",
		},
		{
			name:    "missing authorization header",
			wantErr: true,
		},
		{
			name:    "wrong authentication scheme",
			header:  "Bearer abc123",
			wantErr: true,
		},
		{
			name:    "missing API key",
			header:  "ApiKey",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}
			if tt.header != "" {
				headers.Set("Authorization", tt.header)
			}

			got, err := GetAPIKey(headers)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected an error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != tt.wantKey {
				t.Errorf("got %q, want %q", got, tt.wantKey)
			}
		})
	}
}
