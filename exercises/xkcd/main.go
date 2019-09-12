package main

import (
	"os"
	"xkcd/index/JSON"
	"xkcd/search/comic_index"
)

func main() {
	jsonIndex := &JSON.Index{Name: "test"}
	searchIndex := &comic_index.ComicIndex{}

	action := os.Args[1]
	if action == "create-offline-index" {
		createOfflineIndex(*jsonIndex)
	} else if action == "create-search-index" {
		createSearchIndexFromOfflineIndex(*searchIndex, jsonIndex)
	} else if action == "search" {
		query := os.Args[2]
		foundComics := Search(searchIndex, jsonIndex, query)
		printComics(foundComics)
	} else {
		panic("Available commands: create-offline-index, create-search-index, search")
	}
}
