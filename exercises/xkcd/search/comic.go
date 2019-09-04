package search

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

func CreateComicIndexItem(comic comic.Comic) *ComicIndexItem {
	//todo: remove empty strings
	excluded := []string{"a", "the", "an"}
	titleKeyWords := getKeyWords(comic.Title, excluded)
	transcriptKeyWords := getKeyWords(comic.Transcript, excluded)
	sort.Strings(titleKeyWords)
	sort.Strings(transcriptKeyWords)
	return &ComicIndexItem{TitleKeyWords: titleKeyWords, TranscriptKeyWords: transcriptKeyWords, Num: comic.Num}
}

type ComicIndexItem struct {
	TitleKeyWords      []string
	TranscriptKeyWords []string
	Num                int
}
