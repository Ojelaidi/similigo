package utils

import (
	"github.com/kljensen/snowball"
	"strings"
)

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

var FrenchStopWords = map[string]bool{
	"alors": true, "au": true, "aucuns": true, "aussi": true, "autre": true, "avant": true, "avec": true,
	"avoir": true, "bon": true, "car": true, "ce": true, "cela": true, "ces": true, "ceux": true,
	"chaque": true, "ci": true, "comme": true, "comment": true, "dans": true, "des": true, "du": true,
	"dedans": true, "dehors": true, "depuis": true, "devrait": true, "doit": true, "donc": true,
	"dos": true, "début": true, "elle": true, "elles": true, "en": true, "encore": true, "essai": true,
	"est": true, "et": true, "eu": true, "fait": true, "faites": true, "fois": true, "font": true,
	"force": true, "haut": true, "hors": true, "ici": true, "il": true, "ils": true, "je": true,
	"juste": true, "la": true, "le": true, "les": true, "leur": true, "là": true, "ma": true,
	"maintenant": true, "mais": true, "mes": true, "mine": true, "moins": true, "mon": true, "mot": true,
	"même": true, "ni": true, "nommés": true, "notre": true, "nous": true, "ou": true, "où": true,
	"par": true, "parce": true, "pas": true, "peut": true, "peu": true, "plupart": true, "pour": true,
	"pourquoi": true, "quand": true, "que": true, "quel": true, "quelle": true, "quelles": true,
	"quels": true, "qui": true, "sa": true, "sans": true, "ses": true, "seulement": true, "si": true,
	"sien": true, "son": true, "sont": true, "sous": true, "soyez": true, "sujet": true, "sur": true,
	"ta": true, "tandis": true, "tellement": true, "tels": true, "tes": true, "ton": true, "tous": true,
	"tout": true, "trop": true, "très": true, "tu": true, "voient": true, "vont": true, "votre": true,
	"vous": true, "vu": true, "ça": true, "étaient": true, "état": true, "étions": true, "été": true,
	"être": true,
}

func IsStopWord(word string) bool {
	_, exists := FrenchStopWords[word]
	return exists
}

func StemWord(word string) (string, error) {
	stemmed, err := snowball.Stem(word, "french", true)
	if err != nil {
		return "", err
	}
	return stemmed, nil
}
