package main

import (
	"log"
	"xkcd/comic"
	"xkcd/index/JSON"
	"xkcd/search"
	"xkcd/search/comic_index"
)

func CreateSearchIndexFromOfflineIndex() {
	var comicsChunk []comic.Comic

	jsonIndex := &JSON.Index{Name: "test"}
	searchIndex := comic_index.NewEmpty()

	for iterator, hasNext := jsonIndex.AllComicsIterator(10); hasNext; comicsChunk, hasNext = iterator() {
		searchIndex.Append(comicsChunk)
	}
}

func Search() {
	searchIndex := comic_index.ComicIndex{}
	nums := search.Search(searchIndex, "alt")
	log.Println(nums)

}

func main() {
	Search()
	//CreateSearchIndexFromOfflineIndex()
}
