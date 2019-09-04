package search

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type ComicIndex struct {
}

const name = "xkcd"

const fileNamePrefix = "search_index_"

func (index *ComicIndex) getSourceFilePath() string {
	return fileNamePrefix + name + ".json"
}

func (index *ComicIndex) ReadAll() []ComicIndexItem {
	file, err := os.OpenFile(index.getSourceFilePath(), os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(file)
	indexItems := make([]ComicIndexItem, 0, 0)
	err = json.Unmarshal(content, &indexItems)
	if err != nil {
		log.Fatal(err)
	}
	return indexItems
}

func (index *ComicIndex) Write(comicIndexItems []ComicIndexItem) {
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
