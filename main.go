package similigo

import (
	"math"
	"strings"
)

// SimilarityOptions represents optional settings for hybrid similarity calculation.
// - NgramSize: The n-gram size used for n-gram similarity calculation.
// - WordSimWeight: Weight for word similarity in the final score.
// - NgramSimWeight: Weight for n-gram similarity in the final score.
// - ContainmentSimWeight: Weight for containment similarity in the final score.
type SimilarityOptions struct {
	NgramSize            int
	WordSimWeight        float64
	NgramSimWeight       float64
	ContainmentSimWeight float64
}

const (
	DefaultNgramSize            = 3
	DefaultWordSimWeight        = 0.5
	DefaultNgramSimWeight       = 0.3
	DefaultContainmentSimWeight = 0.2
)

func getFrequencyMap(text string) map[string]int {
	freqMap := make(map[string]int)
	tokens := strings.Fields(strings.ToLower(text))
	for _, token := range tokens {
		freqMap[token]++
	}
	return freqMap
}

func getNgramFrequencyMap(text string, n int) map[string]int {
	freqMap := make(map[string]int)
	runes := []rune(text)
	for i := 0; i <= len(runes)-n; i++ {
		ngram := string(runes[i : i+n])
		freqMap[ngram]++
	}
	return freqMap
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < a && b < c {
		return b
	}
	return c
}

func levenshteinDistance(a, b string) int {

	if len(a) == 0 {
		return len(b)
	}

	if len(b) == 0 {
		return len(a)
	}

	matrix := make([][]int, len(a)+1)
	for i := 0; i <= len(a); i++ {
		matrix[i] = make([]int, len(b)+1)
		matrix[i][0] = i
	}

	for j := 0; j <= len(b); j++ {
		matrix[0][j] = j
	}

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			cost := 0
			if a[i-1] != b[j-1] {
				cost = 1
			}

			matrix[i][j] = min(
				matrix[i-1][j]+1,
				matrix[i][j-1]+1,
				matrix[i-1][j-1]+cost,
			)
		}
	}

	return matrix[len(a)][len(b)]
}

func cosineSimilarity(text1, text2 string) float64 {
	freqMap1 := getFrequencyMap(text1)
	freqMap2 := getFrequencyMap(text2)

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

func ngramCosineSimilarity(text1, text2 string, n int) float64 {
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

func containmentSimilarity(text1, text2 string) float64 {
	keywordTokens := strings.Fields(strings.ToLower(text1))
	titleTokens := strings.Fields(strings.ToLower(text2))

	matchCount := 0.0
	for _, kToken := range keywordTokens {
		for _, tToken := range titleTokens {
			if strings.Contains(kToken, tToken) || strings.Contains(tToken, kToken) {
				matchCount++
			}
		}
	}
	return matchCount
}

// CalculateHybridSimilarity calculates a hybrid similarity score between two text strings.
// It combines different similarity measures (word similarity, n-gram similarity, and containment similarity)
// with custom weightings to provide an overall similarity score between the two texts.
//
// Parameters:
// - text1: The first text string for comparison.
// - text2: The second text string for comparison.
// - options: An optional struct that allows customization of n-gram size and weights.
//
// Returns:
// The hybrid similarity score, which is a weighted combination of the three similarity measures.
func CalculateHybridSimilarity(text1, text2 string, options *SimilarityOptions) float64 {
	if options == nil {
		options = &SimilarityOptions{
			NgramSize:            DefaultNgramSize,
			WordSimWeight:        DefaultWordSimWeight,
			NgramSimWeight:       DefaultNgramSimWeight,
			ContainmentSimWeight: DefaultContainmentSimWeight,
		}
	}

	if options.NgramSize == 0 {
		options.NgramSize = DefaultNgramSize
	}

	if options.WordSimWeight == 0.0 {
		options.WordSimWeight = DefaultWordSimWeight
	}

	if options.NgramSimWeight == 0.0 {
		options.NgramSimWeight = DefaultNgramSimWeight
	}

	if options.ContainmentSimWeight == 0.0 {
		options.ContainmentSimWeight = DefaultContainmentSimWeight
	}

	wordSim := cosineSimilarity(text1, text2)
	ngramSim := ngramCosineSimilarity(text1, text2, options.NgramSize)
	containmentSim := containmentSimilarity(text1, text2)

	return options.WordSimWeight*wordSim + options.NgramSimWeight*ngramSim + options.ContainmentSimWeight*containmentSim
}
