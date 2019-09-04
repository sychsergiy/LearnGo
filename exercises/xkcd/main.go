package main

import (
	"xkcd/comic"
	"xkcd/index"
	"xkcd/index/JSON"
	"xkcd/search/comic_index"
)

func CreateSearchIndexFromOfflineIndex(index index.Index) {
	searchIndex := comic_index.NewEmpty()

	var comicsChunk, comics []comic.Comic
	for iterator, hasNext := index.AllComicsIterator(10); hasNext; comicsChunk, hasNext = iterator() {
		comics = append(comics, comicsChunk...)
		searchIndex.Append(comics)
	}
}

func main() {
	jsonIndex := &JSON.Index{Name: "test"}
	CreateSearchIndexFromOfflineIndex(jsonIndex)

	//_ = jsonIndex.Drop()
	//_ = jsonIndex.Create()
	//failed := comic_index.BulkFill(jsonIndex, 100)
	//log.Println(fmt.Sprintf("Finish filling comic_index, failed items: %d", failed))
	//_ = jsonIndex.Drop()

}
