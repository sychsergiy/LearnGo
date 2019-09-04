package index_item

import (
	"log"
	"regexp"
	"sort"
	"strings"
	"xkcd/comic"
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

type Comic struct {
	TitleKeyWords      []string
	TranscriptKeyWords []string
	Num                int
}

func CreateComic(c comic.Comic) *Comic {
	//todo: remove empty strings
	excluded := []string{"a", "the", "an"}
	titleKeyWords := getKeyWords(c.Title, excluded)
	transcriptKeyWords := getKeyWords(c.Transcript, excluded)
	sort.Strings(titleKeyWords)
	sort.Strings(transcriptKeyWords)
	return &Comic{TitleKeyWords: titleKeyWords, TranscriptKeyWords: transcriptKeyWords, Num: c.Num}
}
