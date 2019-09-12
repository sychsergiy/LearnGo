package main

import (
	"log"
	"xkcd/comic"
	"xkcd/index"
	"xkcd/index/JSON"
	"xkcd/search"
	"xkcd/search/comic_index"
)

func createOfflineIndex(jsonIndex JSON.Index) {
	jsonIndex.Create()
	index.BulkFill(&jsonIndex, 100)
}

func createSearchIndexFromOfflineIndex(searchIndex comic_index.ComicIndex, offlineIndex index.Index) {
	searchIndex.Create()
	var comicsChunk []comic.Comic
	for iterator, hasNext := offlineIndex.AllComicsIterator(10); hasNext; comicsChunk, hasNext = iterator() {
		searchIndex.Append(comicsChunk)
	}
}

func Search(searchIndex *comic_index.ComicIndex, offlineIndex index.Index, query string) []comic.Comic {
	nums := search.Search(*searchIndex, query)
	comics := offlineIndex.BulkRetrieveComic(nums)
	return comics
}

func printComics(comics []comic.Comic) {
	for _, comic_ := range comics {
		log.Print("\n\n")
		log.Println(comic_)
	}
}
