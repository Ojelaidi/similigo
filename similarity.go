package similigo

import "github.com/Ojelaidi/similigo/similarity"

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
