package search

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
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

type ComicIndexItem struct {
	TitleKeyWords      []string
	TranscriptKeyWords []string
	Num                int
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

func ReadComicIndex() []ComicIndexItem {
	file, err := os.OpenFile("search_index.json", os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(file)
	indexItems := make([]ComicIndexItem, 0, 0)
	err = json.Unmarshal(content, &indexItems)
	if err != nil {
		log.Fatal(err)
	}
	return indexItems
}

func WriteComicIndex(comicIndexItems []ComicIndexItem) {
	file, err := os.OpenFile("search_index.json", os.O_WRONLY|os.O_CREATE, 0660)
	if err != nil {
		log.Fatal(err)
	}
	jsonData, err := json.Marshal(comicIndexItems)

	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}
}
