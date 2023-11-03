package similarity

import (
	"testing"
)

func TestNgramCosineSimilarity(t *testing.T) {
	tests := []struct {
		text1   string
		text2   string
		n       int
		wantSim float64
	}{
		{"night", "nacht", 2, 0.25},
		{"night", "night", 2, 1.0},
	}

	for _, tt := range tests {
		t.Run("Ngram_"+tt.text1+"_"+tt.text2, func(t *testing.T) {
			if gotSim := ngramCosineSimilarity(tt.text1, tt.text2, tt.n); gotSim != tt.wantSim {
				t.Errorf("ngramCosineSimilarity() = %v, want %v", gotSim, tt.wantSim)
			}
		})
	}
}
