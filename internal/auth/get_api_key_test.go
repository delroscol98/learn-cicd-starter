package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		input http.Header
		want  string
	}

	tests := []test{
		{input: http.Header{"Authorization": []string{"ApiKey 1234567890abc"}}, want: "1234567890abc"},
		{input: http.Header{"Authorization": []string{"ApiKey"}}, want: ""},
		{input: http.Header{"Authorization": []string{"NotApiKey"}}, want: "hello"},
		{input: http.Header{"Authorization": []string{""}}, want: ""},
	}

	for _, tc := range tests {
		got, _ := GetAPIKey(tc.input)
		if got != tc.want {
			t.Fatalf("Expected: %v, got: %v", tc.want, got)
		}
	}
}
