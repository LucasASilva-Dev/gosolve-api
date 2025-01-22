package version

import (
	"testing"
)

func TestVersion(t *testing.T) {
	tests := []struct {
		name       string
		IDENTIFIER string
		COMMIT     string
		expected   string
	}{
		{"no identifier and no commit", "", "", "0.0.1"},
		{"identifier and no commit", "test-identifier", "", "0.0.1-test-identifier"},
		{"no identifier and commit", "", "test-commit", "0.0.1+commit.test-commit"},
		{"identifier and commit", "test-identifier", "test-commit", "0.0.1-test-identifier+commit.test-commit"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IDENTIFIER = tt.IDENTIFIER
			COMMIT = tt.COMMIT
			actual := Version()
			if actual != tt.expected {
				t.Errorf("Version() = %s, want %s", actual, tt.expected)
			}
		})
	}
}
