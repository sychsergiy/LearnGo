package main

import (
	"encoding/json"
	"log"
	"xkcd/index"
	"xkcd/index/JSON"
	"xkcd/search/comic_index"
)

func CreateSearchIndexFromOfflineIndex(index index.Index) {
	comics := index.RetrieveAllComics()
	comic_index.New(comics)
}

func main() {
	jsonIndex := &JSON.Index{Name: "test"}

	comics := jsonIndex.BulkRetrieveComic([]int{1, 2, 3})
	for _, comic := range comics {
		data, _ := json.Marshal(comic)
		log.Println(string(data))
	}

	CreateSearchIndexFromOfflineIndex(jsonIndex)

	//_ = jsonIndex.Drop()
	//_ = jsonIndex.Create()
	//failed := comic_index.BulkFill(jsonIndex, 100)
	//log.Println(fmt.Sprintf("Finish filling comic_index, failed items: %d", failed))
	//_ = jsonIndex.Drop()

}
