package similigo

import "github.com/Ojelaidi/similigo/similarity"

// CalculateHybridSimilarity calculates a hybrid similarity score between two text strings.
// It combines different similarity measures (word similarity, n-gram similarity, and containment similarity)
// with custom weightings to provide an overall similarity score between the two texts.
//
// Parameters:
// - text1: The first text string for comparison.
// - text2: The second text string for comparison.
// - opts: An optional variadic parameter that allows customization of n-gram size and weights
//
// Returns:
// The hybrid similarity score, which is a weighted combination of the three similarity measures.
func CalculateHybridSimilarity(text1, text2 string, opts ...Option) float64 {
	options := DefaultSimilarityOptions()
	for _, opt := range opts {
		opt(options)
	}

	wordSim := similarity.CosineSimilarity(text1, text2)
	ngramSim := similarity.NgramCosineSimilarity(text1, text2, options.NgramSize)
	containmentSim := similarity.ContainmentSimilarity(text1, text2)

	return options.WordSimWeight*wordSim + options.NgramSimWeight*ngramSim + options.ContainmentSimWeight*containmentSim
}

// FindBestMatchInList takes a target text and a slice of texts, calculates the similarity for each,
// and returns the text with the highest similarity score.
func FindBestMatchInList(targetText string, texts []string, opts ...Option) (bestMatch string, highestScore float64) {
	highestScore = -1
	bestMatch = ""

	for _, text := range texts {
		score := CalculateHybridSimilarity(targetText, text, opts...)
		if score > highestScore {
			highestScore = score
			bestMatch = text
		}
	}

	return bestMatch, highestScore
}
