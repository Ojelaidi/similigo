package similarity

import (
	"github.com/Ojelaidi/similigo/utils"
	"math"
	"strings"
)

func CosineSimilarity(text1, text2 string) float64 {
	freqMap1 := utils.GetFrequencyMap(text1)
	freqMap2 := utils.GetFrequencyMap(text2)

	dotProduct := 0
	magnitude1 := 0
	magnitude2 := 0

	for _, freq := range freqMap1 {
		magnitude1 += freq * freq
	}

	for token1, freq1 := range freqMap1 {
		for token2, freq2 := range freqMap2 {
			if levenshteinDistance(strings.ToLower(token1), strings.ToLower(token2)) <= 2 {
				dotProduct += freq1 * freq2
			}
		}
	}

	for _, freq := range freqMap2 {
		magnitude2 += freq * freq
	}

	magnitude := math.Sqrt(float64(magnitude1)) * math.Sqrt(float64(magnitude2))
	if magnitude == 0 {
		return 0
	}
	return float64(dotProduct) / magnitude
}
