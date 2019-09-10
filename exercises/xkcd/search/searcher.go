package search

import (
	"sort"
	"strings"
	"xkcd/search/comic_index"
)

type searchItem struct {
	Num          int
	MatchPercent int
	Field        string
}

func Search(index comic_index.ComicIndex, query string) []int {
	// todo: create set of words from query(only one word now)
	var searchItems []searchItem
	items := index.ReadAll()

	for _, item := range items {

		for _, titleKeyWord := range item.TitleKeyWords {
			if strings.Contains(titleKeyWord, query) {
				searchItems = append(searchItems, searchItem{item.Num, 100, "title"})
			}
		}

		for _, transcriptKeyWord := range item.TranscriptKeyWords {
			if strings.Contains(transcriptKeyWord, query) {
				searchItems = append(searchItems, searchItem{item.Num, 100, "description"})
			}
		}
	}

	sort.Slice(searchItems, func(i, j int) bool {
		if searchItems[i].Field == "description" && searchItems[j].Field == "title" {
			return true
		} else if searchItems[i].Field == "title" && searchItems[j].Field == "description" {
			return false
		} else {
			return false
		}
	})

	sort.Slice(searchItems, func(i, j int) bool {
		return searchItems[i].MatchPercent < searchItems[j].MatchPercent
	})

	numsSet := make(map[int]bool)
	for _, searchItem := range searchItems {
		numsSet[searchItem.Num] = true
	}

	nums := make([]int, 0, len(searchItems))
	for key, _ := range numsSet {
		nums = append(nums, key)
	}
	return nums
}
