package item

import (
	"log"
	"regexp"
	"strings"
)

func contains(slice []string, item string) bool {
	for _, excludeWord := range slice {
		if strings.Compare(item, excludeWord) == 0 {
			return true
		}
	}
	return false
}

func cleanPunctuation(text string) string {
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(text, " ")
}

func getKeyWords(text string, wordsToExclude []string) []string {
	text = strings.ToLower(text)
	words := strings.Split(cleanPunctuation(text), " ")
	filteredWords := make([]string, 0, len(words))
	for _, word := range words {
		if !contains(wordsToExclude, word) {
			filteredWords = append(filteredWords, word)
		}
	}
	return filteredWords
}
