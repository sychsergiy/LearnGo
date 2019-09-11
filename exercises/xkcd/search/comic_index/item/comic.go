package item

import (
	"sort"
	"xkcd/comic"
	"xkcd/search/util"
)

type Comic struct {
	TitleKeyWords      []string
	TranscriptKeyWords []string
	Num                int
}

func New(c comic.Comic) *Comic {
	//todo: remove empty strings
	excluded := []string{"a", "the", "an"}
	titleKeyWords := util.GetKeyWords(c.Title, excluded)
	transcriptKeyWords := util.GetKeyWords(c.Transcript, excluded)
	sort.Strings(titleKeyWords)
	sort.Strings(transcriptKeyWords)
	return &Comic{TitleKeyWords: titleKeyWords, TranscriptKeyWords: transcriptKeyWords, Num: c.Num}
}
