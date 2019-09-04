package search

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"xkcd/comic"
	"xkcd/search/index_item"
)

type ComicIndex struct {
}

const name = "xkcd"

const fileNamePrefix = "search_index_"

func (index *ComicIndex) getSourceFilePath() string {
	return fileNamePrefix + name + ".json"
}

func (index *ComicIndex) ReadAll() []index_item.Comic {
	file, err := os.OpenFile(index.getSourceFilePath(), os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(file)
	indexItems := make([]index_item.Comic, 0, 0)
	err = json.Unmarshal(content, &indexItems)
	if err != nil {
		log.Fatal(err)
	}
	return indexItems
}

func (index *ComicIndex) Write(comicIndexItems []index_item.Comic) {
	file, err := os.OpenFile(index.getSourceFilePath(), os.O_WRONLY|os.O_CREATE, 0660)
	if err != nil {
		log.Fatal(err)
	}
	jsonData, err := json.Marshal(comicIndexItems)

	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}
}

func (index *ComicIndex) Search() {

}

func NewIndex(comics []comic.Comic) ComicIndex {
	comicIndexItems := make([]index_item.Comic, 0, len(comics))

	for _, comic_ := range comics {
		comicIndexItems = append(comicIndexItems, *index_item.CreateComic(comic_))
	}

	comicIndex := ComicIndex{}
	comicIndex.Write(comicIndexItems)
	return comicIndex
}
