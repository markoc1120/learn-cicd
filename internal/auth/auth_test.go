package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedErr   error
		expectedToken string
	}{
		{
			name:          "no header",
			headers:       http.Header{},
			expectedErr:   ErrNoAuthHeaderIncluded,
			expectedToken: "",
		},
		{
			name:          "wrong authorization in the header",
			headers:       http.Header{"Authorization": []string{"ApiKeys TOKEN_STRING"}},
			expectedErr:   ErrMalformedHeader,
			expectedToken: "",
		},
		{
			name:          "check correct token",
			headers:       http.Header{"Authorization": []string{"ApiKey TOKEN_STRING"}},
			expectedErr:   nil,
			expectedToken: "TOKEN_STRING",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := GetAPIKey(tt.headers)
			assertEqual(t, err, tt.expectedErr)
			assertEqual(t, token, tt.expectedToken)
		})
	}
}

func assertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
