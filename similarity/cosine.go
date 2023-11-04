package similarity

import (
	"github.com/Ojelaidi/similigo/utils"
	"math"
)

func CosineSimilarity(text1, text2 string) float64 {
	freqMap1 := utils.GetFrequencyMap(text1)
	freqMap2 := utils.GetFrequencyMap(text2)

	dotProduct := 0.0
	magnitude1 := 0.0
	magnitude2 := 0.0

	for token, freq1 := range freqMap1 {
		if freq2, exists := freqMap2[token]; exists {
			dotProduct += float64(freq1 * freq2)
		}
		magnitude1 += float64(freq1 * freq1)
	}

	for _, freq := range freqMap2 {
		magnitude2 += float64(freq * freq)
	}

	magnitude := math.Sqrt(magnitude1) * math.Sqrt(magnitude2)
	if magnitude == 0 {
		return 0
	}
	return dotProduct / magnitude
}
