package lib_test

import (
	"testing"

	"github.com/pilinux/gorest/lib"
)

func TestValidateEmail(t *testing.T) {
	testCases := []struct {
		email string
		want  bool
	}{
		{"test@example.com", false}, // RFC 7505 test case
		{"in", false},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@gmail.com", false},
		{"invalid", false},
		{"invalid@", false},
		{"invalid@[127.0.0.1]", false},
		{"me@no-destination.pilinux.me", false},
		{"@missing-local-part.com", false},
		{"hello@world@google.com", false},
		{"user name@google.com", false},
		{"security@google.com", true},
	}

	for i := range testCases {
		tc := testCases[i]
		got := lib.ValidateEmail(tc.email)
		if got != tc.want {
			t.Errorf("lib.ValidateEmail(%q) = %v, want %v", tc.email, got, tc.want)
		}
	}
}
