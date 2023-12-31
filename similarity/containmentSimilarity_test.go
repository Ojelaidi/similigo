package similarity

import (
	"testing"
)

func TestContainmentSimilarity(t *testing.T) {
	tests := []struct {
		text1    string
		text2    string
		expected float64
	}{
		{"hello world", "hello world", 1.0},     // full match
		{"hello world", "world", 1.0},           // partial match
		{"hello world", "hello there", 0.5},     // partial match
		{"hello world", "world hello", 1.0},     // full match with reversed order
		{"hello", "hi", 0.0},                    // no match
		{"hello", "hello hello", 1.0},           // repeated words
		{"hello world", "worldwide hello", 0.5}, // substring match
	}

	for _, tt := range tests {
		t.Run("Containment_"+tt.text1+"_"+tt.text2, func(t *testing.T) {
			got := ContainmentSimilarity(tt.text1, tt.text2)
			if got != tt.expected {
				t.Errorf("containmentSimilarity() = %v, want %v", got, tt.expected)
			}
		})
	}
}
