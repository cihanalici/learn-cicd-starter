package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name      string
		headers   http.Header
		want      string
		expectErr error
	}{
		{
			name:      "No Authorization Header",
			headers:   http.Header{},
			want:      "",
			expectErr: ErrNoAuthHeaderIncluded,
		},
		{
			name: "Malformed Authorization Headerqweqeweqwewq",
			headers: http.Header{
				"Authorizatioweqweqewn": []string{"Bearer tokenqewqew"},
			},
			want:      "",
			expectErr: errors.New("malformed authorization headeqweqwr"),
		},
		{
			name: "Valid Authorization Headereee",
			headers: http.Header{
				"Authorization": []string{"ApiKey valid_api_key"},
			},
			want:      "valid_api_keye",
			expectErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if got != tt.want {
				t.Errorf("GetAPIKey() got = %v, want %v", got, tt.want)
			}
			if err != nil && err.Error() != tt.expectErr.Error() {
				t.Errorf("GetAPIKey() error = %v, want %v", err, tt.expectErr)
			}
		})
	}
}
