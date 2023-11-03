package similarity

import (
	"testing"
)

func TestLevenshteinDistance(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected int
	}{
		{"kitten", "sitting", 3},
		{"", "", 0},
		{"book", "back", 2},
		{"hello", "hello", 0},
	}

	for _, tt := range tests {
		t.Run("Levenshtein_"+tt.a+"_"+tt.b, func(t *testing.T) {
			if got := levenshteinDistance(tt.a, tt.b); got != tt.expected {
				t.Errorf("levenshteinDistance() = %v, want %v", got, tt.expected)
			}
		})
	}
}
