package similarity

import "math"

func getNgramFrequencyMap(text string, n int) map[string]int {
	freqMap := make(map[string]int)
	runes := []rune(text)
	for i := 0; i <= len(runes)-n; i++ {
		ngram := string(runes[i : i+n])
		freqMap[ngram]++
	}
	return freqMap
}
func NgramCosineSimilarity(text1, text2 string, n int) float64 {
	freqMap1 := getNgramFrequencyMap(text1, n)
	freqMap2 := getNgramFrequencyMap(text2, n)

	dotProduct := 0
	magnitude1 := 0
	magnitude2 := 0

	for ngram1, freq1 := range freqMap1 {
		magnitude1 += freq1 * freq1
		if freq2, exists := freqMap2[ngram1]; exists {
			dotProduct += freq1 * freq2
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
