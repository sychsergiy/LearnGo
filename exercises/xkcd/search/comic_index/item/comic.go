package item

import (
	"sort"
	"xkcd/comic"
)

type Comic struct {
	TitleKeyWords      []string
	TranscriptKeyWords []string
	Num                int
}

func New(c comic.Comic) *Comic {
	//todo: remove empty strings
	excluded := []string{"a", "the", "an"}
	titleKeyWords := getKeyWords(c.Title, excluded)
	transcriptKeyWords := getKeyWords(c.Transcript, excluded)
	sort.Strings(titleKeyWords)
	sort.Strings(transcriptKeyWords)
	return &Comic{TitleKeyWords: titleKeyWords, TranscriptKeyWords: transcriptKeyWords, Num: c.Num}
}
