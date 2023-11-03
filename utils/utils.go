package utils

import "strings"

func GetFrequencyMap(text string) map[string]int {
	freqMap := make(map[string]int)
	tokens := strings.Fields(strings.ToLower(text))
	for _, token := range tokens {
		freqMap[token]++
	}
	return freqMap
}

func Min(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < a && b < c {
		return b
	}
	return c
}
