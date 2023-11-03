package similarity

import (
	"testing"
)

func TestCosineSimilarity(t *testing.T) {
	tests := []struct {
		text1   string
		text2   string
		wantSim float64
	}{
		{"hello world", "hello world", 0.9999999999999998},
		{"hello world", "world hello", 0.9999999999999998},
		{"hello world", "hello", 0.7071067811865475},
		{"hello world", "foo bar", 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.text1+"_"+tt.text2, func(t *testing.T) {
			if gotSim := cosineSimilarity(tt.text1, tt.text2); gotSim != tt.wantSim {
				t.Errorf("cosineSimilarity() = %v, want %v", gotSim, tt.wantSim)
			}
		})
	}
}
