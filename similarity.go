package similigo

import (
	"github.com/Ojelaidi/similigo/similarity"
	"github.com/Ojelaidi/similigo/utils"
	"strings"
)

// PreprocessText processes the input text for similarity comparison by performing several steps:
//  1. Tokenization: Splitting the text into words (tokens) based on whitespace.
//  2. Normalization: Converting all words to lowercase to ensure case insensitivity.
//  3. Stop word removal: Eliminating common words (stop words) that are unlikely to be useful
//     in the similarity comparison. It uses both a predefined set of stop words and any custom
//     stop words provided in the SimilarityOptions.
//  4. Stemming: Reducing words to their base or root form (stem).
//
// Parameters:
//   - text: The input text to preprocess.
//   - opts: A pointer to SimilarityOptions which contains settings for the preprocessing,
//     including any custom stop words to consider.
//
// Returns:
// A preprocessed version of the input text with all words stemmed and stop words removed,
// joined into a single string separated by spaces.
func PreprocessText(text string, opts *SimilarityOptions) string {
	words := strings.Fields(text)
	var preprocessedWords []string

	for _, word := range words {
		word = strings.ToLower(word)
		if !utils.IsStopWord(word) && !opts.CustomStopWords[word] {
			stemmedWord, err := utils.StemWord(word)
			if err != nil {
				return ""
			}
			preprocessedWords = append(preprocessedWords, stemmedWord)
		}
	}

	return strings.Join(preprocessedWords, " ")
}

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

	preprocessedText1 := PreprocessText(text1, options)
	preprocessedText2 := PreprocessText(text2, options)

	wordSim := similarity.CosineSimilarity(preprocessedText1, preprocessedText2)
	ngramSim := similarity.NgramCosineSimilarity(preprocessedText1, preprocessedText2, options.NgramSize)
	containmentSim := similarity.ContainmentSimilarity(preprocessedText1, preprocessedText2)

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
