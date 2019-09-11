package search

import (
	"sort"
	"strings"
	"xkcd/search/comic_index"
	"xkcd/search/util"
)

type searchItem struct {
	Num          int
	MatchPercent int
}

func calcMatchPercent(words []string, queryWords []string) int {
	var matchPercentPerWord, matchPercentSum int
	matchPercentPerWord = 100 / len(queryWords)

	matchPercentSum = 0
	for _, word := range words {
		for _, queryWord := range queryWords {
			if strings.Contains(word, queryWord) {
				matchPercentSum += matchPercentPerWord
			}
		}
	}
	return matchPercentSum
}

func Search(index comic_index.ComicIndex, query string) []int {
	var matchPercent int
	var titleMatchedItems, descriptionMatchedItems []searchItem
	items := index.ReadAll()

	queryWords := util.GetKeyWords(query, []string{})

	for _, item := range items {
		matchPercent = calcMatchPercent(item.TitleKeyWords, queryWords)
		if matchPercent > 0 {
			titleMatchedItems = append(titleMatchedItems, searchItem{item.Num, matchPercent})
			break
		}

		matchPercent = calcMatchPercent(item.TranscriptKeyWords, queryWords)
		if matchPercent > 0 {
			descriptionMatchedItems = append(descriptionMatchedItems, searchItem{item.Num, matchPercent})
		}
	}

	sort.Slice(titleMatchedItems, func(i, j int) bool {
		return titleMatchedItems[i].MatchPercent < titleMatchedItems[j].MatchPercent
	})
	sort.Slice(descriptionMatchedItems, func(i, j int) bool {
		return descriptionMatchedItems[i].MatchPercent < descriptionMatchedItems[j].MatchPercent
	})

	nums := make([]int, 0, len(titleMatchedItems)+len(descriptionMatchedItems))
	for _, searchItem := range titleMatchedItems {
		nums = append(nums, searchItem.Num)
	}

	for _, searchItem := range descriptionMatchedItems {
		nums = append(nums, searchItem.Num)
	}
	return nums
}
