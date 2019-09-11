package main

import (
	"log"
	"xkcd/comic"
	"xkcd/index/JSON"
	"xkcd/search"
	"xkcd/search/comic_index"
)

func CreateSearchIndexFromOfflineIndex() {
	jsonIndex := &JSON.Index{Name: "test"}
	searchIndex := comic_index.NewEmpty()

	var comicsChunk, comics []comic.Comic
	for iterator, hasNext := jsonIndex.AllComicsIterator(10); hasNext; comicsChunk, hasNext = iterator() {
		comics = append(comics, comicsChunk...)
		searchIndex.Append(comics)
	}
}

func Search() {
	searchIndex := comic_index.ComicIndex{}
	nums := search.Search(searchIndex, "back")
	log.Println(nums)

}

func main() {
	//Search()
	CreateSearchIndexFromOfflineIndex()
}
