package similarity

import "strings"

func ContainmentSimilarity(text1, text2 string) float64 {
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
