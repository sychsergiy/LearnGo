package main

import (
	"log"
	"xkcd/comic"
	"xkcd/index"
	"xkcd/index/JSON"
	"xkcd/search"
	"xkcd/search/comic_index"
)

func createOfflineIndex() *JSON.Index {
	jsonIndex := &JSON.Index{Name: "test"}
	jsonIndex.Create()
	index.BulkFill(jsonIndex, 100)
	return jsonIndex
}

func createSearchIndexFromOfflineIndex(offlineIndex index.Index) *comic_index.ComicIndex {
	var comicsChunk []comic.Comic

	searchIndex := comic_index.NewEmpty()
	for iterator, hasNext := offlineIndex.AllComicsIterator(10); hasNext; comicsChunk, hasNext = iterator() {
		searchIndex.Append(comicsChunk)
	}
	return &searchIndex
}

func Search(searchIndex *comic_index.ComicIndex, offlineIndex index.Index) []comic.Comic {
	nums := search.Search(*searchIndex, "back ball")
	comics := offlineIndex.BulkRetrieveComic(nums)
	return comics
}

func printComics(comics []comic.Comic) {
	for _, comic_ := range comics {
		log.Print("\n")
		log.Println(comic_)
	}
}

func main() {
	offlineIndex := createOfflineIndex()
	searchIndex := createSearchIndexFromOfflineIndex(offlineIndex)
	foundComics := Search(searchIndex, offlineIndex)
	printComics(foundComics)
}
